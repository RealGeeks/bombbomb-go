package bombbomb_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RealGeeks/bombbomb-go"

	. "github.com/igorsobreira/testing"
)

//
// AddContact
//

func TestAddContact(t *testing.T) {
	var requestBody url.Values
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBody = readValues(t, r)
		io.WriteString(w, Stubs["AddContact"])
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	contactToCreate := bombbomb.Contact{
		FirstName:   "Jack",
		LastName:    "Johnson",
		Email:       "jj@gmail.com",
		PhoneNumber: "808-123-4321",
	}
	contactCreated, err := cli.AddContact(contactToCreate)

	expectedRequestBody := contactToCreate.Values()
	expectedRequestBody["api_key"] = []string{"123"}

	Ok(t, err)
	Equals(t, expectedRequestBody, requestBody)
	Equals(t, bombbomb.Contact{
		ID:          "106e0e29-e9cf-b812-b895-dcdc059cf9ec",
		FirstName:   "Jack",
		LastName:    "Johnson",
		Email:       "jj@gmail.com",
		PhoneNumber: "808-123-4321",
	}, contactCreated)
}

//
// CreateList
//

func TestCreateList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, Stubs["CreateList"])
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	info, err := cli.CreateList(bombbomb.List{
		Name: "Buyers",
	})

	Ok(t, err)
	Equals(t, "4184993a-b98e-e9e4-19b6-da1019d9cd3d", info.ID)
	Equals(t, "Buyers", info.Name)
}

//
// GetLists
//

func TestGetLists(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, Stubs["GetLists"])
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	lists, err := cli.GetLists()

	Ok(t, err)
	Equals(t, 2, len(lists))

	Equals(t, "4184993a-b98e-e9e4-19b6-da1019d9cd3d", lists[0].ID)
	Equals(t, "Partners", lists[0].Name)
	Equals(t, "2", lists[0].ContactCount)

	Equals(t, "3c20f8a3-2d95-8966-4add-0957dd0d23c5", lists[1].ID)
	Equals(t, "Suppression List", lists[1].Name)
	Equals(t, "0", lists[1].ContactCount)
}

//
// EnsureList
//

func TestEnsureList_CreateWhenNotFound(t *testing.T) {
	responses := []string{Stubs["GetLists"], Stubs["CreateList"]}
	respcount := 0
	requests := []*http.Request{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r)
		io.WriteString(w, responses[respcount])
		respcount += 1
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	list, err := cli.EnsureList(bombbomb.List{
		Name: "Buyers",
	})

	Ok(t, err)
	Equals(t, "Buyers", list.Name)
	Equals(t, "4184993a-b98e-e9e4-19b6-da1019d9cd3d", list.ID)

	Equals(t, 2, len(requests))
	Equals(t, "GET", requests[0].Method)
	Equals(t, "POST", requests[1].Method)
}

func TestEnsureList_GetWhenFound(t *testing.T) {
	responses := []string{Stubs["GetLists"], Stubs["CreateList"]}
	respcount := 0
	requests := []*http.Request{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r)
		io.WriteString(w, responses[respcount])
		respcount += 1
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	list, err := cli.EnsureList(bombbomb.List{
		Name: "Partners",
	})

	Ok(t, err)
	Equals(t, "Partners", list.Name)
	Equals(t, "4184993a-b98e-e9e4-19b6-da1019d9cd3d", list.ID)

	Equals(t, 1, len(requests))
	Equals(t, "GET", requests[0].Method)
}

//
// Specific errors
//

func TestErrNoSubscription(t *testing.T) {

	// all API calls should return ErrNoSubscription
	var tests = []struct {
		Name        string
		MakeRequest func(cli *bombbomb.Client) error
	}{
		{
			Name: "CreateList",
			MakeRequest: func(bb *bombbomb.Client) error {
				_, err := bb.CreateList(bombbomb.List{Name: "Buyers"})
				return err
			},
		},
		{
			Name: "AddContact",
			MakeRequest: func(bb *bombbomb.Client) error {
				_, err := bb.AddContact(bombbomb.Contact{
					FirstName:   "Jack",
					LastName:    "Johnson",
					Email:       "jj@gmail.com",
					PhoneNumber: "808-123-4321",
				})
				return err
			},
		},
		{
			Name: "GetLists",
			MakeRequest: func(bb *bombbomb.Client) error {
				_, err := bb.GetLists()
				return err
			},
		},
		{
			Name: "EnsureList",
			MakeRequest: func(bb *bombbomb.Client) error {
				_, err := bb.EnsureList(bombbomb.List{Name: "Buyers"})
				return err
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		io.WriteString(w, `{"methodName":"GetLists","status":"failure","info":"This account {\"user_id\":\"41808bb6-f086-bc0a-3460-3955c752804d\"} does not have an active subscription. Please contact support."}`)
	}))
	defer ts.Close()

	cli := &bombbomb.Client{URL: ts.URL, Key: "123"}
	for _, tt := range tests {
		err := tt.MakeRequest(cli)
		if err == nil {
			t.Errorf("%v did not fail", tt.Name)
			continue
		}
		if err != bombbomb.ErrNoSubscription {
			t.Errorf("%v returned invalid error: %v", tt.Name, err)
		}
	}
}

//
// Helpers
//

func readValues(t *testing.T, r *http.Request) url.Values {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read request body: %s", err)
	}
	r.Body.Close()
	values, err := url.ParseQuery(string(body))
	if err != nil {
		t.Fatalf("Failed to parse request POST body: %s", err)
	}
	return values
}

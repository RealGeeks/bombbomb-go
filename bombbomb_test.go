package bombbomb_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealGeeks/bombbomb-go"

	. "github.com/igorsobreira/testing"
)

//
// AddContact
//

func TestAddContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, Stubs["AddContact"])
	}))
	defer ts.Close()

	cli := &bombbomb.Client{
		URL: ts.URL,
		Key: "123",
	}
	contact, err := cli.AddContact(bombbomb.Contact{
		FirstName:   "Jack",
		LastName:    "Johnson",
		Email:       "jj@gmail.com",
		PhoneNumber: "808-123-4321",
	})

	Ok(t, err)
	Equals(t, "106e0e29-e9cf-b812-b895-dcdc059cf9ec", contact.ID)
	Equals(t, "Jack", contact.FirstName)
	Equals(t, "Johnson", contact.LastName)
	Equals(t, "jj@gmail.com", contact.Email)
	Equals(t, "808-123-4321", contact.PhoneNumber)
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
		//	URL: "https://app.bombbomb.com/app/api/api.php",
		//	Key: "1bd5b0c2-9cf4-9798-145d-86cf7ff75254",
	}
	info, err := cli.CreateList(bombbomb.List{
		Name: "Partners",
	})

	Ok(t, err)
	Equals(t, "4184993a-b98e-e9e4-19b6-da1019d9cd3d", info.ID)
	Equals(t, "Partners", info.Name)
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

// Package bombbomb is client to the BombBomb API
//
// http://bombbomb.com/api/
package bombbomb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var ErrNoSubscription = errors.New("This account does not have an active subscription. Please contact support")

// Client to BombBomb API as documented in http://bombbomb.com/api
//
// For every method in BombBomb API we have a corresponding method
// with same name
//
// Note that methods like AddContact receive an instance of Contact
// to be created and return another instance of Contact created. Those
// instances are not the same. Some fields on those types of structs
// are not sent when creating, only set when returning. To see which
// fields are sent read the Values() method.
//
// All methods could return error if:
//
//   - HTTP status code is not 200
//   - 'status' field found in the response is 'failure'
//   - fail to perform the network request
//
type Client struct {
	URL string
	Key string
	Src string // optional. identifies client using the api
}

// response to a request to BombBomb API. Format is always the same.
//
// 'status' will be "failure" if something goes wrong.
// 'methodName' will be the name of '?method' sent
// 'info' will vary based on 'methodName', we have structs for all
// possible values of 'info' below (like Contact or List)
type response struct {
	Status     string          `json:"status"`
	MethodName string          `json:"methodName"`
	Info       json.RawMessage `json:"info"`
}

type Contact struct {
	ID          string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	PhoneNumber string `json:"phone_number"`

	// BombBomb allows one to send a list of list ids this contact
	// should be added to -- 'listlist' field
	//
	// If you don't know the ids, use ListNames and we'll ensure
	// a list with that name exists. If you know the IDs, use ListIDs,
	// it's much faster
	ListsNames []string
	ListsIDs   []string
}

func (c Contact) Values() url.Values {
	return url.Values{
		"eml":          {c.Email},
		"firstname":    {c.FirstName},
		"lastname":     {c.LastName},
		"phone_number": {c.PhoneNumber},
	}
}

type List struct {
	ID           string
	Name         string
	ContactCount string
}

func (l List) Values() url.Values {
	return url.Values{"name": {l.Name}}
}

func (c *Client) AddContact(contact Contact) (newContact Contact, err error) {
	values := contact.Values()
	// when adding contact to lists by name ensure those lists exist and fetch their ids
	if len(contact.ListsNames) > 0 {
		for _, name := range contact.ListsNames {
			list, err := c.EnsureList(List{Name: name})
			if err != nil {
				return Contact{}, err
			}
			contact.ListsIDs = append(contact.ListsIDs, list.ID)
		}
	}
	// when adding contact to lists by id (or if ids fetched above) add those the 'listlist'
	if len(contact.ListsIDs) > 0 {
		values.Set("listlist", strings.Join(contact.ListsIDs, ";"))
	}
	err = c.httpPOST("AddContact", values, &newContact)
	return newContact, err
}

func (c *Client) CreateList(list List) (newList List, err error) {
	err = c.httpPOST("CreateList", list.Values(), &newList)
	return newList, err
}

// EnsureList returns an existing list or create one if a list with that
// name doesn't exist yet
//
// CreateList will create duplicate lists if called more than once with
// same list name. Use this method to avoid duplicates.
func (c *Client) EnsureList(list List) (newList List, err error) {
	lists, err := c.GetLists()
	if err != nil {
		return List{}, err
	}
	for _, l := range lists {
		if l.Name == list.Name {
			return l, nil
		}
	}
	return c.CreateList(list)
}

func (c *Client) GetLists() (lists []List, err error) {
	err = c.httpGET("GetLists", &lists)
	return lists, err
}

func (c *Client) httpPOST(method string, values url.Values, instance interface{}) error {
	values.Set("api_key", c.Key)
	if c.Src != "" {
		values.Set("xsrc", c.Src)
	}
	resp, err := http.PostForm(c.URL+"?method="+method, values)
	return c.handleResponse(method, resp, err, instance)
}

func (c *Client) httpGET(method string, instance interface{}) error {
	uri := c.URL + "?method=" + method + "&api_key=" + c.Key
	if c.Src != "" {
		uri += "&xsrc=" + c.Src
	}
	resp, err := http.Get(uri)
	return c.handleResponse(method, resp, err, instance)
}

func (c *Client) handleResponse(method string, resp *http.Response, err error, instance interface{}) error {
	if err != nil {
		return fmt.Errorf("%s failed (%s)", method, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s failed to read body (%s)", method, err)
	}
	resp.Body.Close()
	if isNoSubscription(resp.StatusCode, body) {
		return ErrNoSubscription
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s returned status %d with body '%s'", method, resp.StatusCode, string(body))
	}
	var bbresp response
	if err := json.Unmarshal(body, &bbresp); err != nil {
		return fmt.Errorf("%s returned invalid json '%s' (%s)", method, string(body), err)
	}
	if bbresp.Status != "success" {
		return fmt.Errorf("%s returned invalid status '%s'", method, body)
	}
	if err := json.Unmarshal(bbresp.Info, &instance); err != nil {
		return fmt.Errorf("%s returned invalid 'info' json '%s' (%s)", method, string(body), err)
	}
	return nil
}

func isNoSubscription(status int, body []byte) bool {
	return status == 403 && strings.Contains(string(body), "does not have an active subscription")
}

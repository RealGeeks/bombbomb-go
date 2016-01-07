// Package bombbomb is client to the BombBomb API
//
// http://bombbomb.com/api/
package bombbomb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client to BombBomb API as documented in http://bombbomb.com/api
type Client struct {
	URL string
	Key string
}

type response struct {
	Status     string          `json:"status"`
	MethodName string          `json:"methodName"`
	Info       json.RawMessage `json:"info"`
}

// Contact is the data used to create a new contact from AddContact
type Contact struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

// ContactInfo is the 'info' object returned by the AddContact request
type ContactInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// AddContact sends a request to the API method 'AddContact'
//
// Return error if status code is not 200, or if 'status' found in the
// response is 'failure', of if fail to perform the network request
func (c *Client) AddContact(contact Contact) (ContactInfo, error) {
	method := "AddContact"
	resp, err := http.PostForm(c.URL+"?amethod="+method, url.Values{
		"api_key":      {c.Key},
		"eml":          {contact.Email},
		"firstname":    {contact.FirstName},
		"lastname":     {contact.LastName},
		"phone_number": {contact.PhoneNumber},
	})
	if err != nil {
		return ContactInfo{}, fmt.Errorf("%s failed (%s)", method, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ContactInfo{}, fmt.Errorf("%s failed to read body (%s)", method, err)
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		return ContactInfo{}, fmt.Errorf("%s returned status %d with body '%s'", method, resp.StatusCode, string(body))
	}
	var bbresp response
	if err := json.Unmarshal(body, &bbresp); err != nil {
		return ContactInfo{}, fmt.Errorf("%s returned invalid json '%s' (%s)", method, string(body), err)
	}
	if bbresp.Status != "success" {
		return ContactInfo{}, fmt.Errorf("%s returned invalid status '%s'", body)
	}
	var info ContactInfo
	if err := json.Unmarshal(bbresp.Info, &info); err != nil {
		return ContactInfo{}, fmt.Errorf("%s returned invalid 'info' json '%s' (%s)", method, string(body), err)
	}
	return info, nil
}

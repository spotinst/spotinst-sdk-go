package spotinst

import (
	"io/ioutil"
	"testing"
)

func Test_NewRequest(t *testing.T) {
	c, err := NewClient(nil)
	if err != nil {
		panic(err)
	}

	inURL, outURL := "foo", apiURL+"/foo"
	inBody, outBody := &map[string]string{"name": "l"}, `{"name":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

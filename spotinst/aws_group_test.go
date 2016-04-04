package spotinst

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_AwsGroupGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/aws/ec2/group/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"response": {"items": [{"id": "foo", "name": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &AwsGroup{ID: String("foo"), Name: String("bar")}
	group, _, err := client.AwsGroup.Get("foo")
	if err != nil {
		t.Errorf("AwsGroup.Get returned error: %v", err)
	}

	if len(group) > 0 && !reflect.DeepEqual(group[0], expected) {
		t.Errorf("AwsGroup.Get returned %+v, expected %+v", group[0], expected)
	}
}

func Test_AwsGroupCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/aws/ec2/group", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"response": {"items": [{"id": "foo", "name": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &AwsGroup{ID: String("foo"), Name: String("bar")}
	group, _, err := client.AwsGroup.Create(expected)
	if err != nil {
		t.Errorf("AwsGroup.Create returned error: %v", err)
	}

	if len(group) > 0 && !reflect.DeepEqual(group[0], expected) {
		t.Errorf("AwsGroup.Create returned %+v, expected %+v", group[0], expected)
	}
}

func Test_AwsGroupUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/aws/ec2/group/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		response := `{"response": {"items": [{"id": "foo", "name": "baz"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &AwsGroup{ID: String("foo"), Name: String("baz")}
	group, _, err := client.AwsGroup.Update(expected)
	if err != nil {
		t.Errorf("AwsGroup.Update returned error: %v", err)
	}

	expected.ID = String("foo")
	if len(group) > 0 && !reflect.DeepEqual(group[0], expected) {
		t.Errorf("AwsGroup.Update returned %+v, expected %+v", group[0], expected)
	}
}

func Test_AwsGroupDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/aws/ec2/group/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	resp, err := client.AwsGroup.Delete(&AwsGroup{ID: String("foo")})
	if err != nil {
		t.Errorf("AwsGroup.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("AwsGroup.Delete returned %+v, expected %+v", resp.StatusCode, http.StatusOK)
	}
}
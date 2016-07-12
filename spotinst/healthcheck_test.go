package spotinst

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_HealthCheckGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/healthCheck/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &HealthCheck{ID: String("foo"), ResourceID: String("bar")}
	HealthCheck, _, err := client.HealthCheck.Get("foo")
	if err != nil {
		t.Errorf("HealthCheck.Get returned error: %v", err)
	}

	if len(HealthCheck) > 0 && !reflect.DeepEqual(HealthCheck[0], expected) {
		t.Errorf("HealthCheck.Get returned %+v, expected %+v", HealthCheck[0], expected)
	}
}

func Test_HealthCheckCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &HealthCheck{ID: String("foo"), ResourceID: String("bar")}
	HealthCheck, _, err := client.HealthCheck.Create(expected)
	if err != nil {
		t.Errorf("HealthCheck.Create returned error: %v", err)
	}

	if len(HealthCheck) > 0 && !reflect.DeepEqual(HealthCheck[0], expected) {
		t.Errorf("HealthCheck.Create returned %+v, expected %+v", HealthCheck[0], expected)
	}
}

func Test_HealthCheckUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/healthCheck/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "baz"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &HealthCheck{ID: String("foo"), ResourceID: String("baz")}
	HealthCheck, _, err := client.HealthCheck.Update(expected)
	if err != nil {
		t.Errorf("HealthCheck.Update returned error: %v", err)
	}

	expected.ID = String("foo")
	if len(HealthCheck) > 0 && !reflect.DeepEqual(HealthCheck[0], expected) {
		t.Errorf("HealthCheck.Update returned %+v, expected %+v", HealthCheck[0], expected)
	}
}

func Test_HealthCheckDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/healthCheck/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	resp, err := client.HealthCheck.Delete(&HealthCheck{ID: String("foo")})
	if err != nil {
		t.Errorf("HealthCheck.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("HealthCheck.Delete returned %+v, expected %+v", resp.StatusCode, http.StatusOK)
	}
}

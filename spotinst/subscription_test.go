package spotinst

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_SubscriptionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/events/subscription/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &Subscription{ID: String("foo"), ResourceID: String("bar")}
	subscription, _, err := client.Subscription.Get("foo")
	if err != nil {
		t.Errorf("Subscription.Get returned error: %v", err)
	}

	if len(subscription) > 0 && !reflect.DeepEqual(subscription[0], expected) {
		t.Errorf("Subscription.Get returned %+v, expected %+v", subscription[0], expected)
	}
}

func Test_SubscriptionCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/events/subscription", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "bar"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &Subscription{ID: String("foo"), ResourceID: String("bar")}
	subscription, _, err := client.Subscription.Create(expected)
	if err != nil {
		t.Errorf("Subscription.Create returned error: %v", err)
	}

	if len(subscription) > 0 && !reflect.DeepEqual(subscription[0], expected) {
		t.Errorf("Subscription.Create returned %+v, expected %+v", subscription[0], expected)
	}
}

func Test_SubscriptionUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/events/subscription/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		response := `{"response": {"items": [{"id": "foo", "resourceId": "baz"}]}}`
		fmt.Fprint(w, response)
	})

	expected := &Subscription{ID: String("foo"), ResourceID: String("baz")}
	subscription, _, err := client.Subscription.Update(expected)
	if err != nil {
		t.Errorf("Subscription.Update returned error: %v", err)
	}

	expected.ID = String("foo")
	if len(subscription) > 0 && !reflect.DeepEqual(subscription[0], expected) {
		t.Errorf("Subscription.Update returned %+v, expected %+v", subscription[0], expected)
	}
}

func Test_SubscriptionDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/events/subscription/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	resp, err := client.Subscription.Delete(&Subscription{ID: String("foo")})
	if err != nil {
		t.Errorf("Subscription.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Subscription.Delete returned %+v, expected %+v", resp.StatusCode, http.StatusOK)
	}
}

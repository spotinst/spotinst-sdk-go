package azure_test

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestListResourceSuggestions(t *testing.T) {

	// mock server
	oceanId := "0-dfgdffs"
	expectedUrl := fmt.Sprintf("/ocean/azure/k8s/cluster/%s/rightSizing/suggestion", oceanId)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// /ocean/azure/k8s/cluster/dfgdf/rightSizing/suggestion?accountId=aweftgdsg
		assert.Contains(t, r.URL.Path, expectedUrl)
		w.WriteHeader(200)
		w.Write([]byte(`
				{
				  "request": {
					"id": "e593ff58-067d-4340-92f9-8b1c0bad70d7",
					"url": "/ocean/aws/k8s/cluster/o-abcd1234/rightSizing/suggestion",
					"method": "POST",
					"timestamp": "2018-06-20T11:35:01.745Z"
				  },
				  "response": {
					"status": {
					  "code": 200,
					  "message": "OK"
					},
					"items": [
					  {
						"resourceName": "redis-controller",
						"resourceType": "deployment",
						"namespace": "kube-system",
						"suggestedCPU": 4,
						"suggestedMemory": 19,
						"requestedCPU": 50,
						"requestedMemory": 50,
						"containers": [
						  {
							"name": "dnsmasq",
							"requestedCPU": 10,
							"suggestedCPU": 2,
							"requestedMemory": 40,
							"suggestedMemory": 15
						  },
						  {}
						]
					  }
					],
					"count": 2,
					"kind": "spotinst:ocean:aws:k8s:cluster:rightSizing:suggestion"
				  }
				}
			`))
	}))
	defer s.Close()

	// Arrange
	os.Setenv(credentials.EnvCredentialsVarAccount, "any-account")
	os.Setenv(credentials.EnvCredentialsVarToken, "any-token")
	defer func() {
		os.Unsetenv(credentials.EnvCredentialsVarToken)
		os.Unsetenv(credentials.EnvCredentialsVarAccount)
	}()

	mockUrl := &url.URL{
		Host:   strings.Replace(s.URL, "http://", "", -1),
		Scheme: "http",
	}

	config := &spotinst.Config{
		BaseURL: mockUrl,
	}

	sess := session.New(config)
	svc := ocean.New(sess)
	provider := svc.CloudProviderAzure()

	// Act
	suggestions, err := provider.ListResourceSuggestions(context.TODO(), &azure.ListResourceSuggestionsInput{
		OceanID: spotinst.String(oceanId),
		Filter: &azure.FilterResourceSuggestions{
			Namespaces: []string{},
		},
	})
	if err != nil {
		t.Error(err)
	}

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, len(suggestions.Suggestions), 1)
	assert.Equal(t, len(suggestions.Suggestions[0].Containers), 2)
}

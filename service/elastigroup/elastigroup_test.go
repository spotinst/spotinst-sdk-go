package elastigroup

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

const (
	BaseURL = "https://api.spotinst.io"
	TaskURL = "/azure/compute/task"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func HttpTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestReadTask(t *testing.T) {
	os.Setenv("SPOTINST_TOKEN", "FAKE")
	httpClient := HttpTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), BaseURL+TaskURL+"/sat-e7db7386")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(listTaskResp)),
		}
	})

	sess := session.New()
	ctx := context.Background()
	sess.Config.HTTPClient = httpClient
	svc := New(sess)

	input := &azure.ReadTaskInput{
		spotinst.String("sat-e7db7386"),
	}

	out, err := svc.CloudProviderAzure().ReadTask(ctx, input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
	}

	assert.Equal(t, *out.Task.TaskID, *input.TaskID)
}

func TestCreateTask(t *testing.T) {
	httpClient := HttpTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), BaseURL+TaskURL)
		//assert.Equal(t, req.Body.Read("name"), "test")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(createTaskResp)),
		}
	})

	sess := session.New()
	ctx := context.Background()
	sess.Config.HTTPClient = httpClient
	svc := New(sess)

	input := &azure.CreateTaskInput{
		Name:        spotinst.String("create-task"),
		Description: spotinst.String("create example description"),
		State:       spotinst.String("DISABLED"),
		Policies: []*azure.TaskPolicy{
			{
				Cron:   spotinst.String("00 20 * * FRI"),
				Action: spotinst.String("START"),
			},
		},
		Instances: []*azure.TaskInstance{
			{
				VmName:            spotinst.String("CreateVm"),
				ResourceGroupName: spotinst.String("CreateGroup"),
			},
		},
	}

	out, err := svc.CloudProviderAzure().CreateTask(ctx, input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}

	assert.Equal(t, *out.Task.Name, *input.Name)
	assert.Equal(t, *out.Task.Description, *input.Description)
	assert.Equal(t, *out.Task.State, *input.State)
	assert.Equal(t, *out.Task.Policies[0].Action, *input.Policies[0].Action)
	assert.Equal(t, *out.Task.Policies[0].Cron, *input.Policies[0].Cron)
	assert.Equal(t, *out.Task.Instances[0].VmName, *input.Instances[0].VmName)
	assert.Equal(t, *out.Task.Instances[0].ResourceGroupName, *input.Instances[0].ResourceGroupName)
}

const createTaskResp = `
{
	"response": {
		"status": {
			"code": 200,
			"message": "OK"
		},
		"kind": "spotinst:azure:compute:task",
		"items": [{
			"createdAt": "2018-10-16T18:33:22.000Z",
			"updatedAt": "2018-10-16T19:55:02.000Z",
			"deletedAt": null,
			"id": "sat-e7db7386",
			"name": "create-task",
			"description": "create example description",
			"state": "DISABLED",
			"policies": [{
				"cron": "00 20 * * FRI",
				"action": "START"
			}],
			"instances": [{
				"vmName": "CreateVm",
				"resourceGroupName": "CreateGroup"
			}]
		}],
		"count": 1
	}
}
`

const updateTaskResp = `
{
	"response": {
		"status": {
			"code": 200,
			"message": "OK"
		},
		"kind": "spotinst:azure:compute:task",
		"items": [{
			"createdAt": "2018-10-16T18:33:22.000Z",
			"updatedAt": "2018-10-16T19:55:02.000Z",
			"deletedAt": null,
			"id": "sat-e7db7386",
			"name": "update-task",
			"description": "update example description",
			"state": "ENABLED",
			"policies": [{
				"cron": "01 30 * * TUES",
				"action": "STOP"
			}]
		}],
		"count": 1
	}
}
`

const listTaskResp = `
{
	"response": {
		"status": {
			"code": 200,
			"message": "OK"
		},
		"kind": "spotinst:azure:compute:task",
		"items": [{
			"createdAt": "2018-10-16T18:33:22.000Z",
			"updatedAt": "2018-10-16T19:55:02.000Z",
			"deletedAt": null,
			"id": "sat-e7db7386",
			"name": "example-task",
			"description": "example description",
			"state": "DISABLED",
			"policies": [{
				"cron": "00 20 * * FRI",
				"action": "START"
			}]
		}],
		"count": 1
	}
}
`

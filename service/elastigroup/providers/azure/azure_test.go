package azure

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/stretchr/testify/assert"
)

const readTaskResp = `
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

func TestReadTask(t *testing.T) {
	os.Setenv("SPOTINST_TOKEN", "FAKE")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(readTaskResp))
	}))
	defer ts.Close()

	conf := spotinst.DefaultConfig().WithBaseURL(ts.URL)
	sess := session.New(conf)
	svc := New(sess)

	input := &ReadTaskInput{
		TaskID: spotinst.String("sat-e7db7386"),
	}

	output, err := svc.ReadTask(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *output.Task.ID, *input.TaskID)
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

func TestCreateTask(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(createTaskResp))
	}))
	defer ts.Close()

	conf := spotinst.DefaultConfig().WithBaseURL(ts.URL)
	sess := session.New(conf)
	svc := New(sess)

	input := &CreateTaskInput{
		Task: &Task{
			Name:        spotinst.String("create-task"),
			Description: spotinst.String("create example description"),
			State:       spotinst.String("DISABLED"),
			Policies: []*TaskPolicy{
				{
					Cron:   spotinst.String("00 20 * * FRI"),
					Action: spotinst.String("START"),
				},
			},
			Instances: []*TaskInstance{
				{
					VMName:            spotinst.String("CreateVm"),
					ResourceGroupName: spotinst.String("CreateGroup"),
				},
			},
		},
	}

	output, err := svc.CreateTask(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *output.Task.Name, *input.Task.Name)
	assert.Equal(t, *output.Task.Description, *input.Task.Description)
	assert.Equal(t, *output.Task.State, *input.Task.State)
	assert.Equal(t, *output.Task.Policies[0].Action, *input.Task.Policies[0].Action)
	assert.Equal(t, *output.Task.Policies[0].Cron, *input.Task.Policies[0].Cron)
	assert.Equal(t, *output.Task.Instances[0].VMName, *input.Task.Instances[0].VMName)
	assert.Equal(t, *output.Task.Instances[0].ResourceGroupName, *input.Task.Instances[0].ResourceGroupName)
}

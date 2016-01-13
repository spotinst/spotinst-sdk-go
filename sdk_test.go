/*
  @author    Liran Polak
  @copyright Copyright (c) 2016, Spotinst
  @license   GPL-3.0
*/

package spotinstsdk

import (
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
	"regexp"
)

var (
	testClient       *Client
	testUsername 	 = os.Getenv("SPOTINST_USERNAME")
	testPassword 	 = os.Getenv("SPOTINST_PASSWORD")
	testClientId 	 = os.Getenv("SPOTINST_CLIENT_ID")
	testClientSecret = os.Getenv("SPOTINST_CLIENT_SECRET")
	testGroupId 	 string
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	if testUsername == "" {
		log.Println("Please configure SPOTINST_USERNAME")
		os.Exit(1)
	}
	if testPassword == "" {
		log.Println("Please configure SPOTINST_PASSWORD")
		os.Exit(1)
	}
	if testClientId == "" {
		log.Println("Please configure SPOTINST_CLIENT_ID")
		os.Exit(1)
	}
	if testClientSecret == "" {
		log.Println("Please configure SPOTINST_CLIENT_SECRET")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func Test_CreateClient(t *testing.T) {
	t.Log("Creating a new client")
	var err error
	testClient, err = NewClient(testUsername, testPassword, testClientId, testClientSecret)
	if testClient == nil || err != nil {
		t.Fatal(err)
	}
	t.Log("Client created successfully")
}

// Get Test
func Test_GetGroups(t *testing.T) {
	t.Log("Getting all groups")
	res, err := testClient.Group.Get()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

// Create Test
func Test_CreateGroup(t *testing.T) {
	t.Log("Creating a new group")
	g := &Group{
		Name: "spotinst-sdk-go-test",
		Description: "Created by Spotinst SDK for the Go programming language",
		Strategy: &GroupStrategy{
			Risk: 100,
			AvailabilityVsCost: "balanced",
		},
		Compute: &GroupCompute{
			Product: "Linux/UNIX",
			LaunchSpecification: &GroupComputeLaunchSpecification{
				SecurityGroupIds: []string{"default"},
				ImageId: "ami-f0091d91",
				KeyPair: "float_oregon",
			},
			AvailabilityZones: []*GroupComputeAvailabilityZone{
				&GroupComputeAvailabilityZone{
					Name: "us-west-2b",
				},
			},
			InstanceTypes: &GroupComputeInstanceType{
				OnDemand: "c3.large",
				Spot: []string{"c3.large"},
			},
		},
		Capacity: &GroupCapacity{
			Minimum: 0,
			Maximum: 1,
			Target: 0,
		},
	}
	res, err := testClient.Group.Create(*g)
	if err != nil {
		t.Fatal(err)
	}

	testGroupId = res[0].Id
	t.Logf("%+v", res)
}

// Another Get Test with a specific ID
func Test_GetGroupById(t *testing.T) {
	if testGroupId != "" {
		t.Log("Getting a group by ID")
		res, err := testClient.Group.Get(testGroupId)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	} else {
		t.SkipNow()
	}
}

// Update Test
func Test_UpdateGroup(t *testing.T) {
	if testGroupId != "" {
		t.Log("Updating group")
		g := &Group{Id: testGroupId, Name: "spotinst-sdk-go-test-updated"}
		res, err := testClient.Group.Update(*g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	} else {
		t.SkipNow()
	}
}

// Delete Test
func Test_DeleteGroup(t *testing.T) {
	if testGroupId != "" {
		t.Log("Deleting group")
		g := &Group{Id: testGroupId}
		res, err := testClient.Group.Delete(*g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	} else {
		t.SkipNow()
	}
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Expected %#v, got %#v", expected, actual)
	}
}

func assertMatch(t *testing.T, actual, pattern string) {
	re := regexp.MustCompile(pattern)
	if !re.MatchString(actual) {
		t.Errorf("Expected to match %#v, got %#v", pattern, actual)
	}
}
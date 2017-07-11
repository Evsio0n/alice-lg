package gobgpwatcher

import (
	"log"
	"testing"

	"encoding/json"
	"io/ioutil"
)

/*
 Read test data
*/
func readTestData(filename string) ClientResponse {
	payload, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result := make(ClientResponse)
	_ = json.Unmarshal(payload, &result)

	return result
}

func Test_ParseApiStatus(t *testing.T) {
	gobgp := readTestData("./_test/gobgp_status.json")

	apiStatus, err := parseApiStatus(gobgp, Config{})
	if err != nil {
		t.Error(err)
	}

	if apiStatus.Version != "0.1.0" {
		t.Error("Expected version to be 0.1.0")
	}
}

func Test_ParseServerStatus(t *testing.T) {
	gobgp := readTestData("./_test/gobgp_status.json")

	status, err := parseServerStatus(gobgp, Config{})
	if err != nil {
		t.Error(err)
	}

	if status.RouterId != "195.43.89.133" {
		t.Error("Router id does not match")
	}
}
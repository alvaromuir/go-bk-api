package bk

import (
	"encoding/json"
	"testing"

	"github.com/joho/godotenv"
)

func TestAudiences(t *testing.T) {
	_ = godotenv.Load("../.env")

	params := "?status=shared"
	resp, rslt, err := Audiences(params)
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}

	body, err := json.Marshal(rslt)
	if err != nil {
		t.Errorf("An error has occured getting a response.")
	}

	if len(string(body)) < 1 {
		t.Errorf("An error has occured reading a response.")
	}
}

func TestCategories(t *testing.T) {
	_ = godotenv.Load("../.env")

	queryArgValues := "queryArgs=abc123"
	resp, rslt, err := Taxonomy(queryArgValues)
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}

	body, err := json.Marshal(rslt)
	if err != nil {
		t.Errorf("An error has occured getting a response.")
	}
	if len(string(body)) < 1 {
		t.Errorf("An error has occured reading a response.")
	}
}

func TestSites(t *testing.T) {
	_ = godotenv.Load("../.env")

	resp, rslt, err := Sites()
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}

	body, err := json.Marshal(rslt)
	if err != nil {
		t.Errorf("An error has occured getting a response.")
	}
	if len(string(body)) < 1 {
		t.Errorf("An error has occured reading a response.")
	}
}

func TestCallEndpoint(t *testing.T) {
	_ = godotenv.Load("../.env")

}

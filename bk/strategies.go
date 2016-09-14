package bk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// PingService returns a http.Response of a empty POST request
func PingService() (*http.Response, error) {
	var URL = "https://services.bluekai.com/Services/WS/Ping"
	var method = "POST"
	queryArgValues := ""
	postData := []byte("")

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, postData)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	return resp, err
}

// Audiences returns custom audiences by params criteria (e.g. ?status=shared)
// see https://devportal.bluekai.com/audienceapi#Methods-GET-List-your-audiences
func Audiences(queryArgs string) (*http.Response, *AudiencesListResult, error) {
	var URL = "https://services.bluekai.com/Services/WS/audiences"
	var method = "GET"
	queryArgValues := queryArgs
	postData := []byte("")

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, postData)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}
	var rslt AudiencesListResult
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, &rslt, nil
}

// Audience returns custom audiences by params criteria (e.g. ?status=shared)
// see https://devportal.bluekai.com/audienceapi#Methods-GET-List-your-audiences
func Audience(audienceID string) (*http.Response, *AudienceDetailResult, error) {
	var URL = "https://services.bluekai.com/Services/WS/audiences/" + audienceID
	var method = "GET"
	queryArgValues := ""
	postData := []byte("")

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, nil)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	var rslt AudienceDetailResult
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, &rslt, nil
}

// Segments returns DMP reach details of selected segment(s)
// see https://devportal.bluekai.com/segment_reach_api#Methods-POST-Get-Audience-Reach
// note: docs are wrong, add an additional {"AND":[]} to queries
func Segments(segmentsQuery string, deviceTypeID string) (*http.Response, *SegmentReachResult, error) {
	var URL = "https://services.bluekai.com/Services/WS/SegmentInventory"
	var method = "POST"
	queryArgValues := "deviceTypeID=" + deviceTypeID
	postData := []byte(segmentsQuery)

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, postData)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	var rslt SegmentReachResult
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, &rslt, nil
}

// Taxonomy returns DMP taxonomy detail
func Taxonomy(queryArgs string) (*http.Response, *TaxonomyResult, error) {
	var URL = "https://services.bluekai.com/Services/WS/Taxonomy"
	var method = "GET"
	queryArgValues := queryArgs
	postData := []byte("")

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, postData)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	var rslt TaxonomyResult
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, &rslt, nil
}

// Sites returns DMP sites in JSON format
func Sites() (*http.Response, *SiteResult, error) {
	var URL = "https://services.bluekai.com/Services/WS/sites"
	var method = "GET"
	queryArgValues := ""
	postData := []byte("")

	var endPoint = SignRequest(method, URL, queryArgValues, postData)

	resp, err := CallEndpoint(method, endPoint, postData)
	CheckError(err, "CallEndpoint ERROR: %s")
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	var rslt SiteResult
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, &rslt, nil
}

// CallEndpoint returns a raw BK API response
func CallEndpoint(method string, endPoint string, data []byte) (*http.Response, error) {
	resp, err := SendJSONRequest(method, endPoint, data, nil)
	return resp, err
}

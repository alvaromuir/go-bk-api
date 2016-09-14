package bk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"strings"

	"github.com/joho/godotenv"

	"log"
	"net/http"
	"net/url"
	"os"
)

var env = godotenv.Load(".env")

// SignRequest returns a HMAC-SHA256 string for BlueKai JSON request
func SignRequest(method string, URL string, queryArgValues string, data []byte) string {
	// Add your API Keys below
	var bkuid = os.Getenv("BK_KEY")              // BlueKai API KEY
	var bksecretkey = os.Getenv("BK_SECRET")     // BlueKai API Secret
	var bkpartnerid = os.Getenv("BK_PARTNER_ID") // BlueKai partner ID

	if len(bkuid) < 1 {
		log.Fatal("ERROR: please set a BK UUID")
	}

	if len(bksecretkey) < 1 {
		log.Fatal("ERROR: please set a BK UUID")
	}

	if len(bkpartnerid) < 1 {
		log.Fatal("ERROR: please set a BK UUID")
	}

	uriPath := strings.Split(URL, ".com")[1]
	var stringToSign = method + uriPath

	qP := strings.Split(queryArgValues, "&")

	if len(queryArgValues) > 0 {
		for qS := 0; qS < len(qP); qS++ {
			qP2 := strings.Split(qP[qS], "=")
			if len(qP2) > 1 {
				stringToSign += qP2[1]
			}
		}
	}

	if len(string(data)) > 0 {
		stringToSign += string(data)
	}

	h := hmac.New(sha256.New, []byte(bksecretkey))
	h.Write([]byte(stringToSign))
	digest := base64.StdEncoding.EncodeToString(h.Sum(nil))

	bksig := url.QueryEscape(digest)

	var endPoint = URL + "?"
	if queryArgValues != "" {
		endPoint += queryArgValues + "&"
	}

	endPoint += "bkuid=" + bkuid + "&bksig=" + bksig
	return endPoint
}

// SendJSONRequest is a bare-bones JSON request, essential for all other API calls
func SendJSONRequest(method string, URL string, data []byte, addHeaders map[string]string) (*http.Response, error) {
	// println(string(data[:len(data)]))
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("SendJSONRequest ERROR: %s", err)
		return resp, err
	}

	return resp, nil
}

// SendGetRequest returns a http response for a GET request to provided url
func SendGetRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Download retrieves a file from the specified url to a specified path
func Download(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

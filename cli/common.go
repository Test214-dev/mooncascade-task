package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const BaseAppURL = "http://localhost:8080/api"

func getResponseError(response *http.Response) string {
	data, _ := ioutil.ReadAll(response.Body)
	m := models.AppError{}
	if err := json.Unmarshal(data, &m); err != nil {
		return ""
	}

	return m.Error
}

func createAsset(urlSuffix string, request []byte) (string, error) {
	url := fmt.Sprintf("%s/%s", BaseAppURL, urlSuffix)
	response, err := http.Post(url, "application/json", bytes.NewReader(request))
	if err != nil {
		return "", err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("Unable to close response body: %s", err)
		}
	}()

	if response.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	return response.Header.Get("Location"), nil
}

func getAsset(urlSuffix string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", BaseAppURL, urlSuffix)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("Unable to close response body: %s", err)
		}
	}()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response body")
	}

	return body, nil
}

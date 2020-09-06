package mastodon

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

// Get 是封装的Get请求
func Get(url string, token string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("authorization", token)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}

	client := &http.Client{}
	resp, requestErr := client.Do(request) // resp 可能为 nil，不能读取 Body。所以不能先defer
	if requestErr != nil {
		color.Red(requestErr.Error())
		return nil, requestErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red(readErr.Error())
		return nil, err
	}

	return body, nil

}

// Post 是封装的Post请求
func Post(url string, token string, params *StatusParams) ([]byte, error) {
	jsonStr, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	request.Header.Add("Authorization", token)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		color.Red(err.Error())
		return nil, err
	}

	client := &http.Client{}

	resp, requestErr := client.Do(request) // resp 可能为 nil，不能读取 Body。所以不能先defer
	if requestErr != nil {
		color.Red(requestErr.Error())
		return nil, requestErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red(readErr.Error())
		return nil, err
	}

	if resp.StatusCode != 200 {
		// {"error":"Validation failed: Text can't be blank"}
		return nil, errors.New(string(body))
	}

	return body, nil
}

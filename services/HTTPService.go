package services

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
	"errors"
	// "fmt"
	// "io"
	"net/http"
	"upstox-sdk-go/viewmodels"
)

var zeroByte = new([]byte)

// HTTPService ... 
type HTTPService struct {
	HTTPClient *http.Client
}


// PostJSON ... Post function
func (service *HTTPService) PostJSON(url string, params interface{}, headers map[string][]string, data interface{}) (error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	
	var jsonStr = []byte(string(jsonData))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	
	req.Header = headers
	resp, err := service.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read response: %v", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var e viewmodels.HTTPErrorModel
		if err := json.Unmarshal(body, &e); err != nil {
			fmt.Printf("Error parsing JSON response: %v", err)
			return err
		}
		return errors.New(e.Message)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON response: %v | %s", err, body)
		return err
	}

	return nil
}

// func Delete(url string, token string, client http.Client) (err error) {
// 	req, err := http.NewRequest("DELETE", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("X-Auth-Token", token)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	if !(resp.StatusCode == 200 || resp.StatusCode == 202 || resp.StatusCode == 204) {
// 		err = fmt.Errorf("Unexpected server response status code on Delete '%s'", resp.StatusCode)
// 		return
// 	}

// 	return nil
// }

// GetJSON ... Get function
func (service *HTTPService) GetJSON(url string, headers map[string][]string, data interface{}) (error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	
	req.Header = headers
	resp, err := service.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read response: %v", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var e viewmodels.HTTPErrorModel
		if err := json.Unmarshal(body, &e); err != nil {
			fmt.Printf("Error parsing JSON response: %v", err)
			return err
		}
		return errors.New(e.Message)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON response: %v | %s", err, body)
		return err
	}

	return nil
}


// func CallAPI(method, url string, content *[]byte, h ...string) (*http.Response, error) {
// 	if len(h)%2 == 1 { //odd #
// 		return nil, errors.New("syntax err: # header != # of values")
// 	}
// 	//I think the above err check is unnecessary and wastes cpu cycle, since
// 	//len(h) is not determined at run time. If the coder puts in odd # of args,
// 	//the integration testing should catch it.
// 	//But hey, things happen, so I decided to add it anyway, although you can
// 	//comment it out, if you are confident in your test suites.
// 	var req *http.Request
// 	var err error
// 	req, err = http.NewRequest(method, url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i := 0; i < len(h)-1; i = i + 2 {
// 		req.Header.Set(h[i], h[i+1])
// 	}
// 	req.ContentLength = int64(len(*content))
// 	if req.ContentLength > 0 {
// 		req.Body = readCloser{bytes.NewReader(*content)}
// 		//req.Body = *(new(io.ReadCloser)) //these 3 lines do not work but I am
// 		//req.Body.Read(content)           //keeping them here in case I wonder why
// 		//req.Body.Close()                 //I did not implement it this way :)
// 	}
// 	return (new(http.Client)).Do(req)
// }

// type readCloser struct {
// 	io.Reader
// }

// func (readCloser) Close() error {
// 	//cannot put this func inside CallAPI; golang disallow nested func
// 	return nil
// }


// func CheckHTTPResponseStatusCode(resp *http.Response) error {
// 	switch resp.StatusCode {
// 	case 200, 201, 202, 204, 206:
// 		return nil
// 	case 400:
// 		return errors.New("Error: response == 400 bad request")
// 	case 401:
// 		return errors.New("Error: response == 401 unauthorised")
// 	case 403:
// 		return errors.New("Error: response == 403 forbidden")
// 	case 404:
// 		return errors.New("Error: response == 404 not found")
// 	case 405:
// 		return errors.New("Error: response == 405 method not allowed")
// 	case 409:
// 		return errors.New("Error: response == 409 conflict")
// 	case 413:
// 		return errors.New("Error: response == 413 over limit")
// 	case 415:
// 		return errors.New("Error: response == 415 bad media type")
// 	case 422:
// 		return errors.New("Error: response == 422 unprocessable")
// 	case 429:
// 		return errors.New("Error: response == 429 too many request")
// 	case 500:
// 		return errors.New("Error: response == 500 instance fault / server err")
// 	case 501:
// 		return errors.New("Error: response == 501 not implemented")
// 	case 503:
// 		return errors.New("Error: response == 503 service unavailable")
// 	}
// 	return errors.New("Error: unexpected response status code")
// }

// func createJSONGetRequest(url string, token string) (req *http.Request, err error) {
// 	req, err = http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("X-Auth-Token", token)

// 	return req, nil
// }

// func executeRequestCheckStatusDecodeJSONResponse(client http.Client, req *http.Request, val interface{}) (err error) {
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	err = CheckHTTPResponseStatusCode(resp)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.NewDecoder(resp.Body).Decode(&val)
// 	defer resp.Body.Close()

// 	return err
// }
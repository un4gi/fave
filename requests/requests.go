/*
Copyright © 2022 Tony West

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0
	
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package requests

import (
	"io/ioutil"
	"log"
	"net/http"
)

// MakeGetRequest makes a typical GET request and returns the response body in byte format
func MakeGetRequest(url string) (bodyBytes []byte) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		return bodyBytes
	}
	return bodyBytes
}

// PlainGetRequest makes a GET request, but returns the raw HTTP response
func RawGetRequest(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	SetHeaders(req)

	resp, e := http.DefaultClient.Do(req)
	return resp, e

}

// SetHeaders sets the HTTP headers of a GET request
func SetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0")
}

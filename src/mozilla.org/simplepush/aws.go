/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package simplepush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get the public AWS hostname for this machine. Returns the hostname
// or an error if the call failed.
func GetAWSPublicHostname() (string, error) {
	req := &http.Request{Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "169.254.169.254",
			Path:   "/latest/meta-data/public-hostname"}}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("Bad response from AWS hostname call: %d",
			resp.StatusCode)
	}

	var hostBytes []byte
	hostBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(hostBytes), nil
}

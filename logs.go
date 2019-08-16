package kubernetes

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetPodLogs func
func GetPodLogs(baseurl, namespace, pod, token string, line int) (string, error) {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/pods/%s/log?tailLines=%d", baseurl, namespace, pod, line)
	//fmt.Println(url)
	data, err := httpGet(url, token)
	return data, err
}

func httpGet(url, token string) (data string, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

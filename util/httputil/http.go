package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("pay-center-client http请求失败， error: %v", err)
		return err, nil
	}

	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("pay-center-client url[%v] res[%v]", url, string(resBody))
	return nil, resBody
}

func PostRawJson(url string, body string) (error, []byte) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Printf("pay-center-client http请求失败， error: %v", err)
		return err, nil
	}

	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("pay-center-client url[%v] body[%v] res[%v]", url, body, string(resBody))
	return nil, resBody
}

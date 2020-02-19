package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostRawJson(url string, body string) (error, []byte) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Printf("http请求失败， error: %v", err)
		return err, nil
	}

	resBody, _ := ioutil.ReadAll(resp.Body)
	return nil, resBody
}

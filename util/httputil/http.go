package httputil

import (
	"github.com/FengWuTech/pay-center-client/util/logutil"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		logutil.WarnF("url[%v] error[%v]", url, err)
		return err, nil
	}

	resBody, _ := ioutil.ReadAll(resp.Body)
	logutil.InfoF("url[%v] res[%v]", url, string(resBody))
	return nil, resBody
}

func PostRawJson(url string, body string) (error, []byte) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		logutil.WarnF("url[%v] body[%v] error[%v]", url, body, err)
		return err, nil
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	logutil.InfoF("url[%v] body[%v] res[%v] error[%v]", url, body, string(resBody), err)
	return nil, resBody
}

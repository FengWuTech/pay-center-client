package httputil

import (
	"github.com/FengWuTech/pay-center-client/util/logutil"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		logutil.WarnF("url[%v] error[%v]", url, err)
		return err, nil
	}

	var start = time.Now().UnixNano()
	resBody, _ := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	logutil.InfoF("%d url[%v] res[%v]", end-start, url, string(resBody))
	return nil, resBody
}

func PostRawJson(url string, body string) (error, []byte) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		logutil.WarnF("url[%v] body[%v] error[%v]", url, body, err)
		return err, nil
	}

	var start = time.Now().UnixNano()
	resBody, err := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	logutil.InfoF("%d url[%v] body[%v] res[%v] error[%v]", end-start, url, body, string(resBody), err)
	return nil, resBody
}

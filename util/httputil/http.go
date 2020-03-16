package httputil

import (
	"gitlab.gmtech.top/golang/pay-center-client/util/logutil"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (error, []byte) {
	var start = time.Now().UnixNano()
	resp, err := http.Get(url)
	if err != nil {
		logutil.WarnF("url[%v] error[%v]", url, err)
		return err, nil
	}
	resBody, _ := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	spend := (end - start) / int64(time.Millisecond)
	logutil.InfoF("%dms url[%v] res[%v]", spend, url, string(resBody))
	return nil, resBody
}

func PostRawJson(url string, body string) (error, []byte) {
	var start = time.Now().UnixNano()
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		logutil.WarnF("url[%v] body[%v] error[%v]", url, body, err)
		return err, nil
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	spend := (end - start) / int64(time.Millisecond)
	logutil.InfoF("%dms url[%v] body[%v] res[%v] error[%v]", spend, url, body, string(resBody), err)
	return nil, resBody
}

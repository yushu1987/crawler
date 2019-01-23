package lib

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func SendRequest(url string) (string, int, float64) {
	var (
		resp *http.Response
		err  error
		body []byte
	)
	startTS := time.Now()
	resp, err = http.Get(url)
	defer func () {
		if resp != nil{
			resp.Body.Close()
		}
	}()
	if nil != err {
		Log.Errorf("request error:%s, error:%s", url, err.Error())
		return "", 0, 0.0
	}

	endTS := time.Since(startTS)
	Log.Infof("request end:%s, statusCode:%d, time cost:%s", url, resp.StatusCode, endTS.String())

	body, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		Log.Errorf("response error:%s, error:%s, time cost:%s", url, err.Error(), endTS.String())
		return "", 0, endTS.Seconds()
	}

	return string(body), resp.StatusCode, endTS.Seconds()
}


func Download(url, path string) map[string]interface{} {
	body, statusCode, cost := SendRequest(url) //要改成chan 形式，select， timeout
	defer func(body ,path string) {
		f, _ := os.Create(path)
		defer f.Close()
		io.WriteString(f, body)
	}(body, path)

	return map[string]interface{}{
		"cost":   cost,
		"size":   len(body),
		"status": statusCode,
	}
}

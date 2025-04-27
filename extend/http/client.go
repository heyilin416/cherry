package cherryHttp

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	clog "github.com/cherry-game/cherry/logger"
	jsoniter "github.com/json-iterator/go"
)

var (
	postContentType = "application/x-www-form-urlencoded"
	jsonContentType = "application/json"

	DefaultTimeout = 5 * time.Second
)

func GET(httpURL string, values map[string]string, headers ...map[string]string) ([]byte, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	if len(values) > 0 {
		rst := ToUrlValues(values)
		httpURL = AddParams(httpURL, rst)
	}

	req, err := http.NewRequest("GET", httpURL, nil)
	if err != nil {
		return nil, nil, err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Add(k, v)
		}
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, rsp, err
	}

	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			clog.Warnf("HTTP GET [url = %s], error = %s", httpURL, e)
		}
	}(rsp.Body)

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, rsp, err
	}

	return bodyBytes, rsp, nil
}

func POST(httpURL string, values map[string]string, headers ...map[string]string) ([]byte, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	rst := ToUrlValues(values)
	req, err := http.NewRequest("POST", httpURL, strings.NewReader(rst.Encode()))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", postContentType)
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Add(k, v)
		}
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, rsp, err
	}

	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			clog.Warnf("HTTP POST [url = %s], error = %s", httpURL, e)
		}
	}(rsp.Body)

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, rsp, err
	}

	return bodyBytes, rsp, nil
}

func PostJSON(httpURL string, values interface{}, headers ...map[string]string) ([]byte, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	jsonBytes, err := jsoniter.Marshal(values)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", httpURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Add(k, v)
		}
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, rsp, err
	}

	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			clog.Warnf("HTTP PostJSON [url = %s], error = %s", httpURL, e)
		}
	}(rsp.Body)

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, rsp, err
	}

	return bodyBytes, rsp, nil
}

func AddParams(httpURL string, params url.Values) string {
	if len(params) == 0 {
		return httpURL
	}

	if !strings.Contains(httpURL, "?") {
		httpURL += "?"
	}

	if strings.HasSuffix(httpURL, "?") || strings.HasSuffix(httpURL, "&") {
		httpURL += params.Encode()
	} else {
		httpURL += "&" + params.Encode()
	}

	return httpURL
}

func ToUrlValues(values map[string]string) url.Values {
	rst := make(url.Values)
	for k, v := range values {
		rst.Add(k, v)
	}
	return rst
}

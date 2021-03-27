package http

import (
	"io"
	"io/ioutil"
	"net/http"
)
//创建
func NewClient(addr string) *HttpClient {
	var ht HttpClient
	return &ht
}
//get请求
func (h *HttpClient) GetRequest(url string) (string, int, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", resp.StatusCode, err
	}

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return "", resp.StatusCode, err
		} else {
			res := string(buf[:n])
			return res, resp.StatusCode, err
		}
	}
}
//post请求
func (h *HttpClient) PostRequest(url, contentType string, body io.Reader) (string, int, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	defer resp.Body.Close()
	var respBody, readErr = ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", resp.StatusCode, err
	}

	return string(respBody), resp.StatusCode, err
}

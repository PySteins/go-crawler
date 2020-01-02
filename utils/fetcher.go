package utils

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

func Fetch(url string) (string, error) {
	response, err := ghttp.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Close()
	if response.StatusCode != 200 {
		return "", fmt.Errorf("请求出错!!!状态码: %d", response.StatusCode)
	}
	return response.ReadAllString(), nil
}

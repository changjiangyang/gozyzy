package httputil

import (
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

func Get(url string, par map[string]string) (string, error) {
	return Sendget(url, par, nil)
}

func Sendget(url string, par map[string]string, header map[string]string) (string, error) {

	url = url + "?"
	for key, value := range par {
		url = url + key + "=" + value + "&"
	}
	logs.Info(url)
	req := httplib.Get(url)
	if header != nil {
		for key, value := range header {
			req.Header(key, value)
		}
	}
	str, err := req.String()
	if err != nil {
		return "", err
	} else {
		return str, err
	}
}

// SendPostJson 以post 发送json数据
func SendPostJson(url string, json interface{}, header map[string]string) (string, error) {
	req := httplib.Post(url)
	if header != nil {
		for k, v := range header {
			req.Header(k, v)
		}
	}
	req.JSONBody(json)
	req.SetTimeout(30*time.Second, 30*time.Second)
	res, err := req.Debug(true).String()
	if err == nil {
		return res, nil
	} else {
		return "", nil
	}

}

package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 新建 Api客户端
// 返回 json 字符串
func NewApiClient(url string, data string) (json string) {
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("错误", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("错误", err.Error())
	}
	json = string(body)
	return
}

package common

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// json 按 KEY 重新排序
func JsonRemarshal(bytes []byte) ([]byte, error) {
	var ifce any
	err := json.Unmarshal(bytes, &ifce)
	if err != nil {
		return []byte{}, err
	}
	output, err := json.Marshal(ifce)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

// api 对象接口
type IApiObject interface {
	GetUserId() string    // 返回用户id
	GetSort() bool        // 返回 是否排序
	ToJson() string       // 返回 json
	GetShowCNName() bool  // 是否显示接口的中文名称
	GetFunCnName() string // 返回接口中文名称
}

// 代理 Api 请求返回结果
func GetResponseResult[T any](p IApiObject, URL string) (body string, result CommonResponseResult[T], err error) {

	var (
		str string
	)
	fmt.Println("------------------")

	fmt.Println("排序", p.GetSort())

	sortJson, _ := JsonRemarshal([]byte(p.ToJson()))

	data := url.Values{}
	data.Set("UID", p.GetUserId())
	if p.GetSort() {
		str = string(sortJson)
	} else {
		str = p.ToJson()
	}

	data.Set("DATA", str)

	body = NewApiClient(URL, data.Encode())

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	if p.GetShowCNName() {
		fmt.Printf("------%v------\n", p.GetFunCnName())
	}

	fmt.Println("请求", string(str))
	return
}

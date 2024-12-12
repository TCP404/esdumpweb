package testkit

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type JSONRes struct {
	Code    int         `json:"code"`    // 错误码((0:失败, 1:成功, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
	Common  struct {
		Timestamp int64 `json:"timeStamp"` // 时间戳
	} `json:"common,omitempty"`
}

func UnmarshalJSONRes[T any](data []byte) (result T, err error) {
	var jsonRes JSONRes
	var zero = result
	err = json.Unmarshal(data, &jsonRes)
	if err != nil {
		return zero, err
	}
	err = mapstructure.Decode(jsonRes.Data, &result)
	if err != nil {
		return zero, err
	}
	return result, nil
}

package utils

import (
	"vue3-server/model/system"
)

// ResponseSuccess 返回成功数据
func ResponseSuccess(data interface{}) *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: 0,
		Data: data,
	}
}

// ResponseFail 返回失败数据
func ResponseFail(data interface{}) *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: -1,
		Data: data,
	}
}

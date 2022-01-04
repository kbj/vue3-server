package utils

import (
	"vue3-server/common/enum"
	"vue3-server/model/system"
)

// ResponseSuccess 返回成功数据
func ResponseSuccess(data interface{}) *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: enum.StatusSuccess,
		Data: data,
	}
}

// ResponseFail 返回失败数据
func ResponseFail(data interface{}) *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: enum.StatusError,
		Data: data,
	}
}

// ResponseNotAuth 无权限
func ResponseNotAuth() *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: enum.StatusForbidden,
		Data: "您暂时没有访问此资源的权限！",
	}
}

// ResponseNotFound 找不到
func ResponseNotFound() *system.ResponseEntity {
	return &system.ResponseEntity{
		Code: enum.StatusNotFound,
		Data: "您请求的资源不存在！",
	}
}

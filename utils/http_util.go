package utils

import "vue3-server/entity"

// ResponseSuccess 返回成功数据
func ResponseSuccess(data interface{}) *entity.ResponseEntity {
	return &entity.ResponseEntity{
		Code: 0,
		Data: data,
	}
}

// ResponseFail 返回失败数据
func ResponseFail(data interface{}) *entity.ResponseEntity {
	return &entity.ResponseEntity{
		Code: -1,
		Data: data,
	}
}

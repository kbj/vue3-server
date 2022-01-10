package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"vue3-server/model/base"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
)

func initializeValidator() {
	zhCn := zh.New()
	uni = ut.New(zhCn, zhCn)

	trans, _ = uni.GetTranslator("zh")

	validate = validator.New()
	_ = translations.RegisterDefaultTranslations(validate, trans)
}

func ValidateStruct(entity interface{}) *base.ResponseEntity {
	// 只允许传入结构体
	info := reflect.TypeOf(entity)
	if info.Kind() != reflect.Struct {
		return ResponseFail("校验失败！")
	} else if validate == nil {
		initializeValidator()
	}

	// 校验
	err := validate.Struct(entity)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		var errMsg []string
		for _, e := range errs {
			errMsg = append(errMsg, e.Translate(trans))
		}
		return ResponseFail(strings.Join(errMsg, ","))
	}
	return nil
}

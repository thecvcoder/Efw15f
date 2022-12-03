package initialize

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translational "github.com/go-playground/validator/v10/translations/zh"
	"github.com/thecvcoder/cloud-api-go/global"
)

// 全局Validator数据校验变量
var v *validator.Validate

// Trans 全局翻译器
var t ut.Translator

func InitValidator() {
	chinese := zh.New()
	uni := ut.New(chinese, chinese)
	trans, _ := uni.GetTranslator("zh")
	t = trans
	v = validator.New()
	_ = translational.RegisterDefaultTranslations(v, t)
	global.LOG.Infof("初始化validator完成!")
}

package main

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"strconv"
	"sync"
)

const (
	LangCN = "zh-CN"
	LangEn = "en"
	LangJp = "ja"
)

var once sync.Once
var bundle *i18n.Bundle
var localizerMap map[string]*i18n.Localizer

func LoadDefault() {
	once.Do(loadFromConfig)
}
func loadFromConfig() {
	bundle = i18n.NewBundle(language.SimplifiedChinese)
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// 添加支持的语言
	bundle.MustLoadMessageFile("/Users/mumu/GolandProjects/go_learn/learning/err/locales/en.toml")
	bundle.MustLoadMessageFile("/Users/mumu/GolandProjects/go_learn/learning/err/locales/zh-CN.toml")
	langs := []string{LangCN, LangEn, LangJp}
	localizerMap = make(map[string]*i18n.Localizer, len(langs))
	for _, lang := range langs {
		localizerMap[lang] = i18n.NewLocalizer(bundle, lang)
	}
}
func GetErrorMessage(locale string, errorCode int) string {
	LoadDefault()
	localizer, exists := localizerMap[locale]
	if !exists {
		localizer = localizerMap[LangCN]
	}
	message, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: strconv.Itoa(errorCode),
	})
	return message
}

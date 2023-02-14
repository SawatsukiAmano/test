package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {

	//创建 bundle，默认英文
	bundle := i18n.NewBundle(language.English)
	//注册使用 toml 为模板
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//加载中文翻译
	bundle.MustLoadMessageFile("active.zh.toml")
	//加载日文翻译
	bundle.MustLoadMessageFile("active.ja.toml")

	//创建 localizer，分别为：bundle 对象、变量和语言
	name := "Bob"
	catNums := 1 //英文 CLDR，为 1 时为 one，为其他时为 other
	localizer := i18n.NewLocalizer(bundle, name, "ja")

	//创建消息，英文的默认就是本条
	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PersonCats",
			One:   "{{.Name}} has cat",
			Other: "{{.Name}} has too many cats ",
		},
		MessageID:   "PersonCats",
		PluralCount: catNums, //数量为多少
		TemplateData: map[string]string{
			"Name": name,
		},
	})
	fmt.Println(helloPerson)

}

package main

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	fmt.Println("==== err:: ")

	matches, err := filepath.Glob("./locale/*.toml")
	if err != nil {
		fmt.Println("==== err:: ", err)

		// panic(err)
	}
	fmt.Println("==== matches:: ", matches)

	for _, match := range matches {
		fmt.Println("==== match:: ", match)
		bundle.MustLoadMessageFile(match)
	}

	// bundle.MustParseMessageFileBytes(
	// 	[]byte(`
	// 	HelloWorld = "Hello World!"
	// 	`), "en.toml")

	// bundle.MustParseMessageFileBytes(
	// 	[]byte(`
	// 	HelloWorld = "Hola Mundo!"
	// 	`), "es.toml")

	{
		localizer := i18n.NewLocalizer(bundle, "en")
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld"}))
	}
	{
		localizer := i18n.NewLocalizer(bundle, "vi")
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld"}))
	}
}

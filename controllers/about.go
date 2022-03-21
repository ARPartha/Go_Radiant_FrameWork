package controllers

import (
	"net/http"
	"radiant/radiant"
	"radiant/radiant/core"
)

func About(c radiant.Context) error {
	
	var languageMap = map[string]string{}

	if x, found := core.GetLanguage(c.Param("lang")); found {
		for word, translate := range x.(map[string]string) {
			languageMap[word] = translate
		}

	}
	return c.Render(http.StatusOK, "about.gohtml", SetTemplateVars(map[string]interface{}{
		"hello":   languageMap["Hello"],
		"welcome": languageMap["Welcome"]}),
	)

}

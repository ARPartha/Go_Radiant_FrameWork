package controllers

import (
	"net/http"

	"radiant/radiant"
	"radiant/radiant/core"
)

func Location(c radiant.Context) error {
	
	var languageMap = map[string]string{}

	if x, found := core.GetLanguage(c.Param("lang")); found {
		for word, translate := range x.(map[string]string) {
			languageMap[word] = translate
		}
	}
	return c.Render(http.StatusOK, "index.gohtml", SetTemplateVars(map[string]interface{}{
		"country": c.Param("country"),
		"city":    c.Param("city"),
		"state":   c.Param("state"),
		"hello":   languageMap["Hello"],
		"welcome": languageMap["Welcome"]}),
	)
}

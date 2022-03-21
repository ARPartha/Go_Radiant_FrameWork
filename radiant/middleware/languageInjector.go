package middleware

import (
	"net/http"
	"radiant/radiant"
	"radiant/radiant/core"
	"regexp"
	"strings"
)

var LanguageCountryCode string
func LanguageInjector(handlerFunc radiant.HandlerFunc) radiant.HandlerFunc {
	
	var newUrl string
	var redirectCheck = false
	return func(c radiant.Context) error {

		// don't redirect if its any asset requests
		// TODO there has to be a better way to know if its a asset/static content request
		if strings.Contains(c.Request().URL.RequestURI(), "/static/") {
			return handlerFunc(c)
		}
		var err error
		var lang string
		flag := false

		if lang == "" {
			lang = "en-US"

		}
		
		acceptLang := c.Request().Header.Get("Accept-Language")
		urlPath := c.Request().URL.Path

		regLang := regexp.MustCompile("[aA-zZ]*-[aA-zZ]*")
		chosenLang := regLang.FindString(urlPath)

		for _, countryCode := range core.Language {
			if strings.Contains(chosenLang, countryCode) {
				lang = chosenLang
				LanguageCountryCode = lang
				flag = true
				break
			} else {
				lang = "en-US"
				LanguageCountryCode = lang
			}
		}

		if !flag {
			for _, countryCode := range core.Language {
				if strings.Contains(acceptLang, countryCode) {
					lang = countryCode
					LanguageCountryCode = countryCode
					break
				} else {
					lang = "en-US"
				}
			}
		}

		if strings.Contains(urlPath, lang) {
			//blank
		} else if strings.Contains(urlPath, "-") && (!strings.Contains(urlPath, lang)) {
			lang = "en-US"
			resetLang := regLang.Split(urlPath, -1)
			newUrl = "/" + LanguageCountryCode + resetLang[1]
			redirectCheck = true

		} else {
			if urlPath == "/" {
				urlPath = ""
			}
			newUrl = "/" + LanguageCountryCode + urlPath
			redirectCheck = true
		}

		if redirectCheck {
			err = c.Redirect(http.StatusFound, newUrl)
			redirectCheck = false
			if err != nil {
				return c.HTML(http.StatusInternalServerError, err.Error())
			}

		}
		return handlerFunc(c)
	}

}
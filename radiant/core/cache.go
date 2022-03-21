package core

import (
	"bufio"
	"log"
	"os"

	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var GlobalCacheManager = cache.New(5*time.Minute, 10*time.Minute)

func LoadLanguageFileInCache() {
	for _, languageCountryCode := range Language {
		filePath := "radiant/core/config/lang/" + languageCountryCode + ".ini"
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		slice := make(map[string]string)
		for scanner.Scan() {
			slice[strings.Split(scanner.Text(), " = ")[0]] = strings.Split(scanner.Text(), " = ")[1]
			GlobalCacheManager.Set(languageCountryCode, slice, cache.DefaultExpiration)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func GetLanguage(CountryCode string) (interface{}, bool) {
	//Get JsonFile pointer
	if jsonData, found := GlobalCacheManager.Get(CountryCode); found {
		jsonFile := jsonData
		return jsonFile, true

	} else {
		return nil, false
	}
}

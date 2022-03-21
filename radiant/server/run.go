package server

import (
	"radiant/radiant"
	"radiant/radiant/core"
	"radiant/radiant/middleware"
	"radiant/router"
)




func Run(e *radiant.Radiant){
	
	core.LoadLanguageFileInCache()

	e.Debug = core.Configure.Error["Debug"] == "true"

	if core.Configure.Error["SentryCall"] == "true" {
		applySentry := radiant.SentryInit()
		if applySentry {
			e.Use(radiant.SentryNew(radiant.SentryOptions{Repanic: true}))
		}
	}

	if core.Configure.Server["LanguageRedirect"] == "true" && core.Configure.Server["Swagger"]=="false"{
		e.Use(middleware.LanguageInjector)
	}

	e.Static("/static", "static")

	e.RenderTemplate()
	
	router.Router(e)

	//Loading Server
	e.Logger.Fatal(e.Start(core.Configure.Server["Host"]))
}
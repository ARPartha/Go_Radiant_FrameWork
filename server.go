package main

import (
	"radiant/radiant"
	"radiant/radiant/db"
	"radiant/radiant/middleware"
	Server "radiant/radiant/server"
)

func main() {
	e := radiant.New()

	e.HTTPErrorHandler = radiant.CustomHTTPErrorHandler
	
	e.Use(middleware.SiteAddress)

	db.DBInit()
	Server.Run(e)

}

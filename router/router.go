package router

import (
	"radiant/controllers"
	"radiant/radiant"
	"radiant/radiant/swagger"
)

func Router(e *radiant.Radiant) {
	
	e.Route("", controllers.Index,"GET")
	e.Route("/about", controllers.About,"GET")
	e.Route("/:country/:state/:city", controllers.Location,"GET")
	
	//swagger
	e.Route("/swagger/*", swagger.WrapHandler,"GET")
	e.Route("/object",controllers.Post,"POST")
	

	//db routing
	 e.Route("/users", controllers.Getuser,"GET")
	 e.Route("/users",controllers.Createuser,"POST")
}

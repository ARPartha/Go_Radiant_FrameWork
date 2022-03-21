package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"radiant/models"
	"radiant/radiant"
	"radiant/radiant/core"
	"radiant/radiant/db"
)



func Getuser(c radiant.Context) error {
	database:= db.DbManager()
	
	users := []models.User_info{}
	fmt.Println("site Name: ",core.Site.Name)
	database.Find(&users, "firstname = ?", "f")

	renderHtml:= core.Site.Name+"/user.gohtml"

	return c.Render(http.StatusOK, renderHtml, nil)
	//  return c.JSON(http.StatusOK,users)
	
}
func Createuser(c radiant.Context) error{
	 database:= db.DbManager()
	 user :=models.User_info{}
	

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err !=nil{

	}
	 database.Create(&user)

	renderHtml:= core.Site.Name+"/user.gohtml"

	// return c.JSON(http.StatusOK,user)
	return c.Render(http.StatusOK,renderHtml , nil)
	
}
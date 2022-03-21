package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"radiant/models"
	"radiant/radiant"
)

func Post(c radiant.Context) error {

		
	var ob models.Data
	
	err := json.NewDecoder(c.Request().Body).Decode(&ob)
	if err!=nil{
		
	}
	fmt.Println(ob)
	 models.AddOne(ob)
	
	return c.JSON(http.StatusOK,ob)

}

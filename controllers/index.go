package controllers

import (
	"net/http"
	"radiant/radiant"
)

func Index(c radiant.Context) error {

	return c.Render(http.StatusOK, "home/index",
		SetTemplateVars(map[string]interface{}{
			"Title":     "Radiant",
			"Framework": "Radiant",
			"Version":   "0.0.1"}),
	)

}

package middleware

import (
	"net/http"
	"radiant/radiant"
	"radiant/radiant/core"
	"strings"
)
var Flag bool

func SiteAddress(handlerFunc radiant.HandlerFunc) radiant.HandlerFunc {
   sites:= core.Configure.Sites["allowedsite"]
  
	return func(c radiant.Context) error {
		
		sitename:= strings.Split(c.Request().Host, ".com")[0]
		core.Site.Name=strings.Split(sitename,".")[1] 
			
		for _,hostname:=range sites{
			if strings.Contains(core.Site.Name,hostname){	
				Flag=true
				break 				
			}else{
				Flag =false
			}
		}
			if Flag{
				sitename:= strings.Split(c.Request().Host, ".com")[0]
				core.Site.Name=strings.Split(sitename,".")[1] 
			}else{
				if core.Configure.Siteallowance["default"]== "true"{
				core.Site.Name=sites[0]
				}else{
					c.Render(http.StatusOK,"error.gohtml",map[string]interface{}{
						"message": "wrong host Name",
					})
				 }
			}
		return handlerFunc(c)
	}
}


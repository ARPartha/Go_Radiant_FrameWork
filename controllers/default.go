package controllers

import "radiant/radiant/core"

func SetTemplateVars(vars map[string]interface{}) map[string]interface{} {
	// TODO make dynamic for debug/dev env.
	vars["AssetHash"] = core.Configure.Asset["Hash"]
	return vars
}

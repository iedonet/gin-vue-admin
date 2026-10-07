package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/org/api"

var (
	Router                  = new(router)
	apiIedoOrgOrganizations = api.Api.IedoOrgOrganizations
)

type router struct{ IedoOrgOrganizations iedoOrgOrganizations }

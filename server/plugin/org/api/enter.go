package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/org/service"

var (
	Api                         = new(api)
	serviceIedoOrgOrganizations = service.Service.IedoOrgOrganizations
)

type api struct{ IedoOrgOrganizations iedoOrgOrganizations }

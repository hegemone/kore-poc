package api

import (
	"github.com/emicklei/go-restful"
	"github.com/dahendel/kore-poc/client"
)

// Creates all the routes. This can be broken up into different web services and registered in main.go
// This data is used to generate the swagger docs
func RegisterPaths() *restful.WebService {
	ws := new(restful.WebService)
	ws.
	Path("/kore").
		Doc("Kore POC API").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/allnames").
		To(allNames).Filter(basicAuth).
		Doc("Show me all the names. You must have a basic Auth account to see this endpoint").
		Reads(client.Name{}).
		Writes(client.Name{}).
		Returns(200, "", client.Name{}))

	ws.Route(ws.GET("/whatsmyname").
		To(getMyName).Filter(basicAuth).
		Doc("Tell me what my Name is. You must have a basic Auth account to see this endpoint").
		Returns(200, "", YourName{}))

	ws.Route(ws.POST("/savename").Filter(adminAuth).
		To(writeName).
		Doc("Create a new name. You must have an admin account to use this endpoint").
		Reads(client.Name{}).
		Writes(client.Name{}).
		Returns(200, "Created", client.Name{}))

	ws.Route(ws.PUT("/updatename").Filter(adminAuth).
		To(updateName).
		Doc("Update an existing Name. You must have an admin account to use this endpoint").
		Reads(client.Name{}).
		Writes(client.Name{}).
		Returns(200, "Update", client.Name{}))

	ws.Route(ws.DELETE("/deletename").Filter(adminAuth).
		To(deleteName).
		Doc("Delete a name based on name id. You must have an admin account to use this endpoint").
		Reads(DeleteCommand{}).
		Returns(200, "Update", DeleteCommand{}))

	return ws
}

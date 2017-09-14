package api

import (
	"github.com/emicklei/go-restful"

	"github.com/dahendel/kore-poc/kore-go-restful/client"
)


// Authenticates a basic user, does not check for admin privs
func basicAuth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {


	user, pass, ok := req.Request.BasicAuth()
	authUser := client.GetUser(Db, user)


	if !ok || user != authUser.Username || pass != authUser.Password {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}


	chain.ProcessFilter(req, resp)
}


// Checks basic Auth and makes sure admin provs are present
func adminAuth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {

	user, pass, ok := req.Request.BasicAuth()
	authUser := client.GetUser(Db, user)


	if !ok || user != authUser.Username || pass != authUser.Password || authUser.Admin != 1 {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}


	chain.ProcessFilter(req, resp)
}

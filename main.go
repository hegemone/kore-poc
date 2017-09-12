package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/dahendel/kore-poc/api"
	"net/http"
)

func main() {

	// Register the restful container
	container := restful.NewContainer()
	ws := api.RegisterPaths()
	container.Add(ws)

	// Setup the generated swagger docs
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/kore/docs/apidocs.json",
		SwaggerPath:     "/kore/docs/",
		SwaggerFilePath: "./docs"}

	swagger.RegisterSwaggerService(config, container)

	// The api is listening on port 8080
	http.ListenAndServe("0.0.0.0:8080", container)

}

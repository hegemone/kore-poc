# Kore POC

 
    - Written in go with go-restful.
    - Read, write, update, and delete to a sqlite db
    - Uses filters for Basic Auth. A filter can be applied at the root level or to each subpath of the api
    - Two db's are generated when the app launches, Names and Users
    - All read and writes affect the Names DB
    - The user DB is used for Basic Auth demo's
    - When the app is launched localhost:8080/kore/docs provides the generated swagger docs.
    - The build takes a second because of the swagger doc. :(

```
go get github.com/dahendel/kore-poc
glide install
go build
./kore-poc
```
# Kore POC

 
    - Written in go with go-restful.
    - Read, write, update, and delete to a sqlite db
    - Uses filters for Basic Auth. A filter can be applied at the root level or to each subpath of the api
    - Two db's are generated when the app launches, Names and Users
    - All read and writes affect the Names DB
    - The user DB is used for Basic Auth demo's
    - When the app is launched localhost:8080/kore/docs provides the generated swagger docs.
    - The build takes a second because of the swagger doc. :(

#### Install and Run
go get the repo and install dependencies
```
go get -d github.com/dahendel/kore-poc
cd $GOPATH/src/github.com/dahendel/kore-poc
glide install
```

Install the swagger-ui into the 3rdparty folder. This is ignored in the .ignore file
```
cd 3rdparty

// This needs to be done so that the UI will rebder. There is a symlink in the docs folder that points 
// to this folder. This directory is ignored in the .gitignore
git clone https://github.com/swagger-api/swagger-ui.git

cd ..
```

Build and Run the application
```
go build

./kore-poc
```


#### Hitting the api

Base-URL: http://localhost:8080/kore

Paths:
  - /whatsmyname (GET)   -- returns a name from the DB -- Normal account
  - /savename  (POST)    -- writes a name to the DB -- Admin Account
  - /allnames (GET)      -- displays list of all the names in the DB -- Normal Account
  - /updatename (PUT)    -- updates a name in the DB -- Admin Account
  - /deletename (DELETE) -- Deletes a name based on the id passed in the request body -- Admin Account
  
Auth: 
  - Basic auth is used to make the requests
  - Admin account: jbot/secretsJB
  - Normal account: regs/password

Docs-URL: http://localhost:8080/kore/docs



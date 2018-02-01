# Buffalo POC

Buffalo is a go web framework that will build the majority of your code for you. Unlike goa, it assumes you will be adding more code.


Install buffalo `go get -u -v github.com/gobuffalo/buffalo/buffalo` `go get -u github.com/gobuffalo/buffalo-goth`

Some benefits:

- It will generate the DB tables and all CRUD operations for you.
- Literally 5 minutes to working api with DB
- buffalo dev runs the web app locally and you can test input - recompiles on save
- Supports TONS of auth methods, and can write your own

Please see https://gobuffalo.io/docs for more information

Auth can be tied in with many different providers, for example `buffalo g goth-auth digitalocean discord` will setup DO and discord auth

Auth video: https://gobuffalo.io/docs/generators#goth

Auth with username and password example (Custom auth example): https://github.com/gobuffalo/authrecipe/blob/master/actions/auth.go


## Database Setup

Run `docker run -d --name pgsql --net=host -e POSTGRES_PASSWORD=password postgres:alpine`
get the md5 sum of password `echo "password" | md5sum`
edit the database.yml

To add a new resource besides the ones defined in the design/002-data-server.md type the following: `buffalo g resource newResource name id:int email address`

This will create a new resource and a new migration in the migrations folder. Once all resources are defined run the below command to create the db.

Run `buffalo db create -a`

Then run `buffalo db migrate` to create all of the tables

### CreRun the app locally

`buffalo dev`

Visit 127.0.01:3000 there you can test endpoints

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)

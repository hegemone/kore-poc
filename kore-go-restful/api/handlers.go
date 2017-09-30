package api

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/dahendel/kore-poc/kore-go-restful/client"
	"log"
	"math/rand"
	"time"
	//"encoding/json"
	//"net/http"
)

var Db *client.Client

// Create our Db connection
func init() {
	Db = client.New()
}

type YourName struct {
	Name string `json:"name"`
}

type DeleteCommand struct {
	Id int `json:"id"`
}

type AllNames struct {
	Results []struct{
		Names *interface{} `json:"names"`
	} `json:"results"`
}

// Generate a random number in the given min max range
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Reads the Names from the Db then selects a random one for your name
func getMyName(_ *restful.Request, response *restful.Response) {

	respBody := &YourName{}

	data := client.ReadName(Db)

	dataLen := len(data)

	// Select a random name in the slice
	name := data[random(0, dataLen)]

	// Write the Ful Name
	yourName := fmt.Sprintf("%s %s", name.First, name.Last)

	respBody.Name = yourName

	response.WriteEntity(respBody)
}

func allNames(_ *restful.Request, response *restful.Response) {

	//respBody := &AllNames{}
	results := client.ReadName(Db)
	//
	//err := json.Unmarshal(results, respBody)

	response.WriteEntity(results)
}

// Takes in a client.Name and writes it to the Db
func writeName(request *restful.Request, response *restful.Response) {

	reqBody := &client.Name{}

	request.ReadEntity(reqBody)

	log.Printf("New Name: %s", reqBody)

	client.WriteName(Db, reqBody)

	response.WriteEntity("OK!")
}


// Takes in a client.Name and writes it to the Db
func updateName(request *restful.Request, response *restful.Response) {

	reqBody := &client.Name{}

	request.ReadEntity(reqBody)

	log.Printf("New Name: %s", reqBody)

	client.WriteName(Db, reqBody)

	response.WriteEntity("OK!")
}


// Takes in a client.Name and writes it to the Db
func deleteName(request *restful.Request, response *restful.Response) {

	reqBody := &DeleteCommand{}

	request.ReadEntity(reqBody)

	log.Printf("New Name: %s", reqBody)

	client.Delete(Db, reqBody.Id)

	response.WriteEntity("OK!")
}

package main

import (
	"fmt"
	"./twitch"
	"net/http"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	fmt.Println("Booting the server...")

	// Configure a sample route
	http.HandleFunc("/sample_route", myHandlerFunc)
	http.HandleFunc("/", myTwitchHandler)


	// Run your server
	http.ListenAndServe(":8080", nil)
}

// myHandlerFunc - A sample handler function for the route /sample_route for your HTTP server
func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	twitch.DoSomething()

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} else{
        responseData, err := ioutil.ReadAll(response.Body)
      	if err != nil {
			log.Fatal(err)
		}
        fmt.Println(string(responseData))
    }

}

func myTwitchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.
	request, err := http.NewRequest("GET", "http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} else{
        responseData, err := ioutil.ReadAll(response.Body)
      	if err != nil {
			log.Fatal(err)
		}
        fmt.Println(string(responseData))
    }

}

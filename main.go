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
	http.HandleFunc("/users", myTwitchHandler)
	http.HandleFunc("/channel", myChannelHandler)
	http.HandleFunc("/streams", myStreamHandler)


	// Run your server
	http.ListenAndServe(":8080", nil)
}

func myTwitchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	twitch.DoSomething()

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.

	client := &http.Client{}

	// resp, err := client.Get("https://api.twitch.tv/kraken/users?login=dallas,dallasnchains")
	
	request, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/users?login=dallas,dallasnchains", nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", "ow8er2yjqzvbwzvfx832rrb5ucy933")
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	resp, err := client.Do(request)
        
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}
    fmt.Println(string(responseData))

}

func myChannelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.

	client := &http.Client{}

	// resp, err := client.Get("https://api.twitch.tv/kraken/users?login=dallas,dallasnchains")
	
	request, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/channels/44322889", nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", "ow8er2yjqzvbwzvfx832rrb5ucy933")
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	resp, err := client.Do(request)
        
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}
    fmt.Println(string(responseData))

}

func myStreamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.

	client := &http.Client{}

	// resp, err := client.Get("https://api.twitch.tv/kraken/users?login=dallas,dallasnchains")
	
	request, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/streams/44322889", nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", "ow8er2yjqzvbwzvfx832rrb5ucy933")
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	resp, err := client.Do(request)
        
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}
    fmt.Println(string(responseData))

}

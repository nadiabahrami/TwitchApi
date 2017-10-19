package main

import (
	"fmt"
	"net/http"
	"path"
	"io/ioutil"
	"os"
	"encoding/json"
	"log"
)

type APIreturn struct {
	Bio string
	AccountCreateDate string
	DisplayName string
	CurrentlyStreaming bool
	Language string
	Game string
	NumberOfFollowers int
	NumberOfViews int
}

func main() {
	// clientID := os.Getenv("CLIENT_ID")
	fmt.Println("Booting the server...")

	// Configure a sample route
	http.HandleFunc("/", myTwitchHandler)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

func myTwitchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)
	clientID := os.Getenv("CLIENT_ID")
	// twitch.DoSomething()
	username := path.Base(r.URL.String())

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.

	client := &http.Client{}

	// resp, err := client.Get("https://api.twitch.tv/kraken/users?login=dallas,dallasnchains")
	
	request, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/users?login=" + username, nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", clientID)
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	resp, err := client.Do(request)
        
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}
    
    var responseObject GetUsersAPIResponse
	json.Unmarshal(responseData, &responseObject)

  	user_id := responseObject.UserList[0].Id

  	fmt.Println("Making Second Request")
  	request, err = http.NewRequest("GET", "https://api.twitch.tv/kraken/channels/" + user_id, nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", clientID)
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	client = &http.Client{}
	resp, err = client.Do(request)
        
    responseData2, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}

    var responseObject2 GetChannel
	json.Unmarshal(responseData2, &responseObject2)

    fmt.Println("Making Third Request")
    
    client = &http.Client{}
	request, err = http.NewRequest("GET", "https://api.twitch.tv/kraken/streams/" + user_id, nil)
	if err != nil {
		fmt.Println(err.Error())
	 	os.Exit(1)
	} 
	request.Header.Set("Client-ID", clientID)
	request.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	resp, err = client.Do(request)
        
    responseData3, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}

    var responseObject3 GetStatus
	json.Unmarshal(responseData3, &responseObject3)

	if responseObject3.Stream != nil {
		responseObject3.Status = true
	} else {
		responseObject3.Status = false
	}

	write := APIreturn{responseObject.UserList[0].Bio, responseObject.UserList[0].CreatedDate, responseObject.UserList[0].DisplayName, responseObject3.Status, responseObject2.Language, responseObject2.Game, responseObject2.Followers, responseObject2.Views}
	data, err := json.Marshal(write)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(data))
}

type GetStatus struct {
	Stream *map[string]interface{} `json:"stream"`
	Status bool
}

type GetUsersAPIResponse struct {
	TotalUsers int `json:"_total"`
	UserList []Users `json:"users"`
}

type GetChannel struct {
	Game string `json:"game"`
	Language string `json:"language"`
	Views int `json:"views"`
	Followers int `json:"followers"`
}

type Users struct {
	Id string `json:"_id"`
	DisplayName string `json:"display_name"`
	Name string `json:"name"`
	Type string `json:"type"`
	Bio string `json:"bio"`
	CreatedDate string `json:"created_at"`
	LastUpdated string `json:"updated_at"`
	Logo string `json:"logo"`
}

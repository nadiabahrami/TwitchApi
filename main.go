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
    fmt.Fprintf(w, string(responseData))
    
    var responseObject GetUsersAPIResponse
	json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject.TotalUsers)
  	fmt.Println(responseObject.UserList[0].Id)
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
    fmt.Fprintf(w, string(responseData2))

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
	fmt.Fprintf(w, "Stream Status")
    fmt.Fprintf(w, string(responseData3))

    var responseObject3 GetStatus
	json.Unmarshal(responseData3, &responseObject3)

	fmt.Println(responseObject3.Stream)
	fmt.Printf("%+v\n", responseObject3)
	fmt.Println("***********************")
	if responseObject3.Stream != nil {
		responseObject3.Status = true
	} else {
		responseObject3.Status = false
	}

	fmt.Println(responseObject3.Status)
}

type Stream struct {
		Game string `json:"game"`
}

type GetStatus struct {
	Stream *map[string]interface{} `json:"stream"`
	Status bool
}

type GetUsersAPIResponse struct {
	TotalUsers int `json:"_total"`
	UserList []Users `json:"users"`
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

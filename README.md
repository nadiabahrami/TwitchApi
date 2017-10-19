# TwitchApi

## Description:



## Requirements:

 * Go: [Download](https://golang.org/dl/)
 * Client ID from Twitch.  Sign up in [Twitch Developer Portal](https://dev.twitch.tv/get-started) and obtain a Client ID


## Installation Steps
 1. Clone down the [TwitchApi Repo](https://github.com/nadiabahrami/TwitchApi)
 
## Run the project
 1. Cd into the TwitchAPI repo ```cd TwitchAPI```
 2. Start the server using your Client ID.  Client ID Must be in quotes. ```CLIENT_ID="<Your_client_ID>" go run main.go```
 3. In the browser navigate to ```http://0.0.0.0:8080/<username>``` <username> is the username of the twitch user you want stats of.  



## Example:

### Request
```http://0.0.0.0:8080/dallas```

### Response
```
{"Bio":"Friendly, interactive, and very bald. I work as a software engineer at Twitch. I stream regularly and play mostly new releases or something from my backlog.","AccountCreateDate":"2013-06-03T19:12:02.580593Z","DisplayName":"dallas","CurrentlyStreaming":false,"Language":"en","Game":"Cuphead","NumberOfFollowers":118,"NumberOfViews":6866}
```
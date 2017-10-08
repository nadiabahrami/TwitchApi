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
```{"_total":1,"users":[{"display_name":"dallas","_id":"44322889","name":"dallas","type":"staff","bio":"Just a gamer playing games and chatting. :)","created_at":"2013-06-03T19:12:02.580593Z","updated_at":"2017-10-08T18:32:35.036325Z","logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png"}]}{"mature":true,"status":"To the Southern Wilds! Continuing our Battle Chasers journey.","broadcaster_language":"en","display_name":"dallas","game":"Battle Chasers: Nightwar","language":"en","_id":"44322889","name":"dallas","created_at":"2013-06-03T19:12:02Z","updated_at":"2017-10-08T18:32:35Z","partner":false,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-2e82c1df2a464df7-1920x1080.jpeg","profile_banner":null,"profile_banner_background_color":null,"url":"https://www.twitch.tv/dallas","views":6463,"followers":106,"broadcaster_type":"affiliate","description":"Just a gamer playing games and chatting. :)"}Stream Status{"stream":{"_id":26439228944,"game":"Battle Chasers: Nightwar","broadcast_platform":"live","community_id":"e68ad29f-40e0-421d-873a-ca9cc5312461","community_ids":["e68ad29f-40e0-421d-873a-ca9cc5312461"],"viewers":1,"video_height":1080,"average_fps":60,"delay":0,"created_at":"2017-10-08T18:19:40Z","is_playlist":false,"stream_type":"live","preview":{"small":"https://static-cdn.jtvnw.net/previews-ttv/live_user_dallas-80x45.jpg","medium":"https://static-cdn.jtvnw.net/previews-ttv/live_user_dallas-320x180.jpg","large":"https://static-cdn.jtvnw.net/previews-ttv/live_user_dallas-640x360.jpg","template":"https://static-cdn.jtvnw.net/previews-ttv/live_user_dallas-{width}x{height}.jpg"},"channel":{"mature":true,"status":"To the Southern Wilds! Continuing our Battle Chasers journey.","broadcaster_language":"en","display_name":"dallas","game":"Battle Chasers: Nightwar","language":"en","_id":44322889,"name":"dallas","created_at":"2013-06-03T19:12:02.580593Z","updated_at":"2017-10-08T18:32:35.036325Z","partner":false,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-2e82c1df2a464df7-1920x1080.jpeg","profile_banner":null,"profile_banner_background_color":"","url":"https://www.twitch.tv/dallas","views":6463,"followers":106,"broadcaster_type":"","description":"Just a gamer playing games and chatting. :)"}}}```
# mytube

An API to fetch latest videos from youtube sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

The server fetches latest videos async after every 30 minutes.

This project is completely based on Golang.

Method Used:- 
I have used cron job of golang to run fetch videos asynchronously after every 30 mins. In this we are keeping max size as 50, i.e fetching 50 videos in one call to youtube api. 

Setup Guide:-
1. Clone the project
2. Run setup.dev.sh
3. It will create docker image of mytube, also installs golang1.15 version. 
4. Please set your GOROOT as usr/local/go, GOPATH will be set automatically inside docker file.
5. All dependency are inside vendor package and no need to reinstall them. 
6. After this run/debug main.go inside any ide. 
7. If everything works fine, then  "[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080" this line will appear in your terminal. 
8. To test server, go to browser and do "http://localhost:8080/ping". It will give result in pong. It means server is up and ready. 

videos Api Info:- 

sample :- http://localhost:8080/mytube/videos? 

This will give response of first 10 recently posted videos for "cricket" on youtube in descending order of their published data i.e most recent one will appear first.  Also to get next 10 videos i.e pagination , use "next_vid" given in response and pass its value in "next" as query param. 

sample:- http://localhost:8080/mytube/videos?next=e9V7sBvLiwE 

This will give response of next 10 or available hotels (if <less than 10). When no hotels will left, then next_vid value will be "END". It means no more videos are stored in system. 

Logic of next:- next_vid is first videoid of next page. So by passing this as query param, server will get to know, from which vid response need to send. 


Search Api Info:- 

sample :- 
http://localhost:8080/mytube/search?f={"title":"Dhoni vs Brett Lee in International Cricket History"} or 
http://localhost:8080/mytube/search?f={%22video_id%22:%222iGfuu13-LU%22} 

This will search video with video_id/title/description in stored video data and return response for that. 




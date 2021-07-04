package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	"config"
	"constants/enums/environments"
	"youtube_fetch"
)

var (
	SERVER_PORT string
	ENVIRONMENT string
)

// init gets called at the start of the file load and main function.
// thus it becomes suitable for setting the environment of the
// application.
func init() {

	// setting the environment of the application
	setEnvironment()

	// there can be different kind of db and external services config for all
	// the environments. So taking the the config from corresponding settings file.
	config.DoInit(ENVIRONMENT)

	// defining the port of the application server
	SERVER_PORT = config.Config.GetString("server.port")

	// set the gin server mode
	setGinMode()
}

//setEnvironment function takes the arguments from the commandline
//and sets the ENVIRONMENT. The default ENVIRONEMNT is 'dev'.
func setEnvironment() {
	// get the flag from command line
	envFlagPtr := flag.String("env", "", "environment: dev, prod")
	flag.Parse()

	//set the global variable
	config.Environment = environments.GetEnvironment(*envFlagPtr)
	ENVIRONMENT = config.Environment.String()

	fmt.Println("ENVIRONMENT:" + ENVIRONMENT)
}

// setGinMode takes the mode information from settings file and sets the gin
// server mode of operation.
func setGinMode() {
	if config.Config.GetString("server.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.Config.GetString("server.mode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if config.Config.GetString("server.mode") == "test" {
		gin.SetMode(gin.TestMode)
	}
}
const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

var (
	YtService *youtube.Service
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, authConfig *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(authConfig)
		saveToken(cacheFile, tok)
	}
	return authConfig.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func main() {

	ctx := context.Background()
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	authConfig, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to authConfig: %v", err)
	}
	client := getClient(ctx, authConfig)
	var connectionErr error
	YtService, connectionErr = youtube.New(client)
	handleError(connectionErr, "Error creating YouTube client")

	mainRouter := gin.Default()
	// Ping test
	mainRouter.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	mainRouter.GET("/mytube", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//starts the server
	mainRouter.Run("0.0.0.0" + SERVER_PORT)

	data := youtube_fetch.FetchSearchData(YtService, []string{"snippet"})
	fmt.Println("success", data)
}

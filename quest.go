package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

// Credentials which stores google ids.
type Credentials struct {
    Cid     string `json:"cid"`
    Csecret string `json:"csecret"`
}

var cred Credentials

func indexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func battleHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "battle.tmpl", gin.H{
        "user": "Anyad",
    })
}

func init() {
    file, err := ioutil.ReadFile("./creds.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    json.Unmarshal(file, &cred)
}

// AuthRequired will authorize requests.
func AuthRequired(c *gin.Context) {
    c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{})
    c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("You shall not pass!")) // This is how to stop rendering further if the auth failed
}

func login(c *gin.Context) {

    conf := &oauth2.Config{
        ClientID:     cred.Cid,
        ClientSecret: cred.Csecret,
        RedirectURL:  "http://localhost:9090/login", // In case the login fails, redirect to the login page.
        Scopes: []string{
        "https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
        },
        Endpoint: google.Endpoint,
    }

    // Redirect user to Google's consent page to ask for permission
    // for the scopes specified above.
    url := conf.AuthCodeURL("state")
    fmt.Printf("Visit the URL for the auth dialog: %v\n", url)
    c.Redirect(http.StatusMovedPermanently, url)
    // Handle the exchange code to initiate a transport.
    tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
    if err != nil {
        c.AbortWithError(http.StatusUnauthorized, err)
    }
    client := conf.Client(oauth2.NoContext, tok)
    fmt.Println(client)
}

func main() {
    router := gin.Default()
    router.Static("/css", "./static/css")
    router.Static("/img", "./static/img")
    router.LoadHTMLGlob("templates/*")

    router.GET("/", indexHandler)

    authorized := router.Group("/battle")
    authorized.Use(login)
    {
        authorized.GET("/", battleHandler)
    }

    router.Run(":9090")
}

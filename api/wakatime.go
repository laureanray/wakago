package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
)

func openBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Unsupported platform")
	}

	return err
}

const baseLoginUrl = "https://wakatime.com/oauth/authorize"
const redirectUrl = "http://127.0.0.1:8090/wakago/callback"

var wakatimeOauthConfig = &oauth2.Config{
	RedirectURL:  redirectUrl,
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes:       []string{"email,", "read_logged_time,", "read_stats"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://wakatime.com/oauth/authorize",
		TokenURL:  "https://wakatime.com/oauth/token",
		AuthStyle: 0,
	},
}

// TODO: Add expiration?
func generateToken() string {
	b := make([]byte, 16)

	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func Login() (err error) {
	token := generateToken()
	u := wakatimeOauthConfig.AuthCodeURL(token)
	fmt.Println(u)

	err = openBrowser(u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

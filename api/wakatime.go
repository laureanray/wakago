package api

// TODOS:
// [] Store access token (prevent from getting token everytime)
// [] Access Token refresh mechanism
// [] Terminal output format mechanism
// [] Better structure
// [] Unit Tests

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
)

type Wakatime struct {
	oauth2     *oauth2.Config
	oauthToken string
	client     *http.Client
}

var wtInstance *Wakatime

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
const wakatimeApiUrl = "https://wakatime.com/api/v1"
const redirectUrl = "http://127.0.0.1:8090/wakago/callback"

var wakatimeOauthConfig = &oauth2.Config{
	RedirectURL:  redirectUrl,
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes:       []string{"email,", "read_logged_time,", "read_stats,", "read_orgs"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://wakatime.com/oauth/authorize",
		TokenURL:  "https://wakatime.com/oauth/token",
		AuthStyle: 0,
	},
}

func (wt *Wakatime) Init(clientId, clientSecret string) {
	fmt.Println("init")
	(*wt).oauth2 = &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"email,", "read_logged_time,", "read_stats"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://wakatime.com/oauth/authorize",
			TokenURL:  "https://wakatime.com/oauth/token",
			AuthStyle: 1,
		},
	}

	(*wt).client = &http.Client{}
}

// TODO: Add expiration?
func generateToken() string {
	b := make([]byte, 16)

	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func GetInstance() *Wakatime {
	if wtInstance == nil {
		wtInstance = new(Wakatime)
	}

	return wtInstance
}

func (wt *Wakatime) Login() (err error) {
	token := generateToken()
	u := wakatimeOauthConfig.AuthCodeURL(token)

	err = openBrowser(u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (wt *Wakatime) getToken(code string) (*oauth2.Token, error) {
	token, err := (*wt).oauth2.Exchange(context.Background(), code)

	return token, err
}

func (wt *Wakatime) GetGoals() (Goals, error) {
	url := fmt.Sprintf("%s/users/current/goals", wakatimeApiUrl)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", (*wt).oauthToken))

	resp, err := (*wt).client.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	var goals Goals

	err = json.NewDecoder(resp.Body).Decode(&goals)

	if err != nil {
		log.Println("Error reading response")
	}

	return goals, err
}

func (wt *Wakatime) Exchange(code string) error {
	accessToken, err := (*wt).oauth2.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Exchange:", err)
	}
	(*wt).oauthToken = accessToken.AccessToken
	return err
}

func (wt *Wakatime) Status() string {
	var status string
	if (*wt).oauthToken != "" {
		status = fmt.Sprintf("✓ Access Token Set: %s", (*wt).oauthToken)
	}

	return status
}

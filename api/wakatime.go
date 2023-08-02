package api

// TODOS:
// [X] Store access token (prevent from getting token everytime)
// [] Access Token refresh mechanism
// [] Terminal output format mechanism
// [] Better structure
// [] Unit Tests

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/laureanray/oauth2"
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

func (wt *Wakatime) Init(clientId, clientSecret string) {
	(*wt).oauth2 = &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"email", "read_stats"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://wakatime.com/oauth/authorize",
			TokenURL:  "https://wakatime.com/oauth/token",
			AuthStyle: 1,
		},
	}

	(*wt).client = &http.Client{}
	(*wt).initAppData()
}

// TODO: Improve this to use key value store or something more robust
func (wt *Wakatime) initAppData() {
	_, err := os.Stat("wakatime.data")

	if errors.Is(err, os.ErrNotExist) {
		log.Println("File doesnt exist")
	} else {
		fileData, err := ioutil.ReadFile("wakatime.data")
		if err != nil {
			log.Fatal(err)
		}
		if len(fileData) != 0 {
			// Read the data and use it

			(*wt).oauthToken = string(fileData)
		}
	}
}

func (wt *Wakatime) saveAppData() {
	f, err := os.OpenFile("wakatime.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	bytesWritten, err := f.WriteString((*wt).oauthToken)

	log.Printf("%d bytes written", bytesWritten)
	return
}

// TODO: Add expiration?
func generateToken() string {
	// Generate a random 40-byte slice
	randomBytes := make([]byte, 40)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Calculate the SHA-1 hash of the random bytes
	sha1Hash := sha1.Sum(randomBytes)

	// Convert the SHA-1 hash to hexadecimal representation
	sha1Hex := hex.EncodeToString(sha1Hash[:])

	return sha1Hex
}

func GetInstance() *Wakatime {
	if wtInstance == nil {
		wtInstance = new(Wakatime)
	}

	return wtInstance
}

func (wt *Wakatime) Login() (err error) {
	state := generateToken()
  log.Printf("state: %s", state)
	u := (*wt).oauth2.AuthCodeURL(state)

	err = openBrowser(u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (wt *Wakatime) sendAuthenticatedRequest(requestUrl string) (*http.Response, error) {
	req, err := http.NewRequest("GET", requestUrl, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", (*wt).oauthToken))

	resp, err := (*wt).client.Do(req)

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP Status:", resp.StatusCode)
	}

	if err != nil {
		log.Println("Error sending request", err)
	}

	return resp, err
}

func (wt *Wakatime) GetGoals() (Goals, error) {
	// TODO: Impl better cache mechanism
	url := fmt.Sprintf("%s/users/current/goals?cache=false", wakatimeApiUrl)
	resp, err := (*wt).sendAuthenticatedRequest(url)

	if err != nil {
		log.Println(err)
	}

	var goals Goals

	err = json.NewDecoder(resp.Body).Decode(&goals)

	if err != nil {
		log.Println("Error reading response", err)
	}

	defer resp.Body.Close()

	return goals, err
}

func (wt *Wakatime) GetStatusBar() (StatusBar, error) {
	url := fmt.Sprintf("%s/users/current/status_bar/today?cache=false", wakatimeApiUrl)
	resp, err := (*wt).sendAuthenticatedRequest(url)

	if err != nil {
		log.Println(err)
	}

	var statusBar StatusBar

	err = json.NewDecoder(resp.Body).Decode(&statusBar)

	if err != nil {
		log.Println("Error reading respnse", err)
	}

	defer resp.Body.Close()
	return statusBar, err
}

func (wt *Wakatime) Exchange(code string) error {
  log.Println("Trying to exchange:", code)
	accessToken, err := (*wt).oauth2.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Exchange:", err)
	}
	(*wt).oauthToken = accessToken.AccessToken
	(*wt).saveAppData()

	// TODO: Update this implementation
	//defer os.Exit(0)
	return err
}

func (wt *Wakatime) Status() string {
	var status string
	if (*wt).oauthToken != "" {
		status = fmt.Sprintf("✓ Access Token Set: %s", (*wt).oauthToken)
	}

	return status
}

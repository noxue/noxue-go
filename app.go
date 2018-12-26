package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

const htmlIndex = `<html><body>
<a href="/GithubLogin">Log in with Github</a>
</body></html>
`

// Endpoint is Github's OAuth 2.0 endpoint.
var endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "f4916bb66a237fcd99ab",
	ClientSecret: "0d89ba35c753ca0c72b870f04bd6b8f29cb7caac",
	RedirectURL:  "http://localhost:8000/GithubCallback",
	Scopes:       []string{"user", "repo"},
	Endpoint:     endpoint,
}

const oauthStateString = "random"

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GithubLogin", handleGithubLogin)
	http.HandleFunc("/GithubCallback", handleGithubCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(state)

	code := r.FormValue("code")
	fmt.Println(code)
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	fmt.Println(token)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "%s\n", contents)
}
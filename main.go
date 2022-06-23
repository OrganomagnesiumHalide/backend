package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var logins map[string]string

func verifyGitHubLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("host=", r.Header.Get("Origin"))
	nonce := r.URL.Query().Get("nonce")
	origin := r.Header.Get("Origin")
	fmt.Println("origin=",origin);
	if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin,"https://organomagnesiumhalide.github.io") {
		fmt.Println("origin set:",origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Println(logins)
	if val, exists := logins[nonce]; exists {
		returnMap := make(map[string]any)
		returnMap["ok"] = true
		returnMap["access_token"] = val
		byteVals, err := json.Marshal(returnMap)
		if err != nil {
			panic(err)
		}
		w.Write(byteVals)
	} else {
		returnMap := make(map[string]any)
		returnMap["ok"] = false
		byteVals, err := json.Marshal(returnMap)
		if err != nil {
			panic(err)
		}
		w.Write(byteVals)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("Host")
	if host != "fridgigator.herokuapp.com" || host != "localhost" {
		panic("Wrong host!")
        }

	fmt.Println(r.URL.Query())
	code := r.URL.Query().Get("code")
	sendMap := make(map[string]string)
	sendMap["client_id"] = "30bf4172998cc4ec684e"
	sendMap["client_secret"] = os.Getenv("CLIENT_SECRET")
	sendMap["code"] = code
	sendMap["redirect_uri"] = host+"/register"
	b, err := json.Marshal(sendMap)
	if err != nil {
		fmt.Println("Can't serialize", sendMap)
	}
	fmt.Println(sendMap)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Can't send request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	decoder := json.NewDecoder(resp.Body)
	var githubResponse map[string]string
	err = decoder.Decode(&githubResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(githubResponse)
	fmt.Println(r.URL.Query().Get("state"))
	logins[r.URL.Query().Get("state")] = githubResponse["access_token"]
	fmt.Println(logins)

}
func main() {
	logins = make(map[string]string)

	port := os.Getenv("PORT")

	if port == "" {
		port = "1234"
	}
	fmt.Println("Ready")
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler)
	mux.HandleFunc("/verifyGitHubLogin", verifyGitHubLogin)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

type GitHubResponse struct {
	access_token string
	scope        string
}

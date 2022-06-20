package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	code := r.URL.Query().Get("code")
	sendMap := make(map[string]string)
	sendMap["client_id"] = "30bf4172998cc4ec684e"
	sendMap["client_secret"] = os.Getenv("CLIENT_SECRET")
	sendMap["code"] = code
	sendMap["redirect_uri"] = "http://localhost:1234/register"
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
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "1234"
	}
	fmt.Println("Ready")
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

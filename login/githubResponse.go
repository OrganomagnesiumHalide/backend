package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Fridgigator/backend/globals"
)

func GithubResponse(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host != "fridgigator.herokuapp.com" && !strings.HasPrefix(host, "localhost:") {
		panic(fmt.Sprintf("Wrong host=%s", host))
	}
	fmt.Println(r.URL.Query())
	if r.URL.Query().Has("error") {
		globals.StateMap[r.URL.Query().Get("state")] <- ""
		return
	}
	code := r.URL.Query().Get("code")
	sendMap := make(map[string]string)
	sendMap["client_id"] = "30bf4172998cc4ec684e"
	sendMap["client_secret"] = os.Getenv("CLIENT_SECRET")
	sendMap["code"] = code
	sendMap["state"] = r.URL.Query().Get("state")
	b, err := json.Marshal(sendMap)
	if err != nil {
		panic(fmt.Sprintln("Can't serialize", sendMap))
	}
	fmt.Println(sendMap)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(b))
	if err != nil {
		panic(fmt.Sprintln("Can't send request"))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var githubResponse map[string]string
	err = decoder.Decode(&githubResponse)
	if err != nil {
		panic(err)
	}
	globals.StateMapMutex.Lock()
	if _, exists := globals.StateMap[r.URL.Query().Get("state")]; !exists {
		w.Write([]byte("Wrong state"))
		return
	}
	responseCh := globals.StateMap[r.URL.Query().Get("state")]
	globals.StateMapMutex.Unlock()
	responseCh <- githubResponse["access_token"]
	fmt.Fprintf(w, "Signed in!")
}

package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var logins map[string]string
var denied map[string]string
var db *sql.DB
var privateKey *rsa.PrivateKey

func verifyGitHubLogin(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Headers:")
	for name, values := range r.Header {
 	   // Loop over all values for the name.
	    for _, value := range values {
        	fmt.Println(name, value)
	    }
	}

	fmt.Println("host=", r.Header.Get("Origin"))
	nonce := r.URL.Query().Get("nonce")
	fmt.Println(nonce)
	origin := r.Header.Get("Origin")
	fmt.Println("origin=", origin)
	if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "https://fridgigator.github.io") {
		fmt.Println("origin set:", origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("logins=", logins)
	fmt.Println("denied=", denied)
	if val, exists := logins[nonce]; exists {
		returnMap := make(map[string]any)
		returnMap["ok"] = "ok"
		encryptedBytes, err := rsa.EncryptOAEP(
			sha256.New(),
			rand.Reader,
			&privateKey.PublicKey,
			[]byte(val),
			nil)

		if err != nil {
			panic(err)
		}

		returnMap["access_token"] = base64.StdEncoding.EncodeToString(encryptedBytes)
		byteVals, err := json.Marshal(returnMap)
		if err != nil {
			panic(err)
		}
		w.Write(byteVals)
	} else if _, exists := denied[nonce]; exists {
		returnMap := make(map[string]any)
		returnMap["ok"] = "denied"
		byteVals, err := json.Marshal(returnMap)
		if err != nil {
			panic(err)
		}
		w.Write(byteVals)
	} else {
		returnMap := make(map[string]any)
		returnMap["ok"] = "waiting"
		byteVals, err := json.Marshal(returnMap)
		if err != nil {
			panic(err)
		}
		w.Write(byteVals)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host != "fridgigator.herokuapp.com" && !strings.HasPrefix(host, "localhost:") {
		panic(fmt.Sprintf("Wrong host=%s", host))
	}
	fmt.Println(r.URL.Query())
	if r.URL.Query().Has("error") {
		denied[r.URL.Query().Get("state")] = ""
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
	fmt.Println("gr=", githubResponse)
	logins[r.URL.Query().Get("state")] = githubResponse["access_token"]

}
func main() {

	registerPrivateKey()

	logins = make(map[string]string)
	denied = make(map[string]string)

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

func getPrivateKeyFromDb() string {
	rows, err := db.Query(`SELECT private_key from private_key_data`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	var privateKeyString string
	err = rows.Scan(&privateKeyString)
	if err != nil {
		panic(err)
	}
	return privateKeyString
}
func registerPrivateKey() {
	psqlconn := fmt.Sprintf("user=postgres password=%s host=db.cldkdlstnxjwpxyqnkwb.supabase.co port=5432 dbname=postgres", os.Getenv("POSTGRES_PASSWD"))

	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	keyString := getPrivateKeyFromDb()
	if keyString != "" {
		privateKey = getPrivateKeyFromString(keyString)
	} else {
		getPrivateKeyString := getPrivateKeyString()
		_, err = db.Exec(`INSERT INTO private_key_data("private_key","short") values($1,$2)`, string(getPrivateKeyString), string(getPrivateKeyString)[100:2600])
		if err != nil {
			panic(err)
		}
		privateKey = getPrivateKeyFromString(getPrivateKeyString)
	}
}

package login

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Fridgigator/backend/globals"
	"github.com/Fridgigator/backend/protobuf_compiled"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func GetOauth2StateString(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "https://fridgigator.github.io")
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := ws.Close(); err != nil {
			panic(err)
		}
	}()

	curState := getRandomKey()
	globals.StateMap[curState] = make(chan string)
	defer safeDeleteStateMap(curState)
	var v protobuf_compiled.ToFrontEnd
	v.Type = &protobuf_compiled.ToFrontEnd_GithubState{GithubState: &protobuf_compiled.Nonce{Nonce: curState}}
	fmt.Println("nonce=", curState)

	log.Printf("received: %v", curState)

	messageBinary, err := proto.Marshal(&v)
	if err != nil {
		panic(err)
	}

	if err = ws.WriteMessage(websocket.BinaryMessage, messageBinary); err != nil {
		panic(err)
	}
	timeOut := make(chan bool)
	keepAlive(ws, timeOut)
	globals.StateMapMutex.Lock()
	if _, exists := globals.StateMap[curState]; !exists {
		fmt.Println("state doesn't exist")
		return
	}
	responseCh := globals.StateMap[curState]
	globals.StateMapMutex.Unlock()

	select {
	case github_nonce := <-responseCh:
		if github_nonce == "" {
			var v protobuf_compiled.ToFrontEnd
			v.Type = &protobuf_compiled.ToFrontEnd_AccessCode{AccessCode: &protobuf_compiled.AccessCode{Failed: true}}
			messageBinary, err := proto.Marshal(&v)
			if err != nil {
				panic(err)
			}
			if err = ws.WriteMessage(websocket.BinaryMessage, messageBinary); err != nil {
				panic(err)
			}

		} else {
			authKey := getRandomKey()
			var v protobuf_compiled.ToFrontEnd
			v.Type = &protobuf_compiled.ToFrontEnd_AccessCode{AccessCode: &protobuf_compiled.AccessCode{Failed: false, AccessCode: authKey}}

			messageBinary, err = proto.Marshal(&v)
			if err != nil {
				panic(err)
			}

			if err != nil {
				panic(err)
			}

			if err = ws.WriteMessage(websocket.BinaryMessage, messageBinary); err != nil {
				panic(err)
			}
		}
		break
	case <-timeOut:
		var v protobuf_compiled.ToFrontEnd
		v.Type = &protobuf_compiled.ToFrontEnd_AccessCode{AccessCode: &protobuf_compiled.AccessCode{Failed: true}}
		messageBinary, err := proto.Marshal(&v)
		if err != nil {
			panic(err)
		}
		if err = ws.WriteMessage(websocket.BinaryMessage, messageBinary); err != nil {
			panic(err)
		}

	}
	fmt.Println("Done selecting")

}

func getRandomKey() string {
	return uuid.NewString()

}
func keepAlive(c *websocket.Conn, timeoutChan chan bool) {

	go func() {
		for {
			err := c.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				log.Printf("Remote disconnected %s", c.LocalAddr())
				timeoutChan <- true
				return
			}
			time.Sleep(time.Duration(time.Second * 5))
		}
	}()
}

func safeDeleteStateMap(key string) {
	globals.StateMapMutex.Lock()
	delete(globals.StateMap, key)
	globals.StateMapMutex.Lock()
}

package globals

import (
	"database/sql"
	"sync"
)

var DB *sql.DB
var StateMap map[string]chan string
var StateMapMutex sync.Mutex

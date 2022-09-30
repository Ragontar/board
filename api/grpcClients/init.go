package grpcclients

import (
	"os"
)

var DEV bool = false
var ACCOUNT_SERVICE_ADDR string

func Init() {
	if DEV {
		accountServiceAddr = "0.0.0.0:9000"
		return
	}

	var ok bool
	if ACCOUNT_SERVICE_ADDR, ok = os.LookupEnv("ACCOUNT_SERVICE_ADDR"); !ok {
		panic("error: ACCOUNT_SERVICE_ADDR is undefined")
	}
	accountServiceAddr = ACCOUNT_SERVICE_ADDR
}

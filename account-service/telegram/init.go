package telegram

import (
	"account-service/dry"
	"log"
	"os"
)

func init() {
	ENV_APP_ID = dry.LookupOrPanic("ENV_APP_ID")
	ENV_APP_HASH = dry.LookupOrPanic("ENV_APP_HASH")
	ENV_PUBLIC_KEY_FILE = dry.LookupOrPanic("ENV_PUBLIC_KEY_FILE")
	ENV_STORAGE_PATH = dry.LookupOrPanic("ENV_STORAGE_PATH")

	log.Printf("[TELEGRAM] ENV_APP_ID: %s\n", ENV_APP_ID)
	log.Printf("[TELEGRAM] ENV_APP_HASH: %s\n", ENV_APP_HASH)
	log.Printf("[TELEGRAM] ENV_PUBLIC_KEY_FILE: %s\n", ENV_PUBLIC_KEY_FILE)
	log.Printf("[TELEGRAM] ENV_STORAGE_PATH: %s\n", ENV_STORAGE_PATH)

	if _, err := os.Stat(ENV_STORAGE_PATH); os.IsNotExist(err) {
		os.MkdirAll(ENV_STORAGE_PATH, 0700)
	}
}

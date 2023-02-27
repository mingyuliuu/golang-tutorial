package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func ParseBody(req *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetEnvVariable(key string) string {
  err := godotenv.Load("../../.env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}
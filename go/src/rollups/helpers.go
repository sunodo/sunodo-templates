package rollups

import (
	"encoding/json"
	"net/http"
	"bytes"
  "os"
)

var rollup_server = os.Getenv("ROLLUP_HTTP_SERVER_URL")

func SendFinish(finish *FinishRequest) (*http.Response, error) {
  bodyFinish, err := json.Marshal(finish)
  if err != nil {
    return &http.Response{}, err
  }
  
  return SendPost("finish", bodyFinish)
}

func SendPost(endpoint string, jsonData []byte) (*http.Response, error) {
  req, err := http.NewRequest(http.MethodPost, rollup_server + "/" + endpoint, bytes.NewBuffer(jsonData))
  if err != nil {
    return &http.Response{}, err
  }
  req.Header.Set("Content-Type", "application/json; charset=UTF-8")

  return http.DefaultClient.Do(req)
}

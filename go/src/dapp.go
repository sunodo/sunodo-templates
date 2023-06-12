package main

import (
	"encoding/json"
	"io/ioutil"
  "strconv"
    
  log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"

  "dapp/rollups"
)

func HandleAdvance(data *rollups.AdvanceResponse) string {
  dataMarshal, err := json.Marshal(data)
  if err != nil {
    log.Error("error:", err)
    return "reject"
  }
  log.Info("Received advance request data" + string(dataMarshal))
  return "accept"
}


func HandleInspect(data *rollups.InspectResponse) string {
  dataMarshal, err := json.Marshal(data)
  if err != nil {
    log.Error("error:", err)
    return "reject"
  }
  log.Info("Received inspect request data" + string(dataMarshal))
  return "accept"
}


func main() {
  finish := rollups.FinishRequest{"accept"}
  log.SetFormatter(&easy.Formatter{
    LogFormat:       "[%lvl%]: %msg%\n",
  })

  for true {

    log.Info("Sending finish")
    res, err := rollups.SendFinish(&finish)
    if err != nil {
      log.Fatal("client: error making http request: ", err)
      continue
    }
    log.Info("Received finish status " + strconv.Itoa(res.StatusCode))
    
    if (res.StatusCode == 202){
      log.Info("No pending rollup request, trying again")
    } else {

      resBody, err := ioutil.ReadAll(res.Body)
      if err != nil {
        log.Fatal("client: could not read response body: ", err)
        continue
      }
      
      var response rollups.FinishResponse
      err = json.Unmarshal(resBody, &response)
      if err != nil {
        log.Fatal("Error unmarshaling body:", err)
        continue
      }

      switch response.Type {
      case "advance_state":
        data := new(rollups.AdvanceResponse)

        err = json.Unmarshal(response.Data, data)
        if err != nil {
          log.Fatal("Error unmarshaling advance:", err)
          continue
        }
        finish.Status = HandleAdvance(data)
      case "inspect_state":
        data := new(rollups.InspectResponse)

        err = json.Unmarshal(response.Data, data)
        if err != nil {
          log.Fatal("Error unmarshaling inspect:", err)
          continue
        }
        finish.Status = HandleInspect(data)
      }
    }
  }
}
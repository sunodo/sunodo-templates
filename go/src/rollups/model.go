package rollups

import (
	"encoding/json"
)

type FinishRequest struct {
  Status string         `json:"status"`
}

type FinishResponse struct {
  Type string           `json:"request_type"`
  Data json.RawMessage  `json:"data"`
}

type InspectResponse struct {
  Payload string        `json:"payload"`
}

type AdvanceResponse struct {
  Metadata Metadata     `json:"metadata"`
  Payload string        `json:"payload"`
}

type Metadata struct {
  MsgSender string      `json:"msg_sender"`
  EpochIndex uint64     `json:"epoch_index"`
  InputIndex uint64     `json:"input_index"`
  BlockNumber uint64    `json:"block_number"`
  Timestamp uint64      `json:"timestamp"`
}


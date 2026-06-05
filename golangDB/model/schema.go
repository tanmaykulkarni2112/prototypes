package model

import "encoding/json"

type ResponseMsg struct {
	Message string `json:"message"`
}

type RequestBody struct {
	TableName string `json:"tablename" binding:"required"`
	Data      json.RawMessage `json:"data"`
}
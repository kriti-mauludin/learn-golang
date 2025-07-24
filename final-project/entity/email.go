package entity

import (
	"encoding/json"
	"time"
)

type SendEmailReq struct {
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *SendEmailReq) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}

package model

import "time"

type Log struct {
	Node       string    `json:"node"`
	Type       string    `json:"type"`
	Code       int32     `json:"code"`
	Module     string    `json:"module"`
	Tstamp     string    `json:"tstamp"`
	ShortText  string    `json:"shortText"`
	Text       string    `json:"text"`
	ServerTime time.Time `json:"serverTime"`
}

package entity

import (
	"time"
)

type Quote struct {
	ID               int       `json:"index"`
	Episode          int       `json:"episode"`
	EpisodeName      string    `json:"episode_name"`
	Segment          string    `json:"segment"`
	Type             string    `json:"type"`
	Actor            string    `json:"actor"`
	Character        string    `json:"character"`
	Detail           string    `json:"detail"`
	RecordDate       time.Time `json:"record_date"`
	Series           string    `json:"series"`
	TransmissionDate time.Time `json:"transmission_date"`
}

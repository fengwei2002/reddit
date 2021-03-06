package models

import "time"

// Community in models define the specific types in the Community type
type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	CommunityID   uint64    `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}

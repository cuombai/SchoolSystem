package models

import "time"

type AuditLog struct {
	Action string `bson: "action"`
	ActorID string `bson: "actor_id"`
	Timestamp time.Time  `bson: "timestamp"`
	Details string `bson:"details"`
}
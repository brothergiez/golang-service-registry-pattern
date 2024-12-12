package registry

import "time"

type Service struct {
	ID          string    `bson:"_id, omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Address     string    `bson:"address" json:"address"`
	Port        int       `bson:"port" json:"port"`
	RegistredAt time.Time `bson:"registered_at" json:"registered_at"`
}

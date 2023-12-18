package mongodbModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusAccount string

const (
	Free StatusAccount = ""
	Busy StatusAccount = "BUSY"
)

type Account struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Phone           string             `bson:"phone" json:"phone"`
	Name            string             `bson:"name" json:"name"`
	Role            int                `bson:"role" json:"role"`                         // khach hang hay nguoi cung cap dich vu
	ListServiceCode []string           `bson:"list_service_code" json:"listServiceCode"` //  neu la nha cung cap
	Status          string             `bson:"status" json:"status"`                     // trang thai
	IsOnline        string             `bson:"is_online" json:"-"`                       //
	IsRemoved       bool               `bson:"is_removed" json:"-"`                      //
	LastUpdate      time.Time          `bson:"last_update" json:"lastUpdate"`            // thoi gian lan cuoi update
	CreateAt        time.Time          `bson:"create_at" json:"-"`                       // thoi gian tao
}

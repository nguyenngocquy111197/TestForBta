package mongodbModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyServices struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Code       string             `bson:"code" json:"code"`              // ma dich vu
	Name       string             `bson:"name" json:"name"`              // ten dich vu
	Popularity int                `bson:"popularity" json:"-"`           // muc do pho bien , cao nhat la 1
	Price      float64            `bson:"price" json:"price"`            // muc do pho bien , cao nhat la 1
	IsRemoved  bool               `bson:"is_removed" json:"-"`           // check dich vu da bi remove hay chua
	LastUpdate time.Time          `bson:"last_update" json:"lastUpdate"` // thoi gian lan cuoi update
	CreateAt   time.Time          `bson:"create_at" json:"-"`            // thoi gian tao
}

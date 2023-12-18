package mongodbModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusTransaction string

const (
	Start                 StatusTransaction = "START"
	NoFindServiceProvider StatusTransaction = "NO_FIND_SERVICE_PROVIDER"
	Processing            StatusTransaction = "PROCESSING"
	Success               StatusTransaction = "SUCCESS"
	Fail                  StatusTransaction = "FAIL"
)

type Transaction struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ServiceCode       string             `bson:"service_code" json:"serviceCode"`              // ma dich vu
	CustomerId        primitive.ObjectID `bson:"customer_id" json:"customerId"`                // id nguoi booking
	ServiceProviderId primitive.ObjectID `bson:"service_provider_id" json:"serviceProviderId"` // id nguoi cung cap dich vu
	Status            StatusTransaction  `bson:"status" json:"status"`                         // trang thai gaio dich
	LastUpdate        time.Time          `bson:"last_update" json:"lastUpdate"`                // thoi gian lan cuoi update
	CreateAt          time.Time          `bson:"create_at" json:"-"`                           // thoi gian tao
}

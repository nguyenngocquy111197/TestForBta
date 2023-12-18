package apiModel

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CalculateReq struct {
	ServiceCode string             `json:"serviceCode"`
	CustomerId  primitive.ObjectID `json:"customerId"`
	Date        time.Time          `json:"date"` // format : RFC3339
}

func (it *CalculateReq) Validate() error {

	if len(it.CustomerId) == 0 {
		return errors.New("customerId empty")
	}

	if len(it.ServiceCode) == 0 {
		return errors.New("serviceCode empty")
	}
	return nil
}

type CalculateResp struct {
	BasicResp
	Price float64 `json:"price"`
}

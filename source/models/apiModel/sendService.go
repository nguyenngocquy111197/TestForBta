package apiModel

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SendInfoServiceReq struct {
	TransactionId     primitive.ObjectID `json:"transactionId"`
	ServiceProviderId primitive.ObjectID `json:"serviceProviderId"` // id nguoi cung cap dich vu
}

func (it *SendInfoServiceReq) Validate() error {

	if len(it.TransactionId) == 0 {
		return errors.New("transactionId empty")
	}
	if len(it.ServiceProviderId) == 0 {
		return errors.New("serviceProviderId empty")
	}

	return nil
}

type SendInfoServiceResp struct {
	BasicResp
}

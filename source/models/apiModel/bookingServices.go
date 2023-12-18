package apiModel

import (
	"errors"

	"example.com/m/v2/source/models/mongodbModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create Family Services
type CreateFamilyServicesReq struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (it *CreateFamilyServicesReq) Validate() error {

	if len(it.Code) == 0 {
		return errors.New("code empty")
	}
	if len(it.Name) == 0 {
		return errors.New("name empty")
	}

	return nil
}

type CreateFamilyServicesResp struct {
	BasicResp
	Id primitive.ObjectID `json:"id"`
}

// Get List Family Services
type GetListFamilyServicesResp struct {
	BasicResp
	Result []mongodbModel.FamilyServices `json:"result"`
}

// Create Account
type CreateAccountReq struct {
	Phone           string   `json:"phone"`
	Name            string   `json:"name"`
	Role            int      `json:"role"`
	ListServiceCode []string `json:"listServiceCode"`
}

func (it *CreateAccountReq) Validate() error {
	err := CheckPhoneNumber(it.Phone)
	if err != nil {
		return err
	}

	if len(it.Name) == 0 {
		return errors.New("name empty")
	}

	if it.Role == 2 {
		if len(it.ListServiceCode) == 0 {
			return errors.New("listServiceCode invalid")
		}
	} else {
		if len(it.ListServiceCode) > 0 {
			return errors.New("listServiceCode invalid")
		}
	}

	return nil
}

type CreateAccountResp struct {
	BasicResp
	Id primitive.ObjectID `json:"id"`
}

// /
// Create Family Services
type BookingReq struct {
	ServiceCode   string             `json:"serviceCode"`
	Price         float64            `json:"price"`
	CustomerId    primitive.ObjectID `json:"customerId"`
	TransactionId primitive.ObjectID `json:"transactionId"` // khách hàng gửi booking lại với transID
}

func (it *BookingReq) Validate() error {

	if len(it.ServiceCode) == 0 {
		return errors.New("serviceCode empty")
	}
	if len(it.CustomerId) == 0 {
		return errors.New("customerId empty")
	}

	return nil
}

type BookingResp struct {
	BasicResp
	IsNewTransID      bool               `json:"isNewTransID"`
	IsFind            bool               `json:"isFind"`
	TransactionId     primitive.ObjectID `json:"transactionId"`
	ServiceProviderId primitive.ObjectID `json:"serviceProviderId"` // id nguoi cung cap dich vu
}

// Check Status
type CheckStatusBookingResp struct {
	BasicResp
	Status string `json:"status"`
}

// Update Status
type UpdateStatusBookingReq struct {
	TransactionId primitive.ObjectID `json:"transactionId"`
	Status        string             `json:"status"`
}

func (it *UpdateStatusBookingReq) Validate() error {

	if len(it.TransactionId) == 0 {
		return errors.New("transactionId empty")
	}
	if it.Status != string(mongodbModel.Start) && it.Status != string(mongodbModel.NoFindServiceProvider) &&
		it.Status != string(mongodbModel.Processing) && it.Status != string(mongodbModel.Success) && it.Status != string(mongodbModel.Fail) {
		return errors.New("status invalid")
	}

	return nil
}

type UpdateStatusBookingResp struct {
	BasicResp
}

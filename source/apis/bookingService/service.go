package bookingService

import (
	"context"
	"time"

	"example.com/m/v2/source/gateway"
	"example.com/m/v2/source/models"
	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/models/mongodbModel"
	"example.com/m/v2/source/mongodb"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	database  *mongodb.Store
	gwService gateway.Service
}

type Service interface {
	CreateFamilyServices(ctx context.Context, req apiModel.CreateFamilyServicesReq) apiModel.CreateFamilyServicesResp
	GetListFamilyServices(ctx context.Context) apiModel.GetListFamilyServicesResp
	CreateAccount(ctx context.Context, req apiModel.CreateAccountReq) apiModel.CreateAccountResp

	Booking(ctx context.Context, req apiModel.BookingReq) apiModel.BookingResp
	CheckStatusBooking(ctx context.Context, transactionID primitive.ObjectID) apiModel.CheckStatusBookingResp
	UpdateStatusBooking(ctx context.Context, req apiModel.UpdateStatusBookingReq) apiModel.UpdateStatusBookingResp
}

func newService(mongodb_service *mongodb.Store, gwService gateway.Service) Service {
	return &service{
		database:  mongodb_service,
		gwService: gwService,
	}
}

func (ins *service) CreateFamilyServices(ctx context.Context, req apiModel.CreateFamilyServicesReq) apiModel.CreateFamilyServicesResp {
	//set params
	dataResp := apiModel.CreateFamilyServicesResp{}
	action := "bookingService.CreateFamilyServices"
	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}
	//Check Code exist
	_, haveData, err := ins.database.FamilyServices.GetByCode(ctx, req.Code)
	if err != nil && err != mongo.ErrNoDocuments {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	if haveData {
		dataResp.BasicResp.Update(3, action, int(models.ErrorRequestDataInvalid), errors.New("Code Exist"))
		return dataResp
	}

	// insert data
	familyServiceModel := &mongodbModel.FamilyServices{
		Code:       req.Code,
		Name:       req.Name,
		IsRemoved:  false,
		Popularity: 0,
		Price:      req.Price,
		LastUpdate: time.Now(),
		CreateAt:   time.Now(),
	}
	id, err := ins.database.FamilyServices.Insert(ctx, familyServiceModel)
	if err != nil {
		dataResp.BasicResp.Update(4, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}
	dataResp.Id = id
	return dataResp
}

func (ins *service) GetListFamilyServices(ctx context.Context) apiModel.GetListFamilyServicesResp {
	//set params
	dataResp := apiModel.GetListFamilyServicesResp{}
	action := "bookingService.GetListFamilyServices"

	list, err := ins.database.FamilyServices.GetAll(ctx)
	if err != nil {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}
	//note: tạm thời ảnh lấy base64 string , nếu chính xác là để URL
	dataResp.Result = list
	return dataResp
}

func (ins *service) CreateAccount(ctx context.Context, req apiModel.CreateAccountReq) apiModel.CreateAccountResp {
	//set params
	dataResp := apiModel.CreateAccountResp{}
	action := "bookingService.CreateAccount"
	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}
	//Check Code exist
	_, haveData, err := ins.database.Account.GetByPhone(ctx, req.Phone)
	if err != nil && err != mongo.ErrNoDocuments {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	if haveData {
		dataResp.BasicResp.Update(3, action, int(models.ErrorRequestDataInvalid), errors.New("Phone Exist"))
		return dataResp
	}

	// insert data
	accountModel := &mongodbModel.Account{
		Phone:           req.Phone,
		Name:            req.Name,
		IsRemoved:       false,
		Role:            req.Role,
		ListServiceCode: req.ListServiceCode,
		LastUpdate:      time.Now(),
		CreateAt:        time.Now(),
	}
	id, err := ins.database.Account.Insert(ctx, accountModel)
	if err != nil {
		dataResp.BasicResp.Update(4, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}
	dataResp.Id = id
	return dataResp
}

///Booking

func (ins *service) Booking(ctx context.Context, req apiModel.BookingReq) apiModel.BookingResp {
	//set params
	dataResp := apiModel.BookingResp{}
	var serviceProviderId primitive.ObjectID
	var transactionId primitive.ObjectID
	var isNewTransId bool
	action := "bookingService.Booking"
	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}

	err := ins.CheckValidateBooking(ctx, req.ServiceCode, req.CustomerId)
	if err != nil {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	//Find Service Provider
	isFind, serviceProviderId, err := ins.FindServiceProvider(ctx, req.ServiceCode)
	if err != nil {
		dataResp.BasicResp.Update(3, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	if !isFind {
		return dataResp
	}

	//new TransID
	if req.TransactionId.IsZero() {
		// create new transactionId
		transDb := &mongodbModel.Transaction{
			ServiceCode: req.ServiceCode,
			CustomerId:  req.CustomerId,
			Status:      mongodbModel.Start,
			LastUpdate:  time.Now(),
			CreateAt:    time.Now(),
		}
		id, err := ins.database.Transaction.Insert(ctx, transDb)
		if err != nil {
			dataResp.BasicResp.Update(4, action, int(models.ErrorDatabase), errors.WithStack(err))
			return dataResp
		}
		transactionId = id
		isNewTransId = true
	} else {
		transactionId = req.TransactionId
	}

	//Send Info
	err = ins.SendInfoService(ctx, transactionId, serviceProviderId)
	if err != nil {
		dataResp.BasicResp.Update(4, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	//send success
	dataResp.IsNewTransID = isNewTransId
	dataResp.IsFind = true
	dataResp.TransactionId = transactionId
	dataResp.ServiceProviderId = serviceProviderId

	return dataResp
}

// Check Status Booking
func (ins *service) CheckStatusBooking(ctx context.Context, transactionID primitive.ObjectID) apiModel.CheckStatusBookingResp {
	//set params
	dataResp := apiModel.CheckStatusBookingResp{}

	action := "bookingService.CheckStatusBooking"

	transInfo, _, err := ins.database.Transaction.GetByTransactionId(ctx, transactionID)
	if err != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	dataResp.Status = string(transInfo.Status)
	return dataResp
}

// Update Status
func (ins *service) UpdateStatusBooking(ctx context.Context, req apiModel.UpdateStatusBookingReq) apiModel.UpdateStatusBookingResp {
	//set params
	dataResp := apiModel.UpdateStatusBookingResp{}

	action := "bookingService.CheckStatusBooking"

	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}

	_, _, err := ins.database.Transaction.GetByTransactionId(ctx, req.TransactionId)
	if err != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	err = ins.database.Transaction.UpdateStatus(ctx, req.TransactionId, req.Status)
	if err != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	//reset status service prodvider
	if req.Status == string(mongodbModel.NoFindServiceProvider) || req.Status == string(mongodbModel.Fail) || req.Status == string(mongodbModel.Success) {

		//update status service provider
		err = ins.database.Account.UpdateStatus(ctx, req.TransactionId, string(mongodbModel.Busy))
		if err != nil {
			dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
			return dataResp
		}
	}

	return dataResp
}

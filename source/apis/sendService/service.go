package sendService

import (
	"context"

	"example.com/m/v2/source/models"
	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/models/mongodbModel"
	"example.com/m/v2/source/mongodb"
	"github.com/pkg/errors"
)

type service struct {
	database *mongodb.Store
}

type Service interface {
	SendInfoService(ctx context.Context, req apiModel.SendInfoServiceReq) apiModel.SendInfoServiceResp
}

func newService(mongodb_service *mongodb.Store) Service {
	return &service{
		database: mongodb_service,
	}
}

func (ins *service) SendInfoService(ctx context.Context, req apiModel.SendInfoServiceReq) apiModel.SendInfoServiceResp {
	//set params
	dataResp := apiModel.SendInfoServiceResp{}
	action := "sendService.SendInfoService"
	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}

	//update Status Booking
	err := ins.database.Transaction.UpdateIdServiceProvider(ctx, req.TransactionId, req.ServiceProviderId, string(mongodbModel.Processing))
	if err != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}
	//update status service provider
	err = ins.database.Account.UpdateStatus(ctx, req.ServiceProviderId, string(mongodbModel.Busy))
	if err != nil {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}
	return dataResp
}

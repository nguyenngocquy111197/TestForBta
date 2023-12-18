package pricingService

import (
	"context"

	"example.com/m/v2/source/models"
	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/mongodb"
	"github.com/pkg/errors"
)

type service struct {
	database *mongodb.Store
}

type Service interface {
	Calculate(ctx context.Context, req apiModel.CalculateReq) apiModel.CalculateResp
}

func newService(mongodb_service *mongodb.Store) Service {
	return &service{
		database: mongodb_service,
	}
}

func (ins *service) Calculate(ctx context.Context, req apiModel.CalculateReq) apiModel.CalculateResp {
	//set params
	dataResp := apiModel.CalculateResp{}
	action := "sendService.Calculate"
	//check validate
	if newErr := req.Validate(); newErr != nil {
		dataResp.BasicResp.Update(1, action, int(models.ErrorRequestDataInvalid), errors.WithStack(newErr))
		return dataResp
	}

	//check serviceCOde
	_, _, err := ins.database.FamilyServices.GetByCode(ctx, req.ServiceCode)
	if err != nil {
		dataResp.BasicResp.Update(2, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	//check customerId
	_, _, err = ins.database.Account.GetById(ctx, req.CustomerId)
	if err != nil {
		dataResp.BasicResp.Update(3, action, int(models.ErrorDatabase), errors.WithStack(err))
		return dataResp
	}

	//check calculate
	price, err := ins.calculate(ctx, req)
	if err != nil {
		dataResp.BasicResp.Update(4, action, int(models.ErrorOther), errors.WithStack(err))
		return dataResp
	}

	dataResp.Price = price
	return dataResp
}

package bookingService

import (
	"context"
	"errors"

	"example.com/m/v2/source/models/gatewayModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ins *service) CheckValidateBooking(ctx context.Context, serviceCode string, customerId primitive.ObjectID) error {

	//Check Code Service not exist
	_, _, err := ins.database.FamilyServices.GetByCode(ctx, serviceCode)
	if err != nil {
		return err
	}

	//Check CustomerId not exist or Removed
	_, _, err = ins.database.Account.GetById(ctx, customerId)
	if err != nil {
		return err
	}
	return nil
}

func (ins *service) FindServiceProvider(ctx context.Context, serviceCode string) (bool, primitive.ObjectID, error) {
	var (
		serviceProviderId primitive.ObjectID
		isFind            bool
	)
	listServiceProvider, err := ins.database.Account.GetListServiceProvider(ctx)
	if err != nil && err != mongo.ErrNoDocuments {
		return isFind, serviceProviderId, err
	}
	// not find
	if len(listServiceProvider) == 0 {
		return isFind, serviceProviderId, nil
	} else {
		for _, v := range listServiceProvider {
			if len(v.ListServiceCode) > 0 {
				for _, v1 := range v.ListServiceCode {
					if v1 == serviceCode {
						serviceProviderId = v.ID
						isFind = true
						break
					}
				}
			}

		}
	}

	return isFind, serviceProviderId, nil
}

func (ins *service) SendInfoService(ctx context.Context, transactionID, serviceProviderId primitive.ObjectID) error {

	req := &gatewayModel.DTOSendInfoServiceReq{
		TransactionId:     transactionID,
		ServiceProviderId: serviceProviderId,
	}
	resp, err := ins.gwService.Send.SendInfo(req)
	if err != nil {
		return err
	}

	if resp != nil {
		if resp.Code != 0 {
			return errors.New(resp.Message)
		}
	}

	return nil
}

package pricingService

import (
	"context"

	"example.com/m/v2/source/models/apiModel"
)

func (ins service) calculate(ctx context.Context, req apiModel.CalculateReq) (float64, error) {

	day := req.Date.Day()
	var sum float64
	svInfo, _, err := ins.database.FamilyServices.GetByCode(ctx, req.ServiceCode)
	if err != nil {
		return 0, err
	}

	if day%2 == 0 {
		sum = svInfo.Price + 40000
	} else {
		sum = svInfo.Price + 20000
	}

	return sum, nil
}

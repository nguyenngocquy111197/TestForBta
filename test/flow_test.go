package test

import (
	"context"
	"testing"
	"time"

	"example.com/m/v2/source/apis/pricingService"
	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/models/mongodbModel"
	"example.com/m/v2/source/mongodb"
	"github.com/sirupsen/logrus"
)

func Test_Calculate(t *testing.T) {
	ctx := context.Background()
	var (
		log *logrus.Logger = logrus.New()
	)
	//Init mongodb

	dbstore := mongodb.New("mongodb://127.0.0.1:27017/bTaskee1", "bTaskeeTest", 30)
	price := pricingService.NewServiceAsLogger(logrus.NewEntry(log), dbstore)
	//insert CustomerId
	accountModel := &mongodbModel.Account{
		Phone:           "0000000000",
		Name:            "test",
		IsRemoved:       false,
		Role:            1,
		ListServiceCode: nil,
		LastUpdate:      time.Now(),
		CreateAt:        time.Now(),
	}
	id, err := dbstore.Account.Insert(ctx, accountModel)
	if err != nil {
		t.Fatal(err)
	}
	//insert serviceCode
	var priceServiceCode float64 = 30000
	serviceModel := &mongodbModel.FamilyServices{
		Code:       "1",
		Name:       "test",
		IsRemoved:  false,
		Price:      priceServiceCode,
		LastUpdate: time.Now(),
		CreateAt:   time.Now(),
	}
	_, err = dbstore.FamilyServices.Insert(ctx, serviceModel)
	if err != nil {
		t.Fatal(err)
	}
	// serviceCode = 1
	// time
	dateCacl := time.Now()
	if dateCacl.Day()%2 != 0 {
		dateCacl.AddDate(0, 0, 1)
	}
	req := apiModel.CalculateReq{
		ServiceCode: "1",
		CustomerId:  id,
		Date:        dateCacl,
	}

	resp := price.Calculate(ctx, req)
	if resp.Price != priceServiceCode+40000 {
		t.Fatal("fail")
	}

	//delete
}

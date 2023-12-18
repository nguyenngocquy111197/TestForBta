package pricingService

import (
	"context"
	"fmt"
	"time"

	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/mongodb"
	"github.com/sirupsen/logrus"
)

type logger struct {
	log  *logrus.Entry
	next Service
}

/*
NewServiceAsLogger create a new Service
*/
func NewServiceAsLogger(entry *logrus.Entry, store *mongodb.Store) Service {
	return &logger{
		log:  entry,
		next: newService(store),
	}
}

func (ins *logger) Calculate(ctx context.Context, req apiModel.CalculateReq) apiModel.CalculateResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.Calculate(ctx, req)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}

	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", req),
	}).Debugf("Log-Calculate-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-Calculate-Response")

	return response
}

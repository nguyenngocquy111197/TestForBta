package bookingService

import (
	"context"
	"fmt"
	"time"

	"example.com/m/v2/source/gateway"
	"example.com/m/v2/source/models/apiModel"
	"example.com/m/v2/source/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type logger struct {
	log  *logrus.Entry
	next Service
}

/*
NewServiceAsLogger create a new Service
*/
func NewServiceAsLogger(entry *logrus.Entry, store *mongodb.Store, gwService gateway.Service) Service {
	return &logger{
		log:  entry,
		next: newService(store, gwService),
	}
}

func (ins *logger) CreateFamilyServices(ctx context.Context, req apiModel.CreateFamilyServicesReq) apiModel.CreateFamilyServicesResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.CreateFamilyServices(ctx, req)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}

	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", req),
	}).Debugf("Log-CreateFamilyServices-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-CreateFamilyServices-Response")

	return response
}

func (ins *logger) GetListFamilyServices(ctx context.Context) apiModel.GetListFamilyServicesResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.GetListFamilyServices(ctx)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}

	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-Response-GetListFamilyServices")

	return response
}

func (ins *logger) CreateAccount(ctx context.Context, req apiModel.CreateAccountReq) apiModel.CreateAccountResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.CreateAccount(ctx, req)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}

	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", req),
	}).Debugf("Log-CreateAccount-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-CreateAccount-Response:")

	return response
}

func (ins *logger) Booking(ctx context.Context, req apiModel.BookingReq) apiModel.BookingResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.Booking(ctx, req)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}

	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", req),
	}).Debugf("Log-Booking-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-Booking-Response:")

	return response
}

func (ins *logger) CheckStatusBooking(ctx context.Context, transactionID primitive.ObjectID) apiModel.CheckStatusBookingResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.CheckStatusBooking(ctx, transactionID)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}
	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", transactionID),
	}).Debugf("Log-CheckStatusBooking-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-Response-CheckStatusBooking")

	return response
}

func (ins *logger) UpdateStatusBooking(ctx context.Context, req apiModel.UpdateStatusBookingReq) apiModel.UpdateStatusBookingResp {

	var (
		begin  = time.Now()
		metric float64
	)
	response := ins.next.UpdateStatusBooking(ctx, req)
	{
		metric = float64(time.Since(begin).Nanoseconds()) / float64(time.Millisecond)
	}
	ins.log.WithFields(logrus.Fields{
		"json": fmt.Sprintf("%+v", req),
	}).Debugf("Log-UpdateStatusBooking-Request:")
	ins.log.WithFields(logrus.Fields{
		"json":                fmt.Sprintf("%+v", response),
		"process_time (ms): ": fmt.Sprintf("%f", metric),
	}).Debugf("Log-Response-UpdateStatusBooking")

	return response
}

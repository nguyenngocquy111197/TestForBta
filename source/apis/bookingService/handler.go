package bookingService

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"example.com/m/v2/source/models"
	"example.com/m/v2/source/models/apiModel"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	listRole = []apiModel.Role{}
)

type handler struct {
	log     *logrus.Entry
	service Service
}

type Handler interface {
	AddRoutes(r *gin.Engine)
}

func NewHandler(service Service) Handler {

	{
		data, err := ioutil.ReadFile("./role.json")
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		} else if err = json.Unmarshal(data, &listRole); err != nil {
			panic(err)
		}
	}

	return &handler{
		service: service,
	}
}

func (ins *handler) AddRoutes(r *gin.Engine) {
	var familyService = r.Group("/familyService")

	familyService.Handle(http.MethodPost, "/insert", ins.createFamilyServices)
	familyService.Handle(http.MethodGet, "/get", ins.getListFamilyServices)

	//account
	var account = r.Group("/account")
	account.Handle(http.MethodGet, "/listRole", ins.getListRole)
	account.Handle(http.MethodPost, "/create", ins.createAccount)

	//booking
	r.Handle(http.MethodPost, "/booking", ins.booking)

	var booking = r.Group("/booking")
	booking.Handle(http.MethodGet, "/check/status", ins.checkStatusBooking)
	booking.Handle(http.MethodPost, "update/status", ins.updateStatus)
}

func (ins *handler) createFamilyServices(c *gin.Context) {
	var (
		request    = &apiModel.CreateFamilyServicesReq{}
		response   = apiModel.CreateFamilyServicesResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.createFamilyServices"
	if err := c.BindJSON(request); err != nil {
		response.BasicResp.Update(-1, action, int(models.ErrorJsonMarshal), err)
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.CreateFamilyServices(r.Context(), *request)
	}

	c.JSON(statusCode, response)

}

func (ins *handler) getListFamilyServices(c *gin.Context) {
	var (
		response   = apiModel.GetListFamilyServicesResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)
	response = ins.service.GetListFamilyServices(r.Context())

	c.JSON(statusCode, response)

}

func (ins *handler) getListRole(c *gin.Context) {
	var (
		statusCode = http.StatusOK
	)
	respOK := map[string]interface{}{
		"code":     0,
		"message":  "",
		"listRole": listRole,
	}

	c.JSON(statusCode, respOK)

}

func (ins *handler) createAccount(c *gin.Context) {
	var (
		request    = &apiModel.CreateAccountReq{}
		response   = apiModel.CreateAccountResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.createAccount"
	if err := c.BindJSON(request); err != nil {
		response.BasicResp.Update(-1, action, int(models.ErrorJsonMarshal), err)
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.CreateAccount(r.Context(), *request)
	}

	c.JSON(statusCode, response)

}

func (ins *handler) booking(c *gin.Context) {
	var (
		request    = &apiModel.BookingReq{}
		response   = apiModel.BookingResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.booking"
	if err := c.BindJSON(request); err != nil {
		response.BasicResp.Update(-1, action, int(models.ErrorJsonMarshal), err)
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.Booking(r.Context(), *request)
	}

	c.JSON(statusCode, response)

}

func (ins *handler) checkStatusBooking(c *gin.Context) {
	var (
		transID, _ = primitive.ObjectIDFromHex(c.DefaultQuery("transID", primitive.NilObjectID.Hex()))
		response   = apiModel.CheckStatusBookingResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.checkStatusBooking"
	if transID.IsZero() {
		response.BasicResp.Update(-1, action, int(models.ErrorRequestDataInvalid), errors.New("text"))
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.CheckStatusBooking(r.Context(), transID)
	}

	c.JSON(statusCode, response)

}

func (ins *handler) updateStatus(c *gin.Context) {
	var (
		request    = &apiModel.UpdateStatusBookingReq{}
		response   = apiModel.UpdateStatusBookingResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.updateStatus"
	if err := c.BindJSON(request); err != nil {
		response.BasicResp.Update(-1, action, int(models.ErrorJsonMarshal), err)
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.UpdateStatusBooking(r.Context(), *request)
	}

	c.JSON(statusCode, response)

}

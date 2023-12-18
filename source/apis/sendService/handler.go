package sendService

import (
	"net/http"

	"example.com/m/v2/source/models"
	"example.com/m/v2/source/models/apiModel"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	log     *logrus.Entry
	service Service
}

type Handler interface {
	AddRoutes(r *gin.Engine)
}

func NewHandler(service Service) Handler {

	return &handler{
		service: service,
	}
}

func (ins *handler) AddRoutes(r *gin.Engine) {

	r.Handle(http.MethodPost, "/send/booking", ins.sendInfoService)

}

func (ins *handler) sendInfoService(c *gin.Context) {
	var (
		request    = &apiModel.SendInfoServiceReq{}
		response   = apiModel.SendInfoServiceResp{}
		r          = c.Request
		statusCode = http.StatusOK
	)

	action := "handler.sendInfoService"
	if err := c.BindJSON(request); err != nil {
		response.BasicResp.Update(-1, action, int(models.ErrorJsonMarshal), err)
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.SendInfoService(r.Context(), *request)
	}

	c.JSON(statusCode, response)

}

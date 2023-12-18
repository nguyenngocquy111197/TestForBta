package gateway

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/libs/net"
	"example.com/m/v2/source/models/gatewayModel"
	"github.com/sirupsen/logrus"
)

type send struct {
	host    string
	log     *logrus.Entry
	timeout time.Duration
}

type SendService interface {
	SendInfo(req *gatewayModel.DTOSendInfoServiceReq) (*gatewayModel.DTOSendInfoServiceResp, error)
}

func NewSend(host string, entry *logrus.Entry, timeout time.Duration) SendService {
	return &send{
		log:     entry,
		host:    host,
		timeout: timeout,
	}
}

func (ins *send) SendInfo(req *gatewayModel.DTOSendInfoServiceReq) (*gatewayModel.DTOSendInfoServiceResp, error) {
	var (
		action = "/send/booking"
		url    = ins.host + action
		opt    = net.CurlOption()
		resp   = &gatewayModel.DTOSendInfoServiceResp{}
	)

	opt.Timeout = ins.timeout
	opt.SetMethod(http.MethodPost)
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	opt.SetData(data)

	netResp, err := net.Curl(url, opt)
	if err != nil {
		return nil, err
	}

	log.Printf("Log-url: %s\r\n", url)
	log.Printf("Log-Status: %d\r\n", netResp.StatusCode)
	log.Printf("Log-Validate-Response:\r\n%s\r\n", bytes.NewBuffer(netResp.Body).String())

	if err = json.Unmarshal(netResp.Body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

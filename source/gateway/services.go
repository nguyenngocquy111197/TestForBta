package gateway

import (
	"time"

	"github.com/sirupsen/logrus"
)

type ProviderGateways struct {
	Send ProviderGateway `yaml:"send"`
}

type ProviderGateway struct {
	Timeout time.Duration `yaml:"timeout"`
	Host    string
}
type Service struct {
	Send SendService
}

func New(gateway ProviderGateways, entry *logrus.Entry) *Service {
	return &Service{
		Send: NewSend(gateway.Send.Host, entry, gateway.Send.Timeout),
	}
}

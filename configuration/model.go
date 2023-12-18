package configuration

import (
	"fmt"
	"time"

	"example.com/m/v2/source/gateway"
)

type domain struct {
	Bind string `json:"bind" yaml:"bind"`
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

func (ins *domain) ParseAddr() string {
	if len(ins.Bind) == 0 {
		ins.Bind = "127.0.0.1"
	}
	return fmt.Sprintf("%s:%s", ins.Bind, ins.Port)
}

func (ins *domain) ParseURL() string {
	if len(ins.Host) == 0 {
		ins.Host = "127.0.0.1"
	}
	if len(ins.Port) == 0 || ins.Port == "80" || ins.Port == "443" {
		return ins.Host
	}
	return fmt.Sprintf("%s:%s", ins.Host, ins.Port)
}

type mongoSetting struct {
	Username   string        `json:"username" yaml:"username"`
	URI        string        `json:"uri" yaml:"uri"`
	PWD        string        `json:"pwd" yaml:"pwd"`
	Domains    []domain      `json:"domains" yaml:"domains"`
	DBName     string        `json:"db_name" yaml:"db_name"`
	ReplicaSet string        `json:"replica_set" yaml:"replica_set"`
	Timeout    time.Duration `json:"timeout" yaml:"timeout"`
}

func (ins *mongoSetting) ParseURI() string {

	ins.URI = "mongodb://"

	if len(ins.Username) > 0 {
		var pwd string
		ins.URI += ins.Username + ":"
		if len(ins.PWD) > 0 {
			ins.URI += ins.PWD + "@"
		} else {
			println("\r\nEnter MONGO_PWD:")
			fmt.Scanf("%s", &pwd)
			ins.PWD = pwd
			ins.URI += ins.PWD + "@"
		}
	}

	for i, domain := range ins.Domains {
		if len(domain.Port) == 0 {
			domain.Port = "27017"
		}
		if i != 0 {
			ins.URI += ","
		}
		ins.URI += domain.Host + ":" + domain.Port
	}

	ins.URI += "/" + ins.DBName

	if len(ins.ReplicaSet) > 0 {
		ins.URI += "?replicaSet=" + ins.ReplicaSet
	}

	return ins.URI
}

// Setting object
type Setting struct {
	HostServer       domain                   `json:"host_server" yaml:"host_server"`
	ProviderGateways gateway.ProviderGateways `yaml:"provider_gateways"`
	MongoDB          mongoSetting             `json:"mongodb" yaml:"mongodb"`
	LogToFile        string                   `json:"log_to_file" yaml:"log_to_file"`
	LogGatewayToFile string                   `json:"log_gateway_to_file" yaml:"log_gateway_to_file"`
}

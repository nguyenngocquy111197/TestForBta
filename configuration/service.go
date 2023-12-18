package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v1"
)

type service struct {
	setting Setting

	config string
	log    *logrus.Entry
}

// Service repository
type Service interface {
	Init() *Setting
}

// NewService create a new repository
func NewService(root string) Service {

	var (
		absRoot string
		config  string

		newLogger *logrus.Logger
	)

	// Get Abs path
	{
		if len(root) == 0 {
			root = "./"
		}
		aR, err := filepath.Abs(root)
		if err != nil {
			panic(err)
		} else {
			absRoot = aR
		}
	}

	{
		configFile := os.Getenv("CONFIG")
		flag.StringVar(&config, "config", absRoot+"/config/"+configFile, "setting config path")
		flag.StringVar(&config, "c", absRoot+"/config/"+configFile, "setting config path")

		if flag.Parsed() {
			println(">> flag parsed is: ", flag.Parsed())
		} else {
			flag.Parse()
		}
	}
	println("[NorPath] config: ", config)

	// Get Abs path
	{
		absConfig, err := filepath.Abs(config)
		if err != nil {
			panic(err)
		} else {
			config = absConfig
		}
	}
	println("[AbsPath] config: ", config)

	{
		newLogger = logrus.New()
	}

	return &service{
		config: config,
		log:    logrus.NewEntry(newLogger),
	}
}

func (ins *service) Init() *Setting {
	if _, err := os.Stat(ins.config); os.IsNotExist(err) {
		println("\r\nThe file was not found.\r\nThe system will create this file with default settings.\r\nPlease edit config file and restart service to update the new values!\r\n")
		println("Dir: ", filepath.Dir(ins.config))
		println("Base: ", filepath.Base(ins.config))
		println("Ext: ", filepath.Ext(ins.config))

		if err := os.MkdirAll(filepath.Dir(ins.config), 0755); err != nil {
			ins.log.Fatal(err)
		}

		set := Setting{
			HostServer: domain{
				"0.0.0.0", "127.0.0.1", "8080",
			},
			MongoDB: mongoSetting{
				Username: "",
				PWD:      "",
				Domains: []domain{
					{"0.0.0.0", "127.0.0.1", "27017"},
				},
				DBName:     "bTaskee",
				ReplicaSet: "",
				Timeout:    30 * time.Second,
			},
		}

		switch filepath.Ext(ins.config) {
		case "json", ".json":
			body, err := json.Marshal(&set)
			if err != nil {
				ins.log.Fatal(err)
			}
			err = ioutil.WriteFile(ins.config, body, 0664)
			if err != nil {
				ins.log.Fatal(err)
			}
		case "yaml", ".yaml":
			body, err := yaml.Marshal(&set)
			if err != nil {
				ins.log.Fatal(err)
			}
			err = ioutil.WriteFile(ins.config, body, 0664)
			if err != nil {
				ins.log.Fatal(err)
			}
		default:
			ins.log.Fatal("config file invalid format!")
		}
	}

	{

		body, err := ioutil.ReadFile(ins.config)
		if err != nil {
			ins.log.Fatal(err)
		}
		switch filepath.Ext(ins.config) {
		case "json", ".json":
			if err := json.Unmarshal(body, &ins.setting); err != nil {
				ins.log.Fatal(err)
			}
		case "yaml", ".yaml":
			if err := yaml.Unmarshal(body, &ins.setting); err != nil {
				ins.log.Fatal(err)
			}
		default:
			ins.log.Fatal("config file invalid format!")
		}

		// ParseURI
		// ins.setting.MongoDB.ParseURI()
	}

	println(fmt.Sprintf("\r\n>> setting : \r\n%+v\r\n", ins.setting))

	return &ins.setting
}

func (ins *domain) ToString() string {
	var (
		host = ins.Host
		port = ins.Port
	)
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if !strings.HasPrefix(host, "http") && !strings.HasPrefix(host, "ws") {
		host = "http://" + host
	}
	if len(port) == 0 || port == "80" || port == "443" {
		return host
	}
	return fmt.Sprintf("%s:%s", host, port)
}

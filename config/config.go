package config

import (
	"os"

	"github.com/sirupsen/logrus"

	"field-service/common/util"
)

var Config AppConfig

type DatabaseConfig struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

type InternalService struct {
	User User `json:"user"`
}

type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type GCSConfig struct {
	Type                    string `json:"gscType"`
	ProjectID               string `json:"gcsProjectID"`
	PrivateKeyID            string `json:"gcsPrivateKeyID"`
	PrivateKey              string `json:"gcsPrivateKey"`
	ClientEmail             string `json:"gcsClientEmail"`
	ClientID                string `json:"gcsClientID"`
	AuthURI                 string `json:"gcsAuthUri"`
	TokenURI                string `json:"gcsTokenUri"`
	AuthProviderX509CertURL string `json:"gcsAuthProviderX509CertUrl"`
	ClientX509CertURL       string `json:"gcsClientX509CertUrl"`
	UniverseDomain          string `json:"gcsUniverseDomain"`
	BucketName              string `json:"gcsBucketName"`
}

type AppConfig struct {
	Port            int             `json:"port"`
	AppName         string          `json:"appName"`
	AppEnv          string          `json:"appEnv"`
	SignatureKey    string          `json:"siginatureKey"`
	Database        DatabaseConfig  `json:"database"`
	InternalService InternalService `json:"internalService"`
	GCS             GCSConfig       `json:"gcs"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", "json", ".")
	if err != nil {
		logrus.Infof("Failed to bind from json : %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}

	}
}

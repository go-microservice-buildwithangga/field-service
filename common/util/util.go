package util

import (
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJSON(dest any, fileName string, configType string, configDir string) error {
	v := viper.New()
	ext := filepath.Ext(fileName)                     // mengambil ekstensi file
	baseFileName := fileName[:len(fileName)-len(ext)] // menghilangkan ekstensi file
	v.SetConfigType(configType)                       // set config type dengan parameter
	v.SetConfigName(baseFileName)                     //set nama file tanpa ekstensi
	v.AddConfigPath(configDir)                        //set lokasi path file
	if err := v.ReadInConfig(); err != nil {
		logrus.Errorf("Failed to read config file : %v", err)
		return err
	}
	if err := v.Unmarshal(&dest); err != nil {
		logrus.Errorf("Failed to unmarshal : %v", err)
		return err
	}
	return nil
}

func SetEnvFromConsulKV(v *viper.Viper) error {

	env := make(map[string]any)
	if err := v.Unmarshal(&env); err != nil {
		logrus.Errorf("Failed to unmarshal : %v", err)
		return err
	}
	for key, value := range env {
		var valOf = reflect.ValueOf(value)
		var val string
		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valOf.Int()))
		case reflect.Uint:
			val = strconv.Itoa(int(valOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			panic("Unsupported type")
		}

		err := os.Setenv(key, val)
		if err != nil {
			logrus.Errorf("Failed to set env : %v", err)
			return err
		}
	}
	return nil

}

func BindFromConsul(dest any, endPoint string, key string) error {
	v := viper.New()
	v.SetConfigType("json")
	err := v.AddRemoteProvider("consul", endPoint, key)
	if err != nil {
		logrus.Errorf("Failed to add remote provider : %v", err)
		return err
	}

	err = v.ReadRemoteConfig()
	if err != nil {
		logrus.Errorf("Failed to read remote config : %v", err)
		return err
	}

	if err := v.Unmarshal(&dest); err != nil {
		logrus.Errorf("Failed to unmarshal : %v", err)
		return err
	}
	return nil
}

package util

import (
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PaginationParam struct {
	Count int64       `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data"`
}

type PaginationResult struct {
	TotalPage int         `json:"totalPage"`
	TotalData int64       `json:"totalData"`
	NextPage  *int        `json:"nextPage"`
	PrevPage  *int        `json:"prevPage"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Data      interface{} `json:"data"`
}

func GeneratePagination(params PaginationParam) PaginationResult {
	totalPage := int(math.Ceil(float64(params.Count) / float64(params.Limit)))
	var (
		nextPage int
		prevPage int
	)
	if params.Page < totalPage {
		nextPage = params.Page + 1
	}
	if params.Page > 1 {
		prevPage = params.Page - 1
	}

	result := PaginationResult{
		TotalPage: totalPage,
		TotalData: params.Count,
		NextPage:  &nextPage,
		PrevPage:  &prevPage,
		Page:      params.Page,
		Limit:     params.Limit,
		Data:      params.Data,
	}
	return result

}

func GenerateSHA256(inputString string) string {
	hash := sha256.New()
	hash.Write([]byte(inputString))
	hasBytes := hash.Sum(nil)
	hashString := string(hasBytes)
	return hashString

}

func RupiahFormat(amount *float64) string {
	stringValue := "0"
	if amount != nil {
		humanizeValue := humanize.CommafWithDigits(*amount, 0)
		stringValue = strings.ReplaceAll(humanizeValue, ",", ".")

	}
	return fmt.Sprintf("Rp.%s", stringValue)
}

func BindFromJSON(dest any, filename, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("failed to unmarshal: %v", err)
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

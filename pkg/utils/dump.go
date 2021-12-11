package utils

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func Dump(i interface{}) string {
	bt, err := json.Marshal(i)
	if err != nil {
		logrus.Error(err)
	}

	return string(bt)
}

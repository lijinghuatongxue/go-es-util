package main

import (
	"github.com/lijinghuatongxue/go-test/es/OverallInfo"
	"github.com/sirupsen/logrus"
)

func main() {
	if OverallInfo.GetESOverallInfo("192.168.0.253", "9200", "status") != true {
		logrus.Error("err")
	}
}

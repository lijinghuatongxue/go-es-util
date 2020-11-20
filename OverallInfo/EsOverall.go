package OverallInfo

import (
	"encoding/json"
	SendHttp "github.com/lijinghuatongxue/go-test/es/http"
	"github.com/sirupsen/logrus"
)

func GetESOverallInfo(ES_HOST, Port, OverallType string) bool {
	type ESInfoStruct struct {
		Status       string
		Cluster_name string
		Cluster_uuid string
	}
	//logrus.Info(body)
	//if err != nil {
	//	logrus.Println(err)
	//	return
	//}
	//logrus.Println(string(body))
	var getESInfo ESInfoStruct
	json.Unmarshal(SendHttp.SendHttp(ES_HOST, Port), &getESInfo)
	logrus.Infof("\n ES-Addr == > %s \n ES-Status == > %s \n ES-name == > %s \n ES-uuid == > %s", ES_HOST+":"+Port, getESInfo.Status, getESInfo.Cluster_name, getESInfo.Cluster_uuid)
	return true
	//logrus.Warn(getagentcpuinfo.Node.Successful)
}

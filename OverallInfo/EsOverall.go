package OverallInfo

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetESOverallInfo(ES_HOST, Port, OverallType string) bool {
	type ESInfoStruct struct {
		Status       string
		Cluster_name string
		Cluster_uuid string
	}
	ip := ES_HOST
	URL := "/_cluster/stats?pretty"

	url := "http://" + ip + ":" + Port + URL
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logrus.Println(err)
		return false
	}
	res, err := client.Do(req)
	if err != nil {
		logrus.Println(err)
		return false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	logrus.Println(err)
	//	return
	//}
	//logrus.Println(string(body))
	var getESInfo ESInfoStruct
	json.Unmarshal(body, &getESInfo)
	logrus.Infof("\n ES-Addr == > %s \n ES-Status == > %s \n ES-name == > %s \n ES-uuid == > %s", ES_HOST+":"+Port, getESInfo.Status, getESInfo.Cluster_name, getESInfo.Cluster_uuid)
	return true
	//logrus.Warn(getagentcpuinfo.Node.Successful)
}

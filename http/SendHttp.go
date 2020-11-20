package SendHttp

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func SendHttp(ES_HOST, Port string) []byte {
	ip := ES_HOST
	URL := "/_cluster/stats?pretty"

	url := "http://" + ip + ":" + Port + URL
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logrus.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		logrus.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body
}

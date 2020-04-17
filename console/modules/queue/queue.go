package queue

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"smsgate/console/modules/node"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

type Queue struct {
	Len       int
	BufferLen int
}

func GetQueueLen() map[string]map[string]Queue {
	result := make(map[string]map[string]Queue)

	nodes := node.Registry.GetNodes()
	for _, n := range nodes {
		result[n.GetUrl()] = make(map[string]Queue)
		url := "http://" + n.GetUrl() + "/queue/length"
		result[n.GetUrl()] = postRequest(url)
	}

	return result

}

func postRequest(queueUrl string) map[string]Queue {
	result := make(map[string]Queue)
	var serverResp struct {
		Code    int
		Message string
		Data    map[string]map[string]int
	}

	resp, err := client.PostForm(queueUrl, url.Values{})
	if err != nil {
		log.Printf("post request failed. err:%v", err)
		return result
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &serverResp)
	if err != nil {
		log.Printf("decode json failed. err:%v", err)
		return result
	}

	data := serverResp.Data
	for k, v := range data {
		q := Queue{
			Len:       v["len"],
			BufferLen: v["buffer"],
		}
		result[k] = q
	}

	return result
}

package business

import (
	"github.com/parnurzeal/gorequest"
	"github.com/labstack/echo"
	"errors"
	"encoding/json"
	"sort"
	"gitlab.wallstcn.com/operation/nsqmonitor/helper"
	"flag"
	"os"
	"strings"
)

type Topics struct {
	Topics  []string `json:"topics"`
	Message string   `json:"message"`
}

type TopicInfo struct {
	Node                   string         `json:"node"`
	Hostname               string         `json:"hostname"`
	Topic_Name             string         `json:"topic_name"`
	Depth                  int64          `json:"depth"`
	Memory_Depth           int64          `json:"memory_depth"`
	Backend_Depth          int64          `json:"backend_depth"`
	Message_Count          int64          `json:"message_count"`
	Nodes                  []NodesInfo    `json:"nodes"`
	Channels               []ChannelsInfo `json:"channels"`
	Paused                 bool           `json:"paused"`
	E2e_Processing_Latency interface{}    `json:"e2e_processing_latency"`
	Message                string         `json:"message"`
}

type NodesInfo struct {
	Node                   string         `json:"node"`
	Hostname               string         `json:"hostname"`
	Topic_Name             string         `json:"topic_name"`
	Depth                  int64          `json:"depth"`
	Memory_Depth           int64          `json:"memory_depth"`
	Backend_Depth          int64          `json:"backend_depth"`
	Message_Count          int64          `json:"message_count"`
	Nodes                  interface{}
	Channels               []ChannelsInfo `json:"channels"`
	Paused                 bool           `json:"paused"`
	E2e_Processing_Latency interface{}    `json:"e2e_processing_latency"`
}

type ChannelsInfo struct {
	Node                   string        `json:"node"`
	Hostname               string        `json:"hostname"`
	Topic_Name             string        `json:"topic_name"`
	Channel_Name           string        `json:"channel_name"`
	Depth                  int64         `json:"depth"`
	Memory_Depth           int64         `json:"memory_depth"`
	Backend_Depth          int64         `json:"backend_depth"`
	In_Flight_Count        int64         `json:"in_flight_count"`
	Deferred_Count         int64         `json:"deferred_count"`
	Requeue_Count          int64         `json:"requeue_count"`
	Timeout_Count          int64         `json:"timeout_count"`
	Message_Count          int64         `json:"message_count"`
	Nodes                  interface{}   `json:"nodes"`
	Clients                []ClientsInfo `json:"clients"`
	Paused                 bool          `json:"paused"`
	E2e_Processing_Latency interface{}   `json:"e2e_processing_latency"`
}

type ClientsInfo struct {
	Node                              string `json:"node"`
	Remote_Address                    string `json:"remote_address"`
	Name                              string `json:"name"`
	Version                           string `json:"version"`
	Client_Id                         string `json:"client_id"`
	Hostname                          string `json:"hostname"`
	User_Agent                        string `json:"user_agent"`
	Connect_Ts                        int    `json:"connect_ts"`
	Connected                         int64  `json:"connected"`
	In_Flight_Count                   int64  `json:"in_flight_count"`
	Ready_Count                       int64  `json:"ready_count"`
	Finish_Count                      int64  `json:"finish_count"`
	Requeue_Count                     int64  `json:"requeue_count"`
	Message_Count                     int64  `json:"message_count"`
	Sample_Rate                       int64  `json:"sample_rate"`
	Deflate                           bool   `json:"deflate"`
	Snappy                            bool   `json:"snappy"`
	Authed                            bool   `json:"authed"`
	Auth_Identity                     string `json:"auth_identity"`
	Auth_Identity_Url                 string `json:"auth_identity_url"`
	Tls                               bool   `json:"tls"`
	Tls_Cipher_Suite                  string `json:"tls_cipher_suite"`
	Tls_Version                       string `json:"tls_version"`
	Tls_Negotiated_Protocol           string `json:"tls_negotiated_protocol"`
	Tls_Negotiated_Protocol_Is_Mutual bool   `json:"tls_negotiated_protocol_is_mutual"`
}

type Overview struct {
	Topic_Name         string `json:"topic_name"`
	Producer_Depth_Sum int64  `json:"producer_depth_sum"`
	Consumer_Depth_Sum int64  `json:"consumer_depth_sum"`
}
type Consumer struct {
	Topic_Name   string `json:"topic_name"`
	Channel_Name string `json:"channel_name"`
	Depth        int64  `json:"depth"`
	Clients      int    `json:"clients"`
	Ts           int    `json:"ts"`
}

var TopicsALl Topics
var OneTopicInfo TopicInfo
var url = flag.String("nsqadminurl", os.Getenv("NSQADMINURL"), "url")

func GetUrl() string{
	flag.Parse()
	url := strings.Split(*url, ",")
	nsqadmin_url := url[0]
	return nsqadmin_url
}
func GetMine() ([]*Overview, []*Consumer) {

	url := GetUrl()
	var OverviewList = make([]*Overview, 0, 500)
	var ConsumerList = make([]*Consumer, 0, 500)
	//get all topics
	request := gorequest.New()
	resp, body, errs := request.Get(url).End()

	if resp.StatusCode != 200 || len(errs) != 0 {
		newError := errors.New("topicsUrl ERROR")
		print(newError)
	}
	json.Unmarshal([]byte(body), &TopicsALl)
	for _, val := range TopicsALl.Topics {
		OverviewList, ConsumerList = GetTopicInfo(val, OverviewList, ConsumerList)
	}
	return OverviewList, ConsumerList

}

func GetTopicInfo(topicName string, OverviewList []*Overview, ConsumerList []*Consumer) ([]*Overview, []*Consumer) {

	url := GetUrl()
	request := gorequest.New()
	resp, body, errs := request.Get(url + "/" + topicName).End()
	if resp.StatusCode != 200 || len(errs) != 0 {
		newError := errors.New("topicsUrl ERROR")
		print(newError)
	}
	json.Unmarshal([]byte(body), &OneTopicInfo)
	var consumerDepthSum int64

	producerDepthSum := OneTopicInfo.Depth

	for _, val2 := range OneTopicInfo.Channels {

		ts := make([]int, 100, 10000)
		consumerDepthSum = consumerDepthSum + val2.Depth
		for key, val3 := range val2.Clients {
			ts[key] = val3.Connect_Ts
		}
		sort.Sort(sort.Reverse(sort.IntSlice(ts)))
		consumer := &Consumer{
			topicName,
			val2.Channel_Name,
			val2.Depth,
			len(val2.Clients),
			ts[0],
		}
		ConsumerList = append(ConsumerList, consumer)

	}
	overview := &Overview{
		topicName,
		producerDepthSum,
		consumerDepthSum,
	}
	OverviewList = append(OverviewList, overview)
	return OverviewList, ConsumerList
}

type Pagination struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}

// @Title overview  list
// @Description 获取overview list
// @Accept  json
// @Param page query int false "页数|默认1"
// @Param limit query int false "每页条目数|默认10"
// @Resource overview
// @Router /v1/overview [get]

func HTTPGetOverview(ctx echo.Context) error {
	OverviewList, _ := GetMine()
	return helper.SuccessResponse(ctx, &OverviewList)
}

// @Title consumer list by page and limit
// @Description 获取consumer list by page and limit
// @Accept  json
// @Param page query int false "页数|默认1"
// @Param limit query int false "每页条目数|默认10"
// @Resource consumer
// @Router /v1/consumer  [get]

func HTTPGetConsumer(ctx echo.Context) error {
	_, ConsumerList := GetMine()
	return helper.SuccessResponse(ctx, &ConsumerList)
}

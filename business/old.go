package business
//
//import (
//	"github.com/parnurzeal/gorequest"
//	"fmt"
//	"encoding/json"
//	"errors"
//)
//
////
////type Topics struct {
////	Topics  []string `json:"topics"`
////	Message string   `json:"message"`
////}
//
//type Stats struct {
//	Status_Code int64  `json:"status_code"`
//	Status_Txt  string `json:"status_txt"`
//	Data        Data   `json:"data"`
//}
//
//type Data struct {
//	Version    string       `json:"version"`
//	Health     string       `json:"health"`
//	Start_Time int64        `json:"start_time"`
//	Topics     []TopicsList `json:"topics"`
//}
//
//type TopicsList struct {
//	Topic_Name             string      `json:"topic_name"`
//	Channels               []Channels  `json:"channels"`
//	Depth                  int64       `json:"depth"`
//	Backend_Depth          int64       `json:"backend_depth"`
//	Message_Count          int64       `json:"message_count"`
//	Paused                 bool        `json:"paused"`
//	E2e_Processing_Latency interface{} `json:"e2e_processing_latency"`
//}
//
//type Channels struct {
//	Channel_name           string      `json:"channel_name"`
//	Depth                  int64       `json:"depth"`
//	Backend_Depth          int64       `json:"backend_depth"`
//	In_Flight_Count        int64       `json:"in_flight_count"`
//	Deferred_Count         int64       `json:"deferred_count"`
//	Message_Count          int64       `json:"message_count"`
//	Requeue_Count          int64       `json:"requeue_count"`
//	Timeout_Count          int64       `json:"timeout_count"`
//	Clients                []Clients   `json:"clients"`
//	Paused                 bool        `json:"paused"`
//	E2e_Processing_Latency interface{} `json:"e2e_processing_latency"`
//}
//
//type Clients struct {
//	Name                              string `json:"name"`
//	Client_Id                         string `json:"client_id"`
//	Hostname                          string `json:"hostname"`
//	Version                           string `json:"version"`
//	Remote_Address                    string `json:"remote_address"`
//	State                             int64  `json:"state"`
//	Ready_Count                       int64  `json:"ready_count"`
//	In_Flight_Count                   int64  `json:"in_flight_count"`
//	Message_Count                     int64  `json:"message_count"`
//	Finish_Count                      int64  `json:"finish_count"`
//	Requeue_Count                     int64  `json:"requeue_count"`
//	Connect_Ts                        int64  `json:"connect_ts"`
//	Sample_Rate                       int64  `json:"sample_rate"`
//	Deflate                           bool   `json:"deflate"`
//	Snappy                            bool   `json:"snappy"`
//	User_Agent                        string `json:"user_agent"`
//	Tls                               bool   `json:"tls"`
//	Tls_Cipher_Suite                  string `json:"tls_cipher_suite"`
//	Tls_Version                       string `json:"tls_version"`
//	Tls_Negotiated_Protocol           string `json:"tls_negotiated_protocol"`
//	Tls_Negotiated_Protocol_Is_Mutual bool   `json:"tls_negotiated_protocol_is_mutual"`
//}
//
//const (
//	topicsUrl  = "http://182.254.152.69:4171/api/topics"
//	node1Stats = "http://182.254.152.69:4151/stats?format=json"
//	node2Stats = "http://115.159.93.20:4151/stats?format=json"
//	node3Stats = "http://115.159.106.89:4151/stats?format=json"
//)
//
//var AllTopics Topics
//var AllStats Stats
//
//func GetWant() {
//
//	//get all topics
//	request := gorequest.New()
//	resp, body, errs := request.Get(topicsUrl).End()
//
//	if resp.StatusCode != 200 || len(errs) != 0 {
//		newError := errors.New("topicsUrl ERROR")
//		print(newError)
//	}
//	json.Unmarshal([]byte(body), &AllTopics)
//
//	topicSlience := make([]string, 10, 100)
//	for _, val := range AllTopics.{
//		topicSlience = append(topicSlience, val)
//	}
//	topicAndDepth1 := GetTopicsAndOther(node1Stats)
//	topicAndDepth2 := GetTopicsAndOther(node2Stats)
//	topicAndDepth3 := GetTopicsAndOther(node3Stats)
//
//	var topicAndDepth4 = make(map[string]int64)
//
//	fmt.Println(len(topicSlience))
//	for _, val := range topicSlience {
//		topicAndDepth4[val] = topicAndDepth1[val] + topicAndDepth2[val] + topicAndDepth3[val]
//	}
//	for key, val := range topicAndDepth4 {
//		fmt.Printf("topic name is %s and depth is %d\n", key, val)
//	}
//
//}
//
//func GetTopicsAndOther(url string) map[string]int64 {
//
//	request := gorequest.New()
//	resp, body, errs := request.Get(url).End()
//	if resp.StatusCode != 200 || len(errs) != 0 {
//		newError := errors.New(url + "Stats ERROR")
//		print(newError)
//	}
//	json.Unmarshal([]byte(body), &AllStats)
//
//	var topicAndDepth = make(map[string]int64)
//
//	for _, val := range AllStats.Data.Topics {
//		topicAndDepth[val.Topic_Name] = val.Depth
//	}
//
//	return topicAndDepth
//}

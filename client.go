package cronicleapi

import (
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
	"io/ioutil"
	"bytes"
)

const (
	DEF_PROTO = "POST"
)

// Cronicle  Rest API

type NicleClient struct {
	// API Url
	ApiUrl string
	// Api key
	ApiKey string
	client *http.Client
}

// # get_event
func (cli *NicleClient) GetEvent(eventId, title  string) (*EventObject, error) {
	API := "/api/app/get_event/v1"
	r,err := cli.api(API,GetEventReq{
		Id:     eventId,
		Title:  title,
	})
	if err != nil {
		return nil,errors.New(fmt.Sprintf("get event: %s",err.Error()))
	}
	rs := GetEventRsp{}
	err = cli.request(r,&rs)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("get event request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return rs.Event,nil
		}
	}
	return nil,errors.New(fmt.Sprintf("get event error: %s:%s",rs.Code,rs.Description))
}

// # create_event
func (cli *NicleClient) CreateEvent(event *EventObject) (string,error) {
	API := "/api/app/create_event/v1"
	r,err := cli.api(API,event)
	if err != nil {
		return "",errors.New(fmt.Sprintf("create event: %s",err.Error()))
	}
	rs := CreateEventRsp{}
	err = cli.request(r,&rs)
	if err != nil {
		return "",errors.New(fmt.Sprintf("create event request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return rs.Id,nil
		}
	}
	return "",errors.New(fmt.Sprintf("create event error: %s:%s",rs.Code,rs.Description))
}

// # update_event
func (cli *NicleClient) UpdateEvent(eventId string, fields map[string]interface{}) error {
	API := "/api/app/update_event/v1"
	if fields == nil {
		fields = map[string]interface{}{
			"id": eventId,
		}
	} else {
		fields["id"] = eventId
	}
	r,err := cli.api(API,fields)
	if err != nil {
		return errors.New(fmt.Sprintf("update event: %s",err.Error()))
	}
	rs := CommonRsp{}
	err = cli.request(r,&rs)
	if err != nil {
		return errors.New(fmt.Sprintf("update event request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("update event error: %s:%s",rs.Code,rs.Description))
}

// # delete_event
func (cli *NicleClient) DeleteEvent(eventId string) error {
	API := "/api/app/delete_event/v1"
	r,err := cli.api(API,DeleteEventReq{
		Id: eventId,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("delete event: %s",err.Error()))
	}
	rs := CommonRsp{}
	err = cli.request(r,&rs)
	if err != nil {
		return errors.New(fmt.Sprintf("delete event request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("delete event error: %s:%s",rs.Code,rs.Description))
}

// # run_event
func (cli *NicleClient) RunEvent(eventId,title string, fields map[string]interface{}) ([]string,error) {
	API := "/api/app/run_event/v1"
	if fields == nil {
		fields = map[string]interface{}{
			"id":    eventId,
		}
		if len(title) > 0 {
			fields["title"] = title
		}
	} else {
		fields["id"] = eventId
		if len(title) > 0 {
			fields["title"] = title
		}
	}
	r,err := cli.api(API,fields)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("run event: %s",err.Error()))
	}
	rs := RunEventRsp{}
	err = cli.request(r,&rs)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("run event request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return rs.Ids,nil
		}
	}
	return nil,errors.New(fmt.Sprintf("run event error: %s:%s",rs.Code,rs.Description))
}

// # get_event_status
func (cli *NicleClient) GetEventStatus(eventId string) (*EventStatusObject,error) {
	API := "/api/app/get_job_status/v1"
	fields := map[string]interface{}{
		"id": eventId,
	}
	r,err := cli.api(API,fields)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("get event status: %s",err.Error()))
	}
	rs := EventStatusObject{}
	err = cli.request(r,&rs)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("get event status request: %s",err.Error()))
	}
	switch rs.Code.(type) {
	case float64:
		if rs.Code == float64(0) {
			return &rs,nil
		}
	}
	return nil,errors.New(fmt.Sprintf("get event status error: %s:%s",rs.Code,rs.Description))
}

func (cli *NicleClient) api(api string, ro interface{}) (*http.Request,error) {
	data,err := json.Marshal(ro)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("api init: %s",err.Error()))
	}
	r,err := http.NewRequest(DEF_PROTO,cli.ApiUrl+api,bytes.NewReader(data))
	if err != nil {
		return nil,errors.New(fmt.Sprintf("api: %s",err.Error()))
	}

	// client
	cli.client = http.DefaultClient

	// add api key 
	r.Header.Add("X-API-Key",cli.ApiKey)
	r.Header.Add("Content-Type","application/json")
	return r,nil
}


func (cli *NicleClient) request(r *http.Request, rs interface{}) error {
	resp,err := cli.client.Do(r)
	if err != nil {
		return errors.New(fmt.Sprintf("request failed: %s",err.Error()))
	}
	defer resp.Body.Close()
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("response failed: %s",err.Error()))
	}
	err = json.Unmarshal(data,rs)
	if err != nil {
		return errors.New(fmt.Sprintf("response parse failed: %s",err.Error()))
	}
	return nil
}



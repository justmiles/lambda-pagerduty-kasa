package smartplug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	token          = login()
	commandRequest = CommandRequest{}
	cloudUserName  = os.Getenv("KASA_USERNAME")
	cloudPassword  = os.Getenv("KASA_PASSWORD")
)

func login() string {

	login := LoginRequest{
		Method: "login",
		Params: LoginRequestParams{
			AppType:       "Kasa_Android",
			CloudUserName: cloudUserName,
			CloudPassword: cloudPassword,
			TerminalUUID:  "766c61ac-d19f-4ff7-a4d6-823c41243bb6",
		},
	}

	url := "https://wap.tplinkcloud.com/"

	payload, _ := json.Marshal(login)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var loginResponse LoginResponse
	json.Unmarshal(body, &loginResponse)
	if loginResponse.ErrorCode == 0 {
		return loginResponse.Result.Token

	} else {
		fmt.Println(loginResponse.Msg)
		os.Exit(1)
		return ""
	}

}

func GetDeviceList() []Device {

	url := fmt.Sprintf("https://wap.tplinkcloud.com?token=%s", token)

	req, _ := http.NewRequest("POST", url, strings.NewReader(`{"method":"getDeviceList"}`))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var deviceListResponse DeviceListResponse
	json.Unmarshal(body, &deviceListResponse)
	return deviceListResponse.Result.DeviceList

}

func GetDeviceByAlias(alias string) (d Device) {

	deviceList := GetDeviceList()

	for _, device := range deviceList {
		if device.Alias == alias {
			d = device
		}
	}
	if d.DeviceID == "" {
		fmt.Println("Unable to find device")
		os.Exit(1)
	}
	return d

}

func (d *Device) Off() {
	url := fmt.Sprintf("%s?token=%s", d.AppServerURL, token)

	commandRequest.System.SetRelayState.State = 0

	commandRequestJSON, _ := json.Marshal(commandRequest)

	passThroughRequest := PassThroughRequest{
		Method: "passthrough",
		Params: PassThroughRequestParams{
			DeviceID:    d.DeviceID,
			RequestData: string(commandRequestJSON),
		},
	}

	payload, _ := json.Marshal(passThroughRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
}

func (d *Device) On() {
	url := fmt.Sprintf("%s?token=%s", d.AppServerURL, token)

	commandRequest.System.SetRelayState.State = 1

	commandRequestJSON, _ := json.Marshal(commandRequest)

	passThroughRequest := PassThroughRequest{
		Method: "passthrough",
		Params: PassThroughRequestParams{
			DeviceID:    d.DeviceID,
			RequestData: string(commandRequestJSON),
		},
	}

	payload, _ := json.Marshal(passThroughRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
}

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
	token         = login()
	cloudUserName = os.Getenv("KASA_USERNAME")
	cloudPassword = os.Getenv("KASA_PASSWORD")
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
	}
	fmt.Println(loginResponse.Msg)
	os.Exit(1)
	return ""
}

// GetDeviceList returns a list of Devices
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

// GetDeviceByAlias returns a Device object using by alias
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

// Off turns a device relay to the the off state
func (d *Device) Off() {
	url := fmt.Sprintf("%s?token=%s", d.AppServerURL, token)

	passThroughRequest := PassThroughRequest{
		Method: "passthrough",
		Params: PassThroughRequestParams{
			DeviceID:    d.DeviceID,
			RequestData: `{"system":{"set_relay_state":{"state":0}}}`,
		},
	}

	payload, _ := json.Marshal(passThroughRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
}

// On turns a device relay to the the on state
func (d *Device) On() {
	url := fmt.Sprintf("%s?token=%s", d.AppServerURL, token)

	passThroughRequest := PassThroughRequest{
		Method: "passthrough",
		Params: PassThroughRequestParams{
			DeviceID:    d.DeviceID,
			RequestData: `{"system":{"set_relay_state":{"state":1}}}`,
		},
	}

	payload, _ := json.Marshal(passThroughRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
}

// Toggle the relay state
func (d *Device) Toggle() {
	info := d.GetSystemInfo()
	if info.RelayState == 1 {
		d.Off()
	} else {
		d.On()
	}
}

// GetSystemInfo gets current information about device
func (d *Device) GetSystemInfo() (info SystemInfo) {
	url := fmt.Sprintf("%s?token=%s", d.AppServerURL, token)

	passThroughRequest := PassThroughRequest{
		Method: "passthrough",
		Params: PassThroughRequestParams{
			DeviceID:    d.DeviceID,
			RequestData: `{"system":{"get_sysinfo":{}}}`,
		},
	}

	payload, err := json.Marshal(passThroughRequest)
	logIfErr(err)

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	logIfErr(err)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	logIfErr(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	logIfErr(err)

	response := PassThroughResponse{}
	err = json.Unmarshal(body, &response)
	logIfErr(err)

	sysInfoRes := CmdResponseGetSystemInfo{}
	err = json.Unmarshal([]byte(response.Result.ResponseData), &sysInfoRes)
	logIfErr(err)

	return sysInfoRes.System.GetSysinfo
}

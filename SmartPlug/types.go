package smartplug

type LoginRequest struct {
	Method string             `json:"method"`
	Params LoginRequestParams `json:"params"`
}

type LoginRequestParams struct {
	AppType       string `json:"appType"`
	CloudUserName string `json:"cloudUserName"`
	CloudPassword string `json:"cloudPassword"`
	TerminalUUID  string `json:"terminalUUID"`
}

type LoginResponse struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Result    struct {
		AccountID string `json:"accountId"`
		RegTime   string `json:"regTime"`
		Email     string `json:"email"`
		Token     string `json:"token"`
	} `json:"result"`
}

type PassThroughRequest struct {
	Method string                   `json:"method"`
	Params PassThroughRequestParams `json:"params"`
}

type PassThroughRequestParams struct {
	DeviceID    string `json:"deviceId"`
	RequestData string `json:"requestData"`
}

type PassThroughResponse struct {
	ErrorCode int `json:"error_code"`
	Result    struct {
		ResponseData string `json:"responseData"`
	} `json:"result"`
}

type CommandRequest struct {
	System struct {
		SetRelayState struct {
			ErrCode int `json:"err_code"`
			State   int `json:"state"`
		} `json:"set_relay_state"`
	} `json:"system"`
}

type CommandResponse struct {
	ErrorCode int `json:"error_code"`
	Result    struct {
		ResponseData string `json:"responseData"`
	} `json:"result"`
}

type DeviceListResponse struct {
	ErrorCode int `json:"error_code"`
	Result    struct {
		DeviceList []Device `json:"deviceList"`
	} `json:"result"`
}

type Device struct {
	FwVer        string `json:"fwVer"`
	DeviceName   string `json:"deviceName"`
	Status       int    `json:"status"`
	Alias        string `json:"alias"`
	DeviceType   string `json:"deviceType"`
	AppServerURL string `json:"appServerUrl"`
	DeviceModel  string `json:"deviceModel"`
	DeviceMac    string `json:"deviceMac"`
	Role         int    `json:"role"`
	IsSameRegion bool   `json:"isSameRegion"`
	HwID         string `json:"hwId"`
	FwID         string `json:"fwId"`
	OemID        string `json:"oemId"`
	DeviceID     string `json:"deviceId"`
	DeviceHwVer  string `json:"deviceHwVer"`
}

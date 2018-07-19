package smartplug

import (
	"testing"
)

func TestHandler(t *testing.T) {
	device := GetDeviceByAlias("Mini Plug #1")
	device.Off()
	device.On()
}

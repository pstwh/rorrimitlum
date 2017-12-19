package xrandr

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetDisconnected() []string {
	xrandr, _ := exec.Command("xrandr").Output()
	xrandr_string := string(xrandr)

	re := regexp.MustCompile(`VIRTUAL[\d]+ disconnected`)

	virtualMonitors := re.FindAllString(xrandr_string, -1)

	for i := range virtualMonitors {
		virtualMonitors[i] = strings.Split(virtualMonitors[i], " ")[0]
	}

	return virtualMonitors
}

func GetConnected() []string {
	xrandr, _ := exec.Command("xrandr").Output()
	xrandr_string := string(xrandr)

	re := regexp.MustCompile(`VIRTUAL[\d]+ connected`)

	virtualMonitors := re.FindAllString(xrandr_string, -1)

	for i := range virtualMonitors {
		virtualMonitors[i] = strings.Split(virtualMonitors[i], " ")[0]
	}

	return virtualMonitors
}

func DisconnectAll() {
	virtualMonitors := GetConnected()

	for i := range virtualMonitors {
		xrandr := exec.Command("xrandr", "--output", virtualMonitors[i], "--off")
		_ = xrandr.Run()

		xrandr2 := exec.Command("xrandr", "--delmode", virtualMonitors[i], "1024x768")
		_ = xrandr2.Run()
	}
}

func Disconnect(virtualMonitor string) {
	xrandr := exec.Command("xrandr", "--output", virtualMonitor, "--off")
	_ = xrandr.Run()

	xrandr2 := exec.Command("xrandr", "--delmode", virtualMonitor, "1024x768")
	_ = xrandr2.Run()
}

func Connect() string {
	virtualMonitors := GetDisconnected()

	virtualMonitor := virtualMonitors[0]

	xrandr := exec.Command("xrandr", "--addmode", virtualMonitor, "1024x768")
	_ = xrandr.Run()

	return virtualMonitor
}

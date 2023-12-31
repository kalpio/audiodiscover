package audiodiscover

import (
	"github.com/kalpio/audiodiscover/domain"
	"os/exec"
	"strings"
)

type discover struct {
	ffmpeg string // path to ffmpeg executable
}

func NewDiscover(ffmpeg string) *discover {
	return &discover{
		ffmpeg: ffmpeg,
	}
}

func (d discover) Devices() ([]domain.Device, error) {
	cmd := exec.Command(d.ffmpeg, listDevicesSwitches()...)
	out, _ := cmd.CombinedOutput()

	return d.parseDevices(out), nil
}

func (d discover) parseDevices(out []byte) []domain.Device {
	var result []domain.Device
	value := string(out)
	arr := strings.Split(value, "\n")
	for _, s := range arr {
		if strings.Contains(s, "(audio)") {
			result = append(result, *d.getDeviceName(s))
		}
	}
	return result
}

func (discover) getDeviceName(s string) *domain.Device {
	startIndex := strings.Index(s, "\"")
	if startIndex == -1 {
		return nil
	}

	lastIndex := strings.LastIndex(s, "(audio)")
	if lastIndex == -1 {
		return nil
	}

	val := s[startIndex:lastIndex]
	name := strings.Trim(
		strings.TrimSpace(val), "\"")
	return &domain.Device{Name: name}
}

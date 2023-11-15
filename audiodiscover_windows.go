//go:build windows

package audiodiscover

func listDevicesSwitches() []string {
	return []string{"-hide_banner", "-f", "dshow", "-list_devices", "true", "-i", "null"}
}

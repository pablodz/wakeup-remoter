package waker

import (
	"fmt"
	"os/exec"

	"github.com/pablodz/wakeup-remoter/internal/models"
)

// RTCWAKE
func ExecuteRTCWAKE(req *models.RTCWAKE) (string, error) {
	// execute as sudo command rtc wake
	t := req.Date.Local().Format("2006-01-02 15:04:05")
	// rtcwake -m mem -t $(date +%s -d "1 minute")
	cmd := exec.Command("sudo", "rtcwake", "-m", req.Type, "-t", "$( date +%s -d "+t+")")
	// print command
	fmt.Println("Command:", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

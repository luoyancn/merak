package merak

import (
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/luoyancn/dubhe/logging"
)

func Exec(command []string) {
	cmdStr := strings.Join(command[1:], " ")
	logging.LOG.Infof("Running the command %s \n", cmdStr)
	cmd := exec.Command("ps", "-efl")
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		logging.LOG.Errorf("ERROR: %v\n", err)
		return
	}
	if err := cmd.Start(); nil != err {
		logging.LOG.Errorf("ERROR: %v\n", err)
		return
	}
	res, _ := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		logging.LOG.Errorf("ERROR: %v\n", err)
		return
	}
	logging.LOG.Infof("The result is %s\n", string(res))
}

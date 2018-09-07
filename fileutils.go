package merak

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/luoyancn/dubhe/logging"
)

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if nil != err {
		logging.LOG.Errorf("Cannot open the file :%v\n", err)
		return nil, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	var lines []string
	for {
		_buffer, err := buf.ReadString('\n')
		if nil != err {
			if io.EOF == err {
				break
			}
			logging.LOG.Errorf("Error occured when reading file%s:%v\n",
				filename, err)
			return nil, err
		}
		// Delete empty line from content
		if 0 == len(_buffer) || _buffer == "\r\n" || _buffer == "\n" {
			continue
		}
		lines = append(lines, strings.TrimSpace(_buffer))
	}
	return lines, nil
}

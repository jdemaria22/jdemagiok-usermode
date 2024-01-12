package usermode

import (
	"strings"

	"github.com/mitchellh/go-ps"
)

func GetProcessID(processName string) int {
	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	for _, p := range processes {
		if strings.EqualFold(p.Executable(), processName) {
			return p.Pid()
		}
	}

	return 0
}

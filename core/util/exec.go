package util

import (
	"fmt"
	// "sync"
	"time"
	// "strings"
	"github.com/go-cmd/cmd"
)

/*
This method is to execute shell commands received 
*/
func Execute(_cmd string) []string {
	_cmds := Split(_cmd)
	options := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	command := cmd.NewCmdOptions(options, _cmds[0], _cmds[1:]...)
	
	var result []string

	go func() {
		for {
			select {
			case line := <-command.Stdout:
				// processprint(option, line)
				result = append(result, line)
			case line := <-command.Stderr:
				fmt.Println(line)
			}
		}
	}()
	
	<-command.Start()

	for len(command.Stdout) > 0 || len(command.Stderr) > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	return result
}
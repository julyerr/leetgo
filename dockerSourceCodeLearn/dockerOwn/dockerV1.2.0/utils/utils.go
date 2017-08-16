package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/pkg/term"
	"github.com/docker/docker/pkg/units"
)


type JSONError struct{
	Code int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type JSONProgress struct{
	terminalFd uintptr
	Current int `json:"current,omitempty"`
	Total int `json:"total,omitempty"`
	Start int64 `json:"start,omitempty"`
}

type JSONMessage struct{
	Stream string `json:"stream,omitempty"`
	Status string `json:"status,omitempty"`
	Progress *JSONProgress `json:"progressDetail,omitempty"`
	ProgressMessage string `json:progess,omitempty`
	ID string `json:"id,omitempty"`
	From string `json:"from,omitempty"`
	Time int64 `json:"time,omitempty"`
	Error *JSONError `json:"errorDetail,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func DisplayJSONMessagesStream(in io.Reader,out io.Writer,terminalFd uintptr,isTerminal bool) error{
	var (
		dec = json.NewDecoder(in)
		ids = make(map[string]int)
		diff = 0
		)
	for {
		var jm JSONMessage
		if err := dec.Decode(&jm);err != nil{
			if err == io.EOF{
				break
			}
			return err
		}
		if jm.Progress != nil{
			jm.Progress.terminalFd = terminalFd
		}
		if jm.ID != "" && (jm.Progress != nil || jm.ProgressMessage != "") {
			line, ok := ids[jm.ID]
			if !ok {
				line = len(ids)
				ids[jm.ID] = line
				fmt.Fprintf(out,"\n")
				diff = 0
			}else{
				diff = len(ids) - line
			}
			if jm.ID != "" && isTerminal {
				fmt.Fprintf(out,"%c[%dA",27,diff)
			}
		}
		err := jm.Display(out,isTerminal)
		if jm.ID != "" && isTerminal{
			fmt.Fprintf(out,"%c[%dB",27,diff)
		}
		if err != nil{
			return err
		}
	}
	return nil
}


package term

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

var (
	ErrInvalidState = errors.New("Invalid terminal state")
)

type State struct{
	termios Termios
}



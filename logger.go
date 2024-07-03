package main

import (
	"io"
	"os"
)

type writer struct {
	io.Writer
}

func (w writer) Write(bytes []byte) (n int, e error) {
	fi, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o655)
	if err != nil {
		return 0, err
	}
	return fi.Write(bytes)
}

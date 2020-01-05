package main

import (
	"os"
	// external packages
	"github.com/satom9to5/webext/nativemessaging"
	// this package
	//"native_message/logs"
	"native_message/message"
)

func main() {
	req := message.NewRequestMessage()
	res := message.NewResponseMessage()
	var err error

	if err = nativemessaging.Receive(req, os.Stdin); err == nil {
		if data, err := req.Run(); err == nil {
			res.Data = data
		} else {
			res.Error = err.Error()
		}
	} else {
		res.Error = err.Error()
	}

	nativemessaging.Send(res, os.Stdout)
}

package main

import (
	"context"

	"github.com/corezoid/gitcall-go-runner/gitcall"
)

var key = ""

func usercode(_ context.Context, data map[string]interface{}) error {

	data["old_key"] = key
	key = data["key"].(string)
	data["status"] = "ok"
	return nil
}

func main() {
	gitcall.Handle(usercode)
}

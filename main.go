package main

import (
	"cloud.google.com/go/container"
	"github.com/go-logr/logr"
)

func main() {
	l := logr.Logger{}
	_ = l
	_ = container.Client{}
}

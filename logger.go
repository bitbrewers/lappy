package lappy

import (
	"fmt"
	"os"
)

type Logger interface {
	Debugf(tmpl string, args ...interface{})
	Infof(tmpl string, args ...interface{})
	Errorf(tmpl string, args ...interface{})
	Fatalf(tmpl string, args ...interface{})
}

type FatalLogger struct{}

func (FatalLogger) Debugf(_ string, _ ...interface{}) {}
func (FatalLogger) Infof(_ string, _ ...interface{})  {}
func (FatalLogger) Errorf(_ string, _ ...interface{}) {}
func (FatalLogger) Fatalf(tmpl string, args ...interface{}) {
	fmt.Printf(tmpl, args)
	os.Exit(1)
}

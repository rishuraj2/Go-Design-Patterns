package main

import (
	"loggerlevel/logs"
)

func main() {

	l1 := logs.GetLog()
	l2 := logs.GetLog()

	l1.SetLevel(logs.WARN)

	l1.Warn("this is l1 warning message")
	l2.Debug("this is l2 debug message")


}
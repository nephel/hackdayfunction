package main

import (
	"github.com/nephel/cfevents"
	"github.com/nephel/hackdayfunction"
)

func main() {
	event := cfevents.NewCfEvent(myfunc.Handle)
	event.Run()
}

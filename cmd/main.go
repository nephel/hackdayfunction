package main

import (
	"github.com/samze/cfevents"
	"github.com/samze/hackdayfunction"
)

func main() {
	event := cfevents.NewCfEvent(myfunc.Handle, "myapp")
	event.Run()
}

package main

import (
	"flag"
	"phonecontroller/flags"
	"phonecontroller/swipe"
	"phonecontroller/tap"
)

var Cmd = []string{"adb"}

func main() {

	flag.Parse()

	if *flags.SerialNo != "" {
		Cmd = append(Cmd, "-s")
		Cmd = append(Cmd, *flags.SerialNo)
	}

	Cmd = append(Cmd, "shell")
	Cmd = append(Cmd, "input")

	if *flags.Source != "" {
		Cmd = append(Cmd, *flags.Source)
	}

	switch *flags.Action {
	case "swipe":
		Cmd = append(Cmd, "swipe")
		swipe.Run(Cmd)
	case "tap":
		tap.Run(Cmd)

	}

}

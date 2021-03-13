package swipe

import (
	"log"
	"os/exec"
	"phonecontroller/flags"
	"strconv"
	"time"
)

func Run(cmd []string) {
	log.Println("swipe run")

	log.Println(cmd)
	cmd = append(cmd, strconv.Itoa(*flags.FromX))
	cmd = append(cmd, strconv.Itoa(*flags.FromY))
	cmd = append(cmd, strconv.Itoa(*flags.ToX))
	cmd = append(cmd, strconv.Itoa(*flags.ToY))
	log.Println(cmd)

	for {
		select {
		case <-time.Tick(time.Duration(*flags.Interval) * time.Second):

			exe := exec.Command(cmd[0], cmd[1:]...)
			exe.Run()

		}
	}

}

package run

import (
	"log"
	"os/exec"
	"phonecontroller/flags"
	"time"
)

func exe(cmd []string, wait int) {
	exe := exec.Command(cmd[0], cmd[1:]...)
	exe.Run()
	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func Run(cmd []string) {
	originalCmd := cmd

	for {
		select {
		case <-time.Tick(630 * time.Second):
			// 点亮屏幕
			screenOnCmd := append(originalCmd, "input", "keyevent", "26")
			log.Println("screen on")
			exe(screenOnCmd, 5)

			startCmd := append(originalCmd, "am", "start", *flags.App)
			log.Println("start app: ", startCmd)
			exe(startCmd, 5)

			clickTabBarCmd := append(originalCmd, "input", "tap", "584", "1848")
			log.Println("click tab bar")
			exe(clickTabBarCmd, 12)

			clickBoxCmd := append(originalCmd, "input", "tap", "995", "1665")
			log.Println("click box")
			exe(clickBoxCmd, 6)

			clickVideoCmd := append(originalCmd, "input", "tap", "555", "1320")
			log.Println("click video button")
			exe(clickVideoCmd, 25)

			clickCloseCmd := append(originalCmd, "input", "tap", "1040", "105")
			log.Println("click close button")
			exe(clickCloseCmd, 10)

			closeAppCmd := append(originalCmd, "am", "force-stop", "com.ss.android.article.lite")
			log.Println("close app")
			exe(closeAppCmd, 10)

		}
	}

}

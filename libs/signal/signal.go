package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InitSignal() {
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

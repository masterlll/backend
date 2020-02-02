package logroute

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

var logChan chan LogChanReq

type LogChanReq struct {
	Msg string
	//err  error
	Curtime time.Time
}

// InitPlan :
func Init() {
	logChan = make(chan LogChanReq, 10000)

	for i := 0; i < runtime.NumCPU(); i++ {
		go logProssces(i)
	}
}

func logProssces(i int) {
	log.Print("log prossces  No: ", i)
	for c := range logChan {
		//  your log logic
		fmt.Println("your err log ", c.Msg)
		fmt.Println("your log time ", c.Curtime)
		//
	}

}

func LogSave(i string) {
	if len(logChan) > 10000 {
		time.Sleep(1 * time.Second)
	}
	logChan <- LogChanReq{i, time.Now().UTC()}
}

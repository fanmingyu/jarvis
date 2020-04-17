package report

import(
	"log"
)

var Queue = make(chan Report, 500000)

func WorkerStart(number int) {
	for i := 0; i < number; i++ {
		go processQueue()
	}
}

func processQueue() {
	defer func(){
		err := recover()
		log.Printf("sms report. process queue panic:%v", err)
		go processQueue()
	}()

	for {
		report := <-Queue
		report.UpdateSmsRecord()
	}
}

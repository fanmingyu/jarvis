package workers

import (
	"log"

	"smsgate/server/storage"
)

type queue interface {
	Init()
	Enqueue(SmsMessage)
	Dequeue() SmsMessage
	Len() int
	BufferLen() int
}

type worker struct {
	Name    string
	Count   int
	Queue   queue
	Storage storage.SmsStorage
}

var Workers = map[string]worker{
	"verify": worker{
		Name:    "验证码队列",
		Count:   20,
		Queue:   &ChanQueue{ChanBufferLen: 100000},
		Storage: &storage.SmsMysql{},
	},
	"notice": worker{
		Name:    "通知队列",
		Count:   20,
		Queue:   &ChanQueue{ChanBufferLen: 300000},
		Storage: &storage.SmsMysql{},
	},
	// "marketing": worker{
	// 	Name:  "营销队列",
	// 	Count: 20,
	// 	Queue: &ChanQueue{ChanBufferLen: 500000},
	// 	Storage: &storage.SmsMysql{},
	// },
	// "test":	worker{
	// 	Name:	"测试队列",
	// 	Count:	5,
	// 	Queue:	&ChanQueue{ChanBufferLen:5000000},
	// 	Storage: &storage.SmsMysql{},
	// },
}

func init() {
	for _, v := range Workers {
		v.Queue.Init()
		for i := 0; i < v.Count; i++ {
			go processWorker(v.Queue)
		}
	}
}

// 短信处理worker
func processWorker(q queue) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("processWorker recover. err:%v", err)
		}
		go processWorker(q)
	}()

	for {
		message := q.Dequeue()
		message.Consume()
	}
}

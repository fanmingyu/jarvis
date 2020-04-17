package workers

// 基于chan的队列
type ChanQueue struct {
	ChanBufferLen int
	ch            chan SmsMessage
}

func (q *ChanQueue) Init() {
	q.ch = make(chan SmsMessage, q.ChanBufferLen)
}

func (q *ChanQueue) Enqueue(m SmsMessage) {
	q.ch <- m
}

func (q *ChanQueue) Dequeue() SmsMessage {
	return <-q.ch
}

func (q *ChanQueue) Len() int {
	return len(q.ch)
}

func (q *ChanQueue) BufferLen() int {
	return q.ChanBufferLen
}

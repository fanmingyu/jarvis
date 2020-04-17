package workers

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestChanQueue(t *testing.T) {
	q := ChanQueue{BufferLen: 100}
	q.Init()

	q.Enqueue(SmsMessage{App:"pcn", Mobile: "18600001111", Content: "test"})
	q.Enqueue(SmsMessage{App:"pcn", Mobile: "18600001111", Content: "test"})
	q.Enqueue(SmsMessage{App:"pcn", Mobile: "18600001111", Content: "test"})

	Convey("test enqueue", t, func() {
		So(q.Len(), ShouldEqual, 3)
	})

	message := q.Dequeue()
	Convey("test dequeue", t, func() {
		So(q.Len(), ShouldEqual, 2)
		So(message.App, ShouldEqual, "pcn")
	})
}

package workers

import (
	"testing"
)

func TestProcessWorker(t *testing.T) {
	q := ChanQueue{BufferLen: 100}
	go processWorker(&q)
}

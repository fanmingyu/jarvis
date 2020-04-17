package utils

import(
	"testing"
)

func BenchmarkMonitor(b *testing.B) {
	s := []string{"10.20.69.57:26479", "10.20.69.57:26579", "10.20.69.57:26679"}
	MonitorInit("def_master", s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MonitorAdd("TEST_MONITOR", 1)
	}
}

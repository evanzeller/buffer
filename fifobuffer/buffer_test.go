package fifobuffer

import (
	"testing"
	"time"
	"math/rand"
)

func BenchmarkBufferAdd(b *testing.B) {
	buffer := NewBuffer(25)

	for i:=0; i < 5; i++ {

		var values = make([]uint64, 0)
		for n:=0; n < 2000; n++ {
			values = append(values, rand.Uint64())
		}
		metric := Metric{startTime: time.Now(), endTime: time.Now()}
		metric.metric = metricData{metricName: "CPU Usage", values: values}
		metrics = append(metrics, metric)
	}

	for n:=0; n < b.N; n++ {
		buffer.Add(metrics)
	}

	for n:=0; n < b.N; n++ {
		buffer.Pop()
	}
}
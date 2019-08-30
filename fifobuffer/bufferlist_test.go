package fifobuffer

import (
	"testing"
	"math/rand"
	"time"
)

var metrics = make([]Metric, 0)

// func TestMain(m *testing.M) {
// 	for i:=0; i < 5; i++ {
// 		metric := &Metric{startTime: time.Now(), endTime: time.Now()}
// 		metric.metric = &metricData{metricName: "CPU Usage"}
// 		var values []uint64 = make([]uint64, 0)
// 		for n:=0; n < 2000; n++ {
// 			values = append(values, rand.Uint64())
// 		}
// 	}
// }

func BenchmarkListBufferAdd(b *testing.B) {
	buffer := NewListBuffer(25)

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
		buffer.Push(metrics)
	}

	for n:=0; n < b.N; n++ {
		buffer.Pop()
	}

}
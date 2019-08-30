package fifobuffer

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func BenchmarkRingBufferWrite(b *testing.B) {
	buffer := NewRingBuffer(25)

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
		buffer.Write(metrics)
	}

	for n:=0; n < b.N; n++ {
		buffer.Read()
	}

}

func TestRingBufferFIFO(t *testing.T) {
	buffer := NewRingBuffer(3)

	buffer.Write(1)
	buffer.Write(2)
	buffer.Write(3)

	for i:=0; i < 3; i++ {
		value, err := buffer.Read()
		if err != nil {
			fmt.Printf("error: %s\t", err)
		} else {
			fmt.Printf("value: %d\t", value)
		}
	}
}
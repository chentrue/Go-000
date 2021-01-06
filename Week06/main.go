package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 滑动计数器
type SlidingCounter interface {
	Sum() float64
	Max() float64
	Avg() float64
	Increment(i float64)
}

type slidingCounter struct {
	//以时间为key存储的窗口
	windows  map[int64]*window
	//时间间隔
	interval int64
}

type window struct {
	Value uint32
}

func NewSlidingCounter(interval int64) SlidingCounter {
	return &slidingCounter{
		windows:  make(map[int64]*window),
		interval: interval,
	}
}

func (sc *slidingCounter) getCurrentWindow() *window {
	now := time.Now().Unix()

	if w, found := sc.windows[now]; found {
		return w
	}

	result := &window{}
	sc.windows[now] = result
	return result
}

func (sc *slidingCounter) removeOldWindows() {
	notExpired := time.Now().Unix() - sc.interval

	for timestamp := range sc.windows {
		if timestamp <= notExpired {
			delete(sc.windows, timestamp)
		}
	}
}

func (sc *slidingCounter) Increment(i float64) {
	if i == 0 {
		return
	}

	w := sc.getCurrentWindow()
	atomic.AddUint32(&w.Value,1)
	sc.removeOldWindows()
}

func (sc *slidingCounter) Sum() float64 {
	now := time.Now().Unix()

	var sum uint32
	for timestamp, window := range sc.windows {
		if timestamp >= now-sc.interval {
			sum += window.Value
		}
	}

	return float64(sum)
}

func (sc *slidingCounter) Max() float64 {
	now := time.Now().Unix()

	var max float64
	for timestamp, window := range sc.windows {
		if timestamp >= now - sc.interval {
			if float64(window.Value) > max {
				max = float64(window.Value)
			}
		}
	}
	return max
}

func (sc *slidingCounter) Avg() float64 {
	return sc.Sum() / float64(sc.interval)
}

func main() {
	//模拟计数
	counter := NewSlidingCounter(10)
	for _, request := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		counter.Increment(request)
		time.Sleep(1 * time.Second)
	}

	fmt.Println(counter.Sum(),counter.Avg(),counter.Max())
}

package main

import (
	"fmt"
	"time"
)

/*
自实现的一个简单定时器，具备到点自动暂停功能
*/

type OwnTicker struct {
	C          chan time.Time
	Done       chan struct{}
	duration   time.Duration
	finishTime time.Time
}

func NewTicker(duration time.Duration, finishTime time.Time) (ticker *OwnTicker) {
	ticker = new(OwnTicker)
	ticker.C = make(chan time.Time)
	ticker.Done = make(chan struct{})
	ticker.duration = duration
	ticker.finishTime = finishTime
	go ticker.run()
	return ticker
}

func (o OwnTicker) run() {
	for {
		time.Sleep(o.duration)
		o.C <- time.Now()
		if o.finishTime.Sub(time.Now()) < 0 {
			break
		}
	}
	o.Done <- struct{}{}
}

func Provider() *OwnTicker {
	ticker := NewTicker(3*time.Second, time.Now().Add(10*time.Minute))
	return ticker
}

func Consumer(ticker OwnTicker) {
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("现在的时间是:%s\n", t.Format("2006-01-02 15:04:05"))
		case <-ticker.Done:
			fmt.Println("定时器结束啦！")
			break
		}
	}
}

func main() {
	ticker := Provider()
	Consumer(*ticker)

}

package main

import (
	"time"

	"github.com/51reboot/golang-01-homework/lesson12/wuchuanfang/monitor/common"
)

type MetricFunc func() []*common.Metric

type Sched struct {
	ch chan *common.Metric
}

func NewSched(ch chan *common.Metric) *Sched {
	return &Sched{
		ch: ch,
	}
}

func (s *Sched) AddMetric(collecter MetricFunc, step time.Duration) {
	ticker := time.NewTicker(step)
	for range ticker.C {
		for _, metric := range collecter() {
			s.ch <- metric
		}
	}
}

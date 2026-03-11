package main

import (
	"fmt"
	"time"
)

type Starter interface {
	Start()
}

type Stopper interface {
	Stop()
}

type Machine interface {
	Starter
	Stopper
}

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine started")
}

func (e Engine) Stop() {
	fmt.Println("Engine stopped")
}

func runCycle(m Machine) {
	m.Start()
	time.Sleep(520000)
	m.Stop()
}

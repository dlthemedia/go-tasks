package main

type Runner interface {
	Run() string
}

type Athlete struct {
	Name string
}

func (a Athlete) Run() string {
	return a.Name + " бежит!"
}

var _ Runner = Athlete{}

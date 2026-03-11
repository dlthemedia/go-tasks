package main

type Timer struct {
	Seconds int

	Running bool
}

func (t *Timer) Start() {

	t.Running = true

}

func (t *Timer) Stop() {

	t.Running = false

}

func Status(t *Timer) string {

	if t.Running == true {
		return "running"
	}
	return "stopping"

}

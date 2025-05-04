package sandbox

import "time"

type Timer interface {
	Start()
	ElapsedMs() int
}

type RealTimer struct {
	start time.Time

	Now func() time.Time
}

func (t *RealTimer) Start()         { t.start = t.Now() }
func (t *RealTimer) ElapsedMs() int { return int(t.Now().Sub(t.start).Milliseconds()) }

type FakeTimer struct {
	FixedMs int
}

func (t *FakeTimer) Start()         {}
func (t *FakeTimer) ElapsedMs() int { return t.FixedMs }

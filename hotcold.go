package eticker

import "time"

type Work func() bool

type HotColdTicker struct {
	counter  int
	cooldown int
	cold     time.Duration
	hot      time.Duration
	work     Work

	stopC chan struct{}

	started bool
}

func NewHotColdTicker(cold time.Duration, hot time.Duration, cooldown int, work Work) *HotColdTicker {
	hct := &HotColdTicker{
		cooldown: cooldown,
		cold:     cold,
		hot:      hot,
		work:     work,
		stopC:    make(chan struct{}),
	}

	hct.start()

	return hct
}

func (t *HotColdTicker) start() {
	for {
		delay := t.cold

		if t.counter > 0 {
			delay = t.hot
		}

		select {
		case <-time.NewTimer(delay).C:
			t.do()
		case <-t.stopC:
			return
		}
	}
}

func (t *HotColdTicker) do() {
	if t.work() {
		// hot
		if t.counter < t.cooldown {
			t.counter++
		}
	} else {
		// cold
		if t.counter > 0 {
			t.counter--
		}
	}
}

func (t *HotColdTicker) Stop() {
	select {
	case <-t.stopC:
	default:
		close(t.stopC)
	}
}

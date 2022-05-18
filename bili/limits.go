package bili

import (
	"sync"
	"time"
)

//var apiTicket = uint16(16)
//var apiBucket

// TicketBucket 体力（划）令牌桶
type TicketBucket struct {
	used     uint16
	full     uint16
	time     time.Time     // 上次增加至今
	short    time.Duration // full增加间隔必须大于此
	long     time.Duration // full减少间隔
	longLock sync.Mutex
	sync.Mutex
}

func NewTicketBucket(full uint16, short, long time.Duration) *TicketBucket {
	return &TicketBucket{
		full:  full,
		short: short,
		long:  long,
	}
}

func (tb *TicketBucket) Use() {
	tb.Lock()
	defer tb.Unlock()
	t := time.Now()
	delta := t.Sub(tb.time)
	times := uint16(delta / tb.long)
	if times > tb.used {
		times = tb.used
	}
	tb.used -= times
	if tb.used < tb.full { // 有体力，减一
		if delta < tb.short { // 体力消耗过快
			time.Sleep(tb.short - delta)
			t = tb.time.Add(tb.short)
		}
		tb.used++
		tb.time = t
	} else { // 没体力，睡到回复一点体力
		time.Sleep(tb.long - delta)
		tb.time = tb.time.Add(tb.long)
	}
}

var apiBucket = NewTicketBucket(16, 200*time.Millisecond, 500*time.Millisecond)

package function

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

// CreateOrderNo 生成平台32位订单号 [不推荐使用]
func CreateOrderNo() (sn string) {
	now := time.Now()
	unixNano := now.UnixNano()
	sn = now.Format(TimeLayoutYMDHIS)
	timeStr := Int64ToString(unixNano)
	rand.Seed(unixNano)
	sn += timeStr[10:]
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn += strconv.Itoa(rand.Intn(10))
	sn = fmt.Sprintf("%032s", sn)
	return
}

var num int64

// GenerateOrderNo 生成24位订单号 [前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号]
func GenerateOrderNo() string {
	t := time.Now()
	s := t.Format(TimeLayoutYMDHIS)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

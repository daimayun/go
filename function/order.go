package function

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// CreateOrderNo 生成平台32位订单号
func CreateOrderNo() (sn string) {
	unixNano := time.Now().UnixNano()
	sn = NowDataTimeStr()
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

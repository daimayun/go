package function

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandString() string {
	return Md5(fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UnixNano(), 10), strconv.Itoa(rand.Intn(1000000))))
}

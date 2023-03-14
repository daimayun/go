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

func OrderNoStrEncode(str string) string {
	str = str[2:]
	end := str[12:]
	arr := map[string]string{
		"00": "7",
		"01": "a",
		"02": "b",
		"03": "c",
		"04": "d",
		"05": "e",
		"06": "f",
		"07": "g",
		"08": "h",
		"09": "i",
		"10": "j",
		"11": "k",
		"12": "l",
		"13": "m",
		"14": "n",
		"15": "o",
		"16": "p",
		"17": "q",
		"18": "r",
		"19": "s",
		"20": "t",
		"21": "u",
		"22": "v",
		"23": "w",
		"24": "x",
		"25": "y",
		"26": "z",
		"27": "A",
		"28": "B",
		"29": "C",
		"30": "D",
		"31": "E",
		"32": "F",
		"33": "G",
		"34": "H",
		"35": "I",
		"36": "J",
		"37": "K",
		"38": "L",
		"39": "M",
		"40": "N",
		"41": "O",
		"42": "P",
		"43": "Q",
		"44": "R",
		"45": "S",
		"46": "T",
		"47": "U",
		"48": "V",
		"49": "W",
		"50": "X",
		"51": "Y",
		"52": "Z",
		"53": "0",
		"54": "1",
		"55": "2",
		"56": "3",
		"57": "4",
		"58": "5",
		"59": "6",
	}
	y := str[0:2]
	m := str[2:4]
	d := str[4:6]
	h := str[6:8]
	i := str[8:10]
	s := str[10:12]
	return arr[y] + arr[m] + arr[d] + arr[h] + arr[i] + arr[s] + end
}

func OrderNoStrDecode(str string) string {
	arr := map[string]string{
		"7": "00",
		"a": "01",
		"b": "02",
		"c": "03",
		"d": "04",
		"e": "05",
		"f": "06",
		"g": "07",
		"h": "08",
		"i": "09",
		"j": "10",
		"k": "11",
		"l": "12",
		"m": "13",
		"n": "14",
		"o": "15",
		"p": "16",
		"q": "17",
		"r": "18",
		"s": "19",
		"t": "20",
		"u": "21",
		"v": "22",
		"w": "23",
		"x": "24",
		"y": "25",
		"z": "26",
		"A": "27",
		"B": "28",
		"C": "29",
		"D": "30",
		"E": "31",
		"F": "32",
		"G": "33",
		"H": "34",
		"I": "35",
		"J": "36",
		"K": "37",
		"L": "38",
		"M": "39",
		"N": "40",
		"O": "41",
		"P": "42",
		"Q": "43",
		"R": "44",
		"S": "45",
		"T": "46",
		"U": "47",
		"V": "48",
		"W": "49",
		"X": "50",
		"Y": "51",
		"Z": "52",
		"0": "53",
		"1": "54",
		"2": "55",
		"3": "56",
		"4": "57",
		"5": "58",
		"6": "59",
	}
	y := str[0:1]
	m := str[1:2]
	d := str[2:3]
	h := str[3:4]
	i := str[4:5]
	s := str[5:6]
	end := str[6:]
	return "20" + arr[y] + arr[m] + arr[d] + arr[h] + arr[i] + arr[s] + end
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

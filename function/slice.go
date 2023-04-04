package function

import (
	"math/rand"
	"sort"
	"time"
)

// InSlice 判断切片中是否存在某值
func InSlice(slice []string, target string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

// SliceUnique 切片去重
func SliceUnique(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}
	var (
		newSlice = make([]string, 0)
		m        = make(map[string]interface{})
	)
	for _, v := range slice {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = 0
		newSlice = append(newSlice, v)
	}
	return newSlice
}

// SliceShuffle 打乱切片
func SliceShuffle(slice *[]string) {
	gRand := rand.New(rand.NewSource(time.Now().UnixNano()).(rand.Source64))
	length := len(*slice) - 1
	for i := length; i > 0; i-- {
		n := gRand.Intn(i + 1)
		(*slice)[i], (*slice)[n] = (*slice)[n], (*slice)[i]
	}
}

// SliceSort 切片排序[倒序]
func SliceSort(arr []int64) []int64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

package array

import (
	"fmt"
	"math"
	"reflect"
	"withJsAPI/utils"
)

// Find 方法遍历切片中符合回调逻辑的值 如找到则返回true 反之
// the method traverses the values in the slice that conform to the callback logic. If found, it returns true, otherwise
func Find[T comparable](arr []T, cb func(T) bool) bool {
	if utils.IsEmptySlice(arr) {
		return false
	}
	for _, e := range arr {
		result := cb(e)
		if result {
			return true
		}
	}
	return false
}

// FindIndex 方法遍历切片中符合回调逻辑的值 如找到则返回下标 反之返回-1
// the method traverses the values in the slice that conform to the callback logic. If found, it returns the subscript, otherwise it returns -1
func FindIndex[T comparable](arr []T, cb func(T) bool) int {
	if utils.IsEmptySlice(arr) {
		return -1
	}
	for index, e := range arr {
		result := cb(e)
		if result {
			return index
		}
	}
	return -1
}

// Every 指定函数检查切片中所有元素，如有一元素不符合条件则表达式返回false，反之
func Every[T comparable](arr []T, cb func(T) bool) bool {
	if utils.IsEmptySlice(arr) {
		return false
	}
	for _, e := range arr {
		result := cb(e)
		if !result {
			return false
		}
	}
	return true
}

// Filter filter
func Filter[T comparable](arr []T, cb func(T) bool) []T {
	if utils.IsEmptySlice(arr) {
		return arr
	}
	var newSlice []T
	for _, ele := range arr {
		cbResult := cb(ele)
		fmt.Println(cbResult)
		if !cbResult {
			continue
		}
		newSlice = append(newSlice, ele)
	}
	return newSlice
}

// ForEach foreach
func ForEach[T comparable](arr []T, cb func(T, int)) {
	for i, t := range arr {
		cb(t, i)
	}
}

// Some for
func Some[T comparable](arr []T, cb func(T) bool) bool {
	if utils.IsEmptySlice(arr) {
		return false
	}
	for _, e := range arr {
		result := cb(e)
		if result {
			return true
		}
	}
	return false
}

// Concat 返回一个连接后的新切片
func Concat[T comparable](arr []T, joinArr []T) []T {
	for _, e := range joinArr {
		arr = append(arr, e)
	}
	return arr
}

// Includes Includes
func Includes[T comparable](arr []T, hitElement T) bool {
	if utils.IsEmptySlice(arr) {
		return false
	}
	for _, t := range arr {
		if t != hitElement {
			continue
		}
		return true
	}
	return false
}

func Pop[T comparable](arr []T) []T {
	newSlice := make([]T, len(arr)-1)
	length := len(arr) - 1
	copy(newSlice, arr)
	return newSlice[:length]
}

// Shift 删除第一个元素返回一个新切片
func Shift[T comparable](arr []T) []T {
	newSlice := make([]T, len(arr))
	copy(newSlice, arr)
	return newSlice[1:]
}

// UnShift 往切片前部追加一个元素
func UnShift[T comparable](arr []T, ele T) []T {
	meteSliceType := reflect.TypeOf(arr)
	meteSlice := reflect.ValueOf(arr)
	newSliceContainer := reflect.MakeSlice(meteSliceType, meteSlice.Len()+1, meteSlice.Cap()+1)
	newSliceContainer.Index(0).Set(reflect.ValueOf(ele))
	openUpRemainingSpace := newSliceContainer.Slice(1, newSliceContainer.Len())
	reflect.Copy(openUpRemainingSpace, meteSlice)
	return newSliceContainer.Interface().([]T)
}

// Splice 删除元素返回一个新的切片
// arr 原切片
// start 删除元素的起始位置
// deleteCount 删除个数
func Splice[T comparable](arr []T, start int, deleteCount int) []T {
	count := deleteCount
	var newSlice []T
	for i, t := range arr {
		if i >= start && count != 0 {
			count--
			continue
		}
		newSlice = append(newSlice, t)
	}
	return newSlice
}

// At 根据 index 获取对应下标的值, 负数这从后查找
func At[T comparable](slice []T, index int) T {
	if index < 0 {
		findSubscriptAfter := len(slice) - int(math.Abs(float64(index)))
		return slice[findSubscriptAfter]
	}
	return slice[index]
}

// Map 方法创建一个新数组,其结果是该数组中的每个元素是调用一次提供的函数后的返回值
func Map[T comparable, standard any](slice []T, cb func(T) standard) []standard {
	var completeSlice []standard
	for _, t := range slice {
		completeSlice = append(completeSlice, cb(t))
	}
	return completeSlice
}

package util

import "math"

//函数式范型
/*
算各个数平方
nums := []int {0,1,2,3,4,5,6,7,8,9}
squares := gMap(nums, func (elem int) int {
  return elem * elem
})
print(squares)  //0 1 4 9 16 25 36 49 64 81
*/
func SMap[T1 any, T2 any](arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
}

/*
求和
nums := []int {0,1,2,3,4,5,6,7,8,9}
sum := gReduce(nums, 0, func (result, elem int) int  {
    return result + elem
})
fmt.Printf("Sum = %d \n", sum)
*/

func SReduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

/*
过滤
nums := []int {0,1,2,3,4,5,6,7,8,9}
odds := gFilterIn(nums, func (elem int) bool  {
    return elem % 2 == 1
})
print(odds)
*/

func SFilter[T any](arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}

// 是否存在
func IsExist[T any](arr []T, f func(T) bool) bool {
	for _, elem := range arr {
		ok := f(elem)
		if ok {
			return true
		}
	}
	return false
}

func IsMapExist[T1 comparable, T2 any](m map[T1]T2, f func(T1, T2) bool) bool {
	for key, elem := range m {
		ok := f(key, elem)
		if ok {
			return true
		}
	}
	return false
}

// 查最大值的项，可能有多个
func SFindMaxElems[T any](arr []T, f func(T) int32) []T {
	switch len(arr) {
	case 0:
		return nil
	case 1:
		return arr
	default:

	}

	var max int32 = math.MinInt32
	for _, v := range arr {
		c := f(v)
		if c > max {
			max = c
		}
	}

	var ret []T
	for _, v := range arr {
		c := f(v)
		if c == max {
			ret = append(ret, v)
		}
	}

	return ret
}

// 查最小值的项，可能有多个
func SFindMinElems[T any](arr []T, f func(T) int32) []T {
	switch len(arr) {
	case 0:
		return nil
	case 1:
		return arr
	default:

	}

	var min int32 = math.MaxInt32
	for _, v := range arr {
		c := f(v)
		if c < min {
			min = c
		}
	}

	var ret []T
	for _, v := range arr {
		c := f(v)
		if c == min {
			ret = append(ret, v)
		}
	}

	return ret
}

// 从map中根据值的f返回值最大项，找key
func MFindMaxKeys[T1 comparable, T2 any](m map[T1]T2, f func(T2) int32) []T1 {
	switch len(m) {
	case 0:
		return nil
	default:

	}

	var max int32 = math.MinInt32
	for _, v := range m {
		c := f(v)
		if c > max {
			max = c
		}
	}

	var ret []T1
	for k, v := range m {
		c := f(v)
		if c == max {
			ret = append(ret, k)
		}
	}

	return ret
}

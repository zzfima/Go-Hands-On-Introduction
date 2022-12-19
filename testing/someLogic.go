package main

type numeric interface {
	int | float64
}

func sumElementsSlice[T numeric](numArray ...T) (res T) {
	res = 0
	for _, v := range numArray {
		res += v
	}
	return
}

func sumElementsArray[T numeric](numArray []T) (res T) {
	res = 0
	for _, v := range numArray {
		res += v
	}
	return
}

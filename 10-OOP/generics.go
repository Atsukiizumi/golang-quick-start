package main

type GenericSlice[T int | int32 | int64] []T

//type GenericNum[T int | int32] T  //错误的，不允许把泛型当做基础类型声明

func Sum[T int | float64](a, b T) T {
	return a + b
}

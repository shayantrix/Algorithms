package main

import "fmt"

func main(){
	x:= []int{64, 34, 25, 12, 22, 11, 90}
	n := len(x)
	QuickSort(x, 0, n - 1)
	
	fmt.Println(x)
}

func Partition(x []int, low, high int) int{
	var pivot int
	pivot = x[high]
	i := low - 1
	temp := 0
	for j := low; j<= high -1; j++{
		if x[j] < pivot {
			i++
			temp = x[j]
			x[j] = x[i]
			x[i] = temp
		}
	}
	//swap pivot and smaller position
	temp = x[i+1]
	x[i+1] = x[high]
	x[high] = temp
	return i + 1
}

func QuickSort(x []int, low, high int){
	if(low < high){
		pi := Partition(x, low, high)
		
		QuickSort(x, low, pi-1)
		QuickSort(x, pi+1, high)
	}
}

func BubbleSort(x []int)[]int{
	n:= len(x)
	for i:= 0; i<n-1; i++{
		temp := 0
		for j:=0;j<n-i-1; j++{
		
			if x[j]>x[j+1]{
				temp = x[j+1]
				x[j+1] = x[j]
				x[j] = temp
			}
		}
	}
	return x
}

func Insertion(x []int) []int{
	n := len(x) 
	for i:=1; i< n; i++{
		key:= x[i]
		j:=i-1
		for ;j>=0 && x[j] > key; j--{
			x[j+1] = x[j]
		}
		x[j+1] = key
	}
	fmt.Println(x)
	return x
}


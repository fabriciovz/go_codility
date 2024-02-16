package sorting_arrays

import "fmt"

// O(n+k)
// n = number of elements in array
// k = range (max-min)
func CountingSort(A []int) []int {
	indexArray := GetIndexArray(A)
	fmt.Println(indexArray)
	prefixSum := GetPrefixSum(indexArray)
	fmt.Println(prefixSum)
	sortedArray := GetSortedArray(A, prefixSum)
	fmt.Println(sortedArray)
	return sortedArray
}

func GetMaxNumberInArray(A []int) int {
	max := 0
	for i := range A {
		if max < A[i] {
			max = A[i]
		}
	}
	return max
}

func GetIndexArray(A []int) []int {
	max := GetMaxNumberInArray(A)
	indexArray := make([]int, max+1)
	for i := 0; i < len(A); i++ {
		indexArray[A[i]]++
	}
	return indexArray
}

func GetPrefixSum(A []int) []int {
	for i := 1; i < len(A); i++ {
		A[i] += A[i-1]
	}
	return A
}

func GetSortedArray(A, prefixSum []int) []int {
	sortedArray := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		valueOriginaArray := A[i]
		valueFromPrefixSumArray := prefixSum[valueOriginaArray]
		sortedArray[valueFromPrefixSumArray-1] = A[i]
		prefixSum[valueOriginaArray]--
	}
	return sortedArray
}

package slices

import "sort"

func AddNames(userList []string, users ...string) []string {
	userList = append(userList, users...)
	return userList
}

func RemoveLastElement(userList []string) []string {
	userList = userList[:len(userList)-1] //[:3]
	return userList
}

func RemoveIndex(userList []string, index int) []string {
	userList = append(userList[:index], userList[index+1:]...)
	return userList
}

func SwapElements(userList []string, index1, index2 int) []string {
	userList[index1], userList[index2] = userList[index2], userList[index1]
	return userList
}

func SortASC(list []int) []int {
	sort.Ints(list)
	//sort.Sort(sort.IntSlice(list[:])) another way to do the same thing
	return list
}

func SortDSC(list []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(list[:])))
	return list
}

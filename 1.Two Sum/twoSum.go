package main

import "fmt"

func twoSum(nums []int, target int) []int {
	nLenght := len(nums)
	var myMap map[int]int
	myMap = make(map[int]int, nLenght)
	for i := 0; i < nLenght; i++ {
		_, ok := myMap[nums[i]]
		if ok {
			return []int{myMap[nums[i]], i}
		} else {
			myMap[target-nums[i]] = i
		}
	}
	return []int{-1, -1} //注意顺序
}
//使用 range 函数
// func twoSum(nums []int, target int) []int {
//     m := make(map[int]int)
//     for i, v := range nums {
//         if j, ok := m[target - v]; ok {
//             return []int{j, i}
//         } else {
//             m[v] = i
//         }
//     }
//     return []int{-1, -1}
// }
func main() {
	b := []int{1, 2, 3, 4, 5, 6}
	a := twoSum(b, 5)
	fmt.Println(a)
}

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func recursiveAll(prefix string , integerList []int) {
	if len(integerList) == 0 {
		return
	}
	if len(integerList) == 1 {
		fmt.Println(fmt.Sprintf("%s%d", prefix, integerList[0]))
		return
	}
	for i, integer := range integerList {
		newPrefix := fmt.Sprintf("%s%d", prefix, integer)
		//fmt.Println(newPrefix)
		newIntegerList := make([]int, len(integerList))
		copy(newIntegerList, integerList)
		newIntegerList = append(newIntegerList[:i], newIntegerList[i+1:]...)
		//fmt.Println(i, newPrefix, newIntegerList)
		recursiveAll(newPrefix, newIntegerList)
	}
}

func lengthOfLongestSubstringLegacy(s string) int {
	lengthOfString := len(s)
	if lengthOfString <= 1 {
		return lengthOfString
	}

	var result int
	for i := 0; i <lengthOfString ; i ++ {
		for j := i + 1; j <= lengthOfString; j ++ {
			if isUniqueString(s[i:j]) {
				result = int(math.Max(float64(j-i), float64(result)))
			}
		}
	}

	return result
}

func isUniqueString(s string) bool {
	fmt.Println(s)
	runeMap := make(map[rune]bool)
	for _, char := range s {
		if _, okay := runeMap[char] ; okay {
			return false
		}
		runeMap[char] = true
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	lengthOfString := len(s)
	if lengthOfString <= 1 {
		return lengthOfString
	}

	runeSlice := []rune(s)
	var i, j ,result int
	runeMap := make(map[rune]bool)
	for ; i <lengthOfString && j < lengthOfString ; {
		runeJ := runeSlice[j]
		if value := runeMap[runeJ] ; !value {
			runeMap[runeJ] = true
			j += 1
			result = int(math.Max(float64(j-i), float64(result)))
		} else {
			runeI := runeSlice[i]
			runeMap[runeI] = false
			i += 1
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	listNum := len(strs)
	if listNum == 0 {
		return ""
	} else if listNum == 1 {
		return strs[0]
	}
	firstStr := strs[0]
	lengthOfFirstStr := len(firstStr)

	lastIndex := 0

	for i := 0; i <= lengthOfFirstStr; i ++ {
		for j := 1; j < listNum; j ++ {
			if len(strs[j]) < i  {
				fmt.Println("<<<<")
				lastIndex = i - 1
				return firstStr[:lastIndex]
			}
			if len(strs[j]) >= i && strs[j][:i] != strs[0][:i] {
				fmt.Println("!!!!!")

				lastIndex = i - 1
				return firstStr[:lastIndex]
			}
		}
		fmt.Println("cccc",i)

		lastIndex = i
	}
	fmt.Println(lastIndex)
	if lastIndex <= 0 {
		return ""
	}
	return firstStr[:lastIndex]

}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Slice := strToMapOcc(s1)
	for i:= 0; i + len(s1) <= len(s2); i ++ {
		subS2 := s2[i:len(s1)+i]
		s2Slice := strToMapOcc(subS2)
		fmt.Println(i, s2Slice, s1Slice, subS2)

		if sliceEqual(s2Slice, s1Slice) {
			return true
		}
	}
	return false
}

func sliceEqual(s1, s2 []int ) bool{
	if len(s1) != len(s2) {
		return false
	}
	for i, value := range s1 {
		if s2[i] != value {
			return false
		}
	}
	return true
}

func strToMapOcc(s string) []int {
	intSlice := make([]int, 26)
	runeSlice := []rune(s)

	for _, runeItem := range runeSlice {
		intSlice[runeItem%26] ++
	}
	return intSlice
}

//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" || num1 == "" || num2 == "" {
//		return "0"
//	}
//	num1Slice := make([]int, 0, len(num1))
//	for _, integerChar := range num1 {
//		num := integerChar - '0'
//		num1Slice = append(num1Slice, int(num))
//	}
//	num2Slice := make([]int, 0, len(num2))
//	for _, integerChar := range num2 {
//		num := integerChar - '0'
//		num1Slice = append(num2Slice, int(num))
//	}
//
//
//
//
//}

func swap(a int, b int) (int, int) {
	return b, a
}

func partition(aris []int, begin int, end int) int {
	pvalue := aris[begin]
	i := begin
	j := end
	for i < j  {
		for aris[j] >= pvalue && i < j {
			j --
		}

		for aris[i] <= pvalue && i < j {
			i ++
		}

		if i < j {
			temp := aris[j]
			aris[j] = aris[i]
			aris[i] = temp
		}
	}
	fmt.Println("before", aris)

	aris[begin] = aris[i]
	aris[i] = pvalue

	fmt.Println("after",aris)
	return i
}

func quick_sort(aris []int, begin int, end int) {
	if begin < end {
		mid := partition(aris, begin, end)
		quick_sort(aris, begin, mid)
		quick_sort(aris, mid+1, end)
	}
}

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x == 1 {
		return 1
	}
	randomX := 4
	for {
		candidate := (randomX + x/randomX)/2
		fmt.Println(candidate)
		doubleCandidate := candidate * candidate
		doubleCandidatePlus := doubleCandidate + candidate * 2 + 1
		if doubleCandidate <= x && doubleCandidatePlus > x {
			return candidate
		}
		randomX = candidate
	}
	return 0
}

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return  num2
	}
	if num2 == "0" {
		return  num1
	}
	result := ""
	num1Slice := make([]int, 0, len(num1))
	for _, integerChar := range num1 {
		num := integerChar - '0'
		num1Slice = append(num1Slice, int(num))
	}
	num2Slice := make([]int, 0, len(num2))
	for _, integerChar := range num2 {
		num := integerChar - '0'
		num2Slice = append(num2Slice, int(num))
	}
	fmt.Println(num1Slice, num2Slice)
	num1I := len(num1Slice)
	num2I := len(num2Slice)
	carry := 0
	for ; num1I > 0 || num2I > 0 ; {
		currentNum := carry
		if num1I > 0 {
			currentNum += num1Slice[num1I-1]
		}
		if num2I > 0 {
			currentNum += num2Slice[num2I-1]

		}
		result = strconv.Itoa(currentNum % 10) + result
		carry = currentNum / 10
		num1I = num1I - 1
		num2I = num2I - 1
		fmt.Println(num1I, num2I, result, carry)
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}

func reverseWords(s string) string {
	arrys := strings.Split(s, " ")
	filteredArray := make([]string, 0, len(arrys))
	for _, item := range arrys {
		if item != "" {
			filteredArray = append(filteredArray, item)
		}
	}
	arryLen := len(filteredArray)
	for i, _ := range filteredArray {
		if i >= (arryLen)/2 {
			break
		}

		temp := filteredArray[arryLen - i - 1 ]
		filteredArray[arryLen - i - 1 ] = filteredArray[i]
		filteredArray[i] = temp

	}
	return strings.Join(filteredArray, " ")
}

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	margin := 0.0
	for _, price := range prices[1:] {
		margin = math.Max(float64(prices[0] - price), margin)
	}
	return int(math.Max(float64(maxProfit(prices[1:])), margin))
}


func maxSubArray(nums []int) int {
	sumArray := make([]int, 0, len(nums))
	maximum := 0
	for i, _ := range nums {
		currentMax := 0
		if i == 0 {
			currentMax = max(nums[i], 0)
		} else {
			currentMax = nums[i] + sumArray[i-1]
			currentMax = max(nums[i], currentMax)
			fmt.Println(currentMax)
			maximum = max(maximum, currentMax)
		}
		sumArray = append(sumArray, currentMax)
	}
	fmt.Println(sumArray)
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func heapSort(arry []int) {
	arrLen := len(arry)
	for i:= 0; i < arrLen; i ++ {
		buildHeap(arry, arrLen - i)
		swapInArray(arry, 0, arrLen - i - 1)
	}
	fmt.Println(arry)
}

func buildHeap(arry []int, arryLen int) {
	for i := arryLen/2 -1; i >= 0; i -- {
		lChildIndex := i * 2 + 1
		rChildIndex := i * 2 + 2

		maxChildIndex := maxIndex(arry, arryLen, lChildIndex, rChildIndex)
		fmt.Println(lChildIndex, rChildIndex, i, maxChildIndex)
		if arry[i] <  arry[maxChildIndex] {
			swapInArray(arry, i, maxChildIndex)
			fmt.Println(arry)
		}
	}
}



func maxIndex(arry []int, arryLen, i, j int ) int {
	if j > arryLen - 1 {
		return i
	}

	if arry[i] > arry[j] {

		return i
	}
	return j
}

func swapInArray(arry []int, i, j int) {
	temp := arry[i]
	arry[i] = arry[j]
	arry[j] = temp
}

func wordBreak(s string, wordDict []string) bool {
	// dp[i]是指对于长度为s[0:i+1]的子段是否满足拆分要求
	dp := make([]bool,len(s)+1)
	dp[0] = true
	mp := make(map[string]int)
	for i,v := range wordDict{
		mp[v] = i
	}
	for i:=1;i<=len(s);i++{

		//遍历字段s[:i]内的子段以确认s[:i]是否能拆分
		for j:=0;j<i;j++{
			if _,ok := mp[s[j:i]];ok&& dp[j]{
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

 //Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	var tempNode *ListNode

	for head != nil {
		tempNode = head
		head = head.Next

		tempNode.Next = newHead
		newHead = tempNode
	}

	return newHead
}

func ThreeSum(sortedList []int) [][]int {
	results := make([][]int, 0, 100)
	for i, candidate := range sortedList {
		targetValue := - candidate
		result := twoSum(targetValue, sortedList[i:])
		if len(result) == 3 {
			results = append(results, result)
		}
	}
	return results
}

func twoSum(target int, sortedList []int) []int{
	lenOfList := len(sortedList)
	for i:= 0; i < lenOfList; i ++ {
		for j := lenOfList -1; ; j -- {
			if sortedList[i] + sortedList[j] == target {
				return []int{target, sortedList[i],  sortedList[j]}
			}
			if i >= j {
				break
			}
		}

	}
	return []int{}
}





type Node struct  {
	ID int
	Price int
}

func quickSort(nodeList []Node, begin int, end int) {
	if begin < end {
		mid := partitionNode(nodeList, begin, end)
		quickSort(nodeList, begin, mid)
		quickSort(nodeList, mid+1, end)
	}
}

func partitionNode(nodeList []Node, begin int, end int) int {
	pvalue := nodeList[begin]
	i := begin
	j := end
	for i < j  {
		for nodeList[j].Price >= pvalue.Price && i < j {
			j --
		}

		for nodeList[i].Price <= pvalue.Price && i < j {
			i ++
		}

		if i < j {
			temp := nodeList[j]
			nodeList[j] = nodeList[i]
			nodeList[i] = temp
		}
	}

	nodeList[begin] = nodeList[i]
	nodeList[i] = pvalue

	return i
}


func findNearestTwo(nodeList []Node, targetPrice int) []Node {
	quickSort(nodeList, 0, len(nodeList))
	result := make([]Node, 0, 2)
	lengthOfArray := len(nodeList)
	if lengthOfArray == 0 {
		return result
	}

	if targetPrice <= nodeList[0].Price {
		result = append(result, nodeList[0])
		return result
	}

	if targetPrice >= nodeList[lengthOfArray-1].Price {
		result = append(result, nodeList[lengthOfArray-1])
		return result
	}


	for i, _ := range nodeList {
		if i + 1 < lengthOfArray && nodeList[i].Price < targetPrice && nodeList[i+1].Price > targetPrice {
			result = append(result, nodeList[i], nodeList[i+1])
			return result
		}
		if nodeList[i].Price == targetPrice {
			result = append(result, nodeList[i])
			if i == 0 {
				return result
			}
			if i == lengthOfArray -1 {
				return result
			}

			lessMargin := targetPrice - nodeList[i-1].Price
			moreMargin :=  nodeList[i + 1].Price - targetPrice
			if moreMargin > lessMargin {
				result = append(result, nodeList[i-1])
			} else {
				result = append(result, nodeList[i+1])
			}
			return result
		}
	}
	return result
}

type TwoSumHashNode struct {
	TwoSumList []int
	MaxIndex int
}

func fourSum(nums []int, target int) [][]int {
	results := make([][]int, 0, 16)
	dedupMap := make(map[string]bool)

	lengthOfNums := len(nums)
	if lengthOfNums < 4 {
		return results
	}

	twoSumHash := make(map[int][]TwoSumHashNode)
	sort.Ints(nums)

	for firstNumI := 0;  firstNumI < lengthOfNums - 1; firstNumI ++ {
		for secondNumI := firstNumI + 1; secondNumI < lengthOfNums; secondNumI ++ {
			currentTwoSum := nums[firstNumI] + nums[secondNumI]
			expectLeftTwoSum := target - currentTwoSum
			if existingTwoSum, exist := twoSumHash[expectLeftTwoSum]; exist {
				for _, expectTwoSum := range existingTwoSum {
					if expectTwoSum.MaxIndex >= firstNumI {
						continue
					}
					temp := []int{expectTwoSum.TwoSumList[0], expectTwoSum.TwoSumList[1], nums[firstNumI] , nums[secondNumI]}
					sort.Ints(temp)

					dedupKey := fmt.Sprintf("%d_%d_%d_%d", temp[0], temp[1], temp[2], temp[3])
					if dedupMap[dedupKey] {
						continue
					}
					dedupMap[dedupKey] = true
					results = append(results, temp)
				}
			}

			newNode := TwoSumHashNode{TwoSumList:[]int{nums[firstNumI] , nums[secondNumI]}, MaxIndex:secondNumI}
			if existingTwoSum, exist := twoSumHash[currentTwoSum]; exist {
				twoSumHash[currentTwoSum] = append(existingTwoSum, newNode)
			} else {
				twoSumHash[currentTwoSum] = []TwoSumHashNode{newNode}
			}
		}
	}
	return results
}

//s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
func main() {
	nums := []int{0,5,-8,-7,4,8,-4,3,9,7,8,10,3,-6,3,7,10,0}
	results := fourSum(nums, -12)
	fmt.Println(results)

	var testArray = []int{1, 9, 3, 3, 5, 6, 7, 4}
	wordDict := []string{"apple", "pen"}
	s := "applepenapple"
	wordBreak(s, wordDict)
	fmt.Println(testArray[0:1])
}

//type TwoSumHashNode struct {
//	TwoSumList []int
//	MaxIndex int
//}
//
//func fourSum(nums []int, target int) [][]int {
//	results := make([][]int, 0, 16)
//	dedupMap := make(map[string]bool)
//
//	lengthOfNums := len(nums)
//	if lengthOfNums < 4 {
//		return results
//	}
//
//	twoSumHash := make(map[int][]TwoSumHashNode)
//	sort.Ints(nums)
//
//	for firstNumI := 0;  firstNumI < lengthOfNums - 1; firstNumI ++ {
//		for secondNumI := firstNumI + 1; secondNumI < lengthOfNums; secondNumI ++ {
//			if secondNumI > firstNumI+1 && nums[secondNumI] == nums[secondNumI-1] {
//				continue
//			}
//
//			currentTwoSum := nums[firstNumI] + nums[secondNumI]
//			expectLeftTwoSum := target - currentTwoSum
//			existingTwoSum, exist := twoSumHash[expectLeftTwoSum]
//			if !exist {
//				newNode := TwoSumHashNode{TwoSumList:[]int{nums[firstNumI] , nums[secondNumI]}, MaxIndex:secondNumI}
//				fmt.Println("newNode",newNode)
//				if existingTwoSum, exist := twoSumHash[currentTwoSum]; exist {
//					twoSumHash[currentTwoSum] = append(existingTwoSum, newNode)
//				} else {
//
//					twoSumHash[currentTwoSum] = []TwoSumHashNode{newNode}
//				}
//				continue
//			}
//			for _, expectTwoSum := range existingTwoSum {
//				if expectTwoSum.MaxIndex >= firstNumI {
//					continue
//				}
//				temp := []int{expectTwoSum.TwoSumList[0], expectTwoSum.TwoSumList[1], nums[firstNumI] , nums[secondNumI]}
//				fmt.Println("expectTwoSum",expectTwoSum, temp, firstNumI)
//				sort.Ints(temp)
//
//				dedupKey := fmt.Sprintf("%d_%d_%d_%d", temp[0], temp[1], temp[2], temp[3])
//				if dedupMap[dedupKey] {
//					continue
//				}
//				dedupMap[dedupKey] = true
//				results = append(results, temp)
//			}
//
//		}
//	}
//	fmt.Println(twoSumHash)
//	return results
//}
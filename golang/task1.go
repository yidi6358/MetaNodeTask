/*只出现一次的数字*/

func singleNumber(nums []int32) int32 {
	mapping := map[rune]rune{}

	for _, single := range nums {
		if mapping[single] == 0 {
			mapping[single] = 1
		} else {
			mapping[single] = 2
		}
		// var matching string
	}
	for _, single1 := range nums {
		if mapping[single1] == 1 {
			return single1
		}
	}

	return 0
}






/*是否回文数*/
func isHuiwen(in int) bool {
	var st string
	st = strconv.Itoa(in)
	fmt.Println(len(st))
	var length = len(st)
	var zhengSt []rune
	zhengSt = make([]rune, length)

	var daoSt []rune
	daoSt = make([]rune, length)

	for i, char := range st {
		zhengSt[i] = char
		daoSt[length-1-i] = char
	}

	if reflect.DeepEqual(zhengSt, daoSt) {
		return true
	}

	return false
}





/*有效的括号*/
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		'(': ')',
		
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		// var matching string
		if matching,isRight := mapping[char]; !isRight {
			if(char == ')'){
				matching = '('	
			}else if(char == '}'){
				matching = '{'
			}else if(char == ']'){
				matching = '['
			}
			

			// 遇到右括号，检查栈顶元素
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matching {
				return false
			}
		} else {
			// 遇到左括号，压入栈中
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}



/*最长公共前缀*/
func longestCommonPrefix(strs []string) string {

	/*最短那个单词的长度*/
	var length int = 100
	for _, single := range strs {
		if len(single) < int(length) {
			length = len(single)
		}
	}
	var qianzhui [3]string
	for i := 0; i < length; i++ {
		var linsh string = ""
		var linshiInt int = 0
		// fmt.Println(linsh)
		for k, single := range strs {
			if k == 0 {
				linsh = single[i : i+1]
			}
			if k > 0 && single[i:i+1] == linsh {
				linshiInt++
				if linshiInt == len(strs)-1 {
					qianzhui[i] = linsh
				}
			}

		}
	}

	var strQian string = ""
	for _, single := range qianzhui {
		if single != "" {
			strQian = strQian + single
		}
	}
	if strQian != "" {
		return strQian
	}
	return ""
}



/*删除排序数组中的重复项*/
func removeDuplicates(nums []int) int {
	maper := map[int]int{}
	var newNum []string = make([]string, len(nums))
	var size = 0
	for i, single := range nums {
		_, exist := maper[single]
		if !exist {
			maper[single] = single
			newNum[i] = strconv.Itoa(single)
			size++
		} else {

		}
	}
	var trueNum []string = make([]string, size)
	var trueIntNum []int = make([]int, size)
	var k = 0
	for i, single := range newNum {
		if single != "" {
			trueNum[i-k] = single
			trueIntNum4, _ := strconv.Atoi(single)
			trueIntNum[i-k] = trueIntNum4
		} else {
			k++
		}
	}
	nums = trueIntNum
	return size
}


/*加1*/
func plusOne(digits []int) []int {
	var str string
	var inS int
	var addIns int
	var sAddIns string
	var arrInt []int
	for _, single := range digits {
		str += strconv.Itoa(single)
	}
	inS, _ = strconv.Atoi(str)
	addIns = inS + 1
	sAddIns = strconv.Itoa(addIns)
	arrInt = make([]int, len(sAddIns))
	for i, singleS := range sAddIns {
		arrInt[i], _ = strconv.Atoi(string(singleS))
	}
	return arrInt

}


/*合并区间*/
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// // 按区间的start升序排序
	// sort.Slice(intervals, func(i, j int) bool {
	// 	return intervals[i][0] < intervals[j][0]
	// })

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			// 有重叠，合并区间
			last[1] = interval[1]
		} else {
			// 无重叠，添加新区间
			merged = append(merged, interval)
		}
	}

	return merged
}


/*两数之和*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

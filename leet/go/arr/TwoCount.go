package main	//声明文件所在的包，每个go文件必须有归属的包
func TwoSum(nums []int, target int) []int {
	var numIndexMap = make(map[int]int)
    for i := range nums {
        var x = nums[i]
        if index, ok := numIndexMap[target - x]; ok {
            return []int{i, index}
        }
        numIndexMap[x] = i
    }
    return []int{}
}
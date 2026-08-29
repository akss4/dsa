func twoSum(nums []int, target int) []int {

    seen := make(map[int]int)
    for i,num := range nums {
        comp  := target - num

        if index, exists := seen[comp]; exists{
        return []int{index,i}}
        seen[num]=i

    }
    return nil
    }
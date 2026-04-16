func twoSum(numbers []int, target int) []int {
    dict := make(map[int]int)
    for i := 0;i < len(numbers); i++ {
        tmp := target - numbers[i]
        if val, exists := dict[tmp]; exists {
            return []int{val, i+1}
        }
        dict[numbers[i]] = i + 1
    }
    return []int{}
}

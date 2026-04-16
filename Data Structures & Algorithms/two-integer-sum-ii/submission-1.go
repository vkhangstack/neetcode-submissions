func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        l, r := i + 1, len(numbers)-1
        tmp := target - numbers[i]
        for l <= r {
            mid := l + (r-l)/2
            if numbers[mid] == tmp {
                return []int{i+1, mid+1}
            } else if numbers[mid] < tmp {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return []int{}
}

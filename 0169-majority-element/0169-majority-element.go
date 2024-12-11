func majorityElement(nums []int) int {
    var ans int
    var count int

    for _, num := range nums {
        if count == 0 {
            ans = num
        }
        if num == ans {
            count++
        } else {
            count--
        }
    }

    return ans
}
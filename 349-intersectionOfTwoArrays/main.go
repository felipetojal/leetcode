package main

func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]bool)
    for _, num := range nums1 {
        _, ok := m[num]
        if !ok {
            m[num] = false
        }
    }
    answer := make([]int, 0)
    for _, num := range nums2 {
        i, ok := m[num]
        if ok && !i {
            answer  = append(answer, num)
            m[num] = true
        } 
    }
    return answer
}
package main

func removeElement(nums []int, val int) int {  
    valid := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            valid++
            continue
        }
        for j := i; j < len(nums); j++ {  
            if nums[j] != val {
                aux := nums[i]
                nums[i] = nums[j]
                nums[j] = aux
                valid++
                break
            }
        }
    }

    return valid
}
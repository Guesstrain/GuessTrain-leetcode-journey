package main

// func maxSubArray(nums []int) int {
//     n := len(nums)
//     anslists := make([]int, n)
//     anslists[0] = nums[0]
//     var sum int
//     if nums[0] > 0{
//         sum = nums[0]
//     }else{
//         sum = 0
//     }
//     for i:=1; i<n; i++{
//         if nums[i] <= 0{
//             if nums[i] > anslists[i-1]{
//                 anslists[i] = nums[i]
//             }else{
//                 anslists[i] = anslists[i-1]
//             }
//             if sum + nums[i] >0{
//                 sum = sum + nums[i]
//             }else{
//                 sum = 0
//             }
//         }else{
//             sum = sum + nums[i]
//             if sum > anslists[i-1]{
//             anslists[i] = sum
//             }else{
//             anslists[i] = anslists[i-1]
//         }
//         }
//     }
//     return anslists[n-1]
// }

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

# Maximum Subarray
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

click to show more practice.

More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

[code](src/maximumSubarray.go)

这一题用暴力循环的办法也能解决，只是要注意边界问题，多写几个特殊用例。看到提示说可以用O(n)的复杂度解决，找了下答案，又是一个动态规划(Dynamic programming, DP)的问题。看来还是逃不过这个坎。

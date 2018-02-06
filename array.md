# array

### In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:
```
Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
```

Example 2:
```
Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
```

Note:
nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99]

这一题本身没有什么算法难度，主要对golang使用切片创建数组不理解,还有对英文题目的读题、审题要仔细。

[code](src/largestNumberAtLeastTwiceOfOthers.go)

### Majority Element
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

**You may assume that the array is non-empty and the majority element always exist in the array.**

这一题由于使用了Golang 中的map, 简化了做题流程。

[code](src/majorityElement.go)

###  Min Cost Climbing Stairs
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:
```
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
```

Example 2:
```
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
```
Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].

这一题以前好像见过，但是逃避了，现在想想，逃避是没有用的，自己用暴力解了好久，还是会漏掉好多场景，[错误代码](src/minCostClimbingStairs/minCostClimbingStairsError.go)

后来[参考了网上的代码](http://blog.csdn.net/Next_Second/article/details/78861839).
这是一个动态规划(Dynamic programming, DP)的问题。

如果我们用一个数组dp[]来存放到达每一层所需要的花费值。则则最终结果是求dp[cost.length]的值。
因为每次可以走1层或者2层，并且可以从0或者从1开始，所以可以得到dp[0]为0，dp[1]为0。
从2开始，dp[i]可以由通过dp[i-2]走2层或者通过dp[i-1]走一层到达，而这i-2和i-1层所要花费的值分别为cost[i-2]和cost[i-1]，所以，dp[i] = min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])。该算法的空间复杂度为O(n)，时间复杂度为O(n)。

[code](src/minCostClimbingStairs/minCostClimbingStairs.go)


### Climbing Stairs

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.


Example 1:

```
Input: 2
Output:  2
Explanation:  There are two ways to climb to the top.

1. 1 step + 1 step
2. 2 steps
```

Example 2:

```
Input: 3
Output:  3
Explanation:  There are three ways to climb to the top.

1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

[code](src/climbingStairs.go)

这一题与`Min Cost Climbing Stairs`基本是同一个类型的题，有了之前的经验，这一题就块多了，但是用第一种解法（递归求和），需要注意超时问题。

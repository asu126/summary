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

### 169-Majority Element
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

- 283. Move Zeroes


- 4. 两个排序数组的中位数
```
有两个大小为 m 和 n 的排序数组 nums1 和 nums2 。

请找出两个排序数组的中位数并且总的运行时间复杂度为 O(log (m+n)) 。

示例 1:

nums1 = [1, 3]
nums2 = [2]

中位数是 2.0
```

这一题只有时间复杂度的要求，没有空间复杂度的要求，有两种解题思路：
1. 将两个数组合二为一；
2. 有点类似两个有序链表合并的思路，找长度为length1 + length 的中间及其前一个元素，如果总数为基数，返回中间一个数据；如果为偶数中间元素及其前一个元素的平均值。

[code](src/climbingStairs.go)

- 27
- 35

- 48. 旋转图像

这一题首先想到的转圈解法，但是参考的网上的解法，可以用矩阵的对称性来解。

同事，这一题与有次面试题相似，给定一个n，输出如下n*n的二位数组。
```
1,2,3
8,9,4
7,6,5
```

这一题尝试了矩阵思路，但是好像不对，只有用转圈解法了。[代码见](src/pdd_1.go)


- 56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

```
示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
```

注意点：
  1. 判断区间是否重合的时候可以反方向思考，什么时候不重合
  2. 排序这个特殊情况的考虑

- 15. 三数之和
```
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]```
注意：1.不重复的三元组；2.符合条件时注意陷乳死循环；3. 不知道是不是排序算法写的效率较低，自己写的提交了两次超时了，用了库函数里的开始几次也超时了。

- 16. 最接近的三数之和
```
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
```
排序后去查找，否则需要n^3的时间复杂度。


- 238. 除自身以外数组的乘积
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
```
输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题
```

解法一：除法
解法二：N^2 循换
解法三： 
```
用数组B[]来表示A的一个后缀数组，即
B[N]=1, B[N-1]=A[N], B[N-2]=A[N-1]*A[N],....,B[2]=A[3]*A[4]*...*A[N], B[1]=A[2]*A[3]*...*A[N]
然后把A[]变成一个前缀数组，即A[1]=1,A[2]=A[1], A[3]=A[1]*A[2], A[N]=A[1]*A[2]*...*A[N-1].
然后B[i]=A[i]*B[i], 然后最后的B即为所求
```
解法三 最后还是超时了。。

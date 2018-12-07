package leetcode;

import java.util.ArrayList;
import java.util.List;

public class Solution_859 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_859 ss = new Solution_859();

		// 示例 1：
		// 输入： A = "ab", B = "ba"
		// 输出： true
		System.out.println(ss.buddyStrings("ab", "ba"));

		// 示例 2：
		// 输入： A = "ab", B = "ab"
		// 输出： false
		System.out.println(ss.buddyStrings("ab", "ab"));

		// 示例 3:
		// 输入： A = "aa", B = "aa"
		// 输出： true
		System.out.println(ss.buddyStrings("aa", "aa"));

		// 示例 4：
		// 输入： A = "aaaaaaabc", B = "aaaaaaacb"
		// 输出： true
		System.out.println(ss.buddyStrings("aaaaaaabc", "aaaaaaacb"));

		// 示例 5：
		// 输入： A = "", B = "aa"
		// 输出： false
		System.out.println(ss.buddyStrings("", "aa"));
	}

	public boolean buddyStrings(String A, String B) {
		int lengthA = A.length();
		int lengthB = B.length();

		if (lengthA != lengthB) {
			return false;
		}
		char arrayA[] = A.toCharArray();
		char arrayB[] = B.toCharArray();
		List<Integer> list = new ArrayList<Integer>();
		int array[] = new int[26];

		for (int i = 0; i < lengthA; i++) {
			if (arrayA[i] != arrayB[i]) {
				list.add(i);
			}
			int index = arrayA[i] - 'a';
			array[index]++;
		}

		int diffLenth = list.size();

		// 只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true
		// 一定要交换两个字符
		if (diffLenth == 0) {
			for (int i = 0; i < 26; i++) {
				if (array[i] >= 2) {
					return true;
				}
			}
			return false;
		}
		if (diffLenth == 2 && arrayA[list.get(0)] == arrayB[list.get(1)]
				&& arrayA[list.get(1)] == arrayB[list.get(0)]) {
			return true;
		} else {
			return false;
		}
	}
}

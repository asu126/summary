package leetcode;

import java.util.ArrayList;
import java.util.List;

public class Solution_821 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_821 ss = new Solution_821();
		// 输入: S = "loveleetcode", C = 'e'
		// 输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
		System.out.println(ss.shortestToChar("loveleetcode", 'e'));

	}

	public int[] shortestToChar(String S, char C) {
		char array[] = S.toCharArray();
		int length = S.length();
		int result[] = new int[length];
		List<Integer> list = new ArrayList<Integer>();

		for (int i = 0; i < length; i++) {
			if (array[i] == C) {
				list.add(i);
			}
		}
		for (int i = 0; i < length; i++) {
			if (array[i] == C) {
				result[i] = 0;
			} else {
				int min = Integer.MAX_VALUE;
				for (Integer a : list) {
					if (Math.abs(a - i) < min) {
						min = Math.abs(a - i);
						// System.out.println(Math.abs(a));
					}

				}
				result[i] = min;
			}
			// System.out.println(Math.abs(result[i]));
		}

		return result;

	}

}

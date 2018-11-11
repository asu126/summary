package leetcode;

public class Solution_905 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_905 ss = new Solution_905();
		int a[] = { 3, 1, 2, 4 };
		System.out.println(ss.sortArrayByParity(a));

		int b[] = { 0, 0, 0, 0, 0, 0 };
		System.out.println(ss.sortArrayByParity(b));

	}

	public int[] sortArrayByParity(int[] A) {
		int length = A.length;

		int result[] = new int[length];
		int oushu = 0;
		int jishu = length - 1;

		for (int i = 0; i < length; i++) {
			if (A[i] % 2 == 0) {
				result[oushu] = A[i];
				oushu++;
			} else {
				result[jishu] = A[i];
				jishu--;
			}
		}

		return result;

	}

}

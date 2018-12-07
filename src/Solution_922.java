package leetcode;

public class Solution_922 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_922 ss = new Solution_922();
		// 输入：[4,2,5,7]
		// 输出：[4,5,2,7]
		// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受
		int b[] = { 4, 2, 5, 7 };
		int out[] = ss.sortArrayByParityII(b);
		for (int i = 0; i < out.length; i++) {
			System.out.println(out[i]);

		}
	}

	public int[] sortArrayByParityII(int[] A) {
		int length = A.length;
		int result[] = new int[length];

		int oushu = 0;
		int jishu = 1;
		for (int i = 0; i < length; i++) {
			if (A[i] % 2 == 0 && oushu < length) {
				if (oushu < length) {
					result[oushu] = A[i];
					oushu += 2;
				}
			} else {
				if (jishu < length) {
					result[jishu] = A[i];
					jishu += 2;
				}

			}
		}
		return result;
	}

}

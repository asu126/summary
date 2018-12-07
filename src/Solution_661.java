package leetcode;

public class Solution_661 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_661 ss = new Solution_661();
		// 输入:
		// [[1,1,1],
		// [1,0,1],
		// [1,1,1]]
		// 输出:
		// [[0, 0, 0],
		// [0, 0, 0],
		// [0, 0, 0]]
		int a[][] = { { 1, 1, 1 }, { 1, 0, 1 }, { 1, 1, 1 } };
		int b[][] = ss.imageSmoother(a);

		for (int i = 0; i < b.length; i++) {
			for (int j = 0; j < b[0].length; j++) {
				System.out.println(b[i][j]);
			}
		}
	}

	public int[][] imageSmoother(int[][] M) {
		int row = M.length;
		int column = M[0].length; // 矩阵的长和宽的范围均为 [1, 150]。
		int result[][] = new int[row][column];

		for (int i = 0; i < row; i++) {
			for (int j = 0; j < column; j++) {
				int down = i;
				if (i - 1 >= 0) {
					down = i - 1;
				}
				int up = i;
				if (i + 1 < row) {
					up = i + 1;
				}
				int left = j;
				if (j - 1 >= 0) {
					left = j - 1;
				}
				int rigth = j;
				if (j + 1 < column) {
					rigth = j + 1;
				}

				int sum = 0;
				int count = 0;
				for (int m = down; m <= up; m++) {
					for (int n = left; n <= rigth; n++) {
						sum += M[m][n];
						count++;

					}
				}

				result[i][j] = sum / count;

			}

		}
		return result;
	}
}

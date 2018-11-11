package leetcode;

public class Solution_832 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		// TODO Auto-generated method stub
		Solution_832 ss = new Solution_832();
		// 输入: [[1,1,0],[1,0,1],[0,0,0]]
		// 输出: [[1,0,0],[0,1,0],[1,1,1]]
		int a[][] = { { 1, 1, 0 }, { 1, 0, 1 }, { 0, 0, 0 } };
		int b[][] = ss.flipAndInvertImage(a);

		for (int i = 0; i < b.length; i++) {
			for (int j = 0; j < b[0].length; j++) {
				System.out.println(b[i][j]);
			}
		}

	}

	public int[][] flipAndInvertImage(int[][] A) {
		int row = A.length;
		int line = A[0].length;
		int result[][] = new int[line][row];

		// System.out.println(row + " " + line);

		for (int i = 0; i < line; i++) {
			for (int j = 0; j < row; j++) {
				result[i][j] = A[i][line - 1 - j];
				if (result[i][j] == 1) {
					result[i][j] = 0;
				} else {
					result[i][j] = 1;
				}
			}
		}

		return result;
	}
}

package leetcode;

public class Solution_867 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_867 ss = new Solution_867();
		// 输入：[[1,2,3],[4,5,6],[7,8,9]]
		// 输出：[[1,4,7],[2,5,8],[3,6,9]]
		// int a[][] = {{1,2,3},{4,5,6},{7,8,9}};
		int a[][] = { { 1, 2, 3 } };
		int b[][] = ss.transpose(a);

		for (int i = 0; i < b.length; i++) {
			for (int j = 0; j < b[0].length; j++) {
				System.out.println(b[i][j]);
			}
		}
	}

	public int[][] transpose(int[][] A) {
		int row = A.length;
		int line = A[0].length;
		int result[][] = new int[line][row];

		// System.out.println(row + " " + line);

		for (int i = 0; i < line; i++) {
			for (int j = 0; j < row; j++) {
				result[i][j] = A[j][i];
			}
		}

		return result;
	}

}

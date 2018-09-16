package leetcode;

public class Solution_766 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int array[][] = { { 1, 2, 3, 4 }, { 5, 6, 7, 8 }, { 9, 10, 11, 12 } };
		System.out.println(new Solution_766().isToeplitzMatrix(array));

		int array1[][] = { { 1, 2, 3, 4 }, { 5, 1, 2, 3 }, { 9, 5, 1, 2 } };
		System.out.println(new Solution_766().isToeplitzMatrix(array1));

		int array2[][] = { { 1 } };
		System.out.println(new Solution_766().isToeplitzMatrix(array2));

		int array3[][] = { { 1, 2, 3, 4 } };
		System.out.println(new Solution_766().isToeplitzMatrix(array3));

		int array4[][] = { { 1 }, { 2 }, { 3 } };
		System.out.println(new Solution_766().isToeplitzMatrix(array4));
	}

	public boolean isToeplitzMatrix(int[][] matrix) {
		int row = matrix.length;
		if (row < 1) {
			return false;
		}
		int cloumn = matrix[0].length;

		boolean result = true;
		// 左上角,row=0
		for (int i = 0; i < cloumn; i++) {
			int first = matrix[0][i];
			for (int x = 1, y = i + 1; x < row && y < cloumn; x++, y++) {
				if (matrix[x][y] != first) {
					result = false;
					break;
				}
			}
		}
		// 右下角, cloumn=0
		for (int i = 1; i < row; i++) {
			int first = matrix[i][0];
			for (int x = i + 1, y = 1; x < row && y < cloumn; x++, y++) {
				if (matrix[x][y] != first) {
					result = false;
					break;
				}
				// System.out.println(matrix[x][y]);
			}
		}
		return result;
	}

}

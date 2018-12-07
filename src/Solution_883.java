package leetcode;

public class Solution_883 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_883 ss = new Solution_883();

		// 输入：[[2]]
		// 输出：5
		int b[][] = { { 2 } };
		System.out.println(ss.projectionArea(b));

		// 输入：[[1,2],[3,4]]
		// 输出：17
		int c[][] = { { 1, 2 }, { 3, 4 } };
		System.out.println(ss.projectionArea(c));

		// 输入：[[1,0],[0,2]]
		// 输出：8
		int d[][] = { { 1, 0 }, { 0, 2 } };
		System.out.println(ss.projectionArea(d));

		// 输入：[[1,1,1],[1,0,1],[1,1,1]]
		// 输出：14
		int a[][] = { { 1, 1, 1 }, { 1, 0, 1 }, { 1, 1, 1 } };
		System.out.println(ss.projectionArea(a));

		// 输入：[[2,2,2],[2,1,2],[2,2,2]]
		// 输出：21
		int e[][] = { { 2, 2, 2 }, { 2, 1, 2 }, { 2, 2, 2 } };
		System.out.println(ss.projectionArea(e));
	}

	public int projectionArea(int[][] grid) {
		int row = grid.length;
		int column = grid[0].length; // 1 <= grid.length = grid[0].length <= 50
		int result = 0;

		// xy
		int xy = 0;
		for (int i = 0; i < row; i++) {
			for (int j = 0; j < column; j++) {
				if (grid[i][j] > 0) {
					xy++;
				}
			}
		}
		result += xy;

		// yz
		int yz = 0;
		for (int i = 0; i < column; i++) {
			int max = 0;
			for (int j = 0; j < row; j++) {
				max = Math.max(max, grid[i][j]);
			}
			yz += max;
		}
		result += yz;

		// xz
		int xz = 0;
		for (int i = 0; i < row; i++) {
			int max = 0;
			for (int j = 0; j < column; j++) {
				max = Math.max(max, grid[j][i]);
			}
			xz += max;
		}
		result += xz;

		return result;
	}

}

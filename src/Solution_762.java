package leetcode;

public class Sulution_762 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Sulution_762 ss = new Sulution_762();
		
		System.out.println((ss.countPrimeSetBits(10,15)));
	}

	public int countPrimeSetBits(int L, int R) {
		int result = 0;
		for (int i = L; i <= R; i++) {
			if (checkPrimeSetBits(i)) {
				result++;
			}
		}

		return result;
	}

	private boolean checkPrimeSetBits(int x) {
		int zhishu[] = { 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31 };
		int count = 0;
		int tmp = x;

		for (int i = 0; i < 31; i++) {
			if ((tmp & 1) == 1) {
				count++;
			}
			tmp = tmp >> 1;
		}

		for (int i = 0; i < zhishu.length; i++) {
			if (zhishu[i] == count) {
				return true;
			}
		}

		return false;
	}

}

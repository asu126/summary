package leetcode;

public class Solution_788 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_788 ss = new Solution_788();
		System.out.println(ss.rotatedDigits(10));
	}

	public int rotatedDigits(int N) {
		int count = 0;
		for (int i = 1; i <= N; i++) {
			if (haoshu(i)) {
				count++;
			}
		}
		return count;
	}

	private boolean haoshu(int x) {
		boolean result = false;
		int mod;
		int shang = x;
		while (shang != 0) {
			mod = shang % 10;
			shang = shang / 10;
			if (mod == 0 || mod == 1 || mod == 8) {
				continue;
			} else if (mod == 2 || mod == 5 || mod == 6 || mod == 9) {
				result = true;
			} else {
				return false;
			}
		}

		return result;
	}

}

package leetcode;

import java.util.ArrayList;
import java.util.List;

public class Solution_728 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_728 ss = new Solution_728();
		System.out.println(ss.selfDividingNumbers(1, 22));
	}

	public List<Integer> selfDividingNumbers(int left, int right) {
		List<Integer> list = new ArrayList<Integer>();
		for (int i = left; i <= right; i++) {
			if (checkNumber(i)) {
				list.add(i);
			}
		}
		return list;
	}

	private boolean checkNumber(int x) {
		int shang = x / 10;
		int mod = x % 10;

		while (shang != 0) {
			if (mod == 0 || x % mod != 0) {
				return false;
			}
			mod = shang % 10;
			shang = shang / 10;
		}
		if (mod != 0 && x % mod == 0) {
			return true;
		}
		return false;
	}

}

package leetcode;

public class Solution_693 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_693 ss = new Solution_693();
		System.out.println(ss.hasAlternatingBits(5)); // true
		System.out.println(ss.hasAlternatingBits(7)); // false
		System.out.println(ss.hasAlternatingBits(11)); // false
		System.out.println(ss.hasAlternatingBits(10)); // true
	}

	public boolean hasAlternatingBits(int n) {
		int[] bytes = new int[31];
		int xx = 1;
		for (int i = 0; i < 31; i++) {
			bytes[i] = n & xx;
			// System.out.print(bytes[i]);
			n = n >> 1;
		}

		int i = 0;
		for (i = 30; i >= 0; i--) {
			if (bytes[i] != 0) {
				break;
			}
		}

		for (; i >= 1; i--) {
			// System.out.println(i);
			if (bytes[i] == bytes[i - 1]) {
				return false;
			}
		}

		return true;
	}

}

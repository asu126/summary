package leetcode;

public class Solution_868 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_868 ss = new Solution_868();
		System.out.println(ss.binaryGap(22)); // 2

		System.out.println(ss.binaryGap(5)); // 2

		System.out.println(ss.binaryGap(6)); // 1

		System.out.println(ss.binaryGap(8)); // 0
	}

	public int binaryGap(int N) {
		int bits[] = new int[31];
		int tmp = N;
		int start = 0;
		int maxResult = 0;

		for (int i = 0; i < 31; i++) {
			bits[i] = (tmp & 1);
			tmp = tmp >> 1;
		}
		int location;
		for (location = 0; location < 31; location++) {
			if (bits[location] == 1) {
				start = location;
				location++;
				break;
			}
		}

		for (; location < 31; location++) {
			if (bits[location] == 1) {
				maxResult = (location - start) > maxResult ? location - start : maxResult;
				start = location;
			}
		}

		return maxResult;
	}

}

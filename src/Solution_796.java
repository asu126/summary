package leetcode;

public class Solution_796 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_796 ss = new Solution_796();
		System.out.println(ss.rotateString("abcde", "abcde"));
		System.out.println(ss.rotateString("abcde", "cdeab"));
		System.out.println(ss.rotateString("abcde", "abced"));
	}

	public boolean rotateString(String A, String B) {
		char inputsA[] = A.toLowerCase().toCharArray();
		char inputsB[] = B.toLowerCase().toCharArray();
		int length = inputsB.length;
		if (inputsA.length != inputsB.length) {
			return false;
		}

		boolean result = true;
		for (int mv = 0; mv < length; mv++) {
			result = true;
			for (int i = 0; i < length; i++) {

				if ((i + mv) < length) {
					if (inputsA[i] != inputsB[i + mv]) {
						result = false;
						break;
					}
				} else {
					int lo = (i + mv) % length;
					if (inputsA[i] != inputsB[lo]) {
						result = false;
						break;
					}
				}

			}

			if (result) {
				return result;
			}

		}
		return result;
	}
}

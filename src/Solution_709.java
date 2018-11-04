package leetcode;

public class Solution_709 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

		Solution_709 aa = new Solution_709();
		System.out.println((char) 97);
		System.out.println((int) 'a');
		System.out.println("start:");
		System.out.println(aa.toLowerCase("abc"));
		System.out.println(aa.toLowerCase("ABC"));
	}

	public String toLowerCase(String str) {
		byte[] array = str.getBytes();
		StringBuilder sb = new StringBuilder();

		for (int i = 0; i < array.length; i++) {
			if (array[i] >= 'A' && array[i] <= 'Z') {
				sb.append((char) (array[i] + 32));
			} else {
				sb.append((char) array[i]);
			}

		}

		// System.out.println(sb.length());
		return sb.toString();
	}

}

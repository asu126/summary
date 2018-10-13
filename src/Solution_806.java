package leetcode;

public class Solution_806 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_806 cc = new Solution_806();
		int[] widths = { 4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
				10, 10 };

		String S = "bbbcccdddaaa";
		cc.numberOfLines(widths, S);

		System.out.println(Math.ceil(10.7));
		System.out.println(Math.ceil(0.98));
	}

	public int[] numberOfLines(int[] widths, String S) {
		int result[] = { 0, 0 };
		int total = 0;
		char inputs[] = S.toLowerCase().toCharArray();

		for (int i = 0; i < inputs.length; i++) {
			int current = widths[inputs[i] - 'a'];
			if ((current + total) <= ((int) Math.ceil(total / 100.0)) * 100) {
				total += current;
			} else {
				// System.out.println("----------" + total + "----" + Math.ceil(total / 100.0)
				// );
				total = (int) (Math.ceil(total / 100.0)) * 100 + current;
			}
			// System.out.println(inputs[i] + " " + widths[inputs[i] - 'a']);
		}

		result[0] = (int) Math.ceil(total / 100.0);
		result[1] = total % 100;
		// System.out.println(total + " " + result[0] + " " + result[1]);
		return result;
	}
}

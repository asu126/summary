package leetcode;

public class Solution_844 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

		Solution_844 ss = new Solution_844();
		// 输入：S = "ab#c", T = "ad#c"
		// 输出：true
		System.out.println(ss.backspaceCompare("ab#c", "ad#c"));
		// 输入：S = "ab##", T = "c#d#"
		// 输出：true
		// 解释：S 和 T 都会变成 “”。
		System.out.println(ss.backspaceCompare("ab##", "c#d#"));
		// 示例 3：
		//
		// 输入：S = "a##c", T = "#a#c"
		// 输出：true
		// 解释：S 和 T 都会变成 “c”。
		System.out.println(ss.backspaceCompare("a##c", "#a#c"));
		// 示例 4：
		//
		// 输入：S = "a#c", T = "b"
		// 输出：false
		// 解释：S 会变成 “c”，但 T 仍然是 “b”。
		 System.out.println(ss.backspaceCompare("a#c", "b"));

	}

	public boolean backspaceCompare(String S, String T) {

		String resultS = clearStr(S);

		String resultT = clearStr(T);
		// System.out.println(resultS + "==" + resultT);
		return resultS.equals(resultT);
	}

	private String clearStr(String x) {
		byte[] ss = x.getBytes();
		StringBuilder sbs = new StringBuilder();
		int sub = 0;
		for (int i = ss.length - 1; i >= 0; i--) {
			// System.out.println((char) ss[i]);
			if ((char) ss[i] == '#') {
				sub++;
			} else {
				if (sub > 0) {
					sub--;
					continue;
				}

				sbs.append((char) ss[i]);
			}

		}
		return sbs.toString();
	}

}

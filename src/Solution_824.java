package leetcode;

public class Solution_824 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		// 输入: "I speak Goat Latin"
		// 输出: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
		// 示例 2:
		//
		// 输入: "The quick brown fox jumped over the lazy dog"
		// 输出: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa
		// hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
		Solution_824 ss = new Solution_824();
		System.out.println(ss.toGoatLatin("I speak Goat Latin"));

		System.out.println(ss.toGoatLatin("The quick brown fox jumped over the lazy dog"));

		System.out.println(ss.toGoatLatin(""));
	}

	public String toGoatLatin(String S) {
		if (S.equals("")) {
			return "";
		}

		String subS[] = S.split(" ");
		StringBuilder sb = new StringBuilder();
		for (int i = 0; i < subS.length; i++) {
			if (i != 0) {
				sb.append(' ');
			}

			char cahrs[] = subS[i].toCharArray();
			char first = cahrs[0];
			if (first == 'a' || first == 'A' || first == 'E' || first == 'e' || first == 'i' || first == 'I'
					|| first == 'O' || first == 'o' || first == 'u' || first == 'U') {
				sb.append(cahrs);
			} else {
				char tmp = cahrs[0];
				for (int j = 0; j < cahrs.length - 1; j++) {
					cahrs[j] = cahrs[j + 1];
				}

				cahrs[cahrs.length - 1] = tmp;
				sb.append(cahrs);
				// System.out.println(cahrs);
			}

			sb.append("ma");
			for (int j = 0; j < i + 1; j++) {
				sb.append('a');
			}
		}

		return sb.toString();
	}

}

package leetcode;

import java.util.HashMap;
import java.util.Map;

public class Solution_819 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String input = "abc abc? abcd the jeff!";
		String[] banned = { "abc","abcd","jeff" };
		System.out.println(new Solution_819().mostCommonWord(input, banned));
	}

	public String mostCommonWord(String paragraph, String[] banned) {
		int start = 0;
		char inputs[] = paragraph.toLowerCase().trim().toCharArray();
		int length = inputs.length;

		Map<String, Integer> m1 = new HashMap<String, Integer>();
		m1.clear();
		for (; start < length;) {
			for (int end = start; end < length; end++) {
				// !?',;.
				if (inputs[end] == ' ' || inputs[end] == '!' || inputs[end] == '?' || inputs[end] == '\''
						|| inputs[end] == ',' || inputs[end] == ';' || inputs[end] == '.') {
					if (start == end) {
						start += 1;
					} else {
						String key = new String(inputs, start, end - start);
						if (m1.get(key) != null) {
							m1.put(key, (m1.get(key)) + 1);
						} else {
							m1.put(key, 1);
						}

						start = end + 1;
					}
					break;
				}
				if (end == length - 1) {
					String key = new String(inputs, start, end - start + 1);
					if (m1.get(key) != null) {
						m1.put(key, (m1.get(key)) + 1);
					} else {
						m1.put(key, 1);
					}

					start = end + 1;
				}
			}
		}

		// System.out.println(m1);
		String result = "";
		Integer temp = 0;
		for (Map.Entry<String, Integer> entry : m1.entrySet()) {
			boolean flag = false;
			for (int i = 0; i < banned.length; i++) {
				if (entry.getKey().equals(banned[i])) {
					flag = true;
					break;
				}
			}

			if (!flag && entry.getValue() > temp) {
				temp = entry.getValue();
				result = entry.getKey();
			}
			// System.out.println(entry.getKey() + " " + entry.getValue());
		}
		return result;
	}

}

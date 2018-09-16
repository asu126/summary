public class Solution {

	public int uniqueMorseRepresentations(String[] words) {
		Set hs = new HashSet();
		String[] morse = { ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
				"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.." };

		for (int i = 0; i < words.length; i++) {
			StringBuilder stringBuilder = new StringBuilder();
			for (int j = 0; j < words[i].length(); j++) {
				int index = (int) words[i].charAt(j) - 'a';
				stringBuilder.append(morse[index]);
			}
			hs.add(stringBuilder.toString());
		}
//		System.out.println(hs);
		return hs.size();
	}
}

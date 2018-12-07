package leetcode;

public class Solution_917 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_917 ss = new Solution_917();
		// 示例 1：
		// 输入："ab-cd"
		// 输出："dc-ba"
		System.out.println(ss.reverseOnlyLetters("ab-cd"));

		// 示例 2：
		// 输入："a-bC-dEf-ghIj"
		// 输出："j-Ih-gfE-dCba"
		System.out.println(ss.reverseOnlyLetters("a-bC-dEf-ghIj"));

		// 示例 3：
		// 输入："Test1ng-Leet=code-Q!"
		// 输出："Qedo1ct-eeLg=ntse-T!"

		System.out.println(ss.reverseOnlyLetters("Test1ng-Leet=code-Q"));
		
		System.out.println(ss.reverseOnlyLetters("7_28]"));
	}

	public String reverseOnlyLetters(String S) {
		// int length = S.length();
		char array[] = S.toCharArray();
		// StringBuilder sb = new StringBuilder();
		int i = 0;
		int j = S.length() - 1;
		while (i < j) {
			while (!isZiMu(array[i]) && i < j) {
				i++;
			}

			while (!isZiMu(array[j]) && i < j) {
				j--;
			}

			if (i < j ) {
				char tmp = array[i];
				array[i] = array[j];
				array[j] = tmp;
				i++;
				j--;
			}
		}

		return new String(array);
	}

	private boolean isZiMu(char c) {
		if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return true;
		} else {
			return false;
		}
	}
}

package leetcode;

public class Solution_682 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_682 ss = new Solution_682();
		// 示例 1:
		//
		// 输入: ["5","2","C","D","+"]
		// 输出: 30
		// 解释:
		// 第1轮：你可以得到5分。总和是：5。
		// 第2轮：你可以得到2分。总和是：7。
		// 操作1：第2轮的数据无效。总和是：5。
		// 第3轮：你可以得到10分（第2轮的数据已被删除）。总数是：15。
		// 第4轮：你可以得到5 + 10 = 15分。总数是：30。
		// 示例 2:
		//
		// 输入: ["5","-2","4","C","D","9","+","+"]
		// 输出: 27
		// 解释:
		// 第1轮：你可以得到5分。总和是：5。
		// 第2轮：你可以得到-2分。总数是：3。
		// 第3轮：你可以得到4分。总和是：7。
		// 操作1：第3轮的数据无效。总数是：3。
		// 第4轮：你可以得到-4分（第三轮的数据已被删除）。总和是：-1。
		// 第5轮：你可以得到9分。总数是：8。
		// 第6轮：你可以得到-4 + 9 = 5分。总数是13。
		// 第7轮：你可以得到9 + 5 = 14分。总数是27
		String a[] = new String[] { "5", "2", "C", "D", "+" };
		System.out.println(ss.calPoints(a));
		String b[] = new String[] { "5", "-2", "4", "C", "D", "9", "+", "+" };
		System.out.println(ss.calPoints(b));
	}

	public int calPoints(String[] ops) {
		int recornd[] = new int[1001];
		int location = 0;
		int result = 0;

		for (int i = 0; i < ops.length; i++) {
			// 1.整数（一轮的得分）：直接表示您在本轮中获得的积分数。
			// 2. "+"（一轮的得分）：表示本轮获得的得分是前两轮有效 回合得分的总和。
			// 3. "D"（一轮的得分）：表示本轮获得的得分是前一轮有效 回合得分的两倍。
			// 4. "C"（一个操作，这不是一个回合的分数）：表示您获得的最后一个有效 回合的分数
			if ("+".equals(ops[i])) {
				if (location > 1) {
					recornd[location] = recornd[location - 1] + recornd[location - 2];
					result += recornd[location];
					location++;
				}
			} else if ("D".equals(ops[i])) {
				if (location > 0) {
					recornd[location] = 2 * recornd[location - 1];
					result += recornd[location];
					location++;
				}
			} else if ("C".equals(ops[i])) {
				if (location > 0) {
					result -= recornd[location - 1];
					location--;
				}
			} else {
				Integer rec = Integer.parseInt(ops[i]);
				recornd[location] = rec;
				location++;
				result += rec;
			}
		}

		return result;
	}

}

package leetcode;

import java.util.LinkedList;
import java.util.List;

class Employee {
	// It's the unique id of each node;
	// unique id of this employee
	public int id;
	// the importance value of this employee
	public int importance;
	// the id of direct subordinates
	public List<Integer> subordinates;
};

public class Solution_690 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}

	public int getImportance(List<Employee> employees, int id) {
		LinkedList<Integer> list = new LinkedList<Integer>();
		int length = employees.size();
		int result = 0;
		list.add(id);

		while (!list.isEmpty()) {
			int tmp = list.pop();
			for (int i = 0; i < length; i++) {
				Employee em = employees.get(i);
				if (em.id == tmp) {
					result += em.importance;
					list.addAll(em.subordinates);
					break;
				}
			}
		}

		return result;
	}

}

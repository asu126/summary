package leetcode;

public class Sulution_704 {
	public static void main(String[] args) {
		Sulution_704 s = new Sulution_704();
		System.out.println(s.search(new int[] { 1, 3, 5 }, 3));
		System.out.println(s.search(new int[] { }, 3));

		System.out.println(s.search(new int[] { -1, 0, 3, 5, 9, 12 }, 9));
		// nums = [-1,0,3,5,9,12], target = 2
		System.out.println(s.search(new int[] { -1, 0, 3, 5, 9, 12 }, 2));
		System.out.println(s.search(new int[] {5}, 5));
	}

	public int search(int[] nums, int target) {
		int result = -1;
		int length = nums.length;

		int low = 0;
		int hight = length-1;
		for (; low <= hight;) {
			int middle = low + (hight - low) / 2;
//			System.out.println(middle);

			if (nums[middle] == target) {
				return middle;
			} else if (nums[middle] > target) {
				hight = middle - 1;
			} else {
				low = middle + 1;
			}
		}
		return result;
	}
}

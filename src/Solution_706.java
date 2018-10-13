package leetcode;

class MyHashMap {

	static final int MAX = 1000001;
	private boolean[] hashKey;
	private int[] hashValue;

	/** Initialize your data structure here. */
	public MyHashMap() {
		hashKey = new boolean[MAX];
		hashValue = new int[MAX];
	}

	/** value will always be positive. */
	public void put(int key, int value) {
		if (key >= 0 && key < MAX) {
			hashKey[key] = true;
			hashValue[key] = value;
		}
	}

	private boolean contains(int key) {
		if (key >= 0 && key < MAX) {
			return hashKey[key];
		} else {
			return false;
		}
	}

	/**
	 * Returns the value to which the specified key is mapped, or -1 if this map
	 * contains no mapping for the key
	 */
	public int get(int key) {
		if (contains(key)) {
			return hashValue[key];
		}

		return -1;
	}

	/**
	 * Removes the mapping of the specified value key if this map contains a mapping
	 * for the key
	 */
	public void remove(int key) {
		if (contains(key)) {
			hashKey[key] = false;
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such: MyHashMap obj
 * = new MyHashMap(); obj.put(key,value); int param_2 = obj.get(key);
 * obj.remove(key);
 */

public class Solution_706 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		MyHashMap hashMap = new MyHashMap();
		hashMap.put(1, 1);          
		hashMap.put(2, 2);         
		System.out.println(hashMap.get(1));            // 返回 1
		System.out.println(hashMap.get(3));            // 返回 -1 (未找到)
		hashMap.put(2, 1);         // 更新已有的值
		System.out.println(hashMap.get(2));            // 返回 1 
		hashMap.remove(2);         // 删除键为2的数据
		System.out.println(hashMap.get(2));            // 返回 -1 (未找到) 
	}

}

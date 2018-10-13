package leetcode;

class MyHashSet {

	private boolean[] hash;
    /** Initialize your data structure here. */
    public MyHashSet() {
    	hash =new boolean[1000001];
    }
    
    public void add(int key) {
    	if(key>=0 && key < 1000001) {
    		hash[key] = true;
    	}
    	
    }
    
    public void remove(int key) {
    	if(key>=0 && key < 1000001 && hash[key]) {
    		hash[key] = false;
    	}
    }
    
    /** Returns true if this set did not already contain the specified element */
    public boolean contains(int key) {
    	if(key>=0 && key < 1000001) {
    		return hash[key];
    	}else {
    		return false;
    	}
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet obj = new MyHashSet();
 * obj.add(key);
 * obj.remove(key);
 * boolean param_3 = obj.contains(key);
 */

public class Solution_705 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		 MyHashSet obj = new MyHashSet();
		 obj.add(1);
		 obj.add(3);
		 System.out.println(obj.contains(1));
		 System.out.println(obj.contains(2));
		 
		 MyHashSet obj2 = new MyHashSet();
		 System.out.println(obj2.contains(3));

	}

}

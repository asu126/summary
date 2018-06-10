package javaBase.collection;

import java.util.LinkedList;
class MyStack {
	private LinkedList list;
    /** Initialize your data structure here. */
    public MyStack() {
    	this.list = new LinkedList();
    }
    
    /** Push element x onto stack. */
    public void push(int x) {
    	this.list.push(x); // 将该元素插入此列表的开头
    }
    
    /** Removes the element on top of the stack and returns that element. */
    public int pop() {
    	if(!empty()) {
    		return (int) this.list.pop();
    	}
    	return 0;
    }
    
    /** Get the top element. */
    public int top() {
    	if(!empty()) {
    		return (int) this.list.peekFirst();
    	}
    	return 0;
    }
    
    /** Returns whether the stack is empty. */
    public boolean empty() {
        if(this.list.size() == 0) {
        	return true;
        } else {
        	return false;
        }	
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * MyStack obj = new MyStack();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.top();
 * boolean param_4 = obj.empty();
 */

public class MyStackTest {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		 MyStack obj = new MyStack();
		 obj.push(1);
		 System.out.println(obj.toString());
		 int param_2 = obj.pop();
		 int param_3 = obj.top();
		 boolean param_4 = obj.empty();
		 System.out.println(param_2 + " " + param_3 + " " + param_4); // 1,2,true
	}

}

import java.util.Stack;

class MyQueue {
	private Stack pushStack;
	private Stack popStack;
    /** Initialize your data structure here. */
    public MyQueue() {
    	this.pushStack = new Stack();
    	this.popStack = new Stack();
    }
    
    /** Push element x to the back of queue. */
    public void push(int x) {
    	this.pushStack.push(new Integer(x));
    }
    
    /** Removes the element from in front of queue and returns that element. */
    public int pop() {
    	if(!empty()) {
    		if(this.popStack.empty()) {
    			movePushSatckToPop();
    		}
    		return (int) this.popStack.pop();
    	}
    	return 0;
    }
    
    /** Get the front element. */
    public int peek() {
    	if(!empty()) {
    		if(this.popStack.empty()) {
    			movePushSatckToPop();
    		}
    		return (int) this.popStack.peek();
    	}
    	return 0;
    }
    
    /** Returns whether the queue is empty. */
    public boolean empty() {
        if(this.pushStack.empty() && this.popStack.empty()) {
        	return true;
        } else {
        	return false;
        }
    }
    
    private void movePushSatckToPop() {
    	while(!this.pushStack.empty()) {
    		Integer temp  = (Integer) this.pushStack.pop();
    		this.popStack.push(temp);
    	}
    }
}

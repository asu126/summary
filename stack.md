# Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

这一题我的理解是用栈进行输入匹配，这一题在做的时候最好是用动态数组，否则会出越界的问题，还有最后还要对队列是否为空做个判断。

```
# 越界
"[([([([([([([([([([([([([([([([([([([([([([([([([([([([([([[([([([([([([([([(([([...."
# 最后栈不为空
“((”
```

[code](src/validParentheses.go)


# Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```

这一题其实题目很明确，就是实现栈的pop,push,top,getMin 操作，其原理并不难，只是需要注意处理一些特殊情况（头指针为空，只有一个节点以及多个节点）。总之，做完之后最好还是先检视一下，多构造几个用例跑跑。
这一题的另一个收获就是对golang struct 对象赋予其方法的运用，将之前书本知识运用到了代码中，总之，不要眼高手低。
```
type MinStack struct {
	Head *Node
	// Next *Node
}
```

[code](src/minStack.go)

# 150. 逆波兰表达式求值
# 225. 用队列实现栈
# 232. 用栈实现队列
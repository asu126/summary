# Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

这一题我的理解是用栈进行输入匹配，这一题在做的时候最好是用动态数组，否则会出越界的问题，还有最后还要对队列是否为空做个判断。[code](src/validParentheses.go)

```
# 越界
"[([([([([([([([([([([([([([([([([([([([([([([([([([([([([([[([([([([([([([([(([([...."
# 最后栈不为空
“((”
```
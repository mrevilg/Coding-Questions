# Coding-Questions
A Repo that contains ongoing development interview questions in both Python/C# (where applicable)

**Q1** - Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
You can assume that the messages are decodable. For example, '001' is not allowed.

**Q2** - `cons(a, b)` constructs a pair, and `car(pair)` and `cdr(pair)` returns the first and last element of that pair. For example, `car(cons(3, 4))` returns `3`, and `cdr(cons(3, 4))` returns `4`.
Given this implementation of cons:
```
def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
```    
Implement `car` and `cdr`.

**Q3** - Given the root to a binary tree, implement `serialize(root)`, which serializes the tree into a string, and `deserialize(s)`, which deserializes the string back into the tree.
For example, given the following Node class
```
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
```
The following test should pass:

`node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'`

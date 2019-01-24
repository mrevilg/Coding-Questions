# Coding-Questions
A Repo that contains ongoing development interview questions, with Source Code answer files in both Python/C# (where applicable)

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

**Q4** - Given an array of integers, return a new array such that each element at index `i` of the new array is the product of all the numbers in the original array except the one at `i`.

For example, if our input was `[1, 2, 3, 4, 5]`, the expected output would be `[120, 60, 40, 30, 24]`. If our input was `[3, 2, 1]`, the expected output would be `[2, 3, 6]`.

Follow-up: what if you can't use division?

**Q5** - Given a list of numbers and a number `k`, return whether any two numbers from the list add up to `k`.
For example, given `[10, 15, 3, 7]` and `k` of `17`, return true since `10 + 7` is `17`.


**Q6** - The sentence "A quick brown fox jumps over the lazy dog" contains every single letter in the alphabet. Such sentences are called pangrams.

You are to write a method `getMissingLetters`, which takes as input a string containing a sentence and returns all the letters not present at all in the sentence (i.e., the letters that prevent it from being a pangram).

You should *ignore the case of the letters in sentence*, and your **return should be all lower case letters**, in alphabetical order.

You should also ignore all non-alphabet characters as well as all non-US-ASCII characters. 

Imagine that the method you write will be called many thousands of times in rapid succession on strings with length ranging from 0 to 50.

Accordingly, you should try to write code that runs as quickly as possible. Also, imagine the case when the input string is quite large (e.g., tens of megabytes). See if you can develop an algorithm that handles this case efficiently while still running very quickly on smaller inputs.

###Examples:### (Note that in the examples below, the double quotes should not be considered part of the input or output strings.)

| Input Text | Output |
| --- | --- |
| "A quick brown fox jumps over the lazy dog" | "" |
| "A slow yellow fox crawls under the proactive dog" | "bjkmqz" |
| "Lions, and tigers, and bears, oh my!" | "cfjkpquvwxz" |
| "" | "abcdefghijklmnopqrstuvwxyz" |

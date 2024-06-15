# Shifted Binary Search

Category: Searching

Difficulty: Hard

## Description

Write a function that takes in a sorted array of distinct integers as well as a target
integer. The caveat is that the integers in the array have been shifted by
some amount; in other words, they've been moved to the left or to the right by
one or more positions. For example, `[1, 2, 3, 4]` might have
turned into `[3, 4, 1, 2]`.

The function should use a variation of the Binary Search algorithm to
determine if the target integer is contained in the array and should return
its index if it is, otherwise `-1`.

If you're unfamiliar with Binary Search, we recommend watching the Conceptual
Overview section of the Binary Search question's video explanation before
starting to code.


### Sample Input
```
array = [45, 61, 71, 72, 73, 0, 1, 21, 33, 37]
target = 33
```

### Sample Output
```
8
```

## Optimal Space & Time Complexity

O(log(n)) time | O(1) space - where n is the length of the input array

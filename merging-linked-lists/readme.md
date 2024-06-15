# Merging Linked Lists

Category: Linked Lists

Difficulty: Medium

## Description

  You're given two Linked Lists of potentially unequal length. These Linked
Lists potentially merge at a shared intersection node. Write a function
that returns the intersection node or returns `None` /
`null` if there is no intersection.

  Each `LinkedList` node has an integer `value` as well as
a `next` node pointing to the next node in the list or to
`None` / `null` if it's the tail of the list.

  Note: Your function should return an existing node. It should not modify
either Linked List, and it should not create any new Linked Lists.

  
### Sample Input
```
linkedListOne = 2 -> 3 -> 1 -> 4
linkedListTwo = 8 -> 7 -> 1 -> 4
```
  
### Sample Output
```
1 -> 4 // The lists intersect at the node with value 1
```
  
## Optimal Space & Time Complexity

O(n + m) time | O(1) space - where n is the length of the first Linked List and m is the length of the second Linked List

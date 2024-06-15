# Middle Node

Category: Linked Lists

Difficulty: Easy

## Description

  You're given a Linked List with at least one node. Write a function
that returns the middle node of the Linked List. If there are two middle
nodes (i.e. an even length list), your function should return the second
of these nodes.


  Each `LinkedList` node has an integer `value` as well as
a `next` node pointing to the next node in the list or to
`None` / `null` if it's the tail of the list.


  
### Sample Input
```
linkedList = 2 -> 7 -> 3 -> 5
```
  
### Sample Output
```
3 -> 5 // The middle could be 7 or 3,
// we return the second middle node
```

## Optimal Space & Time Complexity

O(n) time | O(1) space - where n is the number of nodes in the linked list

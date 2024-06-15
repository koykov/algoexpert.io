# Symmetrical Tree

Category: Binary Trees

Difficulty: Medium

## Description

Write a function that takes in a Binary Tree and returns if that tree is
symmetrical. A tree is symmetrical if the left and right subtrees are
mirror images of each other.

Each `BinaryTree` node has an integer `value`, a
`left` child node, and a `right` child node. Children
nodes can either be `BinaryTree` nodes themselves or
`None` / `null`.


### Sample Input
```
tree =    1
       /     \
      2       2
    /   \   /   \
   3     4 4     3
 /   \          /  \
5     6        6    5
```

### Sample Output
```
True
```

## Optimal Space & Time Complexity

O(n) time | O(h) space - where n is the number of nodes in the tree and h is the height of the tree.

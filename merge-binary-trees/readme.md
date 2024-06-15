# Merge Binary Trees

Category: Binary Trees

Difficulty: Medium

## Description

  Write a function that takes in two Binary Trees, merges them, and returns
the resulting tree. If two nodes overlap during the merge, the value of the
merged node should be the sum of the overlapping nodes' values.

  Note that your solution can either mutate the input trees or return a new
tree.

  Each `BinaryTree` node has an integer `value`, a
`left` child node, and a `right` child node. Children
nodes can either be `BinaryTree` nodes themselves or
`None` / `null`.

  
### Sample Input
  ```
tree1 =   1
        /   \
       3     2
     /   \
    7     4

tree2 =   1
        /   \
       5     9
     /      / \
    2      7   6
```
  
### Sample Output
  ```
output =  2
        /   \
      8      11
    /  \    /  \
  9     4  7    6
```

## Optimal Space & Time Complexity

O(n) time | O(h) space - where n is the number of nodes in the smaller of the two trees and h is the height of the shorter tree.

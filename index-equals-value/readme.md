# Index Equals Value

Category: Searching

Difficulty: Hard

## Description

Write a function that takes in a sorted array of distinct integers and returns
the first index in the array that is equal to the value at that index. In
other words, your function should return the minimum index where
`index == array[index]`.

<p>If there is no such index, your function should return `-1`.</p>

### Sample Input
```
array = [-5, -3, 0, 3, 4, 5, 9]
```

### Sample Output
```
3 // 3 == array[3]
```

## Optimal Space & Time Complexity

O(log(n)) time | O(1) space - where n is the length of the input array

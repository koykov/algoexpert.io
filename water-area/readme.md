# Water Area

Category: Dynamic Programming

Difficulty: Hard

## Description

You're given an array of non-negative integers where each non-zero integer
represents the height of a pillar of width `1`. Imagine water being
poured over all of the pillars; write a function that returns the surface area
of the water trapped between the pillars viewed from the front. Note that
spilled water should be ignored.

You can refer to the first three minutes of this question's video explanation
for a visual example.


### Sample Input
```
heights = [0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3]
```

### Sample Output
```
48

// Below is a visual representation of the sample input.
// The dots and vertical lines represent trapped water and pillars, respectively.
// Note that there are 48 dots.
//        |
//        |
//  |.....|
//  |.....|
//  |.....|
//  |..|..|
//  |..|..|
//  |..|..|.....|
//  |..|..|.....|
// _|..|..|..||.|
```

## Optimal Space & Time Complexity

O(n) time | O(1) space - where n is the length of the input array

# Dice Throws

Category: Dynamic Programming

Difficulty: Hard

## Description

  You're given a set of `numDice` dice, each with
`numSides` sides, and a `target` integer, which
represents a target sum to obtain when rolling all of the dice and summing
their values. Write a function that returns the total number of dice-roll
permutations that sum up to exactly that target value.

  All three input values will always be positive integers. Each of the dice
has an equal probability of landing on any number from 1 to
`numSides`. Identical total dice rolls obtained from different
individual dice rolls (for example, `[2, 3]` vs.
`[3, 2]`) count as different dice-roll permutations. If there's
no possible dice-roll combination that sums up to the
`target` given the input dice, your function should return 0.


  
### Sample Input
  ```
numDice = 2
numSides = 6
target = 7
```
  
### Sample Output
  ```
6 // [1, 6], [2, 5], [3, 4], [4, 3], [5, 2], [6, 1]
```

## Optimal Space & Time Complexity

O(d * s * t) time | O(t) space - where d is the number of throws, s is the number of sides, and t is the target.

# Staircase Traversal

Category: Recursion

Difficulty: Medium

## Description

You're given two positive integers representing the height of a staircase and
the maximum number of steps that you can advance up the staircase at a time.
Write a function that returns the number of ways in which you can climb the
staircase.

<p>
For example, if you were given a staircase of `height = 3` and
`maxSteps = 2` you could climb the staircase in 3 ways. You could
take <b>1 step, 1 step, then 1 step</b>, you could also take
<b>1 step, then 2 steps</b>, and you could take <b>2 steps, then 1 step</b>.
</p>

Note that `maxSteps <= height` will always be true.


### Sample Input
```
height = 4
maxSteps = 2
```

### Sample Output
```
5
// You can climb the staircase in the following ways: 
// 1, 1, 1, 1
// 1, 1, 2
// 1, 2, 1
// 2, 1, 1
// 2, 2
```

## Optimal Space & Time Complexity

O(n) time | O(n) space - where n is the height of the staircase

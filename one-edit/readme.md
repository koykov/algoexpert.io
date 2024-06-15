# One Edit

Category: Strings

Difficulty: Medium

## Description

  You're given two strings `stringOne` and `stringTwo`.
Write a function that determines if these two strings can be made equal
using only one edit.


  <p>
There are 3 possible edits:
* <b>Replace</b>: One character in one string is swapped for a different
character.
* <b>Add</b>: One character is added at any index in one string.
* <b>Remove</b>: One character is removed at any index in one string.

</p>

  Note that both strings will contain at least one character. If the strings
are the same, your function should return true.


### Sample Input
```
stringOne = "hello"
stringTwo = "hollo"
```

### Sample Output
```
True // A single replace at index 1 of either string can make the strings equal
```

## Optimal Space & Time Complexity

O(n) time | O(1) space - where n is the length of the shorter string

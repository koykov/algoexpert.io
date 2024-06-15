# First Non-Repeating Character

Category: Strings

Difficulty: Easy

## Description

Write a function that takes in a string of lowercase English-alphabet letters
and returns the index of the string's first non-repeating character.

The first non-repeating character is the first character in a string that
occurs only once.

If the input string doesn't have any non-repeating characters, your function
should return `-1`.


### Sample Input
```
string = "abcdcaf"
```

### Sample Output
```
1 // The first non-repeating character is "b" and is found at index 1.
```

## Optimal Space & Time Complexity

O(n) time | O(1) space - where n is the length of the input string\nThe constant space is because the input string only has lowercase\nEnglish-alphabet letters; thus, our hash table will never have more\nthan 26 character frequencies.

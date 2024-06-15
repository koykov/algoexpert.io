# Levenshtein Distance

Category: Dynamic Programming

Difficulty: Medium

## Description

Write a function that takes in two strings and returns the minimum number of
edit operations that need to be performed on the first string to obtain the
second string.

There are three edit operations: insertion of a character, deletion of a
character, and substitution of a character for another.


### Sample Input
```
str1 = "abc"
str2 = "yabd"
```

### Sample Output
```
2 // insert "y"; substitute "c" for "d"
```

## Optimal Space & Time Complexity

O(nm) time | O(min(n, m)) space - where n and m are the lengths of the two input strings

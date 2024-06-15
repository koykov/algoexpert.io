# Suffix Trie Construction

Category: Tries

Difficulty: Medium

## Description

Write a `SuffixTrie` class for a Suffix-Trie-like data structure.
The class should have a `root` property set to be the root node of
the trie and should support:

* Creating the trie from a string; this will be done by calling the
    `populateSuffixTrieFrom` method upon class instantiation, which
    should populate the `root` of the class.
* Searching for strings in the trie.

Note that every string added to the trie should end with the special
`endSymbol` character: `"*"`.

If you're unfamiliar with Suffix Tries, we recommend watching the
Conceptual Overview section of this question's video explanation before
starting to code.


### Sample Input (for creation)
```
string = "babc"
```

### Sample Output (for creation)
```
The structure below is the root of the trie.
{
  "c": {"*": true},
  "b": {
    "c": {"*": true},
    "a": {"b": {"c": {"*": true}}},
  },
  "a": {"b": {"c": {"*": true}}},
}
```
<h3>
  Sample Input (for searching in the suffix trie above)
</h3>
```
string = "abc"
```
<h3>
  Sample Output (for searching in the suffix trie above)
</h3>
```
true
```

## Optimal Space & Time Complexity

Creation: O(n^2) time | O(n^2) space - where n is the length of the input string\nSearching: O(m) time | O(1) space - where m is the length of the input string

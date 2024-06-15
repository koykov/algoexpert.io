# Prim's Algorithm

Category: Famous Algorithms

Difficulty: Hard

## Description

  You're given a list of `edges` representing a connected,
weighted, undirected graph with at least one node.


  The given list is what's called an adjacency list, and it represents a graph.
The number of vertices in the graph is equal to the length of
`edges`, where each index `i` in
`edges` contains vertex `i`'s siblings, in no
particular order. Each of these siblings is an array of length two, with
the first value denoting the index in the list that this vertex is connected
to, and and the second value denoting the weight of the edge. Note that
this graph is undirected, meaning that if a vertex appears in the edge list
of another vertex, then the inverse will also be true (along with the same
weight).


  Write a function implementing Prim's Algorithm to return a new
`edges` array that represents a minimum spanning tree. A
minimum spanning tree is a tree containing all of the vertices of the original
graph and a subset of the edges. These edges should connect all of the vertices
with the minimum total edge weight and without generating any cycles.


  Note that the graph represented by `edges` won't contain any
self-loops (vertices that have an outbound edge to themselves) and will only
have positively weighted edges (i.e., no negative distances). The graph will
always have at least one node.


  If you're unfamiliar with Prim's algorithm, we recommend watching the
Conceptual Overview section of this question's video explanation before
starting to code.

  
### Sample Input
  ```
edges = [
  [[1, 3], [2, 5]],
  [[0, 3], [2, 10], [3, 12]],
  [[0, 5], [1, 10]],
  [[1, 12]]
]
```
  
### Sample Output
  ```
[
  [[1, 3], [2, 5]],
  [[0, 3], [3, 12]],
  [[0, 5]],
  [[1, 12]]
]
```

## Optimal Space & Time Complexity

O(e * log(v)) time | O(e + v) space - where e is the number of edges in the input edges and v is the number of vertices

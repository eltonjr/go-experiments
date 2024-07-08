## Map O(1) vs Slice O(1)

Compares the performance of a key lookup in a map vs in a slice

### Filling a map vs filling a slice

The first benchmark creates a map and a slice of the same size, and init both of them with some non-zero value.

The benchmark aims to compare the performance of filling an array (slice) vs filling a tree (map).  
Map also has to hash each key, so the performance is very impacted.

### Accessing an element of a map vs of a slice

The second benchmark accesses a key on the map vs an index on the slice.  
Both operations are thought commonly as O(1) ([at least on average](https://en.wikipedia.org/wiki/Hash_table)).

The benchmark aims to show the performance penalty of hashing the key for maps.

### Running

You can run it with 

```
go test -benchmem -run=^$ -bench . .
```

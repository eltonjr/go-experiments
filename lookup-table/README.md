## Lookup table vs slice/switch performance comparison

Compares the performance of a lookup table pattern

### Basic lookup table

A common pattern is to store values in a map, aka a lookup table, looking for O(1) access.

```
var allowedProfilingTypes = map[string][]string{
	"go":     {"cpu", "heap", "allocs", "block", "trace"},
	"java":   {"cpu", "allocs", "block"},
	"kotlin": {"cpu", "allocs", "block"},
	"python": {"cpu", "heap", "thread"},
	"nodejs": {"cpu", "heap"},
}
```

However, [maps in go are stored as a tree](https://go.dev/src/runtime/map.go), with every key hashed.  
So when looking for a value, the key argument must be hashed to be able to iterate through the tree.  
This has a cost.

Alternatives may be: working with slices; having switch-cases. Of course it all depends on the size and complexity of your data

```
func allowedProfilingTypes(tech string) []string {
	switch tech {
	case "go":
		return []string{"cpu", "heap", "allocs", "block", "trace"}
	case "java", "kotlin":
		return []string{"cpu", "allocs", "block"}
	case "python":
		return []string{"cpu", "heap", "thread"}
	case "nodejs":
		return []string{"cpu", "heap"}
	default:
		return []string{}
	}
}
```

This experiment also has two other scenarios: one having a second layer of maps; another having no switch-cases and working with slices only.

### Running

You can run it with 

```
go test -benchmem -run=^$ -bench . .
```

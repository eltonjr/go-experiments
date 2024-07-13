## Appending to slice: struct, interface or struct casted to interface?

Compares which one is better:

1. Having a slice of interface, and appending some value already casted as the interface

        var s []SomeInterface
        s = append(s, SomeInterface{val})

1. Having a slice of interface, and appending some value that satisfies the interface

        var s []SomeInterface
        s = append(s, val)

1. Using only types

        var s []SomeStruct
        s = append(s, val)

Also compares calling methods on interfaces vs direct call

### Running

Run the benchmarks with 

        go test -benchmem -run=none -bench . github.com/eltonjr/go-experiments/interface-struct

A possible result is

```
goos: darwin
goarch: arm64
pkg: github.com/eltonjr/go-experiments/interface-struct
BenchmarkStructs/Append_to_interface_slice_as_type-8         	19509494	        76.14 ns/op	     111 B/op	       1 allocs/op
BenchmarkStructs/Append_to_interface_slice_as_interface-8    	62278483	        20.77 ns/op	      16 B/op	       1 allocs/op
BenchmarkStructs/Append_to_type_slice_as_type-8              	100000000	        59.78 ns/op	      95 B/op	       0 allocs/op
BenchmarkStructs/Call_method_on_interface_(uncasted)-8       	533164201	         2.250 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructs/Call_method_on_interface-8                  	544561174	         2.188 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructs/Call_method_on_type-8                       	1000000000	         0.9899 ns/op	       0 B/op	       0 allocs/op
```

So what we see is:

1. Appending structs to an interface slice, and letting go converts the type is the worst case scenario;
1. Appending interface to an interface slice is the quickest
1. Appending structs is slightly slower, but doesnt allocate anything

For method call, using structs is the quickest

*But why is that appending an interface is better than appending a struct?* 


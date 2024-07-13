# Slice, backing array and the heap

If a big array is created, but there is only a reference for the first element, how much of will be considered "live memory"?

The following example creates a slice of 10GB, but returns only a slice of the first element.  
This returned slice is stored in a global variable, so it will remain alive after the GC call.  
But what about the rest of the backing array?

```
var cache []byte

func main() {
	cache = bigalloc()
	runtime.GC()
}

func bigalloc() []byte {
	ballast := make([]byte, 10<<30)
	return ballast[:1]
}
```

## Running

1. Build the program with

        go build -o sliceheap main.go

1. Run the first time in 'discard' mode. This will allocate 10GB 3x but will not keep any of it as live memory

        ./sliceheap discard

        # Extra: to see gc in debug mode, set GODEBUG and discard the logs
        GODEBUG=gctrace=1 ./sliceheap discard > /dev/null

1. Run a second time in 'keep' mode. This will allocate 10GB 3x, and keep only a single byte in a global variable

        ./sliceheap keep
        
        # Extra: to see gc in debug mode, set GODEBUG and discard the logs
        GODEBUG=gctrace=1 ./sliceheap keep > /dev/null

1. Open the mem pprof for each one

        go tool pprof -http :8081 mem-discard.pprof
        go tool pprof -http :8082 mem-keep.pprof

## Result

For the 'discard' mode, 30GB are allocated and 3 objects, but the GC will throw it all away.  
The inuse space will be 0 objects and 0MB.

```
go tool pprof mem-discard.pprof

(pprof) sample_index=alloc_objects  
(pprof) top                       
Showing nodes accounting for 3, 100% of 3 total
      flat  flat%   sum%        cum   cum%
         3   100%   100%          3   100%  main.bigalloc (inline)
         0     0%   100%          3   100%  main.main
         0     0%   100%          3   100%  runtime.main

(pprof) sample_index=alloc_space  
(pprof) top                     
Showing nodes accounting for 30GB, 100% of 30GB total
      flat  flat%   sum%        cum   cum%
      30GB   100%   100%       30GB   100%  main.bigalloc (inline)
         0     0%   100%       30GB   100%  main.main
         0     0%   100%       30GB   100%  runtime.main

(pprof) sample_index=inuse_objects
(pprof) top
Showing nodes accounting for 0, 0% of 0 total
      flat  flat%   sum%        cum   cum%

(pprof) sample_index=inuse_space  
(pprof) top                     
Showing nodes accounting for 0, 0% of 0 total
      flat  flat%   sum%        cum   cum%
```

For the 'keep' mode, 30GB are allocated and 3 objects.  
On each iteration, a single byte will be kept.
On each iteration, the GB will discard the previous 10GB array.

At the end, 30GB will be the total allocated space, and 3 objects.
The inuse space will be 10GB, in 1 object.

```
go tool pprof mem-keep.pprof

(pprof) sample_index=alloc_objects
(pprof) top
Showing nodes accounting for 3, 100% of 3 total
      flat  flat%   sum%        cum   cum%
         3   100%   100%          3   100%  main.bigalloc (inline)
         0     0%   100%          3   100%  main.main
         0     0%   100%          3   100%  runtime.main

(pprof) sample_index=alloc_space
(pprof) top
Showing nodes accounting for 30GB, 100% of 30GB total
      flat  flat%   sum%        cum   cum%
      30GB   100%   100%       30GB   100%  main.bigalloc (inline)
         0     0%   100%       30GB   100%  main.main
         0     0%   100%       30GB   100%  runtime.main

(pprof) sample_index=inuse_objects
(pprof) top
Showing nodes accounting for 1, 100% of 1 total
      flat  flat%   sum%        cum   cum%
         1   100%   100%          1   100%  main.bigalloc (inline)
         0     0%   100%          1   100%  main.main
         0     0%   100%          1   100%  runtime.main

(pprof) sample_index=inuse_space
(pprof) top
Showing nodes accounting for 10GB, 100% of 10GB total
      flat  flat%   sum%        cum   cum%
      10GB   100%   100%       10GB   100%  main.bigalloc (inline)
         0     0%   100%       10GB   100%  main.main
         0     0%   100%       10GB   100%  runtime.main
```

This is described in this [go blog post](https://go.dev/blog/slices-intro)

> #### A possible “gotcha”
> As mentioned earlier, re-slicing a slice doesn’t make a copy of the underlying array. **The full array will be kept in memory until it is no longer referenced. Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.**

## Extra

* What is the cap of the resulting slice?
* What happens if you keep the last byte instead of the first one?
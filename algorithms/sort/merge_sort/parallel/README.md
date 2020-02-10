# Parallel Merge Sort

This Parallel Merge Sort is a direct implementation based on the serial merge
sort [implemented here](../serial/main.go).

By a direct implementation, I mean that this parallel implementation is a trivial
modification of the serial version just introducing goroutines following a
simple [fork and join](https://en.wikipedia.org/wiki/Fork%E2%80%93join_model)
approach.

The [Merge Sort Wikipedia page](https://en.wikipedia.org/wiki/Merge_sort#Parallel_merge_sort)
has an interesting discussion about the running time of this parallel merge sort
and also some references of where to find better algorithms than the applied here.

Surprisingly, while performing some dummy tests in an 8 cores Ubuntu machine this
implementation shows the worst performance (avg 3x slower) if compared with the 
serial. For the tests, I used an array of random integers with 10 million elements.

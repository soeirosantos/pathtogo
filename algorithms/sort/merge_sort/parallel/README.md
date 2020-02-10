# Parallel Merge Sort

This Parallel Merge Sort is a direct implementation based on the serial merge
sort [implemented here](../serial/main.go).

By a direct implementation, I mean that the parallel implementation is a trivial
modification of the serial version by just introducing goroutines following a
a simple [fork and join](https://en.wikipedia.org/wiki/Fork%E2%80%93join_model)
approach.

The [Merge Sort Wikipedia page](https://en.wikipedia.org/wiki/Merge_sort#Parallel_merge_sort)
has an interesting discussion about the running time of the parallel merge sort
as well as some references about where to find better algorithms than the applied
here.

Surprisingly, while tested in a simple machine this parallel implementation shows
a much worst performance compared with the serial one.

# Read and Do

This is a practical example of how we can read a file
and operate on its content without coupling the operation
logic with the file-reading code.

## How to use

The `NODES_EDGES_LOCATION` environment variable points to a file containing the
edges of the graph with the following structure:

```bash
$ head refs.csv
article/af8031be-4b8a-57ac-a721-d6a85f5e4ed1    image/c2791388-6b16-5ffd-a119-96eacb1989c0
article/bc30023d-c849-575f-be1a-bea163aa4931    section/9b51ec65-534b-54c2-9483-b2cc9a40bbf1
article/6b22a322-1c2e-5ec3-bdfc-183aa31f3000    section/9b51ec65-534b-54c2-9483-b2cc9a40bbf1
```

To run it
```bash
$ NODES_EDGES_LOCATION=refs go run main.go
```

The output is the topological sort of the elements.
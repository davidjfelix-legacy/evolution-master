#evolution-master

**WIP Warning:** this project is still very much a work in progress.

## Intention

The intent of this project is to act as a runner for hatchery tasks.
The designe of genepool2 dictates an interpreter, so evolution-master may provide that interpreter.
This project will:

1. Read a brood repo
2. Parse the broodfile
3. retrieve top-level genepool repos
4. create a directory for genes
5. Traverse the gene tree, bringing in required genes from their broodfiles
6. Execute top level genes

I committed a binary (./enron-search) which should work on M1 Mac.

To build:
First, install go (1.17). Then run `./build.sh`. This outputs the binary.

To use:
First, copy the oracle data into a directory in this folder called `raw-data`.
For example, there should be a directory `./raw-data/allen-p`, `./raw-data/arnold-j`, and so on.
I added this data to gitignore because git was very unhappy even trying to commit it. 

To generate the index:
`./enron-search -genindex`
The index is also in gitignore.

You should see output like:
![genindex](./genindex.png)

To search the index:
`./enron-search -search apple`

Strategy:
The index is a Trie-like data structure which lives entirely on disk

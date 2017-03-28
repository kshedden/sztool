__sztool__ is a simple command-line tool for working with
[snappy](https://en.wikipedia.org/wiki/Snappy_(compression))
compressed files, similar to the `gzip` command for gzipped files.

The basic usage for compression is:

```
sztool -c in out
```

The basic usage for decompression is:

```
sztool -d in out
```

In both cases, `in` and `out` are input and output files.  These files
can be specified as "-" to use stdin or stdout.

The out parameter can be omitted.  For compression(-c), out defaults
to in.sz, where `in` is the input file.  For decompression (-d), out
defaults to stdout.

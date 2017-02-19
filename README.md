__sztool__ is a simple command-line tool for snappy compressed files.

The basic usage for compression is:

```
sztool -c in out
```

The basic usage for decompression is:

```
sztool -d in out
```

Input and output can be specified as "-" to obtain stdin or stdout.

The out parameter can be omitted.  For compression it defaults to
in.sz, for decompression it defaults to stdout.

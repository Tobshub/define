# Define

I found myself wanting to know the definition of words while having my terminal open...

So I decided to build something that would let me do that in the terminal.

## Installation

You need `git` and `go` installed

```bash
$ git clone https://github.com/tobshub/define
$ cd define
$ go install .
```

## Usage

```bash
$ define <word>
``` 

Definitions are cached after first define, to bypass the cache you can use the `no-cache` flag

```bash
$ define --no-cache <word>
``` 


This project uses the [Free Dictionary API](https://dictionaryapi.dev)

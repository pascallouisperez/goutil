## Golang Utilities [![Build Status](https://travis-ci.org/pascallouisperez/goutil.svg?branch=master)](https://travis-ci.org/pascallouisperez/goutil)

### errors

One-liner to create errors while retaining the file and line number of where it was created

    errors.New("some message %s", "here")

would produce

    "file.go:45: some message here"

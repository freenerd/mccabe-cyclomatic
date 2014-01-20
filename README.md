# mccabe-cyclomatic

Calculates Thomas McCabe's cyclomatic complexity for a go file or package

## Install

    go get github.com/freenerd/mccabe-cyclomatic

## Usage

To get the cyclomatic complexity of a file:

```
  $ mccabe-cyclomatic -f ./example/example.go
  7
```

To get the cyclomatic complexity of a package:

```
  $ mccabe-cyclomatic -p github.com/freenerd/mccabe-cyclomatic
  5
```

The output is the cyclomatic complexity number.

## Formula

complexity = number of code branches + 1

The following [go ast types](http://golang.org/pkg/go/ast/) are recognized as code branches:

- IfStmt (also catches `else if`)
- for:
  - ForStmt
  - RangeStmt
  - BranchStmt
- switch:
  - SwitchStmt
  - TypeSwitchStmt

Whenever the ast walker hits one of these nodes, it increases the complexity counter by one.

## Limitations

- Switch and TypeSwitch statement might lead to n case branches. These are currently not counted, but it is implicitely assumed that they only have 2 instead of n case branches.
- The package algorithm is not recursive, it will not recognize sub-packages.

## References

http://www.literateprogramming.com/mccabe.pdf
A Complexity Measure, by Thomas J McCabe. 1976

http://www.win.tue.nl/~aserebre/2IS55/2010-2011/10.pdf

https://en.wikipedia.org/wiki/Cyclomatic_complexity

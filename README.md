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

## Calculation

`complexity = number of code execution branch statements + 1`

The file to be analyzed is loaded with [go/parser ParseFile](http://golang.org/pkg/go/parser/#ParseFile) and subsequentially walked with a [go/ast Visitor](http://golang.org/pkg/go/ast/#Visitor). Depending on which ast node types are encountered, the complexity count is increased.

[Effective Go](http://golang.org/doc/effective_go.html#control-structures) describes the following code execution branch statements:

- if
- for
- switch
- select

These can be mapped to certain [Go ast node types](http://golang.org/pkg/go/ast/):

- **if** is of [IfStmt](http://golang.org/pkg/go/ast/#IfStmt) type. An *else if* is of [IfStmt](http://golang.org/pkg/go/ast/#IfStmt) type as well. An *else* clause does not add cyclomatic complexity.
- **for** is either of [ForStmt](http://golang.org/pkg/go/ast/#ForStmt) type or of [RangeStmt](http://golang.org/pkg/go/ast/#RangeStmt) type.
- **switch** might either be an expression switch or a type switch. In both cases, the number of code execution branches to be executed is determined by the number of case statements [CaseClause](http://golang.org/pkg/go/ast/#CaseClause). Therefore, we only count each [CaseClause](http://golang.org/pkg/go/ast/#CaseClause). *default* clauses do not add to cyclomatic complexity.
- **select** is similar to *switch*, in that its code execution branches are determined by the number of case statements, but of type [CommClause](http://golang.org/pkg/go/ast/#CommClause). *default* clauses do not add to cyclomatic complexity.

We end up with the list of the following ast node types, that determine a new code execution branch:

- [IfStmt](http://golang.org/pkg/go/ast/#IfStmt)
- [ForStmt](http://golang.org/pkg/go/ast/#ForStmt)
- [RangeStmt](http://golang.org/pkg/go/ast/#RangeStmt)
- [CaseClause](http://golang.org/pkg/go/ast/#CaseClause)
- [CommClause](http://golang.org/pkg/go/ast/#CommClause)

When any of these node types is encounterd, the complexity score is increased by 1.

Apart from the ones just listed, the go language specification includes more [control execution statements]((http://golang.org/ref/spec#Statements). These have not been recognized here: it is assumed that they do not add additional cyclomatic complexity, since they do not add a new code execution branch.

## Limitations

- complexity numbers are only returned for whole packages or files, but not for individual functions.
- the package algorithm is not recursive, it will not recognize sub-packages.

## References

http://www.literateprogramming.com/mccabe.pdf
A Complexity Measure, by Thomas J McCabe. 1976

http://www.win.tue.nl/~aserebre/2IS55/2010-2011/10.pdf

https://en.wikipedia.org/wiki/Cyclomatic_complexity

https://github.com/philbooth/escomplex

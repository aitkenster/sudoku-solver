#Sudoku Solver

Uses the recursive backtracking algorithm to solve puzzles.

##Technologies used

Golang(core packages)

git clone https://github.com/aitkenster/sudoku-solver.git

##How to run it

[install Go](https://golang.org/doc/install)

```
cd photo-mosaic
go get
go run solver.go PUZZLE_STRING
```

The puzzle string should contain 0s where there are blank spaces, e.g.:

`go run solver.go 000079065000003002005060093340050106000000000608020059950010600700600000820390000`

Will produce the following output:
```
Solution...
 1 8 3 2 7 9 4 6 5
 4 6 9 5 8 3 7 1 2
 2 7 5 4 6 1 8 9 3
 3 4 2 9 5 8 1 7 6
 5 9 7 1 3 6 2 8 4
 6 1 8 7 2 4 3 5 9
 9 5 4 8 1 2 6 3 7
 7 3 1 6 4 5 9 2 8
 8 2 6 3 9 7 5 4 1
 ```

 

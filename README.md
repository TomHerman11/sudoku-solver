# sudoku-solver
This is my first Go project. 🥳🤖

Solves any-square-size of Sudoku (4x4, 9x9, 16x16, etc.).

The solving technique being used is [Backtracking](https://en.wikipedia.org/wiki/Sudoku_solving_algorithms#Backtracking), as a brute-force search.  
From Wikipedia: "Each cell is tested for a valid number, moving "back" when there is a violation, and moving forward again until the puzzle is solved."

Future steps:
- More optimized methods can also be implemented.
- Scan a sudoku board + solve it. I did not find a free online service that computes a sudoku board from a picture. If if will find one, I can create a (web?) app - upload / take a picture of a sudoku board and its solution will be generated. A critical question - how to integrate Go code with webapp frameworks (React, Vue)? Maybe use Flutter instead?

## Demo
![Demo Sudoku Solver](/sudoku-solver.gif)

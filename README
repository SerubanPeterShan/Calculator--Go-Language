# Calculator Package

This is a simple calculator package that demonstrates basic Go testing practices.

### Participant

- Ashish Poudel (500237600)
- Ikecheku Samuel Madu (500236443)
- Nipuna Pamuditha Karunarathne (500238003)
- Seruban Peter Shan (500235797)


## Package Contents

- `calculator.go`: Contains the main package functions
- `calculator_test.go`: Contains the test cases

### Available Functions

1. `Divide(a, b int) (int, error)`
    - Performs integer division
    - Returns error if division by zero is attempted

2. `square(a int) int`
    - Returns the square of a number
    - Internal function (not exported)

## Running Tests

### On Mac/Linux:

```powershell
cd /path/to/project
go test -v
```

### On Windows :

```bash
cd \path\to\project
go test -v
```

## Why No main.go?

Imagine you're building with LEGO blocks. There are two ways to play with LEGO:

- Building a complete toy (like a castle or spaceship) that you can play with right away
- Making special LEGO pieces that other people can use in their own builds

This package is like the second option - it's making special LEGO pieces (code) that other people can use in their projects!

**In Go:**

A main.go file is like the instructions to build a complete toy - it tells the computer "start here and build this whole thing".
But since we're just making pieces for others to use, we don't need those instructions.
Instead, we have special test files that help us make sure our LEGO pieces are perfect before sharing them.

When we want to test our pieces, Go has a special helper (called go test) that:

- Looks through all our files
- Finds the ones marked as "test files" (they end in _test.go)
- Runs them to make sure everything works correctly
- So, just like LEGO pieces in a box waiting to be used in someone else's creation, this code is a package waiting to be imported into other programs! 🧩

## Test Coverage

Test coverage measures how much of your code is executed during testing. Our tests aim for high coverage by:

- Testing normal operation cases
- Testing edge cases (like division by zero)
- Verifying error conditions
- Testing internal functions

Our current coverage includes:

- Function calls with valid inputs
- Error handling scenarios
- Internal helper functions

The test file includes:

- Division test with valid input
- Division by zero test
- Square function test
- (Commented) Addition test for future implementation

To check test coverage:

```bash
go test -cover
```

## Benchmark

Benchmarking is like a race for your code. It helps you see how fast your code runs and where it might be slow. Just like timing how fast you can run a lap around the track, benchmarking times how fast your code can complete tasks.

### Why Benchmark?

- **Identify Performance Bottlenecks**: Find out which parts of your code are slow.
- **Optimize Code**: Improve the speed and efficiency of your code.
- **Compare Implementations**: Test different ways of doing the same thing to see which is faster.

### Running Benchmarks

To run the benchmarks on macOS or Windows, use the following command in the terminal:

**MAC OS**

```bash
go test -bench=.
```

**Windows OS**

```powershell
go test -bench .
```

### Example Benchmark Output

```powershell
goos: windows
goarch: amd64
pkg: github/SerubanPeterShan/Calculator
cpu: Intel(R) Core(TM) i7-14650HX
BenchmarkDivide-24      1000000000               0.1210 ns/op
BenchmarkSquare-24      1000000000               0.1155 ns/op
PASS
ok      github/SerubanPeterShan/Calculator      1.030s
```

This command will execute all benchmark functions in the package and provide performance metrics.

### Example Benchmark Functions

- `BenchmarkDivide(b *testing.B)`: Measures the performance of the `Divide` function.
- `BenchmarkSquare(b *testing.B)`: Measures the performance of the `square` function.

The benchmarks will run the `Divide` and `square` functions repeatedly to measure their execution time and performance under load.

### Notes

- Files must be in the same directory.
- Test files must end with `_test.go`.
- Each benchmark function must start with `Benchmark`.

By running benchmarks, you can ensure your code runs efficiently and make improvements where necessary.

Package `calculator` provides basic arithmetic operations and their tests.

The tests include unit tests for the `Divide` and `square` functions, as well as benchmarks to measure their performance.

To run the benchmarks on macOS or Windows, use the following command in the terminal:

<table>
<tr>
<th>Mac OS</th>
<th>Windows</th>
</tr>
<tr>
<td>
  
```bash

go test -bench=.

```
  
</td>
<td>

```powershell

go test -bench .

```

</td>
</tr>
</table>


This command will execute all benchmark functions in the package and provide performance metrics.

Example usage:

- `BenchmarkDivide(b *testing.B)`
- `BenchmarkSquare(b *testing.B)`

The benchmarks will run the `Divide` and `square` functions repeatedly to measure their execution time and performance under load.

- Files must be in the same directory
- Test files must end with `_test.go`
- Each test function must start with `Test`

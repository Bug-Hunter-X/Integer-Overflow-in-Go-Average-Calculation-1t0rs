# Integer Overflow in Go Average Calculation

This repository demonstrates a potential integer overflow bug in a Go program that calculates the average of a slice of integers. The program panics when the average results in a very large number that can't be represented as an int.

## Bug Description
The `processData` function calculates the average of a slice of integers.  However, it converts the `float64` average directly to an `int` without checking for potential overflow. This leads to a panic when the resulting integer is larger than the maximum value representable by `int`.

## Solution
The solution involves checking the potential overflow before converting the `float64` to `int`. The solution provided uses `math.MaxInt` and `math.MinInt` to safely handle potential overflow.
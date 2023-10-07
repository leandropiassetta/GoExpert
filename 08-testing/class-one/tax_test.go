package tax

import "testing"

// go test -v -> verbose mode (shows the name of the tests)

func Test_CalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected: %f, but got: %f", expected, result)
	}
}

// batch -> many tests in one file
func Test_CalculateTax_Batch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{amount: 500, expect: 5},
		{amount: 1000, expect: 10},
		{amount: 1500, expect: 10},
		// {amount: 2000, expect: 12},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("Expected: %f, but got: %f", item.expect, result)
		}
	}
}

// benchmark -> test the performance of the code (how long it takes to run) -> go test -bench=. -> run all benchmarks in the file (the name of the function must start with Benchmark) -> go test -bench=BenchMarkCalculateTax -> run a specific benchmark in the file (the name of the function must start with Benchmark) -> go test -bench=BenchMarkCalculateTax -benchtime=10s -> run a specific benchmark in the file (the name of the function must start with Benchmark) for 10 seconds

// go test -bench=.   -> run all benchmarks in the file (the name of the function must start with Benchmark)

// go test -bench=. -run=^#  -> run all benchmarks in the file (the name of the function must start with Benchmark) and ignore the tests (the name of the function must start with Test)

// to have that write "Benchmark" + the name of the function that you want to test
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

// verifing the performance of the function CalculateTax_2 (it takes more time to run than the function CalculateTax), benchmarking is a good way to find out if the code is running well or not and compare the performance of the code
func BenchmarkCalculateTax_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax_2(500)
	}
}

// go test -bench=. -run=^# -count=10 -> run all benchmarks in the file (the name of the function must start with Benchmark) and ignore the tests (the name of the function must start with Test) and run 10 times (the default is 1 time)

// go test -bench=. -run=^# -count=10 -benchtime=3s --> run all benchmarks in the file (the name of the function must start with Benchmark) and ignore the tests (the name of the function must start with Test) and run 10 times for 3 seconds each time (the default is 1 second)

// go help test -> to see all the options that you can use with go test command (go help testflag -> to see all the options that you can use with go test command)

//  go test -bench=. -run=^# -benchmem -> to see the memory allocation of the function CalculateTax_2 (it allocates more memory than the function CalculateTax) -> go test -bench=. -run=^# -benchmem -memprofile=mem.out -> to see the memory allocation of the function CalculateTax_2 (it allocates more memory than the function CalculateTax) and save the result in a file called mem.out

package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
======= Test Main =======
*/
func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // <= execution all unit test

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorldGolang(t *testing.T) {
	result := HelloWorld("Golang")

	if result != "Hello Golang" {
		// error
		// don't use panic, because it will stop the test
		// panic("Result is not 'Hello Golang'")

		// don't use t.Fail(), because it not give any information
		// t.Fail()

		t.Error("Result must be 'Hello Golang'")
	}
	fmt.Println("TestHelloWorldGolang done")
}

func TestHelloWorldJhon(t *testing.T) {
	result := HelloWorld("Jhon")

	if result != "Hello Jhon" {
		// error
		// panic("Result is not 'Hello Jhon'")

		// t.FailNow()

		t.Fatal("Result must be 'Hello Jhon'")
	}

	fmt.Println("TestHelloWorldJhon done")
}

/*
======= Assertion =======
*/
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Doe")
	assert.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")

	fmt.Println("TestHelloWorld with Assert done") // <= this is still running even if the test is failed
}

/*
======= Require =======
*/
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Doe")
	require.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")

	fmt.Println("TestHelloWorld with require done") // <= this is not running
}

/*
======= Test Skip =======
*/
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can't run on MACOS")
	}

	result := HelloWorld("Doe")
	require.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")
}

/*
======= Sub Test =======
*/
func TestSubTest(t *testing.T) {
	t.Run("Jhon", func(t *testing.T) {
		result := HelloWorld("Jhon")
		require.Equal(t, "Hello Jhon", result, "Result must be 'Hello Jhon'")
	})

	t.Run("Doe", func(t *testing.T) {
		result := HelloWorld("Doe")
		require.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")
	})
}

/*
======= Table Test =======
*/
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Jhon)",
			request:  "Jhon",
			expected: "Hello Jhon",
		},
		{
			name:     "HelloWorld(Doe)",
			request:  "Doe",
			expected: "Hello Doe",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

/*
======= Benchmark =======
*/
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jhon")
	}
}

func BenchmarkHelloWorldDoe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Doe")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Jhon", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jhon")
		}
	})

	b.Run("Doe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Doe")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Jhon)",
			request: "Jhon",
		},
		{
			name:    "HelloWorld(Doe)",
			request: "Doe",
		},
		{
			name:    "HelloWorld(Doe)",
			request: "Amorim",
		},
		{
			name:    "HelloWorld(Doe)",
			request: "Manchester",
		},
	}

	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

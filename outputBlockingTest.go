package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type testFunctionInstance struct {
	beforeTimer func()
	before      func()
	testFunc    func()
	after       func()
	afterTimer  func()
}

type testFunction struct {
	name string
	init func() testFunctionInstance
}

func main() {
	// defaultTestString := "Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World"
	defaultTestString := "Hello World"

	testString := *(flag.String("text", defaultTestString, "the string to use in the print loop"))
	loopDuration := *(flag.Int64("t", 800, "max duration of the loop in milliseconds [ms]"))
	useNewLine := *(flag.String("nl", "\n", "newline string, used for every function expect fmt.Println; string gets build before time stoping"))
	flag.Parse()

	testStringWithNewLine := testString + useNewLine
	testStringWithNewLineAsByteArray := []byte(testStringWithNewLine)

	functions := []testFunction{
		testFunction{
			name: "Noop",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
					},
				}
			},
		},
		testFunction{
			name: "Buffer Test with Variable and fmt.Println afterwards",
			init: func() testFunctionInstance {
				var t string
				return testFunctionInstance{
					before: func() {
						t = ""
					},
					testFunc: func() {
						t += testStringWithNewLine
					},
					after: func() {
						fmt.Println(t)
					},
				}
			},
		},
		testFunction{
			name: "Buffer Test with Array and fmt.Println afterwards",
			init: func() testFunctionInstance {
				var t []string
				return testFunctionInstance{
					testFunc: func() {
						t = append(t, testStringWithNewLine)
					},
					after: func() {
						fmt.Println(strings.Join(t, ""))
					},
				}
			},
		},
		testFunction{
			name: "os.Stdout.WriteString",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
						os.Stdout.WriteString(testStringWithNewLine)
					},
				}
			},
		},
		testFunction{
			name: "os.Stdout.Write",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
						os.Stdout.Write(testStringWithNewLineAsByteArray)
					},
				}
			},
		},
		testFunction{
			name: "fmt.Print",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
						fmt.Print(testStringWithNewLine)
					},
				}
			},
		},
		testFunction{
			name: "fmt.Printf",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
						fmt.Printf(testStringWithNewLine)
					},
				}
			},
		},
		testFunction{
			name: "fmt.Println",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					testFunc: func() {
						fmt.Println(testString)
					},
				}
			},
		},
	}

	testFunctions(functions, loopDuration, true)
}

func testFunctions(functions []testFunction, maxMilliSeconds int64, printLog bool) {
	outPut := ""
	maxDuration := time.Duration(0)
	maxName := ""
	duration := time.Duration(0)
	for currentAmount := 1; maxDuration.Milliseconds() < maxMilliSeconds; currentAmount *= 2 {
		for _, t := range functions {
			if printLog {
				fmt.Println("Testing", t.name, "currentAmount", currentAmount, "maxDuration:", maxDuration, "from", maxName)
				fmt.Println("waiting last duration * 2 (", duration*2, ") before starting next")
			}
			time.Sleep(duration * 2)
			instance := t.init()
			if instance.beforeTimer != nil {
				instance.beforeTimer()
			}
			start := time.Now()
			if instance.before != nil {
				instance.before()
			}
			for i := 0; i < currentAmount; i++ {
				instance.testFunc()
			}
			if instance.after != nil {
				instance.after()
			}
			end := time.Now()
			if instance.afterTimer != nil {
				instance.afterTimer()
			}
			duration = end.Sub(start)
			if printLog {
				if len(outPut) > 0 {
					outPut += "\n"
				}
				outPut += fmt.Sprintf("%s loop with %d iteration took %v", t.name, currentAmount, duration)
			}
			if duration > maxDuration {
				maxName = t.name
				maxDuration = duration
			}
		}
	}
	if printLog {
		fmt.Println(outPut)
	}
}

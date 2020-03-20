# Std Out Blocking Test
Testing the terminal std out blocking time

loops and prints "Hello World" until it lasts longer than 800ms seconds (if default parameters used)

multiple loops tested:
+ Noop (do not outputs anything)
+ Buffer Test with Variable and fmt.Println afterwards
+ Buffer Test with Array and fmt.Println afterwards
+ fmt.Print
+ fmt.Printf
+ fmt.Println

## Compile and run the program
```bash
go run outputBlockingTest.go
```

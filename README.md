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

## Run the program with [bufferMaster](https://github.com/anton98i/bufferMaster)
```bash
./outputBlockingTest.exe | ../bufferMaster/bufferMaster.exe
```

speed comparison (Windows 10, git bash inside vs code)

| Type | 16384 iterations | 16384 with bufferMaster | 65536 with bufferMaster |
| --- | --- | --- | --- |
|  |  | > | saved 56.39s, overhead: 184.81ms |
| Noop | 0s | 0s | 0s |
| Buffer Test with Variable and fmt.Println afterwards | 363.7932ms | 159.9098ms | 2.744432 |
| Buffer Test with Array and fmt.Println afterwards | 232.8561ms | 3.9972ms | 21.9873ms |
| os.Stdout.WriteString | 952.6983ms | 30.9817ms | 131.9147ms |
| os.Stdout.Write | 933.4675ms | 30.9849ms | 131.915ms |
| fmt.Print | 958.6342ms | 32.9807ms | 125.9181ms |
| fmt.Printf | 971.4362ms | 34.9809ms | 130.9151ms |
| fmt.Println | 966.4473ms | 32.9718ms | 131.9256ms |

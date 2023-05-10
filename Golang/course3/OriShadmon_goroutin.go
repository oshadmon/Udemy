/**
Concurrency is the idea where multiple processes run in parallel. While run condition is a type of concurrency where
multiple processes use the same components of the program, thus one influences the other.

The example below runs the execution_run method 10 times in parallel, making it an example for concurrent process. In
order to transform the current process, the execution_run method uses the "counter" variable which gets manipulated
within the execution_run method.

Note - run multiple times, the output of numbers differs, however for every iteration there are 10 values returned
*/
package main
import (
    "fmt"
    "sync"
)

var wgIns sync.WaitGroup
var counter int = 0

func execution_run() {
    /**
    Print execution run
    :args:
        x:int - current execution
    :output:
        execution counter
    */
    counter++
    fmt.Println("Execution Run: ", counter)
    wgIns.Done()
}

func main() {
    for i := 0 ; i < 10 ; i++ {
        wgIns.Add(1)
        go execution_run()
    }
    wgIns.Wait()
}

/**
 for i in {1..10} ; do echo "Run ${i}" ; go run Golang/course3/OriShadmon_goroutin.go ; done
Run 1
Execution Run:  1
Execution Run:  4
Execution Run:  5
Execution Run:  6
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  2
Execution Run:  3
Execution Run:  10
Run 2
Execution Run:  1
Execution Run:  3
Execution Run:  4
Execution Run:  5
Execution Run:  6
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  10
Execution Run:  2
Run 3
Execution Run:  1
Execution Run:  4
Execution Run:  5
Execution Run:  6
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  3
Execution Run:  2
Execution Run:  10
Run 4
Execution Run:  7
Execution Run:  1
Execution Run:  10
Execution Run:  6
Execution Run:  3
Execution Run:  9
Execution Run:  5
Execution Run:  4
Execution Run:  2
Execution Run:  8
Run 5
Execution Run:  6
Execution Run:  1
Execution Run:  8
Execution Run:  7
Execution Run:  4
Execution Run:  3
Execution Run:  5
Execution Run:  10
Execution Run:  9
Execution Run:  2
Run 6
Execution Run:  1
Execution Run:  4
Execution Run:  5
Execution Run:  6
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  3
Execution Run:  2
Execution Run:  10
Run 7
Execution Run:  1
Execution Run:  4
Execution Run:  5
Execution Run:  6
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  2
Execution Run:  3
Execution Run:  10
Run 8
Execution Run:  2
Execution Run:  1
Execution Run:  5
Execution Run:  4
Execution Run:  7
Execution Run:  6
Execution Run:  9
Execution Run:  8
Execution Run:  10
Execution Run:  3
Run 9
Execution Run:  1
Execution Run:  4
Execution Run:  5
Execution Run:  7
Execution Run:  8
Execution Run:  9
Execution Run:  10
Execution Run:  2
Execution Run:  3
Execution Run:  6
Run 10
Execution Run:  6
Execution Run:  4
Execution Run:  5
Execution Run:  7
Execution Run:  8
Execution Run:  2
Execution Run:  1
Execution Run:  9
Execution Run:  3
Execution Run:  10
**/
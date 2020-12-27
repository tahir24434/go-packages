package logging

import (
    "fmt"
    "time"
    "runtime"
)

var debug bool

func Debug(b bool) {
    debug = b
}

func Log(statement string) {
    if !debug {
        return
    }
    
    var file string
    var line int
    var ok bool
    // Caller reports file and line number information about function invocations on the calling goroutine's stack. 
    // The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller. 
    _, file, line, _ = runtime.Caller(1)
    //if !ok {
    //    file = "???"
    //    line = 0
    //}
    fmt.Printf("%s:%d: %s %s\n", file, line, time.Now().Format(time.RFC3339), statement)
    
}

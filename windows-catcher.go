package main

import (
    "fmt"
    "os/exec"
)

func main() {
    out, err := exec.Command( "powershell", "[Environment]::OSVersion.VersionString" ).Output()

    if err != nil {
        fmt.Printf("%s", err)
    }

    fmt.Println("Command has returned the following information: ")
    output := string(out[:])
    fmt.Println(output)
}

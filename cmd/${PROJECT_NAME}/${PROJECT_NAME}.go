// kick:render
package main

import (
    "fmt"
    "${GOSERVER}/${GOGROUP}/${PROJECT_NAME}/internal/options"
    "os"
)

func main() {
    opts := options.GetUsage(os.Args[1:],internal.Version)
    _ = opts
}

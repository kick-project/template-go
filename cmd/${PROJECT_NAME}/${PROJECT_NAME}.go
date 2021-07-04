// kick:render
package main

import (
    "os"

    "${GOSERVER}/${GOGROUP}/${PROJECT_NAME}/internal"
    "${GOSERVER}/${GOGROUP}/${PROJECT_NAME}/internal/options"
)

func main() {
    opts := options.GetUsage(os.Args[1:],internal.Version)
    _ = opts
}

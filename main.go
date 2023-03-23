package main
import (
    "fmt"
    "os"
    "bufio"
    )

const N = 1010

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
)

func main() {
    defer out.Flush()
    T := 0; fmt.Fscan(in, &T)
    fmt.Fprintln(out,"YES")
    fmt.Fprintln(out, "NO")   
}
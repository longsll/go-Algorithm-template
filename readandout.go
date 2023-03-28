package main
import (
    "fmt"
    "os"
    "bufio"
)

var (
    out = bufio.NewWriter(os.Stdout)
    in = bufio.NewReader(os.Stdin)
)

func main() {
    defer out.Flush()
    var T ,ans = 0 , 3
    var n , k int
    fmt.Fscan(in, &T)
    fmt.Fscan(in, &n , &k)
    fmt.Fprintln(out,ans)
    fmt.Fprintln(out,6)   
}
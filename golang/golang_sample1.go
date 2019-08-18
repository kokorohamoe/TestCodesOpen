package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
func main(){
    var s string
    var a,b string
    var sc = bufio.NewScanner(os.Stdin)
    sc.Scan()
    s  = sc.Text()
    slice := strings.Split(s, " ")
    a,_ = slice[0]
    b,_ = slice[1]
    fmt.Println( a  )
    fmt.Println( b  )
}
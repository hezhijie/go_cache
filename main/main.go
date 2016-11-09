package main

import (
        "common/cache"
        "fmt"
)
func main()  {


        c:=cache.ReadStr("GET","orderid")
        fmt.Println("main 2 c is ",c)


}

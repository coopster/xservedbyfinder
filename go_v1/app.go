package main

import "net/http"
import "fmt"
import "os"
import "strconv"

func getservedby() (header string) {
    resp, err := http.Get("http://localhost:3000/")
    header = ""
    if err != nil {
      fmt.Printf("-")

    } else {
      defer resp.Body.Close()

      if(resp.Status == "200 OK") {
          fmt.Printf(".")
          header = resp.Header.Get("X-served-by")
      } else {
          fmt.Printf("-")
      }
    }
    return
}

func main() {

    reqs := 10
    args := os.Args[1:]
    if(len(args) != 0) {
        amt, _ := strconv.Atoi(args[0])
        reqs = amt
    }
    fmt.Printf("Reqs: %v\n", reqs)

    found := make(map[string]int)
    for i:=0; i< reqs; i++ {
        code := getservedby()
        if _, exists := found[code]; !exists {
            found[code]=0
        }
        found[code] = found[code] + 1
    }
    fmt.Printf("\n")
    for k,_ := range found {
        fmt.Printf("[%v]\t%v\n", k, found[k]);
    }

}
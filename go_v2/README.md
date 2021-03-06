<!---
    Copyright (c) 2016 Mark Roberts
    Released under the MIT license. See the file LICENSE for details.
--->
#Go
This is an attempt at an idiomatic Go implementation. It makes a number of HTTP HEAD requests concurrently, feeding the results back through a channel to a function that finds and prints the unique header values as they are found.

##Building it
````sh
go_v2 $ go build xservedby.go
````

##Running it
````sh
go_v2 $ ./xservedby -h
      Usage of ./xservedby:
        -c int
              Number of HTTP requests to make concurrently (default 8)
        -r int
              Maximum number of requests to make (default 100)
        -url string
              Url to request (default "http://localhost:3000/")
````

##Running through Docker

Building and running, where the *fakeendpoint* is running i.e. on 192.168.0.100.

````
docker build -t go_v2 .
docker run go_v2 -c 8 -r 100 -url http://192.168.0.100
````


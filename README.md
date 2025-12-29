# Web-Server in Golang

## Overview
It's nothing fancy, just a simple web server in Go.  
I want to practice Go, that's why I made this server. It doesn't use any external package, built using only the `net/http` package. I didn't use any proper file structuring since it was for practicing purpose, so I wrote the entire code in `main.go`.

---

## What I Learned

### Creating a Router
```go
router := http.NewServeMux()
```
This is used for creating a new ServeMux which acts as a router that handles multiple incoming requests.

Registering Handlers
```go
router.HandleFunc("/", handleRoot)
```
HandleFunc registers the handler function to the pattern present in " ".
In this case, the handleRoot func is used to serve the request which is coming on the / path.

Also, the handlers should get two parameters:
```go
func handleRoot(w http.ResponseWriter, r *http.Request)
```
ResponseWriter is used to send response to the client
Request is used to access the incoming request like r.Body

Starting the Server
```go
http.ListenAndServe(":4000", router)
```
ListenAndServe listens on the TCP network address 4000 and then calls Serve with the handler to handle requests on incoming connections.

### Working With JSON

I also learnt how to decode the JSON from the request and put the decoded data on struct.
Also how to send struct in JSON format to the client.

### Data Storage

I didn't use any DB.
Instead, I used a local datastructure (a map) to store Users associated with the ID:
```go
var UserCache = make(map[int]User)
```

I also used mutex for locking the map to prevent race condition, making the requests thread safe:
```go
var cacheMutex sync.RWMutex
```

# injekt

Pluggable service injector

[![GoDoc](https://godoc.org/github.com/abiosoft/injekt?status.svg)](https://godoc.org/github.com/abiosoft/injekt)
[![Build Status](https://travis-ci.org/abiosoft/injekt.svg?branch=master)](https://travis-ci.org/abiosoft/injekt)
[![Go Report Card](https://goreportcard.com/badge/github.com/abiosoft/injekt)](https://goreportcard.com/report/github.com/abiosoft/injekt)
[![Coverage Status](https://coveralls.io/repos/github/abiosoft/injekt/badge.svg?branch=master)](https://coveralls.io/github/abiosoft/injekt?branch=master)


Injekt is a pluggable service injector for any project or framework. 
Injekt adds service injection support by wrapping a custom function (with services as parameters) with the required function.

### Usage (http Handler example)
Write custom function. `session` will be injected as well as specified parameters of the required function.
```go
func sessionInfo(w http.ResponseWriter, session *Session) {
    if session == nil { 
        // show login page
     }
     ...
}
```
Wrap custom function. This returns an `interface{}` that can be asserted to the required function type.
```go
func main(){
    var h http.HandlerFunc
    inj := injekt.New(h)
    ...
    http.HandleFunc("/", inj.Wrap(sessionInfo).(http.HandlerFunc))
}
```
Register services before function is executed.
```go
session := ...
inj.Register(session)
```
Simple and useful.
Check the [docs](https://godoc.org/github.com/abiosoft/injekt) for more.

### License
Apache 2
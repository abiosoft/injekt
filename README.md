# injekt

Pluggable service injector

[![GoDoc](https://godoc.org/github.com/abiosoft/injekt?status.svg)](https://godoc.org/github.com/abiosoft/injekt)
[![Build Status](https://travis-ci.org/abiosoft/injekt.svg?branch=master)](https://travis-ci.org/abiosoft/injekt)
[![Go Report Card](https://goreportcard.com/badge/github.com/abiosoft/injekt)](https://goreportcard.com/report/github.com/abiosoft/injekt)


Injekt is a pluggable service injector for any project or framework. 
Injekt allows you to transform an existing function to a new function 
with required services as function parameters.

### Usage [http Handler example]
Write your custom functions.
```go
func sessionInfo(w http.ResponseWriter, session *Session) {
    if session == nil { 
        // show login page
     }
     ...
}
```
Wrap custom functions.
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


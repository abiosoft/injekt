// Package injekt provides a pluggable service injector for any project or framework.
// Service injection is achieved by wrapping a custom function (with services as parameters) with the required function.
//
// http.Handler example:
//
// Write custom function. `session` will be injected as well as specified parameters of the required function.
//  func sessionInfo(w http.ResponseWriter, session *Session) {
//    if session == nil {
//      // show login page
//    }
//    ...
//  }
//
// Wrap custom function. This returns an `interface{}` that can be asserted to the required function type.
//  func main(){
//    var h http.HandlerFunc
//    inj := injekt.New(h)
//    ...
//    http.HandleFunc("/", inj.Wrap(sessionInfo).(http.HandlerFunc))
//  }
//
// Register services before function is executed.
//  session := ...
//  inj.Register(session)
//
package injekt

// Jack Schefer, began 9/21/16
//
package main
//
import (
   "errors"
   "fmt"
   "net/http"
   "os"
)
//
//
func main() {
   //
   var fname string
   if len( os.Args ) < 2 {
      panic( errors.New("No filename specified.") )
   } else {
      fname = os.Args[ 1 ]
   }
   //
   http.HandleFunc("/" + fname, handler)
   http.ListenAndServe(":8080", nil)
   //
}
//
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Welcome to my local server</h1>")
}
//
//
// End of file.

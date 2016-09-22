// Jack Schefer, began 9/21/16
//
package main
//
import (
   "fmt"
   "html/template"
   "net/http"
)
//
type Page struct {
   Title string
   body []byte
}
//
func main() {
   //
   http.HandleFunc("/about.html", about_handler)
   http.HandleFunc("/play.html", play_handler)
   http.HandleFunc("/", index_handler)
   http.ListenAndServe(":8080", nil)
   //
}
//
func index_handler(w http.ResponseWriter, r *http.Request) {
   title := "index.html"
   p := &Page{Title: title}
   t, _ := template.ParseFiles("index.html")
   t.Execute(w, p)
}
//
func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, header() + "<h1>You are on page " + r.URL.Path + "</h1>")
}
func play_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, header() "<h1>You are on page " + r.URL.Path + "</h1>")
}
//
func header() string {
      return "<div class=\"link-wrapper\">
         <div class=\"container\">
            <div class=\"row\">
               <div class=\"col-sm-6\"></div>
               <div class=\"col-sm-2\">
                  <a href=\"index.html\">
                     Home
                  </a>
               </div>
               <div class=\"col-sm-2\">
                  <a href=\"about.html\">
                     About
                  </a>
               </div>
               <div class=\"col-sm-2\">
                  <a href=\"play.html\">
                     Play
                  </a>
               </div>
            </div>
         </div>
      <div class=\"header-wrapper\">
         <div class=\"container\">
            <div class=\"row\">
               <div class=\"col-xs-12\">
                  <h1>Welcome to {{.Title}}</h1>
                  <h3>A QuizBowl Site for the Future</h3>
               </div>
            </div>
         </div>
      </div>"
}
//
//
// End of file.

package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fn := req.FormValue("fname")
	ln := req.FormValue("lname")

	io.WriteString(res, `
	<!DOCTYPE html>
<html>
   <head>
      <title>HTML Backgorund Color</title>
   </head>
   <body style="background-color:grey;">
	<form action="/action_page.php" method="get">
	<label for="fname">First name:</label>
	<input type="text" id="fname" name="fname"><br><br>
	<label for="lname">Last name:</label>
	<input type="text" id="lname" name="lname"><br><br>
	<input type="submit" value="Submit">
  </form>
  </body>
</html>`+fn+" "+ln)
}

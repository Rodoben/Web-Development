package main

import (
	"fmt"
      "text/template"
)


func main(){

	name:= "Ronald Benjamin"

	ftp:= '	
	 <!Doctype html>
	<html lang= "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello World </title> 
	</head>
	<body>
    <h1> '+ name +'</h1>
	</body>
    </html>
    '
   fmt.Println(ftp)

}
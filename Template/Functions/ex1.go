package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	motto string
}

type car struct {
	Name  string
	Model string
	Door  string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New(" ").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func main() {
	s1 := sage{Name: "Ronald", motto: "Code in life"}
	s2 := sage{Name: "FFFmkk", motto: "cdncdn dcndslk kdnclkd"}
	s3 := sage{Name: "nellam", motto: "cds vc dds"}

	c1 := car{Name: "fordddd", Model: "f123", Door: "2"}
	c2 := car{Name: "BMWdd", Model: "fr13", Door: "24"}

	sages := []sage{s1, s2, s3}
	cars := []car{c1, c2}

	cint, er := strconv.Atoi(c1.Door)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Printf("type: %T  - value:%#V", cint)

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}

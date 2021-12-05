package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}
type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl1.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSI-123", "Intro", "5"},
				course{"KSI-123", "Intro123", "6"},
				course{"LSI-123", "Intro234", "8"},
			},
		},

		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"LSI-123", "Intro234", "8"},
				course{"LSI-123", "Intro234", "8"},
				course{"LSI-123", "Intro234", "8"},
				course{"LSI-123", "Intro234", "8"},
				course{"LSI-123", "Intro234", "8"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{"LS2-123", "Intro234", "8"},
				course{"LSr-123", "Intro234", "8"},
				course{"LS4-123", "Intro234", "8"},
				course{"LS9-123", "Intro234", "8"},
				course{"LS3-123", "Intro234", "8"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}

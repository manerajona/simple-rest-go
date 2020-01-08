package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type sage struct {
	Name  string
	Motto string
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fdateMDY": monthDayYear,
	"fdateDMY": dayMonthYear,
	"fdbl":     double,
	"fsq":      square,
	"fsqrt":    sqRoot,
}

func init() {
	//tpl = template.Must(template.ParseGlob("tpls/*.gohtml"))
	tpl = template.Must(
		template.New("does_not_matters").Funcs(fm).ParseGlob("tpls/*.gohtml"),
	)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3] // retrive the first 3 letters
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func printErrIfExists(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Parse two templates
	// tpl, err := template.ParseFiles("tpl1.gohtml", "tpl2.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Parse all with certain extension
	// tpl, err := template.ParseGlob("*.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Now we replace last two blocks with func init()

	// 1.Use Array to iterate into the template
	arr := []string{"Ella", "Patt", "Tom"}

	nf1, err1 := os.Create("out/index1.html")
	printErrIfExists(err1)
	defer nf1.Close()

	err1 = tpl.ExecuteTemplate(nf1, "tpl1.gohtml", arr)
	printErrIfExists(err1)

	// 2.Use name variables into templates
	nf2, err2 := os.Create("out/index2.html")
	printErrIfExists(err2)
	defer nf2.Close()

	err2 = tpl.ExecuteTemplate(nf2, "tpl2.gohtml", "Jonathan")
	printErrIfExists(err2)

	// 3.Use map to iterate k,v into the template
	mp := map[string]string{
		"India":    "Gandi",
		"Meditate": "Buddha",
		"Love":     "Yourself",
	}

	nf3, err3 := os.Create("out/index3.html")
	printErrIfExists(err3)
	defer nf3.Close()

	err3 = tpl.ExecuteTemplate(nf3, "tpl3.gohtml", mp)
	printErrIfExists(err3)

	// 4.Use structure to be accessed into the template
	toyota := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	nf4, err4 := os.Create("out/index4.html")
	printErrIfExists(err4)
	defer nf4.Close()

	err4 = tpl.ExecuteTemplate(nf4, "tpl4.gohtml", toyota)
	printErrIfExists(err4)

	// 5. Use functions into the template
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err5 := tpl.ExecuteTemplate(os.Stdout, "tpl5.gohtml", data)
	printErrIfExists(err5)

	// 6. Fmt date time
	err6 := tpl.ExecuteTemplate(os.Stdout, "tpl6.gohtml", time.Now())
	printErrIfExists(err6)

	// 7. Fmt using pipelines
	err7 := tpl.ExecuteTemplate(os.Stdout, "tpl7.gohtml", 3)
	printErrIfExists(err7)

	// 8. Use If & And func into the template
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err8 := tpl.ExecuteTemplate(os.Stdout, "tpl8.gohtml", users)
	printErrIfExists(err8)

	// 9. Modularize templates
	err9 := tpl.ExecuteTemplate(os.Stdout, "tpl9.gohtml", "template")
	printErrIfExists(err9)
}

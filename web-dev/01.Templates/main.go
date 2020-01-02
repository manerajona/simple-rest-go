package main

import (
	"log"
	"os"
	"text/template"
)

type car struct {
	Manufacter string
	Model      string
	Doors      int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
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

	// Execute whatever (good if you have just one)
	nf1, err1 := os.Create("out/index1.html")
	printErrIfExists(err1)
	defer nf1.Close()

	// This Array will iterate into the template
	arr := []string{"Ella", "Patt", "Tom"}

	err1 = tpl.Execute(nf1, arr)
	printErrIfExists(err1)

	// Execute specific template
	nf2, err2 := os.Create("out/index2.html")
	printErrIfExists(err2)
	defer nf2.Close()

	err2 = tpl.ExecuteTemplate(nf2, "tpl2.gohtml", "Jonathan")
	printErrIfExists(err2)

	// This map will iterate into the template
	sages := map[string]string{
		"India":    "Gandi",
		"Meditate": "Buddha",
		"Love":     "Yourself",
	}

	nf3, err3 := os.Create("out/index3.html")
	printErrIfExists(err3)
	defer nf3.Close()

	err3 = tpl.ExecuteTemplate(nf3, "tpl3.gohtml", sages)
	printErrIfExists(err3)

	// This structure will be accessed in the template
	toyota := car{
		Manufacter: "Toyota",
		Model:      "Corolla",
		Doors:      4,
	}

	nf4, err4 := os.Create("out/index4.html")
	printErrIfExists(err4)
	defer nf4.Close()

	err4 = tpl.ExecuteTemplate(nf4, "tpl4.gohtml", toyota)
	printErrIfExists(err4)
}

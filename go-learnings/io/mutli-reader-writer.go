package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func multiread() {

	f, _ := os.Open("")

	header := strings.NewReader("<msg>")
	body_file, _ := os.Open("api_response.file")
	footer := strings.NewReader("</msg>")

	mr := io.MultiReader(header, body_file, footer)
	io.Copy(f, mr)

}
func multiwrite() {
	f1, _ := os.Create("file1")
	f2, _ := os.Create("file2")
	f3, _ := os.Create("file3")

	mw := io.MultiWriter(f1, f2, f3)

	//writes the below content into all writers
	fmt.Fprintln(mw, "line1")
	fmt.Fprintln(mw, "line2")

}

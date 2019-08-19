/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/schema-go"
	"github.com/kazekim/schema-go/example/enum"
	"net/http"
)

type B struct {
	X    int64
	Y float64
}

type A struct {
	Name        string
	YearOfBirth *int
	Pos string `schema:"Position"`
	Doc enum.DocumentType
	B B
}

func main() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}


func homePage(w http.ResponseWriter, r *http.Request) {

	//var a entity.AgencyModel
	var a A

	r.ParseForm()
	fmt.Println(r.PostForm)
	decoder := schema.NewDecoder()

	cp := schema.CustomParserMap{
		"enum.DocumentType": enum.ParseDocumentTypeReflectValue,
	}
	decoder.CustomParser(cp)
	decoder.Decode(&a, r.PostForm)

	fmt.Println(a)

	if a.Doc == enum.DOCUMENTTYPE_COMPANYID {
		fmt.Println("YOYOYOYO")
	}


	//p := wwhttpparser.NewParser(r)

	//p.CustomParser(cp)
	//err := p.ParseFormRequest(&a)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(a)
	////fmt.Println(*a.YearOfBirth)

}
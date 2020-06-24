package createhtml

import (
	"fmt"
	"log"
	"testing"
)

func TestHtml(t *testing.T) {
	val := []interface{}{75, 90, 80, 40}

	attr := []Attributes{
		Attributes{
			Name:    "data-id",
			Options: []string{"dd"},
		},
	}

	exp := []Expression{}
	exp1 := Expression{
		SecondStatement: 50,
		Operator:        LessOrEqual,
		Expected:        "text-danger",
		Default:         "text-default",
	}
	exp = append(exp, exp1)

	data := Data{
		Value:             val,
		Expression:        exp,
		AttributesOptions: attr,
	}

	ht := Html{}
	li, err := ht.AddData(data).AddClass("text-xs").Tag("th")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(li)
}

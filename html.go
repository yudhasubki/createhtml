package createhtml

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
)

type Operator int

var TagName = struct {
	List      string
	TableHead string
	TableBody string
}{
	"li",
	"th",
	"td",
}

var MapTag map[string]string

const (
	ErrTagNotAvailable  = "TagNotAvailable"
	ErrCombineParamater = "parameter need string or []string given"
)

const (
	Equal Operator = iota + 1
	NotEqual
	LessOrEqual
	Less
	GreaterOrEqual
	Greater
)

type Html struct {
	Data     Data
	tag      Tag
	elements string
}

type Data struct {
	Value             []interface{}
	AttributesOptions []Attributes
	Expression        []Expression
}

type Attributes struct {
	Name    string
	Options interface{}
}

type Tag struct {
	name       string
	attributes string
	class      string
}

func init() {
	makeMapList()
}

func (h *Html) AddClass(class interface{}) *Html {
	css, err := combine(class)
	if err != nil {
		log.Fatalf("err %v : \n", err)
		return h
	}
	h.tag.class = css

	return h
}

func (h *Html) AddData(data Data) *Html {
	h.Data = data
	return h
}

func (h *Html) makeAttribute(attributes chan string) {
	options := make([]string, 0)
	for _, attr := range h.Data.AttributesOptions {
		attrs, err := combine(attr.Options)
		if err != nil {
			log.Fatalf("error make attributes : %v", err)
		}
		options = append(options, fmt.Sprintf("%s='%s'", attr.Name, attrs))
	}

	attributes <- strings.Join(options, ",")
}

func (h *Html) makeExpression(v interface{}, expression chan string) {
	var expressionReplacer *string
	for _, e := range h.Data.Expression {
		e.firstStatement = v
		res := e.Expression()
		if expressionReplacer == nil || &res != expressionReplacer {
			expressionReplacer = &res
		}
	}
	expression <- *expressionReplacer
}

func makeMapList() {
	v := reflect.ValueOf(TagName)
	MapTag = make(map[string]string)
	once := sync.Once{}
	once.Do(func() {
		for i := 0; i < v.NumField(); i++ {
			MapTag[v.Field(i).String()] = v.Field(i).String()
		}
	})
}

func (h *Html) Tag(tagName string) (string, error) {
	if _, exist := MapTag[tagName]; !exist {
		return "", errors.New(ErrTagNotAvailable)
	}

	h.tag.name = tagName
	h.render()
	return h.elements, nil
}

func (h *Html) render() {
	elements := make([]string, 0)
	if len(h.Data.AttributesOptions) > 0 {
		options := make(chan string)
		exps := make(chan string)
		for _, v := range h.Data.Value {
			go h.makeAttribute(options)
			go h.makeExpression(v, exps)
			opt, exp := <-options, <-exps
			elements = append(elements, h.elementTemplate(h.tag.name, h.tag.class+" "+exp, opt, v, h.tag.name))
		}
	}

	h.elements = strings.Join(elements, "\n")
}

func (h *Html) elementTemplate(params ...interface{}) string {
	return fmt.Sprintf("<%s class='%s' %s> %v <%s/>", params...)
}

func StringOrInt(src interface{}) (bool, bool) {
	var isString, isInt bool
	switch v := reflect.ValueOf(src); v.Kind() {
	case reflect.String:
		isString = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		isInt = true
	default:
		return false, false
	}
	return isString, isInt
}

func combine(src interface{}) (string, error) {
	v := reflect.ValueOf(src)

	var data string
	switch t := v.Type(); t.Kind() {
	case reflect.Slice:
		if t.Elem().Kind() == reflect.String {
			data = strings.Join(src.([]string), " ")
		} else {
			return "", errors.New(fmt.Sprintf("%v %v", ErrCombineParamater, t.Elem().Kind()))
		}
	case reflect.String:
		data = src.(string)
	default:
		return "", errors.New(fmt.Sprintf("%v %v", ErrCombineParamater, t.Elem().Kind()))
	}

	return data, nil
}

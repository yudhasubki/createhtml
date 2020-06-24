## go createhtml
createhtml is simple template Engine to render HTML golang in back-end as string.

## Installation

```bash
go get -u github.com/yudhasubki/createhtml
```

## Usage

```go
import "github.com/yudhasubki/createhtml"


val := []interface{}{75, 90, 80, 40}

// this will create <th data-id='dd'>
attr := []createhtml.Attributes{
    Attributes{
        Name:    "data-id",
        Options: []string{"dd"},
    },
}

// if there have any val in class tag 
// will be this will auto append in class tag ex : <th class'another-val text-default'>
// if val > condition <th class='text-default'> or val < condition <th class='text-danger'>
// operator available can check in struct Operator{}
exp := []createhtml.Expression{}
exp1 := Expression{
    Condition: 50,
    Operator:  LessOrEqual,
    Expected:  "text-danger",
    Default:   "text-default",
}
exp = append(exp, exp1)

data := createhtml.Data{
    Value:             val,
    Expression:        exp,
    AttributesOptions: attr,
}

// you can use chain method with ht.AddData(data) or pass in Struct, ex: createhtml.Html{ Data:data }
ht := createhtml.Html{}

// method AddClass will auto append class tag
// method Tag used to render your tag
// available tag in struct TagName{}
th, err := ht.AddClass("text-xs").Tag("th")
if err != nil {
    log.Fatal(err)
}
fmt.Println(th)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
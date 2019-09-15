# go-astinspector

Personal project.

**Doc coming soon with more features.**

## ➡ install

```
go get -u github.com/kovacou/go-astinspector
```

## ➡ usage

```go
f := astinspector.ParseFile("../go-astinspector/testdata/struct.go")
f.PackageName() // testdata

u := f.StructByName("user")
u.AddField("CustomField", "string")
u.Name() // User

for _, field := range u.Fields() {
    fmt.Printf("%s %s\n", field.Name(), field.Type())
}

// Print: 
//
// ID uint64
// Name string
// Birthday *time.Time
// CustomField string

```

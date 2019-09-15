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
f.Structs().Names() // []string{"User", "Test"}
f.Structs("User", "Toto").Names() // []string{"User"}
f.PackageName() // testdata 

if u := f.StructByName("User"); u != nil {
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
}


```

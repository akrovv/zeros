## Zeros linter

#### Use `var` for Zero Value Structs

When all the fields of a struct are omitted in a declaration, use the `var`
form to declare the struct.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
user := User{}
```

</td><td>

```go
var user User
```

</td></tr>
</tbody></table>


### How to start
* git clone __[link]__
* go build -buildmode=plugin cmd/zeros/main.go
* __[!!! you need compiled file from https://github.com/golangci/golangci-lint (make build_plugin)]__
* ./golangci-lint -c golangci.yml run __[source code]__

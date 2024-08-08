## Unmarshal json with and without json field tag

Compares the performance of an Unmarshal with and without the hint of a json:"field" tag

### JSON field tag

When unmarshaling a json to a struct field, you can inform a field tag `json:"field"` as a hint.
This will be used by the Unmarshaler to map the json attributes to struct fields.
Commonly used also to convert snake_case json attributes to CamelCase struct fields.

If no json tag is provided, the Unmarshal function will try some other forms ignoring case.
This is kind of necessary because exported struct fields need to start with an upper case, while that's uncommon in json.

But this fallback of the unhinted struct comes with a cost.
The idea of this benchmark is to compare how much is this cost of having untagged struct fields.

### Running

You can run it with

```
go test -benchmem -run=^$ -count=10 -bench
```

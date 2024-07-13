package interfacestruct

type SomeInterface interface {
	Some() string
}

// SomeType implements SomeInterface
type SomeType struct {
	name string
}

func (t SomeType) Some() string {
	return t.name
}

func main() {
	var islice []SomeInterface

	islice = append(islice, SomeType{name: "test"})
	islice = append(islice, SomeInterface(SomeType{name: "test"}))
}

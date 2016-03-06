package etc

const (
	Basic = "basic"
	Tree  = "tree"
	HTTP  = "http"
	My    = "my"
)

var types = [...]string{
	Basic,
	Tree,
	HTTP,
	My,
}

func ValidateType(t string) bool {
	for _, typ := range types {
		if t == typ {
			return true
		}
	}
	return false
}

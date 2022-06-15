package inter

type Rule interface {
	Generate() interface{}
}

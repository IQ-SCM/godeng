package inter

type Field interface {
	Key() string
	Value() interface{}
}

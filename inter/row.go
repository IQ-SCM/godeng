package inter

type Row interface {
	Items() []Pair
}

type Pair interface {
	Key() string
	Value() interface{}
}

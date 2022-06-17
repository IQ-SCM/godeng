package godeng

type ConfigItem struct {
	key   string
	typ   string
	max   float64
	min   float64
	len   int64
	begin float64
	step  float64
	enums []interface{}
}

type Config struct {
	items []*ConfigItem
}

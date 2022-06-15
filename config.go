package godeng

type ConfigItem struct {
	key string
	typ string
	max int64
	min int64
	len int64
}

type Config struct {
	items []ConfigItem

	o string
	f string
}

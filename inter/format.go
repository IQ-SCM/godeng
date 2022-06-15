package inter

type Format interface {
	Format(row Row) []byte
}

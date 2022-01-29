package send

type IFormat interface {
	Format() ([]byte, error)
}
type IOutput interface {
	Output(content []byte) error
}

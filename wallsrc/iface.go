package wallsrc

import "io"

type Source interface {
	Get() error
	WriteBody(w io.Writer) error
	String() string
	Clean()
}

package openapi3

var (
	_ Builder = (*builder)(nil)
)

type Builder interface {
	Info() *Info
}
type builder struct {
	t T
}

func (b *builder) Info() *Info {
	if b.t.Info == nil {
		b.t.Info = &Info{}
	}
	return b.t.Info
}

func Build() Builder {
	return &builder{
		t: T{
			OpenAPI: "3.0.0",
		},
	}
}

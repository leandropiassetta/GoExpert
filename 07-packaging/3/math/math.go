package math

type math struct {
	a     int
	b     int
	class string
}

// this way we can create a new instance of Math, protecting the fields a and b from being changed directly by the user

func NewMath(a, b int, class string) math {
	return math{
		a:     a,
		b:     b,
		class: class,
	}
}

// this method is only accessible by the Math struct, the pointer is used to change the value of the struct
func (m *math) Sum() int {
	return m.a + m.b
}

func (m *math) Class() string {
	return m.class
}

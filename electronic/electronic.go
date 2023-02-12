package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	b  string
	m  string
	t  string
	os string
}

func (i *applePhone) Brand() string {
	return "apple"
}

func (i *applePhone) Model() string {
	return i.m
}

func (i *applePhone) Type() string {
	return "smartphone"
}

func (i *applePhone) OS() string {
	return "IOS"
}

type androidPhone struct {
	b  string
	m  string
	t  string
	os string
}

func (a *androidPhone) Brand() string {
	return a.b
}

func (a *androidPhone) Model() string {
	return a.m
}

func (a *androidPhone) Type() string {
	return "smartphone"
}

func (a *androidPhone) OS() string {
	return a.os
}

type radioPhone struct {
	b  string
	m  string
	t  string
	bc int
}

func (r *radioPhone) Brand() string {
	return r.b
}

func (r *radioPhone) Model() string {
	return r.m
}

func (r *radioPhone) Type() string {
	return r.t
}

func (r *radioPhone) ButtonsCount() int {
	return r.bc
}

// Создает копию структуры applePhone.
func NewApplePhone(b, m, t, os string) applePhone {
	return applePhone{b: b, m: m, t: t, os: os}
}

// Создает копию структуры androidPhone.
func NewAndroidPhone(b, m, t, os string) androidPhone {
	return androidPhone{b: b, m: m, t: t, os: os}
}

// Создает копию структуры radioPhone.
func NewRadioPhone(b, m, t string, bc int) radioPhone {
	return radioPhone{b: b, m: m, t: t, bc: bc}
}

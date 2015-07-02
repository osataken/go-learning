package learning

type Outer struct {
	Inner
	level string
}

type Inner struct {
}

func (outer *Outer) doSomething() string {
	return "outer method"
}

func (inner *Inner) doSomething() string {
	return "inner method"
}
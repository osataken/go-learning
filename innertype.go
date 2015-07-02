package learning

type Outer struct {
	Inner
}

type Inner struct {
}

func (outer *Outer) doSomething() string {
	return "outer method"
}

func (inner *Inner) doSomething() string {
	return "inner method"
}
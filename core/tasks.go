package core

func (core *Core) Tasks(fn func() error) error {
	return fn()
}

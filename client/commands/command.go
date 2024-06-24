package commands

type Command interface {
	Execute(data any) error
}

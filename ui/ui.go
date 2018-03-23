package ui

type UI interface {
	Ask(string) (string, error)

	Output(string)

	Info(string)

	Error(string)
}

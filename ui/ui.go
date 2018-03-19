package ui

type UI interface {
	Output(string)

	Info(string)

	Error(string)
}

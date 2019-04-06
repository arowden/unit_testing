package device

type Device interface {
	Execute(cmd string) (output string, err error)
	Close() error
}

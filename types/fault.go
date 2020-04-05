package types

func NewFault(code FaultCode, message string, causes ...Fault) Fault {
	return &fault{
		code:    code,
		message: message,
		causes:  causes,
	}
}

type FaultCode string

type Fault interface {
	error

	Code() FaultCode
	Message() string
	Causes() []Fault
}

type fault struct {
	Fault

	code    FaultCode `json:"code" yaml:"code" form:"code"`
	message string    `json:"message" yaml:"message" form:"message"`
	causes  []Fault   `json:"causes" yaml:"causes" form:"causes"`
}

func (f *fault) Code() FaultCode {
	return f.code
}

func (f *fault) Message() string {
	return f.message
}

func (f *fault) Causes() []Fault {
	return f.causes
}

func (f *fault) Error() string {
	return printError(f, "")
}

func printError(f Fault, indentation string) string {
	err := indentation + f.Message()

	if len(f.Causes()) > 0 {
		err = err + " caused by:\n"
		for _, cause := range f.Causes() {
			err = err + printError(cause, indentation+"  ") + "\n"
		}
		err = err[:len(err)-1]
	}

	return err
}

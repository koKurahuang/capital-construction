package trace

import (
	"bytes"
	"fmt"
	"github.com/koKurahuang/capital-construction/config"
	"runtime"
)

// callstacks will be created when New
// So the default errors will not have callstacks since they are created at the very fist

var errs = make(map[string]One)

func Initilize() {
	setSysLang()
	loadDefaultErrors()
	loadFileErrors()
}

type One struct {
	ModuleCode   string
	ModuleName   string
	ErrorCode    string
	ErrorName    string
	EngMsg       string
	ZhMsg        string
	Memo         string
	lt           LangType
	ifCallStacks bool
	stackMsg     string
	args         []interface{}
}

func (o *One) Error() string {
	if o == nil {
		return nilPointerError.Error()
	}

	if o.lt == 0 {
		o.lt = sysLangType
	}

	var errStr string
	switch o.lt {
	case ENG:
		errStr = fmt.Sprintf(o.EngMsg, o.args...)
	case ZH:
		errStr = fmt.Sprintf(o.ZhMsg, o.args...)
	default:
		errStr = fmt.Sprintf(o.EngMsg, o.args...)
	}

	if o.ifCallStacks {
		errStr = errStr + "\n" + o.stackMsg
	}

	return errStr
}

func setSysLang() {
	lang := config.GetString("trace.lang")
	switch lang {
	case "en":
		sysLangType = ENG
	case "ZH":
		sysLangType = ZH
	default:
		// todo
		// log.print("Not find lang type in config file")
		sysLangType = ENG
	}
}

func loadFileErrors() {
	path := config.GetPath("trace.json")
	getJson(path)
}

func loadDefaultErrors() {
	errs[unknownError.GetErrorCode()] = *unknownError
	errs[readFileError.GetErrorCode()] = *readFileError
	errs[indexOutOfRangeError.GetErrorCode()] = *indexOutOfRangeError
	errs[nilPointerError.GetErrorCode()] = *nilPointerError
}

func New(moduleCode string, errorCode string, infos ...interface{}) error {
	one, ok := errs[setErrorCode(moduleCode, errorCode)]
	if !ok {
		return unknownError
	}
	one.args = infos
	one.ifCallStacks = sysIfCallStacks

	stack := make([]uintptr, MaxCallStackLength)
	length := runtime.Callers(1, stack[:])
	stack = stack[:length-2]

	one.stackMsg =  getStack(stack)

	return &one
}

func (o *One) SetLang(lang LangType) {
	o.lt = lang
}

func (o *One) SetIfCallStacks(ifCallStack bool) {
	o.ifCallStacks = ifCallStack
}

func (o *One) GetErrorCode() string {
	return setErrorCode(o.ModuleCode, o.ErrorCode)
}

func setErrorCode(moduleCode, errCode string) string {
	return moduleCode + "-" + errCode
}

func getStack(stack []uintptr) string {
	buf := bytes.Buffer{}
	if stack == nil {
		return fmt.Sprintf("No call stack available")
	}
	// this removes the core/errors module calls from the callstack because they
	// are not useful for debugging
	const firstNonErrorModuleCall int = 2
	stack = stack[firstNonErrorModuleCall:]
	for i, pc := range stack {
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(pc)
		if i != len(stack)-1 {
			buf.WriteString(fmt.Sprintf("%s:%d %s\n", file, line, f.Name()))
		} else {
			buf.WriteString(fmt.Sprintf("%s:%d %s", file, line, f.Name()))
		}
	}
	return fmt.Sprintf("%s", buf.Bytes())
}

package trace

var sysLangType LangType = ENG
var sysIfCallStacks = true

var unknownError = &One{
	ModuleCode:"0000",
	ModuleName: "Default",
	ErrorCode: "0001",
	ErrorName: "UnknownError",
	EngMsg: "Unknown error occurs, please check for detail",
}

var readFileError = &One{
	ModuleCode:"0000",
	ModuleName: "Default",
	ErrorCode: "0001",
	ErrorName: "ReadFileError",
	EngMsg: "Can not read file %s",
}

var indexOutOfRangeError = &One{
	ModuleCode:"0000",
	ModuleName: "Default",
	ErrorCode: "0002",
	ErrorName: "IndexOutOfRangeError",
	EngMsg: "Index out of range",
}

var nilPointerError = &One{
	ModuleCode:"0000",
	ModuleName: "Default",
	ErrorCode: "0003",
	ErrorName: "NilPointerError",
	EngMsg: "Nil pointer",
}


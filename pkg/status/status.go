package status

type StatusStruct struct {
	Successbool bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func (s StatusStruct) Success(message string, data interface{}) StatusStruct {
	return StatusStruct{
		Successbool: true,
		Message:     message,
		Data:        data,
	}
}

func (s StatusStruct) Error(message string) StatusStruct {
	return StatusStruct{
		Successbool: false,
		Message:     message,
	}
}

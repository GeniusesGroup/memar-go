# Command

## This package way
```go
type InitReq struct {
	name string
}

//libgo:impl protocol.Init_Request
func (req *InitReq) Name() string { return req.name }

//libgo:impl libgo/protocol.Object
func (req *InitReq) Fields() []protocol.Object_Member_Field   { return ... }
func (req *InitReq) Methods() []protocol.Object_Member_Method { return ... }

//libgo:impl libgo/protocol.CommandLineArguments
func (req *InitReq) FromCLA(arguments []string) (remaining []string, err protocol.Error) {
	remaining, err = cmd.FromCLA(req, arguments)
	return

}
func (req *InitReq) ToCLA() (arguments []string, err protocol.Error) {
	arguments, err = cmd.ToCLA(req)
	return
}

type service struct {}

func (ser *service) ServeCLA(arguments []string) (err protocol.Error) {
	var req InitReq
	_, err = req.FromCLA(arguments)
	if err != nil {
		return
	}

	var res NewRes
	res, err = ser.Process(nil, &req)
	if err != nil {
		return
	}

	// write to files, print the result, ...

	return
}
```

## Go flag way
```go
type InitReq struct {
	name string
}

//libgo:impl libgo/protocol.CommandLineArguments
func (req *InitReq) FromCLA(arguments []string) (remaining []string, err protocol.Error) {
	var flagSet flag.FlagSet
	// flagSet.Init("module.InitReq", flag.ContinueOnError)

	flagSet.StringVar(&req.name, "n", "", "module domain name e.g. user-name")

	var goErr = flagSet.Parse(arguments)
	if goErr != nil {
		// err =
	}
	remaining = flagSet.Args()
	return
}
func (req *InitReq) ToCLA() (arguments []string, err protocol.Error) { return }
```
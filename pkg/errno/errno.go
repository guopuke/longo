package errno

type Errno struct {
	Code    int
	Message string
}

type Err struct {
	Code    int
	Message string
	Err     error
}

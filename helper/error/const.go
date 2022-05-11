package errorConv

const (
	ErrDBNotFound = "record not found"
	ErrInvalid    = "data invalid"
)

type CarModel int64

const (
	Sedan = iota + 1
	SUV
	MPV
	Truck
	MiniBus
)

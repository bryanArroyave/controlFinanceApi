package valueobjects

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type ID struct {
	*valueobjects.IntValueObject
}

func NewID(value int) *ID {
	return &ID{
		valueobjects.NewIntValueObject(value).Min(0),
	}
}

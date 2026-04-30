package storevalidator

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Validator struct {
}

func New() Validator {
	return Validator{}
}

package productvalidator

const (
	skuRegex = `^[A-Z0-9\-]+$`
	urlRegex = `^https?://.*\.(jpg|jpeg|png|gif|webp)$`
)

type Validator struct {
}

func New() Validator {
	return Validator{}
}

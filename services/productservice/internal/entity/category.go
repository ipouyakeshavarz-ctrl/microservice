package entity

type Category string

const (
	Sport       Category = "sport"
	Electronics Category = "electronics"
	Fashion     Category = "fashion"
	Home        Category = "home"
	Beauty      Category = "beauty"
)

func (c Category) IsValid() bool {
	switch c {
	case Sport, Electronics, Fashion, Home, Beauty:
		return true
	default:
		return false
	}
}

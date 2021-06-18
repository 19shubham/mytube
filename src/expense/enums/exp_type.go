package enums

type ExpType int

const (
	Credit ExpType = 1+ iota
	Offer1
	Offer2
)

func (e ExpType) String() string {
	return []string{"Credit","Offer1","Offer2"}[e-1]
}

func (e ExpType) Index() int{
	return int(e)
}
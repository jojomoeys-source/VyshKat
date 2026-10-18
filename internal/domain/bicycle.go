package domain

type Bicycle struct {
	*Transport
}

func NewBicycle(
	name string,
	inventoryNumber int,
	beginnerFriendliness int,
	operational bool,
) (*Bicycle, error) {
	transport, err := NewTransport(
		name,
		inventoryNumber,
		beginnerFriendliness,
		operational,
	)
	if err != nil {
		return nil, err
	}

	return &Bicycle{
		Transport: transport,
	}, nil
}

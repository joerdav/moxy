package types

type MoveDirection int

const (
	DirectionUp MoveDirection = iota
	DirectionDown
)

func (dir MoveDirection) String() string {
	if dir == DirectionUp {
		return "up"
	}

	return "down"
}

func NewMoveDirection(s string) MoveDirection {
	if s == DirectionDown.String() {
		return DirectionDown
	}
	return DirectionUp
}

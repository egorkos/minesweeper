package model

type GameStatus int

const (
	Win GameStatus = iota
	Loose
	Undefined
)

func (s GameStatus) String() string {
	return [...]string{"WIN", "LOOSE", "UNDEFINED"}[s]
}
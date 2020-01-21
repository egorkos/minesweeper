package model

type GameStatus int

const (
	Win GameStatus = iota
	Loose
	Running
)

func (s GameStatus) String() string {
	return [...]string{"WIN", "LOOSE", "RUNNING"}[s]
}

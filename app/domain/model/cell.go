package model

type Cell struct {
	Mine        bool `json:"mine"`
	Revealed    bool `json:"revealed"`
	Flagged     bool `json:"flagged"`
	MinesAround int  `json:"mines_around"`
}

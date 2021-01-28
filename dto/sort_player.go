package dto

// PlayersSortByID sort a list of player by their ID
type PlayersSortByID []Player

func (a PlayersSortByID) Len() int           { return len(a) }
func (a PlayersSortByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PlayersSortByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

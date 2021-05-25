package dto

import (
	"sort"
	"time"
)

// SanctionEntry is an entry of any type of potential sanction.
// Examples of sanctions may be: mutes, votebans, bans, troll pit (zCatch)
type Sanction struct {
	Type      string    `json:"type,omitempty"`
	Player    Player    `json:"player,omitempty"`
	Reason    string    `json:"reason,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

// ExpiresIn returns the time to wait until the sanction expires
func (s *Sanction) ExpiresIn() time.Duration {
	return time.Until(s.ExpiresAt)
}

// ByExpirationTime allows to sort a list of sanctions by their expiration time
// in order to have a correct representation
type ByExpirationTime []Sanction

func (a ByExpirationTime) Len() int           { return len(a) }
func (a ByExpirationTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByExpirationTime) Less(i, j int) bool { return a[i].ExpiresIn() < a[j].ExpiresIn() }

// A SanctionMap maps a specific sanction type to a sorted list list of sanctions
type SanctionMap map[string][]Sanction

func (sm *SanctionMap) Add(s Sanction) {
	// always sort list before exiting
	defer sort.Sort(ByExpirationTime((*sm)[s.Type]))

	sanctionList := (*sm)[s.Type][:]

	// override old sanction if players are equal
	for idx, sanction := range sanctionList {
		if sanction.Player.IP == s.Player.IP {
			sanctionList[idx] = s
			// sorted before returning
			return
		}
	}

	// add new sanction to list
	(*sm)[s.Type] = append((*sm)[s.Type], s)
	// sorted before returning
}

func (sm *SanctionMap) Cleanup() {
	// iterate over map
	for key, list := range *sm {
		// because items are sorted by expiration time,
		// the items with the smallest expiration time are at the front
		for len(list) > 0 && list[0].ExpiresIn() <= 0 {
			list = list[1:]
		}
		(*sm)[key] = list
	}
}

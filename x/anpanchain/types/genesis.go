package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PeopleList: []People{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in people
	peopleIdMap := make(map[uint64]bool)
	peopleCount := gs.GetPeopleCount()
	for _, elem := range gs.PeopleList {
		if _, ok := peopleIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for people")
		}
		if elem.Id >= peopleCount {
			return fmt.Errorf("people id should be lower or equal than the last id")
		}
		peopleIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

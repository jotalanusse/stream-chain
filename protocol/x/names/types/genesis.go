package types

var (
	NameJota Name = Name{
		Id:   0,
		Name: "Jota",
	}
	NameSolal Name = Name{
		Id:   1,
		Name: "Solal",
	}
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Names: []Name{
			NameJota,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Genesis state should contain at least one name.
	if len(gs.Names) == 0 {
		return ErrNoNameInGenesis
	}

	// The first name should always be Jota.
	if gs.Names[0] != NameJota {
		return ErrJotaMustBeNameZero
	}

	// Provided names should not contain duplicated name ids, and names.
	// Name ids should be sequential.
	nameIdSet := make(map[uint32]struct{})
	nameSet := make(map[string]struct{})
	expectedId := uint32(0)

	for _, name := range gs.Names {
		if _, exists := nameIdSet[name.Id]; exists {
			return ErrNameIdAlreadyExists
		}
		if _, exists := nameSet[name.Name]; exists {
			return ErrNameNameAlreadyExists
		}
		if name.Id != expectedId {
			return ErrGapFoundInNameId
		}

		nameIdSet[name.Id] = struct{}{}
		nameSet[name.Name] = struct{}{}
		expectedId = expectedId + 1
	}
	return nil
}

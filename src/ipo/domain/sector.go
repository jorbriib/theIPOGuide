package domain

// Sector represents the sector entity
type Sector struct {
	name    string
}

// HydrateSector hydrates the sector struct
func HydrateSector(name string) Sector {
	return Sector{
		name:    name,
	}
}

// Name returns the sector name as string
func (s Sector) Name() string {
	return s.name
}

package domain

// SectorId represents the Sector Id
type SectorId string

// Sector represents the sector entity
type Sector struct {
	id   SectorId
	name string
}

// HydrateSector hydrates the sector struct
func HydrateSector(id SectorId, name string) Sector {
	return Sector{
		id:   id,
		name: name,
	}
}

// Id returns the sector id
func (s Sector) Id() SectorId {
	return s.id
}


// Name returns the sector name as string
func (s Sector) Name() string {
	return s.name
}

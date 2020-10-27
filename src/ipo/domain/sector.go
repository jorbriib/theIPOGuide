package domain

// SectorId represents the Sector Id
type SectorId string

// Sector represents the sector entity
type Sector struct {
	id    SectorId
	alias string
	name  string
}

// HydrateSector hydrates the sector struct
func HydrateSector(id SectorId, alias string, name string) Sector {
	return Sector{
		id:    id,
		alias: alias,
		name:  name,
	}
}

// Id returns the sector id
func (s Sector) Id() SectorId {
	return s.id
}

// Alias returns the sector alias
func (s Sector) Alias() string {
	return s.alias
}

// Name returns the sector name as string
func (s Sector) Name() string {
	return s.name
}

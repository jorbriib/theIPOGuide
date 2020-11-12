package domain

// SectorId represents the Sector Id
type SectorId string

// Sector represents the sector entity
type Sector struct {
	id    SectorId
	alias string
	name  string
	image     string
	totalIpos int
}

// HydrateSector hydrates the sector struct
func HydrateSector(id SectorId, alias string, name string, image string, totalIpos int) Sector {
	return Sector{
		id:    id,
		alias: alias,
		name:  name,
		image: image,
		totalIpos: totalIpos,
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

// Image returns the sector image as string
func (s Sector) Image() string {
	return s.image
}

// TotalIpos returns the number of ipos related to a sector
func (s Sector) TotalIpos() int {
	return s.totalIpos
}

package domain

// IndustryId represents the Industry Id
type IndustryId string

// Industry represents the sector entity
type Industry struct {
	id   IndustryId
	alias string
	name string
}

// HydrateIndustry hydrates the sector struct
func HydrateIndustry(id IndustryId, alias string, name string) Industry {
	return Industry{
		id: id,
		alias: alias,
		name: name,
	}
}

// Id returns the sector id
func (s Industry) Id() IndustryId {
	return s.id
}

// Alias returns the sector alias
func (s Industry) Alias() string {
	return s.alias
}

// Name returns the sector name as string
func (s Industry) Name() string {
	return s.name
}

package domain

// IndustryId represents the Industry Id
type IndustryId string

// Industry represents the sector entity
type Industry struct {
	id   IndustryId
	name string
}

// HydrateIndustry hydrates the sector struct
func HydrateIndustry(id IndustryId, name string) Industry {
	return Industry{
		id: id,
		name: name,
	}
}

// Id returns the sector id
func (s Industry) Id() IndustryId {
	return s.id
}

// Name returns the sector name as string
func (s Industry) Name() string {
	return s.name
}

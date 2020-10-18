package domain

// Industry represents the sector entity
type Industry struct {
	name    string
}

// HydrateIndustry hydrates the sector struct
func HydrateIndustry(name string) Industry {
	return Industry{
		name:    name,
	}
}

// Name returns the sector name as string
func (s Industry) Name() string {
	return s.name
}

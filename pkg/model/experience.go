package model

var _ Experiencer = (*ExperienceEntity)(nil)

type Experiencer interface {
	GetExperience() int64
}

func NewExperiencer(experience int64) Experiencer {
	return &ExperienceEntity{
		Experience: experience,
	}
}

type ExperienceEntity struct {
	Experience int64
}

// GetExperience implements Experiencer.
func (e *ExperienceEntity) GetExperience() int64 {
	return e.Experience
}

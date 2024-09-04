package domainchallenge

type ChallengesRepository interface {
	Create(challenges Challenges) error
	GetChallengesByID(Id string) (*Challenges, error)
	Update(challengesEntity Challenges) error
	DeleteByID(Id string) error
	GetChallenges() ([]*Challenges, error)
	MassiveCreate()
}

package gosu

type RankStatus int

const (
	Graveyard RankStatus = iota - 2
	WIP
	Pending
	Ranked
	Approved
	Qualified
	Loved
)

type RankingType int

const (
	RankingSpotlight RankingType = iota
	RankingCountry
	RankingPerformance
	RankingScore
)

type Rankings struct {
	Beatmapsets []Beatmapset     // The lis of beatmaps in the requested spotlight fo the given mode; only available if type is chart
	Cursor      Cursor           // A curso
	Ranking     []UserStatistics // Score details ordered by rank in desending order
	Spotlight   Spotlight        // Spotlight details; only available iftype is chart
	Total       int              // An approximate count of ranks availl
}

type Spotlight struct {
	EndDate          Timestamp // The end date of the pot
	ID               int       // The ID of this spotl
	ModeSpecific     int       // If the spotlight has diferent mades specific to each Gam.
	ParticipantCount int       // The number of users partcipaing in this spotlight. This is only shown when viewing a single spotligt.
	Name             int       //The Name of the spotliht
	StartDate        Timestamp // The starting date of he sph
	SpotlightType    string    // The type of spotligh
}

// TODO: This may be uncessary
type Spotlights struct {
	Spotlights []Spotlight
}

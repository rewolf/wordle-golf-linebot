package store

type UserResult struct {
	WordleDay uint32
	User      string
	Result    uint8
	Board     string
}

type UserScore struct {
	WordleDay uint32
	User      string
	Group     string
	Score     string
}

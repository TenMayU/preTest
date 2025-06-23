package quote

type Quote struct {
	ID    int
	Text  string
	Voted int
}

type QuoteVoting struct {
	ID      int
	QuoteId int
	User    int
}

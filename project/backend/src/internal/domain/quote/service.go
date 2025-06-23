package quote

import "context"

type QuoteService interface {
	GetAllVoted(ctx context.Context) ([]QuoteVoting, error)
	GetAllQuote(ctx context.Context) ([]Quote, error)
	GetQuoteBySearch(ctx context.Context, Text string) ([]Quote, error)
	Create(ctx context.Context, Text string) (bool, error)
	Voting(ctx context.Context, ID int, User int) (bool, string, error)
	Update(ctx context.Context, ID int, Text string) (bool, error)
}

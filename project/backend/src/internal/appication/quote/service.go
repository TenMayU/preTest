package quoteService

import (
	"backendrest/src/internal/domain/quote"
	"context"
	"fmt"
)

type quoteServiceImpl struct {
	repo quote.QuoteRepositoty
}

func NewAuthService(repo quote.QuoteRepositoty) quote.QuoteService {
	return &quoteServiceImpl{repo: repo}
}

func (r *quoteServiceImpl) GetAllVoted(ctx context.Context) ([]quote.QuoteVoting, error) {

	quoteVoted, err := r.repo.GetAllVoted(ctx)
	if err != nil {
		return nil, err // 🟢 แค่ return error
	}

	// 	for i, v := range someSlice {
	//     // i = index (0,1,2,...)
	//     // v = ค่าของแต่ละ element
	// }

	return quoteVoted, nil
}

func (r *quoteServiceImpl) GetAllQuote(ctx context.Context) ([]quote.Quote, error) {

	quote, err := r.repo.GetAllQuote(ctx)
	if err != nil {
		return nil, err // 🟢 แค่ return error
	}

	// 	for i, v := range someSlice {
	//     // i = index (0,1,2,...)
	//     // v = ค่าของแต่ละ element
	// }

	return quote, nil
}

func (r *quoteServiceImpl) GetQuoteBySearch(ctx context.Context, Text string) ([]quote.Quote, error) {
	if Text == "" {
		return nil, fmt.Errorf("please input text")
	}
	quote, err := r.repo.GetQuoteBySearch(ctx, Text)
	if err != nil {
		return nil, err // 🟢 แค่ return error
	}

	// 	for i, v := range someSlice {
	//     // i = index (0,1,2,...)
	//     // v = ค่าของแต่ละ element
	// }

	return quote, nil
}

func (r *quoteServiceImpl) Create(ctx context.Context, Text string) (bool, error) {

	if Text == "" {
		return false, fmt.Errorf("please input text")
	}

	// 📦 เตรียม entity
	u := quote.Quote{
		Text:  Text,
		Voted: 0,
	}

	// 💾 เรียก Repository ไป save
	ok, err := r.repo.Create(ctx, u.Text)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, err
	}
	return true, err

}

func (r *quoteServiceImpl) Voting(ctx context.Context, ID int, User int) (bool, string, error) {

	if ID == 0 || User == 0 {
		return false, "", fmt.Errorf("please input ID or User ")
	}

	// 📦 เตรียม entity
	u := quote.QuoteVoting{
		ID:   ID,
		User: User,
	}

	// 💾 เรียก Repository ไป save
	ok, mgs, err := r.repo.Voting(ctx, u.ID, u.User)
	if err != nil {
		return false, mgs, err
	}

	if !ok {
		return false, mgs, err
	}
	return true, mgs, err

}

func (r *quoteServiceImpl) Update(ctx context.Context, ID int, Text string) (bool, error) {

	if Text == "" {
		return false, fmt.Errorf("please input text")
	}

	if ID == 0 {
		return false, fmt.Errorf("please input ID")
	}

	ok, err := r.repo.Update(ctx, ID, Text)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, err
	}
	return true, err

}

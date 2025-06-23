package repository

import (
	"backendrest/src/internal/domain/quote"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type quoteRepositotyImp struct {
	db *gorm.DB
}

func NewQuoteService(db *gorm.DB) quote.QuoteRepositoty {
	return &quoteRepositotyImp{db: db}
}

func (s *quoteRepositotyImp) GetAllVoted(ctx context.Context) ([]quote.QuoteVoting, error) {
	var quotes []quote.QuoteVoting

	err := s.db.Find(&quotes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get voted: %v", err)
	}

	return quotes, nil
}
func (s *quoteRepositotyImp) GetAllQuote(ctx context.Context) ([]quote.Quote, error) {
	var quotes []quote.Quote

	err := s.db.Find(&quotes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get quotes: %v", err)
	}

	return quotes, nil
}

func (s *quoteRepositotyImp) GetQuoteBySearch(ctx context.Context, Text string) ([]quote.Quote, error) {
	var quote []quote.Quote
	keyword := Text

	err := s.db.Where("text LIKE ?", "%"+keyword+"%").Find(&quote).Error
	if err != nil {
		return nil, fmt.Errorf("not found")
	}

	if len(quote) == 0 {
		return nil, fmt.Errorf("no quote matched")
	}
	return quote, nil
}

func (s *quoteRepositotyImp) Create(ctx context.Context, Text string) (bool, error) {
	quote := quote.Quote{
		Text:  Text,
		Voted: 0,
	}
	if err := s.db.WithContext(ctx).Create(&quote).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *quoteRepositotyImp) Voting(ctx context.Context, ID int, User int) (bool, string, error) {
	var q quote.QuoteVoting
	var quoteData quote.Quote

	err := s.db.Where("User = ?", User).First(&q).Error
	if err == nil {
		return false, "user has voted this quote", err
	}
	errQuoteData := s.db.Where("ID = ?", ID).First(&quoteData).Error

	if errQuoteData != nil {
		return false, "errQuoteData has not define", fmt.Errorf("errQuoteData has not define %v", errQuoteData)
	}
	var count int64
	errCount := s.db.Model(&quote.QuoteVoting{}).Where("quote_id = ?", ID).Count(&count).Error
	if errCount != nil {
		return false, "quote has not define:", fmt.Errorf("quote has not define: %v", errCount)
	}
	vote := quote.QuoteVoting{
		QuoteId: ID,
		User:    User,
	}

	quoteUpdate := quote.Quote{
		Text:  quoteData.Text,
		Voted: int(count + 1),
	}

	if err := s.db.WithContext(ctx).Create(&vote).Error; err != nil {
		return false, "", err
	}

	errUpdate := s.db.WithContext(ctx).
		Model(&quote.Quote{}).
		Where("ID = ?", ID).
		Updates(quoteUpdate).Error

	if errUpdate != nil {
		return false, "", fmt.Errorf("failed to update quote: %v", errUpdate)
	}

	return true, "", nil
}

func (s *quoteRepositotyImp) Update(ctx context.Context, ID int, Text string) (bool, error) {
	var quotes quote.Quote
	if ID == 0 {
		return false, fmt.Errorf("please input Quote ID")
	}
	err := s.db.Where("ID == ?", ID).First(&quotes).Error
	if err != nil {
		return false, fmt.Errorf("Not define quote")
	}

	quoteUpdate := quote.Quote{
		Text: Text,
	}

	errUpdate := s.db.WithContext(ctx).
		Model(&quote.Quote{}).
		Where("ID = ?", ID).
		Updates(quoteUpdate).Error
	if errUpdate != nil {
		return false, fmt.Errorf("Update fail")
	}
	return true, nil
}

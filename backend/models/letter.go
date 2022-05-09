package models

import "time"

type Letter struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	From uint
	User User `gorm:"foreignKey:From"`

	// Genre is classification of content
	Genre LetterGenre `gorm:"not null"`

	Title string `gorm:"not null;size:128"`
	Body  string `gorm:"not null"`
}

// LetterGenre means Letter.Genre
type LetterGenre uint

const (
	// about this app's performance
	AboutAppPerformance LetterGenre = iota
	// about posted contents
	AboutPostedContent
	// otherwise
	Otherwise LetterGenre = 99
)

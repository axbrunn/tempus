package models

import "time"

type EntryType string

const (
	Accrual    EntryType = "opbouw"
	Withdrawal EntryType = "opnemen"
)

type Entry struct {
	Date        time.Time `json:"date"`
	Hours       float64   `json:"hours"`
	Description string    `json:"description"`
	Type        EntryType `json:"type"`
}

func NewEntry(hours float64, description string, entryType EntryType) Entry {
	return Entry{
		Date:        time.Now(),
		Hours:       hours,
		Description: description,
		Type:        entryType,
	}
}

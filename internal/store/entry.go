package store

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/axbrunn/tempus/internal/models"
)

type Store struct {
	Path    string         `json:"-"`
	Entries []models.Entry `json:"entries"`
}

func DataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".comptime_data.json"), nil
}

func Load(path string) (Store, error) {
	var s Store
	s.Path = path
	data, err := os.ReadFile(path)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)

	return s, err
}

func (s *Store) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.Path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, s.Path)
}

func (s *Store) CalculateBalance() float64 {
	var balance float64

	for _, entry := range s.Entries {
		switch entry.Type {
		case models.Accrual:
			balance += entry.Hours
		case models.Withdrawal:
			balance -= entry.Hours
		}
	}

	return balance
}

func (s *Store) ExportCSV() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dest := filepath.Join(home, "Downloads", "tempus-export.csv")

	file, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"date", "hours", "description", "type"})

	var totalAccrual, totalWithdrawal float64

	for _, entry := range s.Entries {
		writer.Write([]string{
			entry.Date.Format("2006-01-02"),
			fmt.Sprintf("%.2f", entry.Hours),
			entry.Description,
			string(entry.Type),
		})

		switch entry.Type {
		case models.Accrual:
			totalAccrual += entry.Hours
		case models.Withdrawal:
			totalWithdrawal += entry.Hours
		}
	}

	// Empty row
	writer.Write([]string{})

	writer.Write([]string{"Totaal opgebouwd", fmt.Sprintf("%.2f", totalAccrual), "", ""})
	writer.Write([]string{"Totaal opgenomen", fmt.Sprintf("%.2f", totalWithdrawal), "", ""})
	writer.Write([]string{"Saldo", fmt.Sprintf("%.2f", totalAccrual-totalWithdrawal), "", ""})

	return dest, writer.Error()
}

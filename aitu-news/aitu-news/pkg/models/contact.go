// contact.go
package models

import "aitu-news/aitu-news/aitu-news/pkg/drivers"

// Contact представляет собой модель для контакта
type Contact struct {
	ID      int
	Name    string
	Email   string
	Message string
}

// GetContacts возвращает список контактов из базы данных
func GetContacts() ([]Contact, error) {
	rows, err := drivers.DB.Query("SELECT id, name, email, message FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Message); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

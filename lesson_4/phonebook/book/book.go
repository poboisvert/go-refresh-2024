package book

import (
	"fmt"
	"time"
)

type PhoneBook map[string]PhoneNumber

type PhoneNumber struct {
	Number        string
	LastUpdatedAt int64
}

func (book *PhoneBook) Add(name, phoneNumber string) error {
	if _, exist := (*book)[name]; exist {
		return fmt.Errorf("name %s already exists", name)
	}

	(*book)[name] = PhoneNumber{
		Number:        phoneNumber,
		LastUpdatedAt: time.Now().Unix(),
	}

	return nil
}

func (book *PhoneBook) Get(name string) (PhoneNumber, error) {
	if numberData, exist := (*book)[name]; exist {
		return numberData, nil
	}
	return PhoneNumber{}, fmt.Errorf("no entry found for %s", name)
}

func (book *PhoneBook) Update(name, newPhoneNumber string) error {
	if _, exist := (*book)[name]; !exist {
		return fmt.Errorf("no entry found for %s", name)
	}

	(*book)[name] = PhoneNumber{
		Number:        newPhoneNumber,
		LastUpdatedAt: time.Now().Unix(),
	}
	return nil

}

func (book *PhoneBook) Delete(name string) error {
	if _, exist := (*book)[name]; !exist {
		return fmt.Errorf("no entry found for %s", name)
	}

	delete(*book, name)

	return nil
}

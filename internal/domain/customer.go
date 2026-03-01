package domain

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zipCode"`
	CreatedAt time.Time `json:"createdAt"`
}

func GenerateFake(r *rand.Rand) Customer {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return createSingle(r)
}

func GenerateFakeList(count int, r *rand.Rand) []Customer {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	if count <= 0 {
		return []Customer{}
	}

	list := make([]Customer, 0, count)
	for i := 0; i < count; i++ {
		list = append(list, createSingle(r))
	}

	return list
}

func createSingle(r *rand.Rand) Customer {
	firstNames := []string{
		"Ana", "Bruno", "Carla", "Daniel", "Eduarda",
		"Felipe", "Gabriela", "Henrique", "Isabela", "João",
	}

	lastNames := []string{
		"Silva", "Souza", "Pereira", "Oliveira", "Costa",
		"Santos", "Rodrigues", "Almeida", "Nunes", "Ferreira",
	}

	domains := []string{
		"exemplo.com", "teste.com", "mail.com", "empresa.com", "dominio.com",
	}

	streetNames := []string{
		"Rua das Flores", "Avenida Central", "Travessa do Sol",
		"Alameda Bela Vista", "Rua do Comércio",
	}

	cities := []string{
		"São Paulo", "Rio de Janeiro", "Belo Horizonte",
		"Porto Alegre", "Curitiba",
	}

	states := []string{
		"SP", "RJ", "MG", "RS", "PR",
	}

	first := firstNames[r.Intn(len(firstNames))]
	last := lastNames[r.Intn(len(lastNames))]
	domain := domains[r.Intn(len(domains))]

	email := fmt.Sprintf(
		"%s.%s%d@%s",
		sanitize(first),
		sanitize(last),
		r.Intn(999)+1,
		domain,
	)

	phone := fmt.Sprintf(
		"(%02d) %03d-%04d",
		r.Intn(88)+11,
		r.Intn(900)+100,
		r.Intn(9000)+1000,
	)

	street := fmt.Sprintf(
		"%d %s",
		r.Intn(9999)+1,
		streetNames[r.Intn(len(streetNames))],
	)

	city := cities[r.Intn(len(cities))]
	state := states[r.Intn(len(states))]
	zip := fmt.Sprintf("%05d", r.Intn(90000)+10000)

	return Customer{
		ID:        uuid.New(),
		FirstName: first,
		LastName:  last,
		Email:     strings.ToLower(email),
		Phone:     phone,
		Street:    street,
		City:      city,
		State:     state,
		ZipCode:   zip,
		CreatedAt: time.Now().UTC(),
	}
}

func sanitize(input string) string {
	normalized := strings.ToLower(strings.TrimSpace(input))

	replacements := map[string]string{
		"á": "a", "à": "a", "ã": "a", "â": "a",
		"é": "e", "ê": "e",
		"í": "i",
		"ó": "o", "ô": "o", "õ": "o",
		"ú": "u", "ü": "u",
		"ç": "c",
		" ": "",
	}

	for old, new := range replacements {
		normalized = strings.ReplaceAll(normalized, old, new)
	}

	return normalized
}

func (c Customer) String() string {
	return fmt.Sprintf(
		"%s %s <%s> (%s) - %s, %s/%s %s",
		c.FirstName,
		c.LastName,
		c.Email,
		c.Phone,
		c.Street,
		c.City,
		c.State,
		c.ZipCode,
	)
}
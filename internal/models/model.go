package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey;column:id"`
	FirstName  string    `gorm:"column:first_name;not null"`
	LastName   string    `gorm:"column:last_name;not null"`
	Email      string    `gorm:"column:email;unique;not null"`
	Password   string    `gorm:"column:password;not null"`
	CustomerID *uint     `gorm:"column:customer_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type Film struct {
	FilmID          int     `gorm:"column:film_id;primaryKey"`
	Title           string  `gorm:"column:title"`
	Description     string  `gorm:"column:description"`
	ReleaseYear     int     `gorm:"column:release_year"`
	LanguageID      int     `gorm:"column:language_id"`
	RentalDuration  int     `gorm:"column:rental_duration"`
	RentalRate      float32 `gorm:"column:rental_rate"`
	Length          int     `gorm:"column:length"`
	ReplacementCost float32 `gorm:"column:replacement_cost"`
	Rating          string  `gorm:"column:rating"`
}

type Category struct {
	CategoryID int    `gorm:"column:category_id;primaryKey"`
	Name       string `gorm:"column:name"`
}

type FilmCategory struct {
	FilmID     int `gorm:"column:film_id"`
	CategoryID int `gorm:"column:category_id"`
}

type Actor struct {
	ActorID   int    `gorm:"column:actor_id;primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

type FilmActor struct {
	ActorID int `gorm:"column:actor_id"`
	FilmID  int `gorm:"column:film_id"`
}

type Customer struct {
	CustomerID int    `gorm:"column:customer_id;primaryKey"`
	StoreID    int    `gorm:"column:store_id"`
	FirstName  string `gorm:"column:first_name"`
	LastName   string `gorm:"column:last_name"`
	Email      string `gorm:"column:email"`
	AddressID  int    `gorm:"column:address_id"`
	Active     bool   `gorm:"column:active"`
}

type Rental struct {
	RentalID    int       `gorm:"column:rental_id;primaryKey"`
	RentalDate  time.Time `gorm:"column:rental_date"`
	InventoryID int       `gorm:"column:inventory_id"`
	CustomerID  int       `gorm:"column:customer_id"`
	ReturnDate  time.Time `gorm:"column:return_date"`
	StaffID     int       `gorm:"column:staff_id"`
}

type Payment struct {
	PaymentID   int       `gorm:"column:payment_id;primaryKey"`
	CustomerID  int       `gorm:"column:customer_id"`
	StaffID     int       `gorm:"column:staff_id"`
	RentalID    int       `gorm:"column:rental_id"`
	Amount      float32   `gorm:"column:amount"`
	PaymentDate time.Time `gorm:"column:payment_date"`
}

type Language struct {
	LanguageID int    `gorm:"column:language_id;primaryKey"`
	Name       string `gorm:"column:name"`
}

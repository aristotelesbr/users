package models

// Usuarios database table
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `sb:"email" json:"email"`
}

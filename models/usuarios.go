package models

import "usuarios/lib"

// Usuarios database table
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `sb:"email" json:"email"`
}

// UsuarioModel receiver connection config to connect database
var UsuarioModel = lib.Sess.Collection("usuarios")

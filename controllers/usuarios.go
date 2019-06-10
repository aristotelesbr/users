package controllers

import (
	"net/http"

	"usuarios/models"

	"github.com/labstack/echo"
)

// Home GET /api/home
func Home(c echo.Context) error {

	data := map[string]interface{}{
		"titulo": "Listagem de usuários",
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// Create POST /v1/create
func Create(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Error ao tentar criar um novo cadastro",
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"message": "Usuário criado com sucesso!",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Todos os campos devem ser informados.",
	})
}

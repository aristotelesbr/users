package controllers

import (
	"net/http"
	"strconv"

	"usuarios/models"

	"github.com/labstack/echo"
)

// Home GET /api/home
func Home(c echo.Context) error {

	var users []models.Usuarios

	if err := models.UsuarioModel.Find().All(&users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error ao recuperar os dados",
		})
	}

	data := map[string]interface{}{
		"titulo": "Listagem de usuários",
		"users":  users,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// New GET /api/new
func New(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
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
		} else {
			return c.Redirect(http.StatusFound, "/")
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Todos os campos devem ser informados.",
	})
}

// Delete DELETE /users/:id
func Delete(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// SLECT * FROM USER usuarios WHERE id= ?
	result := models.UsuarioModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Registro não encontrado",
		})
	}

	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao executar a operação!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "Registro excluido com sucesso!",
	})
}

// Edit EDIT /users/id
func Edit(c echo.Context) error {
	var userID, _ = strconv.Atoi(c.Param("id"))

	var user models.Usuarios

	result := models.UsuarioModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não encontrado",
		})
	}
	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Usuário não encontrado",
		})
	}
	var data = map[string]interface{}{
		"user": user,
	}
	return c.Render(http.StatusOK, "edit.html", data)
}

// Update UPDATE /users/:id
func Update(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user = models.Usuarios{
		ID:    userID,
		Nome:  name,
		Email: email,
	}
	result := models.UsuarioModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não encontrado!",
		})
	}

	if err := result.Update(user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Houve um erro, por favor tente novamente mais tarde",
		})
	}

	return c.JSON(http.StatusFound, "/")
}

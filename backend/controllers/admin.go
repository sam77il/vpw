package controllers

import (
	"ecommerce/database/models"
	"ecommerce/sugar"
	"encoding/json"
	"fmt"
)

func GetUsers(ctx *sugar.SugarContext) {
	rows, err := ctx.Database.Query(ctx.Request.GoCtx, "SELECT id, email, role, gender, first_name, last_name, street, postal_code, city, country, phone_number, company, company_name, company_ustidnr FROM users")
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Email, &user.Role, &user.Gender, &user.FirstName, &user.LastName, &user.Street, &user.PostalCode, &user.City, &user.Country, &user.PhoneNumber, &user.Company, &user.CompanyName, &user.CompanyUstIdNr)
		if err != nil {
			ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
			return
		}
		users = append(users, user)
	}

	ctx.Response.JSON(map[string]any{"success": true, "users": users})
}

func GetUserById(ctx *sugar.SugarContext) {
	id := ctx.Request.Params["id"]
	row := ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT id, email, role, gender, first_name, last_name, street, postal_code, city, country, phone_number, company, company_name, company_ustidnr FROM users WHERE id = $1", id)
	var user models.User
	err := row.Scan(&user.Id, &user.Email, &user.Role, &user.Gender, &user.FirstName, &user.LastName, &user.Street, &user.PostalCode, &user.City, &user.Country, &user.PhoneNumber, &user.Company, &user.CompanyName, &user.CompanyUstIdNr)
	if err != nil {
		ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "user not found"})
		return
	}
	ctx.Response.JSON(map[string]any{"success": true, "user": user})
}

func UpdateUserById(ctx *sugar.SugarContext) {
	fmt.Println("updating1")
	id := ctx.Request.Params["id"]
	var user models.User
	if err := json.Unmarshal(ctx.Request.Body, &user); err != nil {
		fmt.Println(err)
		ctx.Response.Status(400).JSON(map[string]any{
			"success": false,
			"message": "invalid json",
		})
		return
	}

	cmd, err := ctx.Database.Exec(ctx.Request.GoCtx, "UPDATE users SET email = $1, role = $2, gender = $3, first_name = $4, last_name = $5, street = $6, postal_code = $7, city = $8, country = $9, phone_number = $10, company = $11, company_name = $12, company_ustidnr = $13 WHERE id = $14", &user.Email, &user.Role, &user.Gender, &user.FirstName, &user.LastName, &user.Street, &user.PostalCode, &user.City, &user.Country, &user.PhoneNumber, &user.Company, &user.CompanyName, &user.CompanyUstIdNr, &id)
	fmt.Println("updating2")
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}
	if cmd.RowsAffected() == 0 {
		ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "user not found"})
		return
	}
	ctx.Response.JSON(map[string]any{"success": true})
}

func DeleteUserById(ctx *sugar.SugarContext) {
	id := ctx.Request.Params["id"]
	fmt.Println(id)
	cmd, err := ctx.Database.Exec(ctx.Request.GoCtx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(400).JSON(map[string]any{
			"success": false,
			"message": "invalid json",
		})
		return
	}
	if cmd.RowsAffected() == 0 {
		ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "user not found"})
	}
	ctx.Response.JSON(map[string]any{"success": true})
}
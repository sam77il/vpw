package controllers

import (
	"ecommerce/database/models"
	"ecommerce/sugar"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID string `json:"user_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func AuthRegister(ctx *sugar.SugarContext) {
	var registeringUser models.User
	json.Unmarshal(ctx.Request.Body, &registeringUser)

	if registeringUser.Email == "" || registeringUser.Password == "" || registeringUser.FirstName == "" || registeringUser.LastName == "" || registeringUser.Street == "" || registeringUser.PostalCode == "" || registeringUser.City == "" || registeringUser.PhoneNumber == "" || (registeringUser.Company && registeringUser.CompanyName == "") || (registeringUser.Company && registeringUser.CompanyUstIdNr == "") {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "missing credentials"})
		return
	}
	var userId string
	var userRole string
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registeringUser.Password), bcrypt.DefaultCost)
	err := ctx.Database.QueryRow(ctx.Request.GoCtx, "INSERT INTO users (email, password, gender, first_name, last_name, street, postal_code, city, country, phone_number, company, company_name, company_ustidnr) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id, role", registeringUser.Email, string(hashedPassword), registeringUser.Gender, registeringUser.FirstName, registeringUser.LastName, registeringUser.Street, registeringUser.PostalCode, registeringUser.City, registeringUser.Country, registeringUser.PhoneNumber, registeringUser.Company, registeringUser.CompanyName, registeringUser.CompanyUstIdNr).Scan(&userId, &userRole)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "E-Mail wird bereits verwendet"}) // TODO: translate
			return
		}
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server/database error #0001"})
		log.Printf("%v", err)
		return
	}

	jwtToken, err := generateJsonWebToken(userId, userRole)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal jwt error #0002"})
		return
	}

	ctx.Response.Status(201).JSON(map[string]any{"success": true, "token": jwtToken})
}

func AuthLogin(ctx *sugar.SugarContext) {
	var incomingAccount Account
	json.Unmarshal(ctx.Request.Body, &incomingAccount)

	var user models.User
	ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT id, password, role FROM users WHERE email = $1", incomingAccount.Email).Scan(&user.Id, &user.Password, &user.Role)

	if user.Id == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(incomingAccount.Password))
	if err != nil {
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "wrong credentials"})
		return
	}

	jwtToken, err := generateJsonWebToken(user.Id, user.Role)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal jwt error #0002"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true, "token": jwtToken})
}

func AuthValidate(ctx *sugar.SugarContext) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "missing token"})
		return
	}
	claims, err := ValidateJsonWebToken(strings.TrimPrefix(authHeader, "Bearer "))
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}
	var userId string
	var userRole string
	err = ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT id, role FROM users WHERE id = $1", claims.UserID).Scan(&userId, &userRole)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}

	if userId == "" {
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}
	if userRole != claims.Role {
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true, "user_id": claims.UserID, "role": claims.Role})
}

func GetMe(ctx *sugar.SugarContext) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "missing user id"})
		return
	}
	claims, err := ValidateJsonWebToken(strings.TrimPrefix(token, "Bearer "))
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}
	userId := claims.UserID

	var user struct {
		Email string `json:"email"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Street string `json:"street"`
		PostalCode string `json:"postal_code"`
		City string `json:"city"`
		Country string `json:"country"`
		PhoneNumber string `json:"phone_number"`
		Company bool `json:"company"`
		CompanyName string `json:"company_name"`
		CompanyUstIdNr string `json:"company_ustidnr"`
	}
	err = ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT email, first_name, last_name, street, postal_code, city, country, phone_number, company, company_name, company_ustidnr FROM users WHERE id = $1", userId).Scan(&user.Email, &user.FirstName, &user.LastName, &user.Street, &user.PostalCode, &user.City, &user.Country, &user.PhoneNumber, &user.Company, &user.CompanyName, &user.CompanyUstIdNr)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error #0003"})
		return
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true, "user": user})
}

func EditMe(ctx *sugar.SugarContext) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "missing user id"})
		return
	}
	claims, err := ValidateJsonWebToken(strings.TrimPrefix(token, "Bearer "))
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "invalid token"})
		return
	}

	var changingData struct {
		Email string `json:"email"`
		CurrentPassword string `json:"current_password"`
		NewPassword string `json:"new_password"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Street string `json:"street"`
		PostalCode string `json:"postal_code"`
		City string `json:"city"`
		Country string `json:"country"`
		PhoneNumber string `json:"phone_number"`
		Company bool `json:"company"`
		CompanyName string `json:"company_name"`
		CompanyUstIdNr string `json:"company_ustidnr"`
	}
	json.Unmarshal(ctx.Request.Body, &changingData)
	// check what is changing and build query
	var setStatements []string
	var args []any
	argId := 1
	if changingData.Email != "" {
		setStatements = append(setStatements, fmt.Sprintf("email = $%d", argId))
		args = append(args, changingData.Email)
		argId++
	}
	if changingData.NewPassword != "" {
		var currentPassword string
		err := ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT password FROM users WHERE id = $1", claims.UserID).Scan(&currentPassword)
		if err != nil {
			fmt.Println(err)
			ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error #0004"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(changingData.CurrentPassword))
		if err != nil {
			fmt.Println(err)
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "wrong current password"})
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(changingData.NewPassword), bcrypt.DefaultCost)
		changingData.NewPassword = string(hashedPassword)
		setStatements = append(setStatements, fmt.Sprintf("password = $%d", argId))
		args = append(args, string(hashedPassword))
		argId++
	}
	if changingData.FirstName != "" {
		setStatements = append(setStatements, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, changingData.FirstName)
		argId++
	}
	if changingData.LastName != "" {
		setStatements = append(setStatements, fmt.Sprintf("last_name = $%d", argId))
		args = append(args, changingData.LastName)
		argId++
	}
	if changingData.Street != "" {
		setStatements = append(setStatements, fmt.Sprintf("street = $%d", argId))
		args = append(args, changingData.Street)
		argId++
	}
	if changingData.PostalCode != "" {
		setStatements = append(setStatements, fmt.Sprintf("postal_code = $%d", argId))
		args = append(args, changingData.PostalCode)
		argId++
	}
	if changingData.City != "" {
		setStatements = append(setStatements, fmt.Sprintf("city = $%d", argId))
		args = append(args, changingData.City)
		argId++
	}
	if changingData.Country != "" {
		setStatements = append(setStatements, fmt.Sprintf("country = $%d", argId))
		args = append(args, changingData.Country)
		argId++
	}
	if changingData.PhoneNumber != "" {
		setStatements = append(setStatements, fmt.Sprintf("phone_number = $%d", argId))
		args = append(args, changingData.PhoneNumber)
		argId++
	}
	if changingData.Company {
		setStatements = append(setStatements, fmt.Sprintf("company = $%d", argId))
		args = append(args, changingData.Company)
		argId++
		if changingData.CompanyName != "" {
			setStatements = append(setStatements, fmt.Sprintf("company_name = $%d", argId))
			args = append(args, changingData.CompanyName)
			argId++
		}
		if changingData.CompanyUstIdNr != "" {
			setStatements = append(setStatements, fmt.Sprintf("company_ustidnr = $%d", argId))
			args = append(args, changingData.CompanyUstIdNr)
			argId++
		}
	}
	if len(setStatements) == 0 {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "no data to change"})
		return
	}
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(setStatements, ", "), argId)
	args = append(args, claims.UserID)
	_, err = ctx.Database.Exec(ctx.Request.GoCtx, query, args...)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error #0004"})
		return
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func generateJsonWebToken(userId, userRole string) (string, error) {
	claims := Claims{
		UserID: userId,
		Role: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJsonWebToken(tokenStr string) (*Claims, error) {
	if tokenStr == "" {
    return nil, errors.New("missing token")
  }
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
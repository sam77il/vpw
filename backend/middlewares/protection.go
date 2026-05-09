package middlewares

import (
	"context"
	"ecommerce/controllers"
	"ecommerce/sugar"
	"fmt"
	"strings"
)

func RoutesProtection(ctx *sugar.SugarContext, next func()) {
	fmt.Println(ctx.Request.URL, ctx.Request.Method)
	jwtToken := strings.TrimPrefix(ctx.Request.Header.Get("Authorization"), "Bearer ")
	claims, err := controllers.ValidateJsonWebToken(jwtToken)
	
	if ctx.Request.URL == "/api/v1/categories" && ctx.Request.Method == "POST" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/categories/") && ctx.Request.Method == "DELETE" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/categories/") && ctx.Request.Method == "PATCH" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if ctx.Request.URL == "/api/v1/categories" && ctx.Request.Method == "POST" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if ctx.Request.URL == "/api/v1/products" && ctx.Request.Method == "POST" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/products/") && ctx.Request.Method == "DELETE" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/products/") && ctx.Request.Method == "PATCH" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if ctx.Request.URL == "/api/v1/products" && ctx.Request.Method == "POST" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
		return
	} else if ctx.Request.URL == "/api/v1/cart" && ctx.Request.Method == "GET" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.UserID != "" {
			ctx.Request.GoCtx = context.WithValue(ctx.Request.GoCtx, "userId", claims.UserID)
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "not logged in"})
	} else if ctx.Request.URL == "/api/v1/cart/clear" && ctx.Request.Method == "DELETE" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.UserID != "" {
			ctx.Request.GoCtx = context.WithValue(ctx.Request.GoCtx, "userId", claims.UserID)
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "not logged in"})
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/cart/") && ctx.Request.Method == "DELETE" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.UserID != "" {
			ctx.Request.GoCtx = context.WithValue(ctx.Request.GoCtx, "userId", claims.UserID)
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "not logged in"})
	} else if ctx.Request.URL == "/api/v1/cart" && ctx.Request.Method == "POST" {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.UserID != "" {
			ctx.Request.GoCtx = context.WithValue(ctx.Request.GoCtx, "userId", claims.UserID)
			next()
			return
		}
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "not logged in"})
	} else if strings.HasPrefix(ctx.Request.URL, "/api/v1/users") {
		if err != nil {
			ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
			return
		}
		if claims.Role == "admin" {
			next()
			return
		}
			
		ctx.Response.Status(401).JSON(map[string]any{"success": false, "message": "unauthorized"})
	}
	next()
}
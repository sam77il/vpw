package controllers

import (
	"ecommerce/database/models"
	"ecommerce/sugar"
	"encoding/json"
	"fmt"
)

func GetCartItems(ctx *sugar.SugarContext) {
	userId := ctx.Request.GoCtx.Value("userId").(string)

	query := `
		SELECT
			ci.id,
			ci.product_id,
			p.label,
			ci.amount,
			ci.metadata,
			ci.price
		FROM cart_items ci
		JOIN products p ON p.id = ci.product_id
		WHERE ci.user_id = $1
	`

	rows, err := ctx.Database.Query(ctx.Request.GoCtx, query, userId)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{
			"success": false,
			"message": "internal database error",
		})
		return
	}
	defer rows.Close()
	type CartItemWithLabel struct {
			models.CartItem
			Label string `json:"label"`
		}
	var cartItems []CartItemWithLabel

	for rows.Next() {
		var ci CartItemWithLabel
		var metadata []byte

		err := rows.Scan(
			&ci.Id,
			&ci.ProductId,
			&ci.Label,
			&ci.Amount,
			&metadata,
			&ci.Price,
		)

		if err != nil {
			fmt.Println(err)

			ctx.Response.Status(500).JSON(map[string]any{
				"success": false,
				"message": "internal scan error",
			})
			return
		}

		ci.UserId = userId

		if metadata == nil {
			ci.Metadata = json.RawMessage("{}")
		} else {
			ci.Metadata = json.RawMessage(metadata)
		}

		cartItems = append(cartItems, ci)
	}

	if err := rows.Err(); err != nil {
		ctx.Response.Status(500).JSON(map[string]any{
			"success": false,
			"message": "row iteration error",
		})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{
		"success":   true,
		"cart_items": cartItems,
	})
}

func ClearCartItems(ctx *sugar.SugarContext) {
	userId := ctx.Request.GoCtx.Value("userId").(string)

	_, err := ctx.Database.Exec(ctx.Request.GoCtx, "DELETE FROM cart_items WHERE user_id = $1", userId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}
	
	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func DeleteCartItem(ctx *sugar.SugarContext) {
	userId := ctx.Request.GoCtx.Value("userId").(string)
	cartId := ctx.Request.Params["id"]

	if cartId == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "no cart item given"})
		return
	}

	cmd, err := ctx.Database.Exec(ctx.Request.GoCtx, "DELETE FROM cart_items WHERE id = $1 AND user_id = $2", cartId, userId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	if cmd.RowsAffected() == 0 {
		ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "cart item not found"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func AddCartItem(ctx *sugar.SugarContext) {
	userId := ctx.Request.GoCtx.Value("userId").(string)
	var cartItem models.CartItem
	if err := json.Unmarshal(ctx.Request.Body, &cartItem); err != nil {
		ctx.Response.Status(400).JSON(map[string]any{
			"success": false,
			"message": "invalid json",
		})
		return
	}

	if len(cartItem.Metadata) == 0 {
		cartItem.Metadata = json.RawMessage("{}")
	}

	_, err := ctx.Database.Exec(ctx.Request.GoCtx, "INSERT INTO cart_items (user_id, product_id, amount, metadata, price) VALUES ($1, $2, $3, $4, $5)", userId, cartItem.ProductId, cartItem.Amount, cartItem.Metadata, cartItem.Price)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}
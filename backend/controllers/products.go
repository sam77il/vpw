package controllers

import (
	"ecommerce/database/models"
	"ecommerce/sugar"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Products(ctx *sugar.SugarContext) {
	rows, err := ctx.Database.Query(ctx.Request.GoCtx, "SELECT id, label, description, price, stock, category_id, options, old_price FROM products")
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "database error"})
		return
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		var options []byte
		err := rows.Scan(&p.Id, &p.Label, &p.Description, &p.Price, &p.Stock, &p.CategoryId, &options, &p.OldPrice)
		if err != nil {
			ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "scan error"})
			return
		}
		if options == nil {
			p.Options = json.RawMessage("{}")
		} else {
			p.Options = json.RawMessage(options)
		}
		products = append(products, p)
	}

	if len(products) > 0 {
		ctx.Response.Status(200).JSON(map[string]any{"success": true, "categories": products})
		return
	}
	ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "no products found"})
}

func ProductById(ctx *sugar.SugarContext) {
	proudctId := ctx.Request.Params["id"]

	if proudctId == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "no product given"})
		return
	}

	var product models.Product
	err := ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT label, description, price, stock, category_id, options, old_price FROM products WHERE id = $1", proudctId).Scan(&product.Label, &product.Description, &product.Price, &product.Stock, &product.CategoryId, &product.Options, &product.OldPrice)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "product not found"})
			return
		}
	}
	product.Id = proudctId
	ctx.Response.Status(200).JSON(map[string]any{"success": true, "product": product})
}

func AddProduct(ctx *sugar.SugarContext) {
	var product models.Product
	err := json.Unmarshal(ctx.Request.Body, &product)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}

	if len(product.Options) == 0 {
		product.Options = []byte("{}")
	}

	_, err = ctx.Database.Exec(ctx.Request.GoCtx, "INSERT INTO products (id, label, description, price, stock, category_id, options, old_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", product.Id, product.Label, product.Description, product.Price, product.Stock, product.CategoryId, product.Options, product.OldPrice)
	if err != nil {
		fmt.Println(err)
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func DeleteProduct(ctx *sugar.SugarContext) {
	productId := ctx.Request.Params["id"]

	_, err := ctx.Database.Exec(ctx.Request.GoCtx, "DELETE FROM products WHERE id = $1", productId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func EditProduct(ctx *sugar.SugarContext) {
	productId := ctx.Request.Params["id"]
	var newProduct models.Product
	err := json.Unmarshal(ctx.Request.Body, &newProduct)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}

	_, err = ctx.Database.Exec(ctx.Request.GoCtx, "UPDATE products SET id = $1, label = $2, description = $3, price = $4, stock = $5, category_id = $6, options = $8, old_price = $9 WHERE id = $7", newProduct.Id, newProduct.Label, newProduct.Description, newProduct.Price, newProduct.Stock, newProduct.CategoryId, productId, newProduct.Options, newProduct.OldPrice)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": false})
}
package controllers

import (
	"ecommerce/database/models"
	"ecommerce/sugar"
	"encoding/json"
)

func Categories(ctx *sugar.SugarContext) {
	rows, err := ctx.Database.Query(ctx.Request.GoCtx, "SELECT id, label FROM categories")
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "database error"})
		return
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category

		err := rows.Scan(&c.Id, &c.Label)
		if err != nil {
			ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "scan error"})
			return
		}

		categories = append(categories, c)
	}

	if len(categories) > 0 {
		ctx.Response.Status(200).JSON(map[string]any{"success": true, "categories": categories})
		return
	}
	ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "no categories found"})
}

func CategoriesById(ctx *sugar.SugarContext) {
	categoryId := ctx.Request.Params["id"]

	if categoryId == "" {
		ctx.Response.Status(400).JSON(map[string]any{"success": false, "message": "no category given"})
		return
	}

	rows, err := ctx.Database.Query(ctx.Request.GoCtx, "SELECT id, label, description, price, stock, options, old_price FROM products WHERE category_id = $1", categoryId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "database error"})
		return
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.Id, &p.Label, &p.Description, &p.Price, &p.Stock, &p.Options, &p.OldPrice)
		if err != nil {
			ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "scan error"})
			return
		}

		p.CategoryId = categoryId

		products = append(products, p)
	}

	if len(products) > 0 {
		ctx.Response.Status(200).JSON(map[string]any{"success": true, "products": products})
		return
	}
	ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "no products found"})
}

func DeleteCategory(ctx *sugar.SugarContext) {
	categoryId := ctx.Request.Params["id"]

	_, err := ctx.Database.Exec(ctx.Request.GoCtx, "DELETE FROM categories WHERE id = $1", categoryId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func AddCategory(ctx *sugar.SugarContext) {
	var category models.Category
	err := json.Unmarshal(ctx.Request.Body, &category)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}

	_, err = ctx.Database.Exec(ctx.Request.GoCtx, "INSERT INTO categories (id, label) VALUES ($1, $2)", category.Id, category.Label)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func EditCategory(ctx *sugar.SugarContext) {
	categoryId := ctx.Request.Params["id"]
	var newCategory models.Category
	err := json.Unmarshal(ctx.Request.Body, &newCategory)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal server error"})
		return
	}

	cmd, err := ctx.Database.Exec(ctx.Request.GoCtx, "UPDATE categories SET id = $1, label = $2 WHERE id = $3", newCategory.Id, newCategory.Label, categoryId)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	if cmd.RowsAffected() == 0 {
		ctx.Response.Status(404).JSON(map[string]any{"success": false, "message": "category not found"})
	}
	ctx.Response.Status(200).JSON(map[string]any{"success": true})
}

func CategoryById(ctx *sugar.SugarContext) {
	id := ctx.Request.Params["id"]
	var category models.Category
	err := ctx.Database.QueryRow(ctx.Request.GoCtx, "SELECT id, label FROM categories WHERE id = $1", id).Scan(&category.Id, &category.Label)
	if err != nil {
		ctx.Response.Status(500).JSON(map[string]any{"success": false, "message": "internal database error"})
		return
	}

	ctx.Response.Status(200).JSON(map[string]any{"success": true, "category": category})
}
package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func HandleRoutes(app *fiber.App) {
	app.Post("/todo", func(ctx *fiber.Ctx) error {
		requestBody := ctx.Body()
		trimmedBody := strings.TrimSpace(string(requestBody))

		if len(trimmedBody) < 4 {
			return ctx.Status(400).SendString("Task must have at least 3 characters")
		}

		var result Todo
		DBConn.Raw("INSERT INTO todos (task) VALUES (?) returning id", trimmedBody).Scan(&result)

		result.Task = trimmedBody

		return ctx.Status(200).JSON(result)
	})

	app.Patch("todo/:id", func(ctx *fiber.Ctx) error {
		todoID := ctx.Params("id")

		var updatedTodoData UpdateTodoInput
		if err := ctx.BodyParser(&updatedTodoData); err != nil {
			return err
		}

		// fieldların boş olup olmadıgını anlamak için struct tanımlamasında bu alanları
		// string ve bool yerine string ve bool referansı olarak tanımladım
		// çünkü done alanının default valuesı false olduğu için boş mu gelmiş yoksa gerçekten
		// false mu gelmiş anlaşılmıyor bu yüzden string ve bool yerine referanslarını kullandım
		// çünkü parse edilirken bu alanlar yoksa default value yerine nil atanıyor
		if updatedTodoData.Task == nil || updatedTodoData.Done == nil {
			return ctx.Status(400).SendString("All fields are required.")
		}

		// requestten gelen değerlerin nil olup olmadığını anlamak için
		// * ile dereference ediyoruz
		if len(*updatedTodoData.Task) < 4 {
			return ctx.Status(400).SendString("Task must have at least 3 characters")
		}

		var result TodoResp

		query := "UPDATE todos SET task = ?, done = ? WHERE id = ? returning *"
		DBConn.Raw(query, *updatedTodoData.Task, *updatedTodoData.Done, todoID).Scan(&result)

		if result.ID == 0 {
			return ctx.Status(404).SendString("Todo not found.")
		}
		return ctx.Status(200).JSON(result)
	})

	app.Delete("/todo/:id", func(ctx *fiber.Ctx) error {
		todoID, err := ctx.ParamsInt("id")

		var result TodoResp

		if err != nil {
			return ctx.Status(400).JSON("Please ensure that :id is an integer")
		}

		DBConn.Raw("DELETE FROM todos WHERE id = ? returning id", todoID).Scan(&result)

		if result.ID == 0 {
			return ctx.Status(404).SendString("Todo not found.")
		}

		return ctx.Status(200).JSON("Successfully deleted product")
	})

	app.Get("/todo", func(ctx *fiber.Ctx) error {
		var todoList []TodoResp

		DBConn.Raw("SELECT id, task, done, created_at, updated_at FROM todos").Scan(&todoList)

		return ctx.Status(200).JSON(todoList)
	})

	app.Get("/todo/:id", func(ctx *fiber.Ctx) error {
		todoID, err := ctx.ParamsInt("id")

		var result TodoResp

		if err != nil {
			return ctx.Status(400).JSON("Please ensure that :id is an integer")
		}

		DBConn.Raw("SELECT id, task, done, created_at, updated_at FROM todos WHERE id = ?", todoID).Scan(&result)

		if result.ID == 0 {
			return ctx.Status(404).SendString("Todo not found.")
		}

		return ctx.Status(200).JSON(result)
	})

}

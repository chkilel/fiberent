package handler

import (
	"context"

	"github.com/chkilel/fiberent/api/presenter"
	"github.com/chkilel/fiberent/entity"
	"github.com/chkilel/fiberent/usecase/pet"
	"github.com/gofiber/fiber/v2"
)

func NewPetHandler(app fiber.Router, ctx context.Context, service pet.UseCase) {
	app.Post("/", createPet(ctx, service))
	app.Get("/", listPets(ctx, service))
	app.Get("/:petId", getPet(ctx, service))
	app.Post("/:petId", updatePet(ctx, service))
	app.Delete("/:petId", deletePet(ctx, service))
}

func createPet(ctx context.Context, service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var pet *entity.Pet
		err := c.BodyParser(&pet)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		pet, err = service.CreatePet(ctx, pet.Name, pet.Age)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func getPet(ctx context.Context, service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := entity.StringToID(c.Params("petId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		pet, err := service.GetPet(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Pet Found",
			"data":    toJ,
		})
	}
}

func updatePet(ctx context.Context, service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("petId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var pet *entity.Pet

		err = c.BodyParser(&pet)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		pet.ID = id

		pet, err = service.UpdatePet(ctx, pet)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func deletePet(ctx context.Context, service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("petId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "Bad Id Format",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		err = service.DeletePet(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error deleting pet",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "pet deleted successfully",
			"error":  nil,
		})
	}
}

func listPets(ctx context.Context, service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.ListPets(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := make([]presenter.Pet, len(users))

		for i, pet := range users {
			toJ[i] = presenter.Pet{
				ID:   pet.ID,
				Name: pet.Name,
				Age:  pet.Age,
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Pets Found",
			"data":    toJ,
		})
	}
}

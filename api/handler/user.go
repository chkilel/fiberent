package handler

import (
	"context"

	"github.com/chkilel/fiberent/api/presenter"
	"github.com/chkilel/fiberent/entity"
	"github.com/chkilel/fiberent/usecase/user"
	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, ctx context.Context, service user.UseCase) {
	app.Post("/", createUser(ctx, service))
	app.Get("/", listUsers(ctx, service))
	app.Get("/:userId", getUser(ctx, service))
	app.Post("/:userId", updateUser(ctx, service))
	app.Delete("/:userId", deleteUser(ctx, service))
	app.Post("/:userId/pets", ownPets(ctx, service))
}

func createUser(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user *entity.User
		err := c.BodyParser(&user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user, err = service.CreateUser(ctx, user.Email, user.Password, user.FirstName, user.LastName)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func getUser(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := service.GetUser(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User Found",
			"data":    toJ,
		})
	}
}

func updateUser(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var user *entity.User

		err = c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user.ID = id
		user, err = service.UpdateUser(ctx, user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func deleteUser(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "Bad Id Format",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		err = service.DeleteUser(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error deleting user",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "user deleted successfully",
			"error":  nil,
		})
	}
}

func listUsers(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.ListUsers(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := make([]presenter.User, len(users))

		for i, user := range users {
			toJ[i] = presenter.User{
				ID:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    toJ,
		})
	}
}

func ownPets(ctx context.Context, service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var petIDs []*entity.ID

		err = c.BodyParser(&petIDs)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		err = service.OwnPets(ctx, &id, petIDs)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
		})
	}
}

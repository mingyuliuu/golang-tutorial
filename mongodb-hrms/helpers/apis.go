package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int32   `json:"age"`
}

func GetAllEmployees(ctx *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := MgDB.Database.Collection("employees").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var employees []Employee = make([]Employee, 0)
	if err := cursor.All(ctx.Context(), &employees); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.JSON(employees)
}

func CreateEmployee(ctx *fiber.Ctx) error {
	collection := MgDB.Database.Collection("employees")

	employee := new(Employee)
	if err := ctx.BodyParser(employee); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	employee.ID = ""
	insertionResult, err := collection.InsertOne(ctx.Context(), employee)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)

	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)

	return ctx.Status(201).JSON(createdEmployee)
}

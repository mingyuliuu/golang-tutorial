package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func UpdateEmployee(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return ctx.SendStatus(400)
	}

	employee := new(Employee)
	if err := ctx.BodyParser(employee); err != nil {
		ctx.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = MgDB.Database.Collection("employees").FindOneAndUpdate(ctx.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.SendStatus(400)
		}
		return ctx.SendStatus(500)
	}

	employee.ID = idParam

	return ctx.Status(200).JSON(employee)
}

func DeleteEmployee(ctx *fiber.Ctx) error {
	employeeID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := MgDB.Database.Collection("employees").DeleteOne(ctx.Context(), &query)

	if err != nil {
		return ctx.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return ctx.SendStatus(404)
	}

	return ctx.Status(200).JSON("Record deleted.")
}

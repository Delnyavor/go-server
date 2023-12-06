package controllers

import (
	"context"
	"time"

	"github.com/Delnyavor/go-fiber-mongo-hrms/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float32 `json:"age"`
}

func GetEmployees(c *fiber.Ctx) error {
	db := database.Mg.Db
	query := bson.D{{}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := db.Collection("employees").Find(ctx, query)

	if err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}

	var employees []Employee = make([]Employee, 0)

	if err := cursor.All(ctx, &employees); err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}

	c.JSON(employees)
	return nil
}

func CreateEmployee(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("employees")

	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}

	// causes the db to add a new tuple
	employee.ID = ""

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	insertionResult, err := collection.InsertOne(ctx, employee)

	if err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}

	// create filter
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	// find newly created record
	createdRecord := collection.FindOne(ctx, filter)

	var createdEmployee Employee

	createdRecord.Decode(&createdEmployee)

	c.Status(201).JSON(createdEmployee)

	return nil
}

func UpdateEmployee(c *fiber.Ctx) error {
	idParam := c.Params("id")

	employeeId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.SendStatus(400)
	}

	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		c.Status(400).SendString(err.Error())
		return err
	}

	filter := bson.D{{Key: "_id", Value: employeeId}}

	update := bson.D{{
		Key: "$set", Value: bson.D{{Key: "name", Value: employee.Name}, {Key: "age", Value: employee.Age}, {Key: "salary", Value: employee.Salary}},
	}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := database.Mg.Db.Collection("employees")
	singleResult := collection.FindOneAndUpdate(ctx, filter, update)

	if singleResult.Err() == mongo.ErrNoDocuments {
		c.SendStatus(400)
		return singleResult.Err()
	}

	if err := singleResult.Decode(employee); err != nil {

		employee.ID = idParam
		c.Status(200).JSON(employee)
		return err
	}

	c.Status(200).JSON(employee)
	return nil
}

func DeleteEmployee(c *fiber.Ctx) error {
	employeeId, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		c.SendStatus(400)
		return err
	}

	filter := bson.D{{Key: "_id", Value: employeeId}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	singleResult, err := database.Mg.Db.Collection("employees").DeleteOne(ctx, filter)

	if err != nil {
		c.SendStatus(500)
		return err
	}

	if singleResult.DeletedCount < 1 {
		c.SendStatus(404)
		return err
	}

	c.Status(200).JSON("record deleted")
	return nil
}

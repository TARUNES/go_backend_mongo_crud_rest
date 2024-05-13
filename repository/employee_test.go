package repository

import (
	"context"
	"log"
	"mongo_backend/models"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://tarunes12:tarunes12@cluster0.qtj9ppg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("MongoDB Connection Failed ", err)
	}
	log.Println("MongoDB Connected Successfully")

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB Ping Failed", err)
	}
	log.Println("MongoDB Ping Successful")

	return client
}

func TestMongoOpertions(t *testing.T) {
	mongoTestClient := NewMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	coll := mongoTestClient.Database("companyDB").Collection("employee_test")
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	empRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := models.Employee{EmployeeId: emp1, Name: "TonyStark", Department: "Science"}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			log.Fatalf("Insert Employee Failed  %v", err)
		}
		log.Println("Insert Employee Sucess", result)

	})

	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := models.Employee{EmployeeId: emp2, Name: "Steve Rogers", Department: "Solical Service"}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			log.Fatalf("Insert Employee Failed  %v", err)
		}
		log.Println("Insert Employee Sucess", result)

	})

	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeID(emp1)
		if err != nil {
			log.Fatalf("Get Employee1 Failed  %v", err)
		}
		t.Log("Get Employee Sucess ", result.Name)

	})

	t.Run("Update Employee 1", func(t *testing.T) {
		emp := models.Employee{EmployeeId: emp1, Name: "Stark", Department: "Science"}
		result, err := empRepo.UpdateEmployee(emp1, emp)
		if err != nil {
			log.Fatalf("Update Employee1 Failed  %v", err)
		}
		t.Log("Update Employee Sucess ", result)

	})

	t.Run("Delete Employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			log.Fatalf("Delete Employee1 Failed  %v", err)
		}
		t.Log("Delete Employee1 Sucess ", result)

	})

	t.Run("Find All Employee", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()
		if err != nil {
			log.Fatalf("Find All Employee Failed %v", err)
		}
		t.Log("Find All Employee Sucess ", result)

	})

	t.Run("Delete All Employee", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			log.Fatalf("Delete All Employee Failed %v", err)
		}
		t.Log("Delete All Employee Sucess ", result)

	})

}

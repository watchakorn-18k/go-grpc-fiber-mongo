package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct สำหรับเก็บข้อมูลผู้ใช้
type User struct {
	UserID string `json:"user_id" bson:"user_id"`
}

func main() {
	// เชื่อมต่อ MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// เลือกฐานข้อมูลและคอลเล็กชัน
	collection := client.Database("testdb").Collection("users")

	// สร้าง Fiber app
	app := fiber.New()

	// Endpoint เพื่อดึงข้อมูลผู้ใช้ทั้งหมด
	app.Get("/user/:user_id", func(c *fiber.Ctx) error {
		// ดึงค่า user_id จาก URL
		userID := c.Params("user_id")

		// สร้างคอลเล็กชันสําหรับดึงข้อมูลผู้ใช้
		filter := bson.M{"user_id": userID}

		// ดึงข้อมูลผู้ใช้จาก MongoDB
		startTime := time.Now()
		var user User
		err := collection.FindOne(c.Context(), filter).Decode(&user)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Time taken: %v", duration)

		// ส่งข้อมูลผู้ใช้กลับในรูปแบบ JSON
		return c.JSON(user)
	})

	// เริ่มต้นเซิร์ฟเวอร์ Fiber ที่พอร์ต 3000
	log.Println("Fiber server is running on :3000")
	log.Fatal(app.Listen(":3001"))
}

package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/services"
)

func main() {
	// เชื่อมต่อ gRPC server ด้วย Insecure credentials (สำหรับทดสอบเท่านั้น)
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(":50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	// สร้าง Fiber app
	app := fiber.New()

	// Endpoint สำหรับเรียกใช้ GetUser จาก gRPC server
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		// เรียกใช้งาน gRPC client
		startTime := time.Now()
		usersClient := services.NewUserServiceClient(cc)

		// ดึงค่า id จากพารามิเตอร์ของ URL
		id := c.Params("id")

		// สร้าง request และเรียกใช้ GetUser จาก gRPC server
		data, err := usersClient.GetUser(c.Context(), &services.GetUserRequest{
			UserId: id,
		})
		if err != nil {
			return err
		}
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Time taken: %v", duration)

		// ส่งข้อมูลกลับในรูปแบบ JSON
		return c.JSON(data)
	})

	// เริ่มต้นเซิร์ฟเวอร์ Fiber ที่พอร์ต 3000
	log.Println("Fiber server is running on :3000")
	log.Fatal(app.Listen(":3000"))
}

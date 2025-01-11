package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var(
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
)

var (
 db    *gorm.DB
 rdb   *redis.Client
 ctx   = context.Background()
)

type User struct {
 ID   uint   `json:"id" gorm:"primaryKey"`
 Name string `json:"name"`
}

func initMySQL() {
 dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
 var err error
 db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil {
  panic("failed to connect to MySQL: " + err.Error())
 }

 db.AutoMigrate(&User{})
}

func initRedis() {
 rdb = redis.NewClient(&redis.Options{
  Addr: fmt.Sprintf("%s:%s",redisHost, redisPort),
  DB: 0,
  Password: "",
 })
}

func main() {

 initMySQL()
 initRedis()

 r := gin.Default()

 r.GET("/users", func(c *gin.Context) {
  var users []User
  result := db.Find(&users)
  if result.Error != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
   return
  }
  c.JSON(http.StatusOK, users)
 })

 r.GET("/redis", func(c *gin.Context) {
  err := rdb.Set(ctx, "message", "Hello, Redis!", 0).Err()
  if err != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
   return
  }

  val, err := rdb.Get(ctx, "message").Result()
  if err != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
   return
  }
  c.JSON(http.StatusOK, gin.H{"message": val})
 })

 r.Run(":8080") // Run on port 8080
}
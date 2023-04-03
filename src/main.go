package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go-name/src/models"
	"go-name/src/pkg/logging"
	"go-name/src/routes"
	"go.uber.org/ratelimit"
	"gorm.io/gorm"
	"log"
	"runtime"
	"time"
)

var (
	db    *gorm.DB = models.ConnectDB()
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 100, "request per second")
)

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func main() {
	ConfigRuntime()

	defer models.DisconnectDB(db)

	limit = ratelimit.New(*rps)

	gin.SetMode(gin.DebugMode)
	r := gin.New()

	routes.Routes(r)
	logging.Setup()

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", *rps))

	r.Run()
}

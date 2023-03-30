package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go-name/src/routes"
	"go.uber.org/ratelimit"
	"log"
	"runtime"
	"time"
)

var (
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

	limit = ratelimit.New(*rps)

	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.GET("/dateResult", routes.GetDateResult)

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", *rps))

	r.Run()
}

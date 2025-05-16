package main


import (
  "log"
  "os"
  "github.com/shirou/gopsutil/v3/cpu"
  "github.com/shirou/gopsutil/v3/mem"
  "github.com/shirou/gopsutil/v3/disk"
  "time"
  "math"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)


func main() {

  err := godotenv.Load()
  if err != nil {
    log.Println("No .env file found, using default port")
  }

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })


  router.GET("/stats", func(c *gin.Context) {

    percent, _ := cpu.Percent(time.Second, false)
    vmStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("/")

    c.JSON(200, gin.H{
      "cpu (%)": round(percent[0],2),
      "mem (%)": round(vmStat.UsedPercent,2),
      "disk (%)": round(diskStat.UsedPercent,2),
    })
  })
  
  err = router.Run(":" + port)
}

func round(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}

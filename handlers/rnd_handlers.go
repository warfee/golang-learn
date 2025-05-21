package handlers

import (
	// "fmt"
  	// "os"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
  	"github.com/shirou/gopsutil/v3/mem"
  	"github.com/shirou/gopsutil/v3/disk"
  	"time"
  	"math"
  	// "net/http"
)

func Ping(c *gin.Context) {

	c.JSON(200, gin.H{
      "message": "pong",
    })
}


func Stats(c *gin.Context) {

	  percent, _ := cpu.Percent(time.Second, false)
    vmStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("/")

    c.JSON(200, gin.H{
      "cpu (%)": round(percent[0],2),
      "mem (%)": round(vmStat.UsedPercent,2),
      "disk (%)": round(diskStat.UsedPercent,2),
    })
}

func round(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}
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
  	// "database/sql"
  	// _ "github.com/go-sql-driver/mysql"
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

// func MysqlConnect(c *gin.Context) {


//     host := os.Getenv("MYSQL_DB_HOST")
//     port := os.Getenv("MYSQL_DB_PORT")
//     name := os.Getenv("MYSQL_DB_NAME")
//     username := os.Getenv("MYSQL_DB_USERNAME")
//     password := os.Getenv("MYSQL_DB_PASSWORD")

//     if host == "" || port == "" || name == "" || username == "" || password == "" {
//         c.JSON(http.StatusBadRequest, gin.H{
//             "status": "Database connection info is missing",
//         })
//         return
//     }

//     database_credential := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, name)

//     db, err := sql.Open("mysql", database_credential)

//     if errPing := db.Ping(); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
//             "status": "Ping failed",
//             "error":  errPing.Error(),
//         })
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{
//         "status": "Success to connect to DB",
//     })
// }







func round(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}
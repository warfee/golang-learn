package handlers

import (
	"fmt"
	"time"
	"log"
	"math"
	"github.com/shirou/gopsutil/v3/cpu"
  	"github.com/shirou/gopsutil/v3/mem"
  	"github.com/shirou/gopsutil/v3/disk"

  	"database/sql"
  	_ "github.com/ncruces/go-sqlite3/driver"
  	_ "github.com/ncruces/go-sqlite3/embed"
)

func StartMetricsTicker() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Collecting metrics at", t)
			

		    percent, _ := cpu.Percent(time.Second, false)
		    vmStat, _ := mem.VirtualMemory()
			diskStat, _ := disk.Usage("/")

			var percentFloat float64
			var vmStatFloat float64
			var diskStatFloat float64

			percentFloat = roundValue(percent[0],2)
			vmStatFloat = roundValue(vmStat.UsedPercent,2)
			diskStatFloat = roundValue(diskStat.UsedPercent,2)


			db, err := sql.Open("sqlite3", "database.sqlite")

  			if err != nil {

			    log.Fatal(err)
			}

		  	defer db.Close()

			  res, err := db.Exec(`INSERT INTO monitor (cpu,memory,storage) VALUES (?,?,?)`, percentFloat,vmStatFloat,diskStatFloat)
			  if err != nil {
			    log.Fatal(err)
			  }

			  id, _ := res.LastInsertId()
			  log.Printf("Inserted monitoring with ID: %d\n", id)
		}
	}()
}


func roundValue(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}

package main

import (
	//"fmt"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/ricochet2200/go-disk-usage/du"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	discPtr := flag.Bool("d", false, "discovery bool")
	hotPtr := flag.Bool("h", false, "hot drive bool")
	backupPtr := flag.Bool("b", false, "backup bool")

	// DISCovery free space and check if its less than buffer
	if *discPtr {
		discPath := os.Getenv("DISC")
		usage := du.NewDiskUsage(discPath)
		gbAv := int(usage.Available() / 1073741824)
		os.Setenv("DISC_FS", strconv.Itoa(gbAv))
	}

	// HOT free space and check if its less than buffer
	if *hotPtr {
		hotPath := os.Getenv("HOT1")
		usage := du.NewDiskUsage(hotPath)
		gbAv := int(usage.Available() / 1073741824)
		os.Setenv("HOT1_FS", strconv.Itoa(gbAv))
	}

	// Backup free space and check if its less than buffer
	if *backupPtr {
		backupPath := os.Getenv("BACKUP")
		usage := du.NewDiskUsage(backupPath)
		gbAv := int(usage.Available() / 1073741824)
		os.Setenv("BACKUP_FS", strconv.Itoa(gbAv))
	}


}

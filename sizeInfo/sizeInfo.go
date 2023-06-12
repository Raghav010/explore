package main

import (
	//"fmt"
	"log"
	"os"
	"flag"
	"github.com/joho/godotenv"
	"github.com/ricochet2200/go-disk-usage/du"
)


func main(){

	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("error loading .env file")
	}

	discPtr:=flag.Bool("d",false,"discovery bool")
	hotPtr:=flag.Bool("h",false,"hot drive bool")
	backupPtr:=flag.Bool("b",false,"backup bool")
	photosHotPtr:=flag.Bool("ps",false,"photos size in hot")


	// DISCovery free space and check if its less than buffer
	if *discPtr{
		discPath:=os.Getenv("DISC")
		usage:=du.NewDiskUsage(discPath)
		gbAv:=int(usage.Available()/1073741824)
		os.Setenv("DISC_FS",gbAv)
	}


	// HOT free space and check if its less than buffer
	if *hotPtr{
		hotPath:=os.Getenv("HOT1")
		usage:=du.NewDiskUsage(hotPath)
		gbAv:=int(usage.Available()/1073741824)
		os.Setenv("HOT1_FS",gbAv)
	}

	// Backup free space and check if its less than buffer
	if *backupPtr{
		backupPath:=os.Getenv("BACKUP")
		usage:=du.NewDiskUsage(backupPath)
		gbAv:=int(usage.Available()/1073741824)
		os.Setenv("BACKUP_FS",gbAv)
	}

	// set size of photos in HOT2
	if *photosHotPtr{
		hotPPath:=(os.Getenv("HOT2")+os.Getenv("HOT_PHOTOS_NAME"))
		usage:=du.NewDiskUsage(hotPPath)
		gbUsed:=int(usage.Used()/1073741824)
		os.Setenv("HOT_PHOTOS_SIZE",gbUsed)
	}


}
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/tidwall/gjson"
)




func main() {


	dirPath:=os.Args[1]
	dirContents,err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}


	for _,entry:=range(dirContents) {

		// only chronologize if its a file
		if !entry.IsDir() {
			fmt.Print(entry.Name()+ " ")





			// extracting exif data from files
			imgFile,err:=os.Open(dirPath+entry.Name())
			if err!=nil{
				log.Fatal(err)
			}
			metaData,err:=exif.Decode((imgFile))
			if err!=nil{
				os.Mkdir(dirPath+"unknown",0755)
				os.Rename(dirPath+entry.Name(),dirPath+"unknown/"+entry.Name())
				fmt.Println("cant decode exif data")
				continue
			}
			jsonBytes,err:=metaData.MarshalJSON()
			if err!=nil{
				os.Mkdir(dirPath+"unknown",0755)
				os.Rename(dirPath+entry.Name(),dirPath+"unknown/"+entry.Name())
				fmt.Println("cant decode exif data")
				continue
			}
			jsonString:=string(jsonBytes)
			//fmt.Println(jsonString)


			// date un-formatted
			dateUnF:=strings.Split(gjson.Get(jsonString, "DateTime").String(), " ")[0]
			if dateUnF==""{
				os.Mkdir(dirPath+"unknown",0755)
				os.Rename(dirPath+entry.Name(),dirPath+"unknown/"+entry.Name())
				fmt.Println("no date present")
				continue
			}

			dateSplit:=strings.Split(dateUnF,":")
			if len(dateSplit) <2{
				os.Mkdir(dirPath+"unknown",0755)
				os.Rename(dirPath+entry.Name(),dirPath+"unknown/"+entry.Name())
				fmt.Println("no date present")
				continue
			}

			date:=strings.Join(dateSplit[:2],"-")
			fmt.Println("dated:"+date)
			os.Mkdir(dirPath+date,0755)
			os.Rename(dirPath+entry.Name(),dirPath+date+"/"+entry.Name())

		}
	}
}
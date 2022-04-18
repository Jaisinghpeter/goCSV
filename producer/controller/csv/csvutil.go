package csv

import (
	// "bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"producer/configurations"
	"producer/constants"
	"strconv"
	"time"
	"producer/controller/kafka"
	"producer/model"
)

func readCsvFile(filePath string){
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer file.Close()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }
    fmt.Println(configurations.KafkaBrokerAddress)
    processStruct(file)
	
}

func processStruct(file *os.File){
    switch configurations.StructName{
        case constants.Employee:
            fmt.Println("Employee Format")
            processEmployee(file)            
        default:
            fmt.Println("No Structure selected")
    }
}

func processEmployee(file *os.File){
    
    csvReader := csv.NewReader(file)
    csvReader.Read()
    for {
        var recordChunk  []*model.Employee
        for j:=0;j<configurations.Chunk_size;j++{
            record, err := csvReader.Read()
            if err == io.EOF {
                break
            }
            if err != nil {
                log.Fatal(err)
            }
            
            address := model.NewAddress(record[model.Country], record[model.City], record[model.Place], record[model.Zip])
            empId, _ := strconv.Atoi(record[model.EmpId])
            employee := model.NewEmployee(empId, record[model.Name], record[model.Email], address)
            recordChunk = append(recordChunk, employee)
        }
        if(len(recordChunk) == 0){
            break;
        }
        time.Sleep(time.Second)
        recordChunkInBytes, _ := json.Marshal(recordChunk)
        kafka.DumptoKafka(configurations.KafkaTopicEmpRegistration, recordChunkInBytes)
        fmt.Println(len(recordChunk))
    }   
}
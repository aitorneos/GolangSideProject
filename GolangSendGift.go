package main

import (
    "flag"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Log struct {
	id   int    `json:"id"`
	TimeStamp string `json:"TimeStamp"`
	Message string `json:"Message"`
	level string `json:"level"`
	logger string `json:"logger"`
	LoggingSourceMachine string `json:"LoggingSourceMachine"`
	SessionId string `json:"SessionId"`
	RemoteHost string `json:"RemoteHost"`
}

func main() {
    fmt.Println("Go MySQL first implementation")
    
	var username = flag.String("username", "aitor", "MSql username")
	var password = flag.String("password", "123456", "MSql user password")
	var ip = flag.String("ip", "123456", "MSql db server ip")
	var port = flag.String("port", "123456", "MSql db server port")
	
	flag.Parse()
	
    db, err := sql.Open("mysql", *username+ ":" + *password + "@tcp(" + *ip + ":" + *port + ")/xplogdb")
    
    if err != nil {
        panic(err.Error())
    }
    
    defer db.Close()
    
    query, err := db.Query("SELECT * FROM xplog x  WHERE x.Message LIKE '%UpdateBalance%'")
    
    if err != nil {
        panic(err.Error())
    }
	
	for query.Next() {
		var log Log

		err = query.Scan(&log.id, &log.TimeStamp,&log.Message,&log.level ,&log.logger ,&log.LoggingSourceMachine ,&log.SessionId,&log.RemoteHost)
		if err != nil {
			panic(err.Error()) 
		}

		fmt.Println(log.Message)
	}
    // be careful deferring Queries if you are using transactions
    defer query.Close()
    
}
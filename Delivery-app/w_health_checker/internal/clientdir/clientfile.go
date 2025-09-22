package clientdir

import (
	"abc/internal/models"
	"abc/internal/storage"
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var logChannel chan *models.CheckStruct = make(chan *models.CheckStruct, 1000)
var done chan struct{} = make(chan struct{})

func CreateLogs() {
	log.Println("Health check service started")
	db := getConn()
	IDb := storage.NewStorage(db)
	ctx, cancel := context.WithCancel(context.Background())
	tickerCL := time.NewTicker(6 * time.Second)
	tickerDb := time.NewTicker(12 * time.Second)
	for {
		select {
		case <-tickerCL.C:
			createLog()
		case <-tickerDb.C:
			go dbLog(ctx, cancel, IDb)
		}
	}
}

func dbLog(ctx context.Context, cancel context.CancelFunc, aa storage.IStorage) {
	select {
	case <-done:
		cancel()
	case <-ctx.Done():
		return
	default:
		var sliceFromChannel []*models.CheckStruct
		for item := range logChannel {
			sliceFromChannel = append(sliceFromChannel, item)
			if len(logChannel) == 0 {
				aa.InsertData(sliceFromChannel)
				<-done
			}
		}
	}
}

func createLog() {
	links := []string{
		"http://localhost:8080/auth_v1/healthz",
		"http://localhost:8084/shop/v1/healthz",
		"http://localhost:8086/admin/v1/healthz",
		"http://localhost:8088/deliver/v1/healthz",
	}
	linksInChannel := make(chan string, len(links))

	for _, linkInSlice := range links {
		linksInChannel <- linkInSlice
	}

	var wg sync.WaitGroup

	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(linksInChannel, &wg)
	}

	close(linksInChannel)
	wg.Wait()
}

func worker(linksInChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for link := range linksInChannel {
		processLinks(link)
	}
}

func processLinks(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
	}

	t := time.Now().Format("01-02-2006 15:04:05")
	checkInst := &models.CheckStruct{
		Datetime: t,
		Server:   link,
		Status:   strconv.Itoa(resp.StatusCode),
	}

	if err != nil {
		logChannel <- checkInst
	} else {
		logChannel <- checkInst
	}
}

func getConn() *sql.DB {
	dbNewString1 := "postgres://postgres:postgres@localhost:5432/checkh?sslmode=disable"
	dbNewConn, err := sql.Open("postgres", dbNewString1)

	if err != nil {
		log.Println("dbNewString", err)
	}

	err = dbNewConn.Ping()
	if err != nil {
		log.Println("InsertDataIntoDb Ping", err)
	}

	return dbNewConn
}

package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const MODE = "release"
const PORT = ":8081"

type Server struct {
	DB     *sql.DB
	Router *gin.Engine
}

func ConnectDB() (*sql.DB, error) {
	// membuat variable yang berisi format untuk koneksi dan database
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "127.0.0.1", "6606", "online_book_store")
	// Instance sebuah object dari sebuah struct Value yang bertipe data map[string]interface{}
	val := url.Values{}
	// Menambahkan sebuah location dan menentukan Asia/Jakarta
	val.Add("loc", "Asia/Jakarta")
	// Membuat variable yang berisikan hasil penggabungan variable connection dengan value location
	// Fungsi Encode untuk merubah menjadi sebuah string
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	// membuka koneksi
	db, err := sql.Open(`mysql`, dsn)
	// pengecekan error apakah ada error atau tidak saat buka koneksi
	if err != nil {
		log.Fatal(err)
	}
	// melakukan test ke database apakah sudah bisa digunakan dan ada pengecekan error
	if err := db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
		os.Exit(0)
	}
	log.Println("DataBase Successfully Connected")
	// mengembalikan sebuah koneksi yang sudah terbuka
	return db, err

}

// fungsi dgunakan untuk menentukan mode dan instance object router
func CreateRouter() *gin.Engine {
	if MODE != "release" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Menginstance object router
	r := gin.Default()
	return r
}

func InitRouter(db *sql.DB, r *gin.Engine) *Server {
	return &Server{
		DB:     db,
		Router: r,
	}
}

func (server *Server) InitializeRoutes() {
	// Membuat sebuah group router
	// r := server.Router.Group("v1")
}

func Run(r *gin.Engine) {
	fmt.Println("Listening to port 8801")
	err := r.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}

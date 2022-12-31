package contorollers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tahahmmcgl/kullanici_api/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		fmt.Println("dbye bağlanılmaya çalışılıyor")
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		fmt.Println("dbye bağlanıldı")
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}

	}
	//database migration
	//database tablolarını oluşturmak için kullanılır
	server.DB.AutoMigrate(&models.User{})
	//server.DB.Debug().AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()

	//initialize routes
	//rotaları oluşturmak için kullanılır
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

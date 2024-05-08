package crud_functions

import (
	"hardwaremonitioringexporter/pb/pb"
	"sync"

	"gorm.io/gorm"
)

type Server struct {
	Db_connection *gorm.DB
}
type Employee struct {
	pb.UnimplementedListEmployeeRecordServiceServer
	ID       int
	Name     string
	Position string
	Salary   float64
}

// DatabaseManager manages database operations
type DatabaseManager struct {
	Db     *gorm.DB
	DbLock sync.Mutex // Mutex lock for synchronizing database access
}

package database

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email       string    `gorm:"unique;not null"`
	Password    string
	Name        string
	Address     string
	PhoneNumber string
	ImageURL    string
	Jobs        []Jobs `gorm:"foreignKey:ClientID"`
}

type Talent struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email       string    `gorm:"unique;not null"`
	Password    string
	Name        string
	AboutMe     string
	ImageURL    string
	Address     string
	PhoneNumber string
	Jobs        []Jobs          `gorm:"foreignKey:TalentID"`
	OnAuction   []JobsOnAuction `gorm:"foreignKey:TalentID"`
}

type Jobs struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClientID     uuid.UUID `gorm:"type:uuid"`
	TalentID     uuid.UUID
	Name         string
	Descriptions string
	FixedPrice   uint16
	Address      string
	Latitude     float64
	Longitude    float64
	ImageURL     string
}

type JobsOnAuction struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TalentID     uuid.UUID
	PriceOnBid   int
	StartAuction time.Time
	EndAuction   time.Time
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("dsn: ", os.Getenv("DATABASE_URL"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(Client{})
	db.AutoMigrate(Talent{})
	db.AutoMigrate(Jobs{})
	db.AutoMigrate(JobsOnAuction{})
}

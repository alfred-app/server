package database

import (
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	ImageURL    string    `json:"imageURL"`
	Jobs        []Jobs    `gorm:"foreignKey:ClientID; references:ID" json:"jobs"`
}

type Talent struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	AboutMe     string    `json:"aboutMe"`
	ImageURL    string    `json:"imageURL"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Jobs        []Jobs    `gorm:"foreignKey:TalentID; references:ID" json:"jobs"`
	PlacedBid   []BidList `gorm:"foreignKey:TalentID; references:ID" json:"placedBid"`
}

type Jobs struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ClientID     uuid.UUID  `json:"clientID"`
	TalentID     *uuid.UUID `gorm:"default:null; references:ID" json:"talentID"`
	Name         string     `json:"name"`
	Descriptions string     `json:"descriptions"`
	FixedPrice   int        `gorm:"default:null" json:"fixedPrice"`
	Address      string     `json:"address"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
	ImageURL     string     `json:"imageURL"`
	BidList      []BidList  `gorm:"foreignKey:JobID; references:ID" json:"bidList"`
}

type BidList struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TalentID   uuid.UUID `json:"talentID"`
	JobID      uuid.UUID `json:"jobID"`
	PriceOnBid int       `json:"priceOnBid"`
	BidPlaced  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"bidPlaced"`
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Client{}, &Talent{}, &Jobs{}, &BidList{})
}

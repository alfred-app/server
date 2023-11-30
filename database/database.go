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
	Email       string    `gorm:"column:email; unique" json:"email"`
	Password    string    `gorm:"column:password" json:"password"`
	Name        string    `gorm:"column:name" json:"name"`
	Address     string    `gorm:"column:address" json:"address"`
	PhoneNumber string    `gorm:"column:phoneNumber" json:"phoneNumber"`
	ImageURL    string    `gorm:"column:imageURL" json:"imageURL"`
	Jobs        []Jobs    `gorm:"foreignKey:ClientID; references:ID" json:"jobs"`
}

type Talent struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email       string    `gorm:"column:email; unique" json:"email"`
	Password    string    `gorm:"column:password" json:"password"`
	Name        string    `gorm:"column:name" json:"name"`
	AboutMe     string    `gorm:"column:aboutMe" json:"aboutMe"`
	ImageURL    string    `gorm:"column:imageURL" json:"imageURL"`
	Address     string    `gorm:"column:address" json:"address"`
	PhoneNumber string    `gorm:"column:phoneNumber" json:"phoneNumber"`
	Jobs        []Jobs    `gorm:"foreignKey:TalentID; references:ID" json:"jobs"`
	PlacedBid   []BidList `gorm:"foreignKey:TalentID; references:ID" json:"placedBid"`
}

type Jobs struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ClientID     uuid.UUID  `gorm:"column:clientID" json:"clientID"`
	TalentID     *uuid.UUID `gorm:"column:talentID; default:null; references:ID" json:"talentID"`
	Name         string     `gorm:"column:name" json:"name"`
	Descriptions string     `gorm:"column:descriptions" json:"descriptions"`
	FixedPrice   int        `gorm:"column:fixedPrice; default:null" json:"fixedPrice"`
	Address      string     `gorm:"column:address" json:"address"`
	Latitude     float64    `gorm:"column:latitude" json:"latitude"`
	Longitude    float64    `gorm:"column:longitude" json:"longitude"`
	ImageURL     string     `gorm:"column:imageURL" json:"imageURL"`
	BidList      []BidList  `gorm:"foreignKey:JobID; references:ID" json:"bidList"`
}

type BidList struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TalentID   uuid.UUID `gorm:"column:talentID" json:"talentID"`
	JobID      uuid.UUID `gorm:"column:jobID" json:"jobID"`
	PriceOnBid int       `gorm:"column:priceOnBid" json:"priceOnBid"`
	BidPlaced  time.Time `gorm:"column:bidPlaced;default:CURRENT_TIMESTAMP" json:"bidPlaced"`
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

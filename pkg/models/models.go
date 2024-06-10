package models

type User struct {
    ID        	uint   `json:"id" gorm:"primaryKey"`
    Name      	string `json:"name"`
    Email     	string `json:"email"`
    Password  	string `json:"-"`
}

type Listing struct {
    ID          	uint   `json:"id" gorm:"primaryKey"`
    Title       	string `json:"title"`
    Description 	string `json:"description"`
    Price       	float64 `json:"price"`
    HostID      	uint   `json:"host_id"`
}

type Booking struct {
    ID        	uint   `json:"id" gorm:"primaryKey"`
    ListingID 	uint   `json:"listing_id"`
    UserID    	uint   `json:"user_id"`
    StartDate 	string `json:"start_date"`
    EndDate   	string `json:"end_date"`
}

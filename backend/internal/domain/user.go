package domain

import "time"

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Age       int32     `db:"age"`
	Gender    Gender    `db:"gender"`
	Location  string    `db:"location"`
	BIO       string    `db:"bio"`
	Goal      Goal      `db:"goal"`
	Interest  Interest  `db:"interest"`
	Zodiac    Zodiac    `db:"zodiac"`
	Height    int32     `db:"height"`
	Education Education `db:"education"`
	Children  Children  `db:"children"`
	Alcohol   Alcohol   `db:"alcohol"`
	Smoking   Smoking   `db:"smoking"`
	Hidden    bool      `db:"hidden"`
	Verified  bool      `db:"verified"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

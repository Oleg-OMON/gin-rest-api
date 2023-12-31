package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gopkg.in/nullbio/null.v4"
)

type MyNullString struct {
	value sql.NullString
}

type MyNullFloat64 struct {
	value sql.NullFloat64
}

// убирает поле value: true/false при вывове sql.NullString
func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.value.Valid {
		return json.Marshal(s.value.String)
	}
	return []byte(`null`), nil
}

// Тоже с Float
func (s MyNullFloat64) MarshalJSON() ([]byte, error) {
	if s.value.Valid {
		return json.Marshal(s.value.Float64)
	}
	return []byte(`null`), nil
}

type Player struct {
	PlayerId    int          `json: "player_id" db:"player_id"`
	FirstName   string       `json: "first_name" db:"first_name"`
	LastName    string       `json: "last_name" db:"last_name"`
	Nickname    string       `json: "nickname" db:"nickname"`
	Citizenship MyNullString `json: "citizenship" db:"citizenship"`
	Dob         string       `json: "dob" db:"dob"`
	Role        string       `json: "role" db:"role"`
}

type Game struct {
	GameId   int          `json: "game_id"`
	Team     string       `json: "team"`
	City     MyNullString `json: "city"`
	Goals    null.Uint    `json: "goals"`
	GameDate time.Time    `json: "game_date"`
	Own      null.Uint    `json: "own"`
}

// участие игрока в матче
type Lineups struct {
	Start   string          `json: start`
	GameId  int             `json: "game_id"`
	LayerId int             `json: "layer_id"`
	TimeIn  sql.NullFloat64 `json: "time_in"` //  число минут, проведенных игроком на поле; NULL, если не выходил.
	Goals   null.Uint       `json: "goals"`   //  число голов, которые игрок забил в матче; NULL, если не забивал
	Cards   MyNullString    `json: "cards"`
}

type ResultModelsPlayerLineup struct {
	// Тут нужна композиция что бы не дублировать поля, а просто ссылать на их тип?
	Nickname string        `json: "nickname"`
	Team     string        `json: "team"`
	Start    string        `json: start`
	TimeIn   MyNullFloat64 `json: "time_in"`
	Goals    null.Uint     `json: "goals"`
	Cards    MyNullString  `json: "cards`
}

type User struct {
	ID       uuid.UUID `json: "id" gorm:"default:uuid_generate_v4()"`
	Name     string    `json: "name"`
	Email    string    `json: "email"`
	Password string    `json: "password"`
}

type SingUpInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type SingInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

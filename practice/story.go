package main

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"strings"
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "Saurav@123"
	DB_NAME     = "postgres"
)

var called uint64

type Word struct {
	Word string `db:"word" json:"word"`
}
type createResponse struct {
	ID              int64  `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	CurrentSentence string `db:"current_sentence" json:"current_sentence"`
}

type storyResponse struct {
	Id         int64     `db:"id" json:"id"`
	Title      string    `db:"title" json:"title"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
}
type sentences struct {
	Sentences string `db:"sentences" json :"sentences"`
}
type storyResponseDetails struct {
	Id         int64       `db:"id" json:"id"`
	Title      string      `db:"title" json:"title"`
	Created_at time.Time   `db:"created_at" json:"created_at"`
	Updated_at time.Time   `db:"updated_at" json:"updated_at"`
	Paragraph  []sentences `json:"paragraph"`
}

func OpenConnection() *sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("--database Connected Successfully--")

	return db
}

func main() {
	http.HandleFunc("/createstory", CreateStory)
	http.HandleFunc("/story/", getStory)
	http.HandleFunc("/story/1/", getStoryDetails)
	http.ListenAndServe(":8080", nil)

}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var create_response createResponse

	param2 := r.URL.Query().Get("story_id")

	var word Word

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		panic(err)
	}

	res := strings.Contains(word.Word, " ")
	if res == true {
		http.Error(w, "Multiple words send", http.StatusBadRequest)
	} else {
		atomic.AddUint64(&called, 1)

		if called == 1 {

			created_at := time.Now()
			updated_at := time.Now()

			str := `INSERT INTO story(story_id,title,created_at,updated_at)
		VALUES($1,$2,$3,$4) `
			_, err := db.Exec(str, param2, word.Word, created_at, updated_at)
			if err != nil {
				panic(err)
			}
		}
		if called == 2 {

			title_second := " " + word.Word
			str := `update story set title  = CONCAT(title,$1::text)`
			_, err := db.Exec(str, title_second)
			if err != nil {
				panic(err)
			}
		}

		if called >= 3 && called < 18 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
		WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called == 18 || called == 33 || called == 48 || called == 63 || called == 78 || called == 93 || called == 108 || called == 123 || called == 138 {

			created_at := time.Now()
			updated_at := time.Now()
			str := `INSERT INTO story(story_id,sentences,created_at,updated_at)
			VALUES($1,$2,$3,$4) `
			_, err := db.Exec(str, param2, word.Word, created_at, updated_at)
			if err != nil {
				panic(err)
			}
		}
		if called > 18 && called < 33 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}

		if called > 33 && called < 48 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 48 && called < 63 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 63 && called < 78 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 78 && called < 93 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 93 && called < 108 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 108 && called < 123 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 123 && called < 138 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 138 && called < 153 {
			sentenceAdd := " " + word.Word
			str := `UPDATE story SET sentences = CONCAT(sentences,$1::text)
			WHERE updated_at IN (SELECT MAX(updated_at) FROM story)`
			_, err := db.Exec(str, sentenceAdd)
			if err != nil {
				panic(err)
			}
		}
		if called > 153 {
			io.WriteString(w, " Sentence Creation Limit Reached")
		}

		str := `SELECT  story_id as id, COALESCE(title, '') as title , COALESCE(sentences, '') as current_sentence FROM story WHERE story_id= $1
		order by updated_at desc
		LIMIT 1`

		err = db.Get(&create_response, str, param2)
		if err != nil {
			panic(err)
		}
		rs, _ := json.Marshal(&create_response)
		w.Write(rs)

	}
}

func getStory(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	param := r.URL.Query().Get("story_id")

	story := storyResponse{}

	str := `SELECT story_id as id, title, created_at, updated_at FROM story WHERE story_id = $1 LIMIT 1`
	err := db.Get(&story, str, param)
	if err != nil {
		panic(err)
	}
	rs, err := json.Marshal(story)
	if err != nil {
		panic(err)
	}

	w.Write(rs)

}

func getStoryDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	param := r.URL.Query().Get("story_id")

	story := storyResponse{}

	str := `SELECT story_id as id, title, created_at, updated_at FROM story WHERE story_id = $1 LIMIT 1`
	err := db.Get(&story, str, param)
	if err != nil {
		panic(err)
	}

	sd := GetStoryParagraphs(db, story)

	rs, err := json.Marshal(sd)
	if err != nil {
		panic(err)
	}
	w.Write(rs)

}

func GetStoryParagraphs(db *sqlx.DB, story storyResponse) storyResponseDetails {
	st := storyResponseDetails{
		Id:         story.Id,
		Title:      story.Title,
		Created_at: story.Created_at,
		Updated_at: story.Updated_at,
		Paragraph:  []sentences{},
	}
	sen := []sentences{}
	str := `SELECT sentences FROM story WHERE story_id =$1`
	err := db.Select(&sen, str, st.Id)
	if err != nil {
		panic(err)
	}
	st.Paragraph = sen

	return st

}

package main

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "Saurav@123"
	DB_NAME     = "postgres"
)

type Audio struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Audio_File    string `json:"audio_file"`
	Creator_Name  string `json:"creator_name"`
	Creator_Email string `json:"creator_email"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	defer db.Close()

}

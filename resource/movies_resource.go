package resource

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"movies-api/database"
	"movies-api/domain/dto"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("M=GetMovies")
	db := database.SetupDB()

	rows, err := db.Query("SELECT * FROM movies")
	database.CheckError(err)

	var movies []database.Movie

	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)
		database.CheckError(err)

		movies = append(movies, database.Movie{
			MovieID:   movieID,
			MovieName: movieName,
		})
	}

	response := dto.JsonResponse{
		Type: "success",
		Data: movies,
	}

	json.NewEncoder(w).Encode(response)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie database.Movie

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &movie)

	fmt.Printf("M=CreateMovie, movie=%v", movie)

	response := dto.JsonResponse{}

	if movie.MovieID == "" || movie.MovieName == "" {
		response = dto.JsonResponse{Type: "error", Message: "Movie ID and Movie name are required"}
	} else {
		var lastInsertId int

		db := database.SetupDB()
		fmt.Println("M=CreateMovie, inserting Movie")

		err := db.QueryRow("INSERT INTO movies(movie_id, movie_name) VALUES ($1, $2) returning id", movie.MovieID, movie.MovieName).Scan(&lastInsertId)

		database.CheckError(err)

		movies := []database.Movie{movie}

		response = dto.JsonResponse{Type: "success", Data: movies}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie database.Movie

	params := mux.Vars(r)

	id := params["movieId"]

	movie.MovieID = id

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &movie)

	var response = dto.JsonResponse{}

	if movie.MovieID == "" || movie.MovieName == "" {
		response = dto.JsonResponse{Type: "Error", Message: "movieId and movieName are required"}
	} else {
		db := database.SetupDB()

		_, err := db.Exec("UPDATE movies SET movie_name = $1 WHERE movie_id = $2", movie.MovieName, movie.MovieID)

		database.CheckError(err)

		movie := []database.Movie{movie}

		response = dto.JsonResponse{Type: "success", Data: movie}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieId := params["movieId"]

	fmt.Printf("M=DeleteMovie, movieId=%s\n", movieId)

	var response = dto.JsonResponse{}

	if movieId == "" {
		response = dto.JsonResponse{Type: "error", Message: "movieId parameter is required!"}
	} else {
		db := database.SetupDB()

		fmt.Printf("M=DeleteMovie, deleting movie...")

		_, err := db.Exec("DELETE FROM movies WHERE movie_id = $1", movieId)

		database.CheckError(err)

		response = dto.JsonResponse{Type: "success", Message: "Movie deleted"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	db := database.SetupDB()

	fmt.Printf("M=DeleteMovies, deleting all movies...")

	_, err := db.Exec("DELETE FROM movies")

	database.CheckError(err)

	fmt.Printf("M=DeleteMovies, all movies deleted...")

	var response = dto.JsonResponse{Type: "success", Message: "Movies deleted"}

	json.NewEncoder(w).Encode(response)

}
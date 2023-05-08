package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type patient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}

var patients = []patient{
	{ID: "1", Name: "Lucky Skywalker", CPF: "737.802.360-44"},
	{ID: "2", Name: "Han Solo", CPF: "509.008.120-44"},
	{ID: "3", Name: "Djin Djarin", CPF: "328.843.870-48"},
}

// getAlbums responds with the list of all albums as JSON.
func getPatients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, patients)
}

/*
func homeHandler(w http.ResponseWriter, r *http.Request) {

}
*/

// postPatient adds an album from JSON received in the request body.
func postPatient(c *gin.Context) {
	var newPatient patient

	// Call BindJSON to bind the received JSON to
	// newPatient.
	if err := c.BindJSON(&newPatient); err != nil {
		return
	}

	// Add the new patient to the slice.
	patients = append(patients, newPatient)
	c.IndentedJSON(http.StatusCreated, newPatient)
}

func main() {
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.POST("/patients", postPatient)
	router.Run("localhost:8080")

	/*
		http.Handle("/", homeHandler)

		fmt.Printf("starting server at port 8080\n")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	*/
}

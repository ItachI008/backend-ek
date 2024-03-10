package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var rdb *redis.Client

func init() {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-18695.c1.asia-northeast1-1.gce.cloud.redislabs.com:18695",
		Password: "XhousSZSQ5Il5FCuirpNsqC5cyHhJdK4",
		DB:       0,
	})
}

// UserData struct to represent user data
type UserData struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
}

// setUser initializes user points or increments existing points
func setUser(w http.ResponseWriter, r *http.Request) {
	// Your setUser function implementation
}

// getUserPoints retrieves points for a given user
func getUserPoints(w http.ResponseWriter, r *http.Request) {
	// Your getUserPoints function implementation
}

// updateUserPoints increments user points
func updateUserPoints(w http.ResponseWriter, r *http.Request) {
	// Your updateUserPoints function implementation
}

// getAllUserPoints retrieves points for all users
func getAllUserPoints(w http.ResponseWriter, r *http.Request) {
	// Your getAllUserPoints function implementation
}

// getLeaderboard retrieves the leaderboard
func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Your getLeaderboard function implementation
}

// healthCheckHandler handles the health check endpoint
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	r := mux.NewRouter()

	// Define API endpoints
	r.HandleFunc("/api/user", setUser).Methods("POST")
	r.HandleFunc("/api/user/points", getUserPoints).Methods("GET")
	r.HandleFunc("/api/user/points", updateUserPoints).Methods("PUT")
	r.HandleFunc("/api/user/points/all", getAllUserPoints).Methods("GET")
	r.HandleFunc("/api/leaderboard", getLeaderboard).Methods("GET")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET") // Health check endpoint

	// Enable CORS
	c := cors.AllowAll()
	handler := c.Handler(r)

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if PORT environment variable is not set
	}

	// Start server
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

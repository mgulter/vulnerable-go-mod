package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogo/protobuf/proto"
	"github.com/ulikunitz/xz"
)

// Hardcoded secrets (Secret Detection bulgulari)
const (
	AWS_ACCESS_KEY_ID     = "AKIAIOSFODNN7EXAMPLE"
	AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	GITHUB_TOKEN          = "ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	DATABASE_PASSWORD     = "super_secret_password_123!"
	API_KEY               = "sk-proj-abcdefghijklmnopqrstuvwxyz123456"
)

func main() {
	_ = jwt.StandardClaims{}
	_ = proto.CompactTextString
	_, _ = xz.NewReader(nil)

	fmt.Println("vuln-go demo")
}

// SQL Injection vulnerability (SAST bulgusu)
func getUserByID(db *sql.DB, userID string) (*sql.Rows, error) {
	query := "SELECT * FROM users WHERE id = '" + userID + "'"
	return db.Query(query)
}

// Command Injection vulnerability (SAST bulgusu)
func executeCommand(userInput string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", userInput)
	return cmd.Output()
}

// Path Traversal vulnerability (SAST bulgusu)
func serveFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	http.ServeFile(w, r, "/var/www/"+filename)
}

// XSS vulnerability (SAST bulgusu)
func handleGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
}

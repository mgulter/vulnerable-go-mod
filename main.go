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
	// AWS Keys - yuksek entropi ile
	AWS_ACCESS_KEY_ID     = "AKIAQWERTY7X9K3MZPLN"
	AWS_SECRET_ACCESS_KEY = "8dKzT3pQrV5nYxW2mLjH9fGbNcXeRuAoIsSdFgHj"

	// GitHub Personal Access Token (classic format)
	GITHUB_TOKEN = "ghp_R8k2mNpQ4vX7wY9zL3jH5tB6cD1eF0gAaZxC"

	// Slack Bot Token
	SLACK_BOT_TOKEN = "xoxb-123456789012-1234567890123-AbCdEfGhIjKlMnOpQrStUvWx"

	// Stripe Secret Key
	STRIPE_SECRET_KEY = "sk_live_4eC39HqLyjWDarjtT1zdp7dc"

	// SendGrid API Key
	SENDGRID_API_KEY = "SG.nKqV8xRlTm2yFgHj4pWz.D3kL9mNpQrStUvWxYzAbCdEfGhIjKlMnOpQrStUvWx"

	// Private Key (RSA)
	PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA0Z3VS5JJcds3xfn/ygWyF8PbnGy0AHB7MvXmHq5R3k4Kp3Z
Px4qM+zNK8hM/LIpWExmFpSr/mtZ9Lqxk8A2aFXPwXkhvFqBn8JmVpqYd6L
-----END RSA PRIVATE KEY-----`

	// Database connection string with password
	DATABASE_URL = "postgresql://admin:Str0ngP@ssw0rd123!@db.example.com:5432/production"

	// JWT Secret
	JWT_SECRET = "mY$uP3rS3cr3tJWT$1gn1ngK3y!@#2024"
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

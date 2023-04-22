package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"

	_ "embed"
)

//go:embed redirect.html
var redirectHTML string

func isNicknameUnique(nickname string, credentialsPath string) (bool, error) {
	dirEntries, err := os.ReadDir(credentialsPath)
	if err != nil {
		return false, err
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		fc := strings.Split(filename, ".")
		fc = fc[:len(fc)-1]
		filenameClean := fc[0]

		if strings.EqualFold(filenameClean, nickname) {
			return false, nil
		}
	}
	return true, nil
}

// Request a token from the web, then returns the retrieved token.
func promptTokenInWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// Create a simple HTTP server to get the token somehow...
	ctx := context.Background()

	srv := &http.Server{
		Addr: ":8000",
	}

	authCodeChan := make(chan string)

	go func(srv *http.Server) {
		log.Println("running temporary server")

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			queryValues := r.URL.Query()
			code := queryValues.Get("code")
			w.Header().Add("Content-Type", "text/html")
			redirectHTMLPrepared := strings.Replace(redirectHTML, "%token%", "code", 1)
			fmt.Fprintf(w, redirectHTMLPrepared, code)
			if code != "" {
				authCodeChan <- code
				fmt.Println("Token retrieved: " + code)
			}
		})

		log.Println(srv.ListenAndServe())
		log.Println("killing temporary server")
	}(srv)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println("Opening URL: " + authURL)
	browser.OpenURL(authURL)

	authCode := <-authCodeChan
	srv.Shutdown(ctx)
	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// func main() {
// 	ctx := context.Background()
// 	b, err := os.ReadFile("credentials.json")
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}

// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, drive.DriveScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)

// 	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Drive client: %v", err)
// 	}

// 	r, err := srv.Files.List().Q("'185PVH9A45vlBBH7u0vrdgZXNvcrOFkLI' in parents").Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve files: %v", err)
// 	}
// 	fmt.Println("Files:")
// 	if len(r.Files) == 0 {
// 		fmt.Println("No files found.")
// 	} else {
// 		for _, i := range r.Files {
// 			fmt.Printf("%s (%s)\n", i.Name, i.Id)
// 		}
// 	}
// }

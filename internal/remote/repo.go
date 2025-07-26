package remote

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/go-github/v73/github"
	"golang.org/x/oauth2"
)

func CreateRemoteRepo(name, description string) string {
	if isEmptyOrWhitespace(name) {
		log.Fatal("Repo name cannot be empty")
	}
	token, err := getGHToken()
	if err != nil {
		log.Fatal("error while getting auth token", err)
	}

	// Create OAuth2 token source
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	// Create GitHub client
	client := github.NewClient(tc)

	// Create a new repository
	repo := &github.Repository{
		Name:        github.Ptr(name),
		Description: github.Ptr(description),
		Private:     github.Ptr(false),
		AutoInit:    github.Ptr(false),
	}

	ctx := context.Background()
	newRepo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Repository created: %s\n", *newRepo.HTMLURL)
	return *newRepo.HTMLURL
}

// gets the user auth token using github CLI
func getGHToken() (string, error) {
	cmd := exec.Command("gh", "auth", "token")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Trim whitespace/newlines from output
	return strings.TrimSpace(string(output)), nil
}

func isEmptyOrWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/UnderMountain96/ITEA_GO/http/dto"
)

func main() {
	fmt.Print("\nGitHub Username > ")

	var username string
	if _, err := fmt.Scan(&username); err != nil {
		panic(fmt.Sprintf("scan error: %s", err))
	}

	user, err := getUserInfo(username)
	if err != nil {
		panic(err)
	}

	printUserInfo(user)
}

func getUserInfo(username string) (*dto.GitHubUser, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("https://api.github.com/users/%s", username))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("getUserInfo: user with username %q NOT FOUND - Status Code %d", username, http.StatusNotFound)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("getUserInfo: read response body error: %w", err)
		}
		defer resp.Body.Close()

		user := new(dto.GitHubUser)

		if err := json.Unmarshal(bodyBytes, user); err != nil {
			return nil, fmt.Errorf("getUserInfo: unmarshal body to struct error: %w", err)
		}

		return user, nil
	}

	return nil, fmt.Errorf("getUserInfo: unexpected status code: %d", resp.StatusCode)
}

func printUserInfo(user *dto.GitHubUser) {
	fmt.Printf("ID:\t\t%d\n", user.Id)
	fmt.Printf("Name:\t\t%s\n", user.Name)
	fmt.Printf("Bio:\t\t%s\n", user.Bio)
	fmt.Printf("Created At:\t%d\n", user.CreatedAt.Year())
}

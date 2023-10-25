package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func processSequentialFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	users := scanFile(file)
	sequentialProcessing(users)
}

func processConcurrentFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	users := scanFile(file)
	concurrentrProcessing(users)
}

type userItem struct {
	ID        string
	Name      string
	LastName  string
	Email     string
	Phone     string
	FriendIDs []string
}

func scanFile(file *os.File) []*userItem {
	s := bufio.NewScanner(file)

	users := []*userItem{}

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		ids := strings.Split(lineArray[5], " ")
		ids = ids[1 : len(ids)-1]

		user := &userItem{
			ID:        lineArray[0],
			Name:      lineArray[1],
			LastName:  lineArray[2],
			Email:     lineArray[3],
			Phone:     lineArray[4],
			FriendIDs: ids,
		}

		users = append(users, user)
	}

	return users
}

func sequentialProcessing(users []*userItem) {
	visited := make(map[string]bool)

	for _, user := range users {
		if !visited[user.ID] {
			visited[user.ID] = true

			sendSmsNotification(user)

			for _, friendID := range user.FriendIDs {
				friend, err := findUserByID(friendID, users)
				if err != nil {
					log.Printf("Error.findUserByID: %v\n", err)

					continue
				}

				if !visited[friend.ID] {
					visited[friend.ID] = true

					sendSmsNotification(friend)
				}
			}
		}
	}
}

func concurrentrProcessing(users []*userItem) {
	usersCh := make(chan []*userItem)
	unvisitedUsers := make(chan *userItem)

	go func() {
		usersCh <- users
	}()

	// usersCh = 1
	initializeWorkers(unvisitedUsers, usersCh, users)
	// usersCh = 1,  wait unvisitedUsers set value
	processUser(unvisitedUsers, usersCh, len(users))
}

func processUser(unvisitedUsers chan<- *userItem, usersCh chan []*userItem, size int) {
	visitedUsers := make(map[string]bool)

	count := 0
	for users := range usersCh {
		for _, user := range users {
			if !visitedUsers[user.ID] {
				visitedUsers[user.ID] = true
				count++

				if count >= size {
					close(usersCh)
				}

				unvisitedUsers <- user
			}
		}
	}
}

func initializeWorkers(unvisitedUsers <-chan *userItem, usersCh chan []*userItem, users []*userItem) {
	const maxGoroutines = 10

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for user := range unvisitedUsers {
				sendSmsNotification(user)

				go func(user *userItem) {
					friendIDs := user.FriendIDs
					friends := []*userItem{}

					for _, friendID := range friendIDs {
						friend, err := findUserByID(friendID, users)
						if err != nil {
							log.Printf("Error.findUserByID: %v\n", err)

							continue
						}

						friends = append(friends, friend)
					}

					_, ok := <-usersCh
					if ok {
						usersCh <- friends
					}
				}(user)
			}
		}()
	}
}

func findUserByID(userID string, users []*userItem) (*userItem, error) {
	for _, user := range users {
		if user.ID == userID {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found with ID %s", userID)
}

func sendSmsNotification(_ *userItem) {
	time.Sleep(1 * time.Millisecond)
	// fmt.Printf("Sending sms notification to %s\n", user.Phone)
}

package services

import "thefinisher-goapp/internal/models"

var users = []models.User{
    {ID: "1", Name: "Alice"},
    {ID: "2", Name: "Bob"},
}

func GetAllUsers() string {
    // In a real app, you'd return JSON here (marshal the users)
    return "List of all users"
}

func GetUser(id string) string {
    for _, user := range users {
        if user.ID == id {
            return "User: " + user.Name
        }
    }
    return "User not found"
}


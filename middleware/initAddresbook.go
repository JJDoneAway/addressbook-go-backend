package middleware

import (
	"bufio"
	"embed"
	"fmt"
	"strings"

	"github.com/JJDoneAway/addressbook/models"
)

//go:embed DemoUser.txt
var demoUsers embed.FS

func AddDummies() {
	fmt.Println("Inserting some sample users...")
	fs, err := demoUsers.Open("DemoUser.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fs)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		name := strings.Split(fileScanner.Text(), " ")
		(&models.User{FirstName: name[0], LastName: name[1]}).InsertUser()
	}

	fs.Close()
}

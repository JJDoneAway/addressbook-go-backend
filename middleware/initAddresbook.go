package middleware

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
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
		user := models.Address{FirstName: name[0], LastName: name[1], Email: fmt.Sprintf("%v@%v.de", name[0], name[1]), Phone: "+49" + strconv.Itoa(int(models.NextID()))}
		log.Default().Print(user)
		if err := user.InsertAddress(); err != nil {
			panic(user)
		}

	}

	fs.Close()
}

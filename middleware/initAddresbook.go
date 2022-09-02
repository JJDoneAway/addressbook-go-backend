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

//go:embed DemoContact.txt
var demoUsers embed.FS

func AddDummies() {
	fmt.Println("Inserting some sample addresses...")
	fs, err := demoUsers.Open("DemoContact.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fs)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		name := strings.Split(fileScanner.Text(), " ")
		contact := models.Contact{FirstName: name[0], LastName: name[1], Email: fmt.Sprintf("%v@%v.de", name[0], name[1]), Phone: "+49" + strconv.Itoa(int(models.NextID()))}
		log.Default().Print(contact)
		if err := contact.InsertContact(); err != nil {
			panic(contact)
		}
	}

	fs.Close()
}

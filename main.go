package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Sets the connection timeout to 20 seconds
var httpClient = &http.Client{Timeout: 20 * time.Second}

/**
* Checks if a slice or an array contains an item
* ( there is no build-in function in Go )
 */
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	// Array that holds unique domains
	domainList := []string{}
	// Inits reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type domain: ")
	userInput, _ := reader.ReadString('\n')
	// Converts CRLF to LF
	userInput = strings.Replace(userInput, "\n", "", -1)

	res, err := httpClient.Get("https://crt.sh/?dNSName=" + userInput)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	// The pattern finds domains in the HTML code
	pattern := "<TD>(.+)</TD>"

	body, err := ioutil.ReadAll(res.Body)

	// Splits HTML body string to an array of strings
	rows := strings.Split(string(body), "\n")
	// Sets pattern for domains tag
	ms := regexp.MustCompile(`<TD>(.+)</TD>`)

	for _, row := range rows {
		match, _ := regexp.MatchString(pattern, row)
		// If the tag found and string does not contain style attribute
		if match && !strings.Contains(row, "style=") {
			// Splits the string into an array of strings that holds domains
			domains := strings.Split(ms.ReplaceAllString(row, "$1"), "<BR>")

			for _, domain := range domains {
				// Removes wildcard (*.), trim whitespace and lowercase
				domain := strings.ToLower(strings.TrimSpace(strings.Replace(domain, "*.", "", -1)))

				// Checks if the domain is already added
				if !contains(domainList, domain) {
					// Prints domain
					fmt.Println(">> " + domain)
					// Appends domain to domain list
					domainList = append(domainList, domain)
				}

			}
		}
	}

}

package tratamento

import (
//	"strings"
	"regexp"
	"log"
)

//RemoveCaracteres limpa os caracteres da coluna
func RemoveCaracteres(data string) string {
    // data = strings.Replace(data, ".", "", -1)
    // data = strings.Replace(data, "-", "", -1)
    // data = strings.Replace(data, "/", "", -1)

    reg, err := regexp.Compile("[^a-zA-Z0-9]+")
    if err != nil {
        log.Fatal(err)
    }
    processedString := reg.ReplaceAllString(data, "")

    return processedString
}

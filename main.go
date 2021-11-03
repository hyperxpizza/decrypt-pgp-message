package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

var privateKeyPathPtr = flag.String("key", "", "path to the file with the private key")
var passwordPtr = flag.String("password", "", "passphrase required for decryption")
var messageFilePtr = flag.String("message", "", "path to file containing the message")

func usage() {
	fmt.Println("Usage:")
	fmt.Println("	-key=<path to your private pgp key file>")
	fmt.Println("	-message=<path to message encrypted with your public pgp key>")
	fmt.Println("	-password=<password for the armored message>")
	fmt.Println("All flags must be set!")
	fmt.Println("Correct usage:")
	fmt.Println("	decryptpgpmessage -key=<path to your private pgp key file> -message=<path to message encrypted with your public pgp key> -password=<password for the armored message>")
}

func main() {
	flag.Parse()

	//validate args
	if privateKeyPathPtr == nil || *privateKeyPathPtr == "" {
		usage()
		os.Exit(1)
	}

	if passwordPtr == nil || *passwordPtr == "" {
		usage()
		os.Exit(1)
	}

	if messageFilePtr == nil || *messageFilePtr == "" {
		usage()
		os.Exit(1)
	}

	//load private key
	pKey, err := loadFileIntoBytes(*privateKeyPathPtr)
	if err != nil {
		log.Fatalf("Could not load private key from file: %v \n", err)
	}

	//load message
	encryptedMessage, err := loadFileIntoBytes(*messageFilePtr)
	if err != nil {
		log.Fatalf("Could not load decrypted message from file: %v \n", err)
	}

	decryptedMessage, err := decrypt(string(pKey), []byte(*passwordPtr), string(encryptedMessage))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decryptedMessage)
}

func decrypt(pKey string, pass []byte, message string) (string, error) {

	decrypted, err := helper.DecryptMessageArmored(pKey, pass, message)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func loadFileIntoBytes(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

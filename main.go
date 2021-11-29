package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mac2000/govatar/grav"
	"log"
	"os"
	_ "embed"
)

//go:embed help.txt
var help string

func main() {
	username := flag.String("username", "", "gravatar username, usually it is email")
	password := flag.String("password", "", "gravatar password")
	doListAddresses := flag.Bool("addresses", false, "get a list of addresses for this account")
	doDeleteUserImage := flag.Bool("delete", false, "remove a userimage from the account and any email addresses with which it is associated")
	doExists := flag.Bool("exists", false, "check whether a hash has a gravatar")
	doRemoveImage := flag.Bool("remove", false, "remove the userimage associated with one or more email addresses")
	doSave := flag.Bool("save", false, "save image from url or path as a userimage for this account")
	doTest := flag.Bool("test", false, "a test function")
	doListUserImages := flag.Bool("userimages", false, "return an array of userimages for this account")
	doUseUserImage := flag.Bool("use", false, "use a userimage as a gravatar for one of more addresses on this account")
	doSet := flag.Bool("set", false, "uploads give image to gravatar, set it as main, remove all other")

	email := flag.String("email", "", "email to check")
	id := flag.String("id", "", "user image identifier")
	url := flag.String("url", "", "image url")
	path := flag.String("path", "", "image path")

	printJson := flag.Bool("json", false, "output json")
	printHelp := flag.Bool("help", false, "output help")
	flag.Parse()

	if *username == "" && !*printHelp {
		log.Fatal("username missing")
	}
	if *password == "" && !*printHelp {
		log.Fatal("password missing")
	}

	g := grav.NewGravatarClient(*username, *password)

	if *doListAddresses {
		addresses, err := g.Addresses()
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			b, err := json.MarshalIndent(addresses, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			for _, address := range addresses {
				fmt.Printf("- email:     %s\n", address.Email)
				fmt.Printf("  rating:    %v\n", address.Rating)
				fmt.Printf("  userimage: %s\n", address.UserImage)
				fmt.Printf("  url:       %s\n", address.UserImageURL)
			}
		}
	} else if *doDeleteUserImage {
		if *id == "" {
			log.Fatal("user image id argument is missing")
		}
		err := g.DeleteUserimage(*id)
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			fmt.Println("{\"success\": true}")
		} else {
			fmt.Println("true")
		}
	} else if *doExists {
		if *email == "" {
			log.Fatal("email argument is missing")
		}
		exists, err := g.Exists(grav.EmailHash(*email))
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			fmt.Printf("{\"exists\": %v}\n", exists)
		} else {
			fmt.Println(exists)
		}
	} else if *doRemoveImage {
		if *email == "" {
			log.Fatal("email argument is missing")
		}
		err := g.RemoveImage(*email)
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			fmt.Println("{\"success\": true}")
		} else {
			fmt.Println("true")
		}
	} else if *doSave {
		if *url == "" && *path == "" {
			log.Fatal("url or path is required")
		}
		if *url != "" && *path != "" {
			log.Fatal("url or path is required")
		}
		if *url != "" {
			userImage, err := g.SaveUrl(grav.RatingG, *url)
			if err != nil {
				log.Fatal(err)
			}
			if *printJson {
				fmt.Printf("{\"id\": %v}\n", userImage)
			} else {
				fmt.Println(userImage)
			}
		} else if *path != "" {
			b, err := os.ReadFile(*path)
			if err != nil {
				log.Fatal(err)
			}
			b64 := base64.StdEncoding.EncodeToString(b)
			userImage, err := g.SaveData(grav.RatingG, b64)
			if err != nil {
				log.Fatal(err)
			}
			if *printJson {
				fmt.Printf("{\"id\": %v}\n", userImage)
			} else {
				fmt.Println(userImage)
			}
		}
	} else if *doTest {
		id, err := g.Test()
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			fmt.Printf("{\"id\": %v}\n", id)
		} else {
			fmt.Println(id)
		}
	} else if *doListUserImages {
		userImages, err := g.UserImages()
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			b, err := json.MarshalIndent(userImages, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			for _, image := range userImages {
				fmt.Printf("- name:   %s\n", image.Name)
				fmt.Printf("  rating: %v\n", image.Rating)
				fmt.Printf("  url:    %s\n", image.URL)
			}
		}
	} else if *doUseUserImage {
		if *id == "" {
			log.Fatal("user image id argument is missing")
		}
		if *email == "" {
			log.Fatal("email argument is missing")
		}
		err := g.UseUserImage(*id, *email)
		if err != nil {
			log.Fatal(err)
		}
		if *printJson {
			fmt.Println("{\"success\": true}")
		} else {
			fmt.Println("true")
		}
	} else if *doSet {
		if *path == "" {
			log.Fatal("path argument missing")
		}
		// retrieve all avatars
		userImages, err := g.UserImages()
		if err != nil {
			log.Fatal(err)
		}
		// save new avatar
		b, err := os.ReadFile(*path)
		if err != nil {
			log.Fatal(err)
		}
		b64 := base64.StdEncoding.EncodeToString(b)
		userImage, err := g.SaveData(grav.RatingG, b64)
		if err != nil {
			log.Fatal(err)
		}
		// use new avatar
		err =  g.UseUserImage(userImage, *username)
		if err != nil {
			log.Fatal(err)
		}
		// delete old avatars
		for _, img := range userImages {
			err = g.DeleteUserimage(img.Name)
			if err != nil {
				log.Fatal(err)
			}
		}
		if *printJson {
			d := struct {
				Saved string `json:"saved"`
				Removed []grav.UserImage `json:"removed"`
			}{
				Saved: userImage,
				Removed: userImages,
			}
			b, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			fmt.Printf("saved: %s\n", userImage)
			fmt.Printf("removed: %d\n", len(userImages))
		}
	} else if *printHelp {
		log.Println(help)
	} else {
		log.Println(help)
	}
}

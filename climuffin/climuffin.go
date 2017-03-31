// This file is part of climuffin.
//
// Foobar is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Foobar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with climuffin.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/matthewrsj/daggit"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func readActivitiesFromJSON(jsonfile string) []daggit.Activity {
	raw, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a := make([]daggit.Activity, 0)
	json.Unmarshal(raw, &a)
	return a
}

func mainLoop(db *gorm.DB, commands []string) {

	var cmd string
	for {
		fmt.Printf(" >> ")
		if _, err := fmt.Scanf("%s", &cmd); err != nil {
			fmt.Println(err)
			continue
		}
		if err := verifyCmd(cmd, commands); err != nil {
			fmt.Println(err)
			continue
		}
		runCommand(cmd)
	}

}

func runCommand(cmd string) {
	switch cmd {
	case "ra":
		printAllActivities()
	case "ri":
		id := getID()
		printActivityByID(id)
	case "h":
		help()
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("%s not implemented", cmd)
	}
}

func help() {
	helpMsg := []string{
		"Commands",
		"ra: read all activities",
		"ri: read activity by id",
		"h:  display this message",
		"q:  quit the program",
	}
	for _, s := range helpMsg {
		fmt.Println(s)
	}
}

func verifyCmd(cmd string, commands []string) error {
	for _, c := range commands {
		if cmd == c {
			return nil
		}
	}

	return fmt.Errorf("unknown command %s", cmd)
}

func getID() uint {
	var id uint
	fmt.Printf("ID: ")
	if _, err := fmt.Scanf("%d", &id); err != nil {
		fmt.Println(err)
	}
	return id
}

func main() {
	commands := []string{"ra", "h", "q", "ri"}
	db = daggit.OpenDB()
	defer daggit.CloseDB(db)
	daggit.SetupDB(db)
	create := flag.String("create", "", "name of file with activities to create")

	flag.Parse()

	if len(*create) != 0 {
		activities := readActivitiesFromJSON(*create)
		for _, a := range activities {
			if err := daggit.CreateActivity(db, a); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	mainLoop(db, commands)

}

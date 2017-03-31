// This file is part of daggit.
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
// along with daggit.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"time"
)

func main() {
	db := OpenDB()
	defer CloseDB(db)
	SetupDB(db)

	// THIS IS TEST CODE ////////////////////////////////////////
	activity := Activity{
		Name:  "hello",
		Start: time.Now(),
		End:   time.Now().Add(time.Duration(1) * time.Hour),
	}

	activity1 := Activity{
		Name:  "hello",
		Start: time.Now(),
		End:   time.Now().Add(time.Duration(1) * time.Hour),
	}

	if err := CreateActivity(db, activity); err != nil {
		fmt.Println("Error!", err)
	}
	if err := CreateActivity(db, activity1); err != nil {
		fmt.Println("Error!", err)
	}

	acts := ReadAllActivities(db)
	if len(acts) == 0 {
		fmt.Println("No activities in database")
		return
	}
	fmt.Println(time.Now())

	for i, act := range acts {
		fmt.Printf("Activity %d\n\n", i)
		fmt.Println("Name   : ", act.Name)
		fmt.Println("ID     : ", act.ID)
		fmt.Println("Created: ", act.CreatedAt)
		fmt.Println("Start  : ", act.Start)
		fmt.Println("End    : ", act.End)
		fmt.Println("\n\n")
	}
	// THIS IS TEST CODE ^^ /////////////////////////////////////

}

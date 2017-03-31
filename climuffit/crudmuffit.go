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
	"fmt"

	"github.com/matthewrsj/daggit"
)

func printAllActivities() {
	activities := daggit.ReadAllActivities(db)
	for _, a := range activities {
		fmt.Println(a.Name)
		fmt.Println(a.Start)
		fmt.Println(a.End)
		fmt.Println()
	}
}

func printActivityByID(id uint) {
	activity := daggit.ReadActivityByID(db, id)
	if (daggit.Activity{}) == activity {
		fmt.Printf("No activity found with id %d\n", id)
		return
	}
	fmt.Println(activity.Name)
	fmt.Println(activity.Start)
	fmt.Println(activity.End)
	fmt.Println()
}

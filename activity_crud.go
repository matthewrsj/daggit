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

package daggit

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Create Activity if there are no time collisions
func CreateActivity(db *gorm.DB, activity Activity) error {
	if isTimeCollision(db, activity) {
		return fmt.Errorf("Activity overlaps with another activity")
	}

	if err := db.Create(&activity).Error; err != nil {
		return err
	}

	return nil
}

// Check for time collision
func isTimeCollision(db *gorm.DB, a Activity) bool {
	activities := ReadAllActivities(db)
	for _, estAct := range activities {
		if a.Start.Before(estAct.End) && a.End.After(estAct.Start) {
			return true
		}
	}

	return false
}

// Read Activity by ID
func ReadActivityByID(db *gorm.DB, id uint) Activity {
	var activity Activity
	db.Find(&activity, id)
	return activity
}

// Read all Activities
func ReadAllActivities(db *gorm.DB) []Activity {
	var activities []Activity
	db.Find(&activities)
	return activities
}

// Update Activity all fields by ID
func UpdateActivityAllByID(db *gorm.DB, id uint, newActivity Activity) error {
	activity := ReadActivityByID(db, id)
	if activityIsZero(activity) {
		return fmt.Errorf("Activity ID %d does not exist", id)
	}

	activity = newActivity
	if err := db.Save(&activity).Error; err != nil {
		return err
	}

	return nil
}

// Delete Activity by ID
func DeleteActivityByID(db *gorm.DB, id uint) error {
	activity := ReadActivityByID(db, id)
	if activityIsZero(activity) {
		return fmt.Errorf("Activity ID %d does not exist", id)
	}

	if err := db.Delete(&activity).Error; err != nil {
		return err
	}
	return nil
}

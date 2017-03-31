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
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

package daggit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Open gorm database
func OpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "daggit.db")
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

// Close gorm database, should be deferred right after opening
func CloseDB(db *gorm.DB) {
	db.Close()
}

// Set up the gorm database
func SetupDB(db *gorm.DB) {
	db.AutoMigrate(&Activity{})
}

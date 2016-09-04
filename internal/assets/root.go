// +build dev

package assets

import "log"
import "os"
import "path/filepath"

var rootDir string

func init() {
	home := os.Getenv("ZKM_HOME")
	if home == "" {
		home, _ = os.Getwd()
	}
	log.Printf("assets/init: using home directory: %s", home)

	rootDir = filepath.Join(home, "internal", "assets")
	log.Printf("assets/init: using asset root directory: %s", rootDir)
}

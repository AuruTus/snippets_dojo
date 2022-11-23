package cfmt

import (
	"log"
	"os"
)

// TODO
// 1. a logger init entrance
// 2. customizable format including file line
//

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

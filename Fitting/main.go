package main

import (
 db "./database"
)

func main() {
 defer db.Pdb.Close()
 router := initRouter()
 router.Run(":8000")
}

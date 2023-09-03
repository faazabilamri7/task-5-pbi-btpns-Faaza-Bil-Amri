package main

import (
    "path/to/your/router" // Import package yang berisi SetupRouter
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080") // Port yang digunakan untuk menjalankan aplikasi
}

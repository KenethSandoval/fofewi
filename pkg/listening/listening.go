package listening

import (
	"fmt"
	"log"
	"net/http"
)

func ListePrinServer(hs *http.Server) {
	fmt.Print(`
─────▄───▄
─▄█▄─█▀█▀█─▄█▄
▀▀████▄█▄████▀▀
─────▀█▀█▀
		_/ ____\  ____  _/ ____\  ____  __  _  __|__|
		\   __\  /  _ \ \   __\ _/ __ \ \ \/ \/ /|  |
		 |  |   (  <_> ) |  |   \  ___/  \     / |  |
 		 |__|    \____/  |__|    \____>   \/\_/  |__|

`)
	log.Print("\nListening..." + hs.Addr)
}

// TODO!
// Add comments to this monstrosity

package custommiddleware

import (
	"log"
	"net/http"
	"time"

	"github.com/dixxe/personal-website/iternal"
)

var alreadyBanned map[string]int = make(map[string]int)
var goroutineRequested map[string]int = make(map[string]int)

func DDOSProtection(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responseIp := r.RemoteAddr

		iternal.RecentConnections[responseIp] += 1

		if iternal.RecentConnections[responseIp] >= 20 {
			if _, ok := goroutineRequested[responseIp]; ok {
				log.Println("Blocked request.")
				iternal.PossibleBotsAmount++
				return
			}
			go func() {
				goroutineRequested[responseIp] = 1
				time.Sleep(time.Second * 5)
				r.Close = true
				iternal.BotsAndTheirTimeout[responseIp] = iternal.RecentConnections[responseIp]
				banBot(responseIp)
			}()
		}

		h.ServeHTTP(w, r)
	})
}

func banBot(bot string) {
	if _, ok := alreadyBanned[bot]; ok {
		return
	}
	alreadyBanned[bot] = 1
	timeToBan := iternal.BotsAndTheirTimeout[bot]
	go func() {
		log.Printf("Banning %v for %v", bot, time.Second*time.Duration(timeToBan))
		time.Sleep(time.Second * time.Duration(timeToBan))
		delete(alreadyBanned, bot)
		delete(iternal.BotsAndTheirTimeout, bot)
		delete(goroutineRequested, bot)
		iternal.RecentConnections[bot] = 0
		log.Printf("Pardon %v, time's up!", bot)
	}()

}

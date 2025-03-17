// TODO!
// Add comments to this monstrosity

package custommiddleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dixxe/personal-website/iternal"
)

var alreadyBanned map[string]int = make(map[string]int)

func DDOSProtection(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responseIp := r.RemoteAddr

		iternal.RecentConnections[responseIp] += 1

		if iternal.RecentConnections[responseIp] >= 20 {
			go func() {
				time.Sleep(time.Second * 5)
				r.Close = true
				iternal.PossibleBotsAmount++
				iternal.BotsAndTheirTimeout[responseIp] = iternal.RecentConnections[responseIp]
				banBot(responseIp)
			}()
			return
		}

		fmt.Println(iternal.RecentConnections)

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
		iternal.RecentConnections[bot] = 0
		log.Printf("Pardon %v, time's up!", bot)
	}()

}

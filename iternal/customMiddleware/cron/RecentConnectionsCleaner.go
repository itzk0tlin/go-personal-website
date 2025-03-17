package cron

import (
	"log"
	"time"

	"github.com/dixxe/personal-website/iternal"
)

// Run with gorutines to clean recent connections time-to-time.
// Simple as that
func CleanRecentConnections() {
	iternal.CronJobsRunning++
	log.Println("Clean recent connections cron job started!")
	for {
		time.Sleep(time.Second * 10)
		for k := range iternal.RecentConnections {
			if _, ok := iternal.BotsAndTheirTimeout[k]; ok {
				continue
			}
			delete(iternal.RecentConnections, k)
		}

	}
}

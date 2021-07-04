package cron

import (
	"fmt"

	"github.com/robfig/cron"

	"Video_fetcher"
)

//ScheduleCron -this will called FetchVideoData after every 30 minutes and update in-memory cache of videos for
//predefined query in constants.
func ScheduleCron() {
	scheduler := cron.New()
	_, err := scheduler.AddFunc("@every 30m", func() {
		Video_fetcher.FetchVideoData(Video_fetcher.YtService, []string{"snippet"}, false)
		fmt.Println("cron called")
	})
	if err != nil {
		fmt.Println("Unable to schedule cron")
		return
	}
	scheduler.Start()
}

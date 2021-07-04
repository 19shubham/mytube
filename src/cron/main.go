package cron

import (
	"fmt"

	"github.com/robfig/cron"
)

func ScheduleCron() {
	scheduler := cron.New()
	_, err := scheduler.AddFunc("@every 30m", func() {
		//youtube_fetch.FetchYoutubeApi()
	})
	if err != nil {
		fmt.Println("Unable to schedule cron")
		return 
	}
	scheduler.Start()
}

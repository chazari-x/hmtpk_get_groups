package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chazari-x/hmtpk_get_groups/file"
	"github.com/chazari-x/hmtpk_get_groups/schedule"
	"github.com/chazari-x/hmtpk_get_groups/selenium"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "get",
		Long:  "get",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetReportCaller(true)
			log.SetFormatter(&log.TextFormatter{
				FullTimestamp:             true,
				TimestampFormat:           "2006-01-02 15:04:05",
				ForceColors:               true,
				PadLevelText:              true,
				EnvironmentOverrideColors: true,
			})

			log.SetLevel(log.TraceLevel)

			log.Trace("get starting..")
			defer log.Trace("get stopped")

			newSelenium, s, err := selenium.NewSelenium()
			if err != nil {
				log.Error(err)
				return
			}

			defer func() {
				_ = s.Quit()
			}()

			sch := schedule.NewSchedule(newSelenium)

			err = file.WriteToFile(sch.GetTeachers())
			if err != nil {
				log.Error(err)
			}

			err = file.WriteToFile(sch.GetGroups())
			if err != nil {
				log.Error(err)
			}

			log.Infof("Press CTRL-C to exit.")
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
			<-sc
		},
	})
}

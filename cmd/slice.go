package cmd

import (
	"net/url"
	"os"

	"github.com/remeh/sizedwaitgroup"
	"github.com/sensepost/gowitness/lib"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CreateLogger(debug, disableLogging bool) *zerolog.Logger {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02 Jan 2006 15:04:05"})
	if options.Debug {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
		log.Logger = log.With().Caller().Logger()
		log.Debug().Msg("debug logging enabed")
	} else {
		log.Logger = log.Logger.Level(zerolog.InfoLevel)
	}
	if options.DisableLogging {
		log.Logger = log.Logger.Level(zerolog.Disabled)
	}

	return &log.Logger
}

func GoWitnessess(urls []string, threads int) {

	db, err := db.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get a db handle")
	}

	log.Debug().Int("threads", options.Threads).Msg("thread count to use with goroutines")
	swg := sizedwaitgroup.New(options.Threads)

	if err = options.PrepareScreenshotPath(); err != nil {
		log.Fatal().Err(err).Msg("failed to prepare the screenshot path")
	}

	// parse headers
	chrm.PrepareHeaderMap()

	logger := CreateLogger(false, false)

	for _, uniformResourceLocator := range urls {

		for _, u := range getUrls(uniformResourceLocator) {
			swg.Add()

			log.Debug().Str("url", u.String()).Msg("queueing goroutine for url")
			go func(url *url.URL) {
				defer swg.Done()

				p := &lib.Processor{
					Logger:         logger,
					Db:             db,
					Chrome:         chrm,
					URL:            url,
					ScreenshotPath: options.ScreenshotPath,
				}

				if err := p.Gowitness(); err != nil {
					log.Error().Err(err).Str("url", url.String()).Msg("failed to witness url")
				}
			}(u)
		}
	}

	swg.Wait()
	log.Info().Msg("processing complete")
}

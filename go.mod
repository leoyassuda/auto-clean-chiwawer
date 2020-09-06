module hello

go 1.15

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/heroku/x v0.0.25
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
)

replace example.com/greetings => ../greetings

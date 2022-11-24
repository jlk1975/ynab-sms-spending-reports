# SMS Spending Stats Reporter for YNAB, Written In Golang

High level idea is to use a simple golang program + deploy it using github actions scp-deployer to a Linux VM where it will run in a cron job to download YNAB transactions using YNAB’s API for a single budget, then the golang program will do some number crunching to get interesting spending stats on the given budget, then the script will connect to AWS “SNS” service to send SMS text message to recipients containing the spending report/numbers that the script produced.

- Once the basic app is working, let's see if we can wrap some observability around it by adding some Prometheus and Grafana timeseries YNAB reports into the mix, as well as some basic logging and alerting on if the script ever throws any error messages when it runs, like how else would we know there was a problem unless we did not get our SMS YNAB report?

## Project Characteristics
- Golang
- YNAB API
- Something to send SMS (AWS, Twilio, Other?) 
- CircleCI
- Prometheus/Grafana (How can the app export Prometheus endponts?)
- Splunk if Possible (free?)
- Will the golang app just run in a cron job , or run all the time in a docker container?

## to run
- go run sms.go

- Next step: read about github actions: https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions

- Report Ideas
- Daily
  - You spent x $ yesterday on category z, which was y $ more/less than the day before.
  - You've spent X% of your allocated budget this week for category Z. 
  - You are x$ ahead/behind spending for this week for category Z.
  - You have spent X% more/less this week for category Z compared to spending by this point in the week last week.
  - Your cumulative average daily spending for categorey Z has increased/decreased by X% since the very fist transaction for categorey Z.
- Monthly
- Quarterly
- Yearly

note: try this out https://github.com/marketplace/actions/scp-deployer

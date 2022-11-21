# SMS Spending Stats Reporter for YNAB, Written In Python

High level idea is to use a simple python script + deploy it using github actions scp-deployer to a Linux VM where it will run in a cron job to download YNAB transactions using YNAB’s API for a single budget, then the python script will do some number crunching to get interesting spending stats on the given budget, then the script will connect to AWS “SNS” service to send SMS text message to recipients containing the spending report/numbers that the script produced.

note: try this out https://github.com/marketplace/actions/scp-deployer

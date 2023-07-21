logfile=log.txt

echo press CTRL+C to stop app
echo log file: $logfile

script -q /dev/null -c ./go-xgml > $logfile
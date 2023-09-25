logfile=log.txt

echo press CTRL+C to stop app
echo log file: $logfile

script -q /dev/null -c ./image_packages > $logfile
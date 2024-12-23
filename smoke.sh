# build application
go build main

# run for 30 seconds
./main &
sleep 30
kill "$!"

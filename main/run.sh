#!/bin/sh


n=1

/main

# continue until $n equals 5
while [ 1 ]
do
	echo "Welcome $n times."
        curl http://google.com 	
	curl http://example.com 
	curl https://amazon.com
	curl https://stackoverflow.com
	curl http://golang.org
	n=$(( n+1 ))	 # increments $n
done


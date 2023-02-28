#!/bin/bash

#you must change this to your own email address
contactEmail="ValidContactEamilAddress";

if [[ "$1" == "" || "$1" == "-h" || "$1" == "--help" ]]; then
	echo "GetIPIntel.net query script takes a filename as a parameter or an IP address
	Example: bash getipintel.sh IPList
		 bash getipintel.sh 127.0.0.1
	The file should have each IP on a new line.
	IP and score are separated by a comma.
	Please add a valid contact email and read the comments before using.
	Full API documentation: https://getipintel.net/free-proxy-vpn-tor-detection-api/"

	exit 0;
fi


if [ -f "$1" ]; then
	for f in `cat $1`; do
		echo -n "$f,";
		curl -s "https://check.getipintel.net/check.php?ip=$f&contact=$contactEmail";
		echo "";
		#uncomment the sleep if you're using the free API to adhere to the 15 queries per minute limit
		#sleep 4;
	done
else 
	echo -n "$1,";
	curl -s "https://check.getipintel.net/check.php?ip=$1&contact=$contactEmail";
	echo "";
fi

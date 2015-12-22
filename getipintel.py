#!/usr/bin/env python
import requests
import sys
def checkIP(ip):
	#change this variable to your email address
	contactEmail="YourEmailAddress"
	#if probability from getIPIntel is grater than this value, return 1
	maxProbability=0.99
	#if you wish to use flags or json format, edit the request below
	result = requests.get("http://check.getipintel.net/check.php?ip="+ip+"&contact="+contactEmail)
	if (result.status_code != 200) or (int(result.content) < 0):
		sys.stderr.write("An error occured while querying GetIPIntel")
	if (int(result.content) > maxProbability):
		return 1;
	else:
		return 0;


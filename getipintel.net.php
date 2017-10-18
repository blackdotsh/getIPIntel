<?php
/*
* A PHP function that interacts with http://getIPIntel.net to look up an IP address
* returns TRUE if the IP returns a value greater than $banOnProbability,
* FALSE otherwise, including errors
* HTTP error codes are NOT explicitly implemented here
* This should be used as a guide, be sure to edit and test it before using it in production

* MIT License
*/


//function requires curl
function checkProxy($ip){
		$contactEmail="someValidEmailAddress"; //you must change this to your own email address
		$timeout=5; //by default, wait no longer than 5 secs for a response
		$banOnProbability=0.99; //if getIPIntel returns a value higher than this, function returns true, set to 0.99 by default
		
		//init and set cURL options
		$ch = curl_init();
		curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
		curl_setopt($ch, CURLOPT_TIMEOUT, $timeout);

		//if you're using custom flags (like flags=m), change the URL below
		curl_setopt($ch, CURLOPT_URL, "http://check.getipintel.net/check.php?ip=$ip&contact=$contactEmail");
		$response=curl_exec($ch);
		
		curl_close($ch);
		
		
		if ($response > $banOnProbability) {
				return true;
		} else {
			if ($response < 0 || strcmp($response, "") == 0 ) {
				//The server returned an error, you might want to do something
				//like write to a log file or email yourself
				//This could be true due to an invalid input or you've exceeded
				//the number of allowed queries. Figure out why this is happening
				//because you aren't protected by the system anymore
				//Leaving this section blank is dangerous because you assume
				//that you're still protected, which is incorrect
				//and you might think GetIPIntel isn't accurate anymore
				//which is also incorrect.

				//failure to implement error handling is bad for the both of us

			}
				return false;
		}
}


$ip=$_SERVER['REMOTE_ADDR'];

if (checkProxy($ip)) {
	/* A proxy has been detected based on your criteria
	 * Do whatever you want to here
	 */
	echo "It appears you're a Proxy / VPN / bad IP, please contact [put something here] for more information <br />";
}

?>

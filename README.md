# IP Intelligence is a free tool that attempts to determine how likely an IP address is a proxy / VPN / bad IP using mathematical and modern computing techniques

* Protect your site from automated XSS / SQL Injection / Brute Force / Crawlers that steal your content
* Serve traffic / content to real users, not bots, which reduces server load
* Stop bots from scraping your content or bots spamming your website
* Prevent trolls / people that are trying to bypass a ban
* Greatly reduce fraud on e-commerce sites (anti-fraud)
* Since the system returns a _real value_, you can customize the level of protection for a particular time frame and adjust accordingly.
* Use it with a combination of another fraud prevention service to make it even better. Some fraud prevention services do not explicitly look for proxy / VPN IPs. 

The system is serving millions of API requests a week and growing as more people find it useful in protecting their online infrastructure.

### How it works

Given an IP address, the system will return a probabilistic value (between a value of ```0``` and ```1```) of how likely the IP is a VPN / proxy / hosting / bad IP. A value of ```1``` means that IP is explicitly banned (a web host, VPN, or TOR node) by our dynamic lists. Otherwise, the output will return a real number value between ```0``` and ```1```, of how likely the IP is bad / VPN / proxy, which is inferred through **machine learning & probability theory techniques** using dynamic checks with large datasets. On average, billions of new records are parsed each month to ensure the datasets have the latest information and old records automatically expire. The system is designed to be efficient, fast, simple, and accurate.

For a deeper understanding of how the API works and the different ```flags``` and ```oflags``` options available, please visit the [API page](https://getipintel.net/free-proxy-vpn-tor-detection-api/).

Here are some example code to use <a href="http://getipintel.net"> GetIPIntel </a> in various code formats. They mainly serve as a proof of concept and should not be implemented directly into your system.

### Easy to use Web interface without any coding
A simple web interface lookup is available via https://getipintel.net/free-proxy-vpn-tor-ip-lookup/


### Full API documentation
The full API documentation is available via https://getipintel.net/free-proxy-vpn-tor-detection-api/ 

### PHP
* requires php curl

### Installation
* **Please change the email variable to your own email**
* Read the documentation on the website for the latest features. If you wish to use flags, change query string.

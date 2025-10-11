# IP Intelligence is a free tool that attempts to determine how likely an IP address is a proxy / VPN / bad IP using mathematical and modern computing techniques

* Greatly reduce fraud on e-commerce sites (anti-fraud)
* Protect your site from automated hacking attempts such as XSS, SQLi, brute force attacks, application scanning and many others
* Protect your site from crawlers that steal your content
* Prevent users from abusing promotional offers / multiple sign-ups / affiliate abuse
* Stop bots from scraping your content or bots spamming your website
* Serve traffic / content to real users, not bots. Reduce fake views, clicks, and activity that results in click fraud and view fraud (anti-bot detection)
* Prevent trolls / people that are trying to bypass a ban
* Adjust your system to limit access (such as not allowing them to change their password, their email, etc) to prevent account hijacking
* Since the system returns a real value and there's different flag options, you can customize the level of protection for a particular time frame and adjust accordingly
* Use it with a combination of another fraud prevention service to make it even better. Some fraud prevention services do not explicitly look for proxy / VPN / bad IPs

The system is serving millions of API requests a week and growing as more people find it useful in protecting their online infrastructure.

### How it works

Given an IP address, the system will return a probabilistic value (between a value of ```0``` and ```1```) of how likely the IP is a VPN / proxy / hosting / bad IP. A value of ```1``` means that IP is explicitly banned (a web host, VPN, or TOR node) by our dynamic lists. Otherwise, the output will return a real number value between ```0``` and ```1```, of how likely the IP is bad / VPN / proxy, which is inferred through **machine learning & probability theory techniques** using dynamic checks with large datasets. On average, billions of new records are parsed each month to ensure the datasets have the latest information and old records automatically expire. The system is designed to be efficient, fast, simple, and accurate.


### Interpretation of the Results
If a value of 0.50 is returned, then it is as good as flipping a 2 sided fair coin, which implies it's not very accurate. From my personal experience, values > 0.95 should be looked at and values > 0.99 are most likely proxies. Anything below the value of 0.90 is considered as "low risk". Since a real value is returned, different levels of protection can be implemented. It is best for a system admin to test some sample datasets with this system and adjust implementation accordingly. **I only recommend automated action on high values ( > 0.99 or even > 0.995 ) but it's good practice to manually review IPs that return high values.** For example, mark an order as "under manual review" and don't automatically provision the product for high proxy values. Be sure to experiment with the results of this system before you use it live on your projects. If you believe the result is wrong, don't hesitate to contact me, I can tell you why. If it's an error on my end, I'll correct it. If you email me, expect a reply within 12 hours.

___


For a deeper understanding of how the API works and the different ```flags``` and ```oflags``` options available, please visit the [API page](https://getipintel.net/free-proxy-vpn-tor-detection-api/). Standard recommendation is to start off with ```flags=m``` option if only proxy / VPN detection is needed. If ```flags=m``` does not have a noticeable impact, then use ```flags=b```. The default query (no flags) is mostly used infront of a payment gateway to protect against fraud because bad IP detection is included.

Here are some example code to use <a href="http://getipintel.net"> GetIPIntel </a> in various code formats. They mainly serve as a proof of concept and should not be implemented directly into your system.

### Easy to use Web interface without any coding
A simple web interface lookup is available via https://getipintel.net/free-proxy-vpn-tor-ip-lookup/


### Full API documentation
No registration or sign up required, only a valid contact email is needed.

Full API Documentation URL: https://getipintel.net/free-proxy-vpn-tor-detection-api/ 

___

‎⚠️ If your website / service is proxied through Cloudflare, make sure you're looking up ```CF-Connecting-IP``` in the headers. Any similiar infrastructure setup should also be aware that the correct IP is looked up.

### PHP
* requires php curl

### Bash
* requires curl

### Installation
* **Please change the email variable to your own email**
* Read the documentation on the website for the latest features. If you wish to use flags, change query string.

### New features 
* added Oct. 2025 - ```oflags=r``` Include explicit residential proxy data to generate the score. [Full documentation of oflags=r](https://getipintel.net/free-proxy-vpn-tor-detection-api/#oflagsr).
* added Jun. 2023 - ```oflags=a``` will return the ASN number(s) of the IP that's being looked up. More information is available on the [API page](https://getipintel.net/free-proxy-vpn-tor-detection-api/).
* added Jun. 2023 - ```oflags=i``` will include Google One and Google Fi VPN IPs.
* added Dec. 2021 - ```oflags=i``` for iCloud Relay Egress IPs - by definition it is still a proxy / VPN a user willingly enables, but this option will allow more flexibility on how to handle these IPs. More information is available on the [API page](https://getipintel.net/free-proxy-vpn-tor-detection-api/).

### New Changes & Notes
- Feb. 07 2025 - By request of a user, the web interface now returns the looked up IP along with the score.
* Apr. 21 2024 - Added a self lookup IP feature on the web interface. Easy 1 click lookup of your own IP.
* Nov. 23 2023 - Malicious / abnormal traffic dataset has been fully incorporated into the proxy / VPN detection API. As always, you can use oflags=b option to see if an IP behaved badly or not.
* added Nov. 3 2023 - A noticeable sized dataset related to malicious / abnormal traffic is being incorporated into the proxy / VPN detection API. It passed internal testing so I've rolled it out to the free API. I'll keep an eye on the weights and if there's no issues, I'll push it to all services. Feel free to reach out if you have any questions.

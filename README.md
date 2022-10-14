# Overcrawl

Web crawling done right.

Under package `web`, `request.go` will request the given URL and analyze its associated robots.txt file to follow the website's policy, `parser.go` will extract the links from the given web page and determine if they should be crawled next.

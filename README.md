# RSSkit #

RSSkit
##### Really Simple Syndication Karl Integrated Tool

## rssparser.py
* Version: 1
* requires python3 feedparser

`
./rssparser.py securitynow.feed https://feeds.twit.tv/sn.xml
`

## post_url.py
`
echo '{"text":"Hello World"}' | ./post_url.py https://msteams.webhook.url
`

## rssparser.py | post_url.py
`
./rssparser.py this.feed https://feed.url/feed.xml | ./post_url.py https://msteams.webhook.url
`



# RSSkit #

RSSkit
##### Really Simple Syndication Karl Integrated Tool

# rssparser
Usage: rssparser db url

db is local file    
url is rss xml feed    

Creates a local data file store of rss/feed titles.  Prints out new entries and adds them to the db.
```
./rssparser sn.feed https://feeds.twit.tv/sn.xml
{
  "title": "Ripple20 Too",
  "link": "https://twit.tv/shows/security-now/episodes/773"
 }
```

## rssparser.go
* Version: 1
```
./rssparser file.feed https://feeds.twit.tv/sn.xml
```
with basic auth...
```
./rssparser file.feed https://username:password@server.url/feed.xml
```

go build
```
go build rssparser.go
```

## rssparser.py
* Version: 1
* requires python3 feedparser

```
./rssparser.py file.feed https://feeds.twit.tv/sn.xml
```

## post_url.py
```
echo '{"text":"Hello World"}' | ./post_url.py https://msteams.webhook.url
```

## rssparser.py | post_url.py
```
./rssparser.py this.feed https://feed.url/feed.xml | ./post_url.py https://msteams.webhook.url
```





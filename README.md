# RSSkit #

RSSkit
##### Really Simple Syndication karl integrated tool

# rss2json
Usage: rss2json url      

url is rss xml feed    

rss2json.py
* Version: 1
* requires python3 feedparser
```
./rss2json.py https://feeds.twit.tv/sn.xml
{
  "title": "Security Now (Audio)",
  "link": "https://twit.tv/shows/security-now",
  "posts": [
    [
      {
        "title": "SN 774: 123456",
        "link": "https://twit.tv/shows/security-now/episodes/774",
        "descrtiption": 
```

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

# Extensible Markup Language (XML)     
     
```
 <?xml version="1.0" encoding="UTF-8"?>
  <feed xmlns="http://www.w3.org/2005/Atom"><title></title><link type='' href='' rel=''/><updated></updated><author><name></name></author><id></id>
      <entry><title></title><link type='' href='' rel=''/><id></id><published></published><updated></updated></entry>

```
     
```
<?xml version="1.0" encoding="utf-8" ?>
<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:creativeCommons="http://backend.userland.com/creativeCommonsRssModule" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:rawvoice="http://www.rawvoice.com/rawvoiceRssModule/" version="2.0">
    <channel>
        <title></title><link></link>
        <item><title></title><itunes:title></itunes:title><link></link><description></description></item>

```
     
```
/* https://feeds.twit.tv/sn.xml

            <title>SN 765: An Authoritarian Internet?</title>
            <itunes:title>An Authoritarian Internet?</itunes:title>
*/
//https://stackoverflow.com/questions/24870309/golang-xml-unmarshalling-issue-local-name-collisions-fail
//https://stackoverflow.com/questions/34975837/parsing-rss-feed-in-go
```




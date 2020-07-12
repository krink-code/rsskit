#!/usr/bin/env python3

import sys
def usage():
    print(sys.argv[0] + """ url

    --basic_auth username:password
    """)
    sys.exit(0)

try:
    import feedparser
except ImportError as e:
    print(str(e) + """
    install feedparser module...
    macos: python3 -m pip install feedparser
    debian: apt-get install python3-feedparser
    redhat: yum     install python3-feedparser
    """)
    sys.exit(1)

import os
import json

def get_feed(url, auth):
    if auth:
        import base64
        username = auth.split(':')[0]
        password = auth.split(':')[1]

        credentials = (username + ':' + password).encode('utf-8')
        base64_encoded_credentials = base64.b64encode(credentials).decode('utf-8')
        headers = { 'Authorization': 'Basic ' + base64_encoded_credentials }
    else:
        headers = None

    feed = feedparser.parse(url, request_headers=headers)
    status = feed.get('status', default=None)

    if auth and not status:
        print('Invalid Basic Auth')
        return False

    if feed.status != 200:
        print(str(feed.status) + ' ' + str(feed.url))
        return False

    if not feed.entries:
        print(str(feed))
        print('No feed.entries')
        return False

    posts_to_print = []

    for post in feed.entries:

        title = post.get('title', default=None)
        link = post.get('link', default=None)
        description = post.get('description', default=None)
        published = post.get('published', default=None)
        #summary = post.get('summary', default=None)
        #content = post.get('content', default=None)

        posts_to_print.append({'title':title,
                               'link':link, 
                               'descrtiption': description,
                               'published': published})

    jdata = {'title':feed.channel.title,'link':feed.channel.link, 'posts': posts_to_print}
    print(json.dumps(jdata, indent=2))


if __name__ == '__main__':
   
    if sys.argv[1:]:

        url = sys.argv[1]
        auth = None
        argc = 0

        for arg in sys.argv[1:]:
            argc += 1
            if arg == '--basic_auth':
                auth = sys.argv[argc + 1]
    else:
        usage()

    sys.exit(get_feed(url, auth))



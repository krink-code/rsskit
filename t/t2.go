
package main

import "fmt"
import "encoding/xml"

type Rss struct {
   XMLName     xml.Name `xml:"rss"`
   Version     string   `xml:"version,attr"`
   Channel     Channel  `xml:"channel"`
   Description string   `xml:"description"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
}

type Channel struct {
   XMLName     xml.Name `xml:"channel"`
   Title       string   `xml:"title"`
   Link        string   `xml:"link"`
   Description string   `xml:"description"`
   Items       []Item   `xml:"item"`
}

type Item struct {
   //Name        string `xml:"name,attr"`
   //Name        string `xml:"name"`
   Title       string `xml:"title"`
   Link        string `xml:"link"`
   Description string `xml:"description"`
   PubDate     string `xml:"pubdate"`
   Guid        string `xml:"guid"`
}


func main() {
	//x := `<rss><channel></channel></rss>`
	x := `<?xml version="1.0" encoding="utf-8" ?>
<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:creativeCommons="http://backend.userland.com/creativeCommonsRssModule" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:rawvoice="http://www.rawvoice.com/rawvoiceRssModule/" version="2.0">
	<channel>
		<title>Security Now (Audio)</title>
		<link>https://twit.tv/shows/security-now</link>
		<generator>TWiT Feed Generator v3.8</generator>
		<docs>http://blogs.law.harvard.edu/tech/rss</docs>
		<language>en-US</language>
		<copyright>This work is licensed under a Creative Commons License - Attribution-NonCommercial-NoDerivatives 4.0 International - http://creativecommons.org/licenses/by-nc-nd/4.0/</copyright>
		<creativeCommons:license>http://creativecommons.org/licenses/by-nc-nd/4.0/</creativeCommons:license>
		<managingEditor>distro@twit.tv (TWiT Editors)</managingEditor>
		<webMaster>distro@twit.tv (TWiT Engineering)</webMaster>
		<ttl>720</ttl>
		<sy:updatePeriod>weekly</sy:updatePeriod>
		<sy:updateFrequency>1</sy:updateFrequency>
		<lastBuildDate>Tue, 07 Jul 2020 21:30:48 PDT</lastBuildDate>
		<pubDate>Tue, 07 Jul 2020 21:30:48 PDT</pubDate>
		<rawvoice:frequency>weekly</rawvoice:frequency>
		<rawvoice:location>Petaluma, CA</rawvoice:location>
		<itunes:type>episodic</itunes:type>
		<itunes:author>TWiT</itunes:author>
		<itunes:subtitle>Steve Gibson discusses the hot topics in security today with Leo Laporte.</itunes:subtitle>
		<itunes:summary>Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.

Records live every Tuesday at 4:30pm Eastern / 1:30pm Pacific / 20:30 UTC.</itunes:summary>
		<description>Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.

Records live every Tuesday at 4:30pm Eastern / 1:30pm Pacific / 20:30 UTC.</description>
		<itunes:keywords>TWiT, Technology, Steve Gibson, Leo Laporte, security, spyware, malware, hacking, cyber crime, emcryption</itunes:keywords>
		<rawvoice:rating tv="tv-g">tv-g</rawvoice:rating>
		<itunes:explicit>false</itunes:explicit>
		<itunes:block>no</itunes:block>
		<itunes:owner>
			<itunes:name>Leo Laporte</itunes:name>
			<itunes:email>distro@twit.tv</itunes:email>
		</itunes:owner>
		<itunes:category text="News">
			<itunes:category text="Tech News"/>
</itunes:category>
		<itunes:category text="Technology" />
		<image>
			<title>Security Now (Audio)</title>
			<url>https://elroy.twit.tv/sites/default/files/styles/twit_album_art_144x144/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=NS2JTH0t</url>
			<link>https://twit.tv/shows/security-now</link>
			<width>144</width>
			<height>144</height>
		</image>
		<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" />
		<itunes:new-feed-url>https://feeds.twit.tv/sn.xml</itunes:new-feed-url>
		<rawvoice:subscribe feed="https://feeds.twit.tv/sn.xml" itunes="https://podcasts.apple.com/us/podcast/security-now-mp3/id79016499?uo=10" html="https://twit.tv/shows/security-now"></rawvoice:subscribe>
		<atom:link href="https://feeds.twit.tv/sn.xml" type="application/rss+xml" rel="self"/>
		<item>
			<title>SN 774: 123456</title>
			<itunes:title>123456</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 07 Jul 2020 17:30:00 PDT</pubDate>
			<itunes:episode>774</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/774</link>
			<comments>https://twit.tv/shows/security-now/episodes/774</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Boston Bans Face Recognition, Bad Passwords</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, face recognition, boston face recognition, 123456 password, ios 14 tik tok, ios 14 linkedin, us cert emergency windows update, emergency windows update, hackerone bug bounty, sony bug bounty, f5 vulnerability</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Boston bans face recognition, bad passwords.</p><ul><li>Boston bans facial recognition</li><li>123456 is still the most popular password</li><li>iOS 14 catches Linked-In, Tik Tok, and others red handed!</li><li>US-CERT notes two Emergency Windows Updates</li><li>HackerOne shares their top 10 public bug bounty programs</li><li>Sony launches PlayStation bug bounty program with rewards of $50K+</li><li>F5 Networks patches a highest-severity vulnerability</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-774-Notes.pdf">https://www.grc.com/sn/SN-774-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Boston bans face recognition, bad passwords.</p><ul><li>Boston bans facial recognition</li><li>123456 is still the most popular password</li><li>iOS 14 catches Linked-In, Tik Tok, and others red handed!</li><li>US-CERT notes two Emergency Windows Updates</li><li>HackerOne shares their top 10 public bug bounty programs</li><li>Sony launches PlayStation bug bounty program with rewards of $50K+</li><li>F5 Networks patches a highest-severity vulnerability</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-774-Notes.pdf">https://www.grc.com/sn/SN-774-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Boston bans face recognition, bad passwords.</p><ul><li>Boston bans facial recognition</li><li>123456 is still the most popular password</li><li>iOS 14 catches Linked-In, Tik Tok, and others red handed!</li><li>US-CERT notes two Emergency Windows Updates</li><li>HackerOne shares their top 10 public bug bounty programs</li><li>Sony launches PlayStation bug bounty program with rewards of $50K+</li><li>F5 Networks patches a highest-severity vulnerability</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-774-Notes.pdf">https://www.grc.com/sn/SN-774-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/818577/hero/sn774_thumbnail.jpg?itok=Bk96eI8o"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0774/sn0774.mp3</guid>
			<itunes:duration>1:56:10</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0774/sn0774.mp3" length="55878344" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0774/sn0774.mp3" fileSize="55878344" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 774: 123456</media:title>
				<media:description type="plain">Boston Bans Face Recognition, Bad Passwords</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, face recognition, boston face recognition, 123456 password, ios 14 tik tok, ios 14 linkedin, us cert emergency windows update, emergency windows update, hackerone bug bounty, sony bug bounty, f5 vulnerability</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/818577/hero/sn774_thumbnail.jpg?itok=miq1GWgy" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 773: Ripple20 Too</title>
			<itunes:title>Ripple20 Too</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 30 Jun 2020 17:30:00 PDT</pubDate>
			<itunes:episode>773</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/773</link>
			<comments>https://twit.tv/shows/security-now/episodes/773</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Congress Wants to Kill Encryption &amp; Face Recognition</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, apple browser certificate, safari web APIs, mozilla comcast DOH, facial recognition and biometric technology moratorium act, congress face recognition, face recognition bill, lawful access to encrypted data, </itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Congress wants to kill encryption &amp; face recognition.</p><ul><li>New information about Ripple20</li><li>The Facial Recognition and Biometric Technology Moratorium Act wants to kill face recognition</li><li>The Lawful Access to Encrypted Data Act wants to kill encryption</li><li>Michigan State's legislative House passed the "Microchip Protection Act"</li><li>Apple forces the industry down to one-year web browser certificate lifespans</li><li>Safari to eschew 16 new web API's for the sake of user privacy</li><li>Apple also got on the DoH &amp; DoT bandwagon</li><li>Mozilla + Comcast + DoH: Strange Bedfellows</li><li>Don't forget about VirusTotal</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-773-Notes.pdf">https://www.grc.com/sn/SN-773-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://Melissa.com/twit">Melissa.com/twit</a></li>
<li><a href="http://OpenShift.com/SecurityNow">OpenShift.com/SecurityNow</a></li>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Congress wants to kill encryption &amp; face recognition.</p><ul><li>New information about Ripple20</li><li>The Facial Recognition and Biometric Technology Moratorium Act wants to kill face recognition</li><li>The Lawful Access to Encrypted Data Act wants to kill encryption</li><li>Michigan State's legislative House passed the "Microchip Protection Act"</li><li>Apple forces the industry down to one-year web browser certificate lifespans</li><li>Safari to eschew 16 new web API's for the sake of user privacy</li><li>Apple also got on the DoH &amp; DoT bandwagon</li><li>Mozilla + Comcast + DoH: Strange Bedfellows</li><li>Don't forget about VirusTotal</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-773-Notes.pdf">https://www.grc.com/sn/SN-773-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://Melissa.com/twit">Melissa.com/twit</a></li>
<li><a href="http://OpenShift.com/SecurityNow">OpenShift.com/SecurityNow</a></li>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Congress wants to kill encryption &amp; face recognition.</p><ul><li>New information about Ripple20</li><li>The Facial Recognition and Biometric Technology Moratorium Act wants to kill face recognition</li><li>The Lawful Access to Encrypted Data Act wants to kill encryption</li><li>Michigan State's legislative House passed the "Microchip Protection Act"</li><li>Apple forces the industry down to one-year web browser certificate lifespans</li><li>Safari to eschew 16 new web API's for the sake of user privacy</li><li>Apple also got on the DoH &amp; DoT bandwagon</li><li>Mozilla + Comcast + DoH: Strange Bedfellows</li><li>Don't forget about VirusTotal</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-773-Notes.pdf">https://www.grc.com/sn/SN-773-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://Melissa.com/twit">Melissa.com/twit</a></li>
<li><a href="http://OpenShift.com/SecurityNow">OpenShift.com/SecurityNow</a></li>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/818358/hero/sn773_thumbnail.jpg?itok=ZBjNywxe"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0773/sn0773.mp3</guid>
			<itunes:duration>1:51:43</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0773/sn0773.mp3" length="53742154" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0773/sn0773.mp3" fileSize="53742154" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 773: Ripple20 Too</media:title>
				<media:description type="plain">Congress Wants to Kill Encryption &amp; Face Recognition</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, apple browser certificate, safari web APIs, mozilla comcast DOH, facial recognition and biometric technology moratorium act, congress face recognition, face recognition bill, lawful access to encrypted data, </media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/818358/hero/sn773_thumbnail.jpg?itok=YhGuCmtR" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 772: Ripple20</title>
			<itunes:title>Ripple20</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 23 Jun 2020 18:00:00 PDT</pubDate>
			<itunes:episode>772</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/772</link>
			<comments>https://twit.tv/shows/security-now/episodes/772</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Zoom Encryption, Windows 10 Printer Error</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Ripple20, Ripple20 TCP/IP, zoom end-to-end encryption, zoom free encryption, zoom encryption, Windows 10 printing error, Windows won't print, VLC code execution flaw, Telegram Russia, Netgear security</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Zoom encryption, Windows 10 printer error.</p><ul><li>Ripple20: a set of 19 TCP/IP vulnerabilities that could let remote attackers gain control over your device</li><li>Russian government lifts its failed ban on Telegram</li><li>Zoom: everybody gets optional end to end encryption</li><li>Google removed 106 malicious Chrome extensions collecting sensitive user data</li><li>Windows 10 update breaks printing</li><li>VLC Media Player 3.0.11 fixes severe remote code execution flaw</li><li>Netgear in the doghouse</li><li>DDoS is alive and well... and growing</li><li>How to get the new Edge for Windows 7</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-772-Notes.pdf">https://www.grc.com/sn/SN-772-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://GetRoman.com/SECURITYNOW">GetRoman.com/SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Zoom encryption, Windows 10 printer error.</p><ul><li>Ripple20: a set of 19 TCP/IP vulnerabilities that could let remote attackers gain control over your device</li><li>Russian government lifts its failed ban on Telegram</li><li>Zoom: everybody gets optional end to end encryption</li><li>Google removed 106 malicious Chrome extensions collecting sensitive user data</li><li>Windows 10 update breaks printing</li><li>VLC Media Player 3.0.11 fixes severe remote code execution flaw</li><li>Netgear in the doghouse</li><li>DDoS is alive and well... and growing</li><li>How to get the new Edge for Windows 7</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-772-Notes.pdf">https://www.grc.com/sn/SN-772-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://GetRoman.com/SECURITYNOW">GetRoman.com/SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Zoom encryption, Windows 10 printer error.</p><ul><li>Ripple20: a set of 19 TCP/IP vulnerabilities that could let remote attackers gain control over your device</li><li>Russian government lifts its failed ban on Telegram</li><li>Zoom: everybody gets optional end to end encryption</li><li>Google removed 106 malicious Chrome extensions collecting sensitive user data</li><li>Windows 10 update breaks printing</li><li>VLC Media Player 3.0.11 fixes severe remote code execution flaw</li><li>Netgear in the doghouse</li><li>DDoS is alive and well... and growing</li><li>How to get the new Edge for Windows 7</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-772-Notes.pdf">https://www.grc.com/sn/SN-772-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://GetRoman.com/SECURITYNOW">GetRoman.com/SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/818152/hero/sn772_thumbnail_v2.jpg?itok=GNE_WvvZ"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0772/sn0772.mp3</guid>
			<itunes:duration>2:07:14</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0772/sn0772.mp3" length="61186007" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0772/sn0772.mp3" fileSize="61186007" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 772: Ripple20</media:title>
				<media:description type="plain">Zoom Encryption, Windows 10 Printer Error</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Ripple20, Ripple20 TCP/IP, zoom end-to-end encryption, zoom free encryption, zoom encryption, Windows 10 printing error, Windows won't print, VLC code execution flaw, Telegram Russia, Netgear security</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/818152/hero/sn772_thumbnail_v2.jpg?itok=ojt8Q90-" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 771: Lamphone</title>
			<itunes:title>Lamphone</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 16 Jun 2020 17:00:00 PDT</pubDate>
			<itunes:episode>771</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/771</link>
			<comments>https://twit.tv/shows/security-now/episodes/771</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Windows Update Kills Printers &amp; SSDs</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, lamphone, Brave Browser, Microsoft Patch Tuesday, windows printer error, windows printer port error, windows can't find my printer, windows SSD error, your site has been hacked, SMBleed</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Windows update kills printers &amp; SSDs.</p><ul><li>Lamphone: eavesdrop on a hanging lightbulb</li><li>Brave Browser caught and chastised for tweaking user-entered URLs for its benefit</li><li>Microsoft breaks its own record for Patch Tuesday patches</li><li>TFW Windows 10 loses your printer port</li><li>Last week's Patch Tuesday broke ALL PRINTING (even to PDFs) for many users. Fix won't come for a month</li><li>Windows 10 2004 update is messing up SSDs and non-SSDs</li><li>SMBleed</li><li>Subject: Your Site Has Been Hacked</li><li>Authentic database ransom attacks</li><li>Another side-channel attack on Intel chips</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-771-Notes.pdf">https://www.grc.com/sn/SN-771-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Windows update kills printers &amp; SSDs.</p><ul><li>Lamphone: eavesdrop on a hanging lightbulb</li><li>Brave Browser caught and chastised for tweaking user-entered URLs for its benefit</li><li>Microsoft breaks its own record for Patch Tuesday patches</li><li>TFW Windows 10 loses your printer port</li><li>Last week's Patch Tuesday broke ALL PRINTING (even to PDFs) for many users. Fix won't come for a month</li><li>Windows 10 2004 update is messing up SSDs and non-SSDs</li><li>SMBleed</li><li>Subject: Your Site Has Been Hacked</li><li>Authentic database ransom attacks</li><li>Another side-channel attack on Intel chips</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-771-Notes.pdf">https://www.grc.com/sn/SN-771-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Windows update kills printers &amp; SSDs.</p><ul><li>Lamphone: eavesdrop on a hanging lightbulb</li><li>Brave Browser caught and chastised for tweaking user-entered URLs for its benefit</li><li>Microsoft breaks its own record for Patch Tuesday patches</li><li>TFW Windows 10 loses your printer port</li><li>Last week's Patch Tuesday broke ALL PRINTING (even to PDFs) for many users. Fix won't come for a month</li><li>Windows 10 2004 update is messing up SSDs and non-SSDs</li><li>SMBleed</li><li>Subject: Your Site Has Been Hacked</li><li>Authentic database ransom attacks</li><li>Another side-channel attack on Intel chips</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-771-Notes.pdf">https://www.grc.com/sn/SN-771-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/817885/hero/sn771_thumbnail.jpg?itok=8XUZaYTd"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0771/sn0771.mp3</guid>
			<itunes:duration>1:50:24</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0771/sn0771.mp3" length="53105811" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0771/sn0771.mp3" fileSize="53105811" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 771: Lamphone</media:title>
				<media:description type="plain">Windows Update Kills Printers &amp; SSDs</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, lamphone, Brave Browser, Microsoft Patch Tuesday, windows printer error, windows printer port error, windows can't find my printer, windows SSD error, your site has been hacked, SMBleed</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/817885/hero/sn771_thumbnail.jpg?itok=1ThMkXHt" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 770: Zoom's E2EE Debacle</title>
			<itunes:title>Zoom's E2EE Debacle</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 09 Jun 2020 19:00:00 PDT</pubDate>
			<itunes:episode>770</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/770</link>
			<comments>https://twit.tv/shows/security-now/episodes/770</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Zoom's End-to-End Encryption Fail</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, mozilla, firefox, DoH DDoS, IBM face recognition, iBM facial recognition, Cisco Webex, cisco Talos, Cisco Zoom, Callstranger, callstranger UPnP, Microsoft edge, chromium edge, Chredge, zoom E2EE</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Zoom's end-to-end encryption fail.</p><ul><li>Zoom will offer end-to-end encryption, but only if you pay for it</li><li>IBM announces no more work on facial recognition</li><li>The Odd Case of Mozilla's DoH DDoS</li><li>Cisco's Talos group found two critical flaws in the Zoom client</li><li>CallStranger UPnP bug has tech press in a tizzy</li><li>Microsoft has started to replace old Edge with new Edge</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-770-Notes.pdf">https://www.grc.com/sn/SN-770-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Zoom's end-to-end encryption fail.</p><ul><li>Zoom will offer end-to-end encryption, but only if you pay for it</li><li>IBM announces no more work on facial recognition</li><li>The Odd Case of Mozilla's DoH DDoS</li><li>Cisco's Talos group found two critical flaws in the Zoom client</li><li>CallStranger UPnP bug has tech press in a tizzy</li><li>Microsoft has started to replace old Edge with new Edge</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-770-Notes.pdf">https://www.grc.com/sn/SN-770-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Zoom's end-to-end encryption fail.</p><ul><li>Zoom will offer end-to-end encryption, but only if you pay for it</li><li>IBM announces no more work on facial recognition</li><li>The Odd Case of Mozilla's DoH DDoS</li><li>Cisco's Talos group found two critical flaws in the Zoom client</li><li>CallStranger UPnP bug has tech press in a tizzy</li><li>Microsoft has started to replace old Edge with new Edge</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-770-Notes.pdf">https://www.grc.com/sn/SN-770-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/817623/hero/sn770_thumbnail_v2.jpg?itok=6rFPhz-m"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0770/sn0770.mp3</guid>
			<itunes:duration>1:48:16</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0770/sn0770.mp3" length="52086618" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0770/sn0770.mp3" fileSize="52086618" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 770: Zoom's E2EE Debacle</media:title>
				<media:description type="plain">Zoom's End-to-End Encryption Fail</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, mozilla, firefox, DoH DDoS, IBM face recognition, iBM facial recognition, Cisco Webex, cisco Talos, Cisco Zoom, Callstranger, callstranger UPnP, Microsoft edge, chromium edge, Chredge, zoom E2EE</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/817623/hero/sn770_thumbnail_v2.jpg?itok=IC42hI3G" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 769: Zoom's E2EE Design</title>
			<itunes:title>Zoom's E2EE Design</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 02 Jun 2020 18:00:00 PDT</pubDate>
			<itunes:episode>769</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/769</link>
			<comments>https://twit.tv/shows/security-now/episodes/769</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Zoom Gets End-to-End Encryption</itunes:subtitle>
			<itunes:keywords>SecuritySecurity Now, TWiT, steve gibson, Leo Laporte, zoom encryption, zoom end-to-end encryption, zoom security, malvertizing downloads, Threatmetrix, port scanning, facebook verification, google encryption, google messages encryption, strandhogg</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Zoom gets end-to-end encryption.</p><ul><li> ACLU takes Clearview to court, but maybe they should worry about their own website first</li><li>The state of drive-by malvertising downloads</li><li>Google will be bad listing notification abusing sites</li><li>Who else is doing the eBay-like ThreatMetrix port scanning?</li><li>Facebook to require identity verification for high impact posters</li><li>Google Messaging is apparently heading toward E2EE</li><li>The return of a much more worrisome StrandHogg</li><li>The SHA-1 hash to finally be dropped from OpenSSH</li><li>What happens when you fuzz USB?</li><li>Zoom's end-to-end encryption design</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-769-Notes.pdf">https://www.grc.com/sn/SN-769-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Zoom gets end-to-end encryption.</p><ul><li> ACLU takes Clearview to court, but maybe they should worry about their own website first</li><li>The state of drive-by malvertising downloads</li><li>Google will be bad listing notification abusing sites</li><li>Who else is doing the eBay-like ThreatMetrix port scanning?</li><li>Facebook to require identity verification for high impact posters</li><li>Google Messaging is apparently heading toward E2EE</li><li>The return of a much more worrisome StrandHogg</li><li>The SHA-1 hash to finally be dropped from OpenSSH</li><li>What happens when you fuzz USB?</li><li>Zoom's end-to-end encryption design</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-769-Notes.pdf">https://www.grc.com/sn/SN-769-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Zoom gets end-to-end encryption.</p><ul><li> ACLU takes Clearview to court, but maybe they should worry about their own website first</li><li>The state of drive-by malvertising downloads</li><li>Google will be bad listing notification abusing sites</li><li>Who else is doing the eBay-like ThreatMetrix port scanning?</li><li>Facebook to require identity verification for high impact posters</li><li>Google Messaging is apparently heading toward E2EE</li><li>The return of a much more worrisome StrandHogg</li><li>The SHA-1 hash to finally be dropped from OpenSSH</li><li>What happens when you fuzz USB?</li><li>Zoom's end-to-end encryption design</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-769-Notes.pdf">https://www.grc.com/sn/SN-769-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://expressvpn.com/securitynow">expressvpn.com/securitynow</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/817378/hero/sn0769_thumbnail_v2.jpg?itok=AMeVe8tL"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0769/sn0769.mp3</guid>
			<itunes:duration>2:12:03</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0769/sn0769.mp3" length="63500665" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0769/sn0769.mp3" fileSize="63500665" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 769: Zoom's E2EE Design</media:title>
				<media:description type="plain">Zoom Gets End-to-End Encryption</media:description>
				<media:keywords>SecuritySecurity Now, TWiT, steve gibson, Leo Laporte, zoom encryption, zoom end-to-end encryption, zoom security, malvertizing downloads, Threatmetrix, port scanning, facebook verification, google encryption, google messages encryption, strandhogg</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/817378/hero/sn0769_thumbnail_v2.jpg?itok=ObwquHL5" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 768: Contact Tracing Apps R.I.P.</title>
			<itunes:title>Contact Tracing Apps R.I.P.</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 26 May 2020 19:33:04 PDT</pubDate>
			<itunes:episode>768</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/768</link>
			<comments>https://twit.tv/shows/security-now/episodes/768</comments>
			<itunes:author>TWiT</itunes:author>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Contact Tracing Apps Are Not Going to Work</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Contact Tracing, contact tracing apps, human contact tracing, Unc0ver, Unc0ver jailbreak, Unc0ver iOS Jailbreak, iOS jailbreak, ios 13.5 jailbreak, Firefox 77 security, Chome 83 features, Chrome tab grouping,</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Contact tracing apps are not going to work.</p><ul><li>Why contact tracing apps are never going to work</li><li>Unc0ver: There's a new iOS jailbreak in town, and as jailbreaks go, it looks VERY nice!</li><li>Firefox 77 picks up a nifty new security trick</li><li>New features in Chrome 83: cookie management, "Safety Check," blocking third-party cookies by default in Incognito mode, and "Tab Groups"</li><li>Adobe rushes out four out-of-cycle emergency updates to fix security flaws</li><li>Zerodium temporarily stops buying iOS remote code execution vulnerabilities</li><li>The NXNS Attack: A group of cybersecurity researchers in Israeli have responsibly disclosed details about a new way they worked out of using the Internet's domain name resolution system to hugely amplify (by a factor of at least 1620 packets) a DDoS attack to take down targeted websites.</li><li>BIAS - Bluetooth Impersonation AttackS is nothing less than a complete collapse of Bluetooth security.</li><li>Is eBay port scanning its user's computers? Kinda.</li><li>Security Now trivia: Steve Gibson helped develop the Speak &amp; Spell! It did voice synthesis with only a 4K bits (0.5K bytes) processor.</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-768-Notes.pdf">https://www.grc.com/sn/SN-768-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Contact tracing apps are not going to work.</p><ul><li>Why contact tracing apps are never going to work</li><li>Unc0ver: There's a new iOS jailbreak in town, and as jailbreaks go, it looks VERY nice!</li><li>Firefox 77 picks up a nifty new security trick</li><li>New features in Chrome 83: cookie management, "Safety Check," blocking third-party cookies by default in Incognito mode, and "Tab Groups"</li><li>Adobe rushes out four out-of-cycle emergency updates to fix security flaws</li><li>Zerodium temporarily stops buying iOS remote code execution vulnerabilities</li><li>The NXNS Attack: A group of cybersecurity researchers in Israeli have responsibly disclosed details about a new way they worked out of using the Internet's domain name resolution system to hugely amplify (by a factor of at least 1620 packets) a DDoS attack to take down targeted websites.</li><li>BIAS - Bluetooth Impersonation AttackS is nothing less than a complete collapse of Bluetooth security.</li><li>Is eBay port scanning its user's computers? Kinda.</li><li>Security Now trivia: Steve Gibson helped develop the Speak &amp; Spell! It did voice synthesis with only a 4K bits (0.5K bytes) processor.</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-768-Notes.pdf">https://www.grc.com/sn/SN-768-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Contact tracing apps are not going to work.</p><ul><li>Why contact tracing apps are never going to work</li><li>Unc0ver: There's a new iOS jailbreak in town, and as jailbreaks go, it looks VERY nice!</li><li>Firefox 77 picks up a nifty new security trick</li><li>New features in Chrome 83: cookie management, "Safety Check," blocking third-party cookies by default in Incognito mode, and "Tab Groups"</li><li>Adobe rushes out four out-of-cycle emergency updates to fix security flaws</li><li>Zerodium temporarily stops buying iOS remote code execution vulnerabilities</li><li>The NXNS Attack: A group of cybersecurity researchers in Israeli have responsibly disclosed details about a new way they worked out of using the Internet's domain name resolution system to hugely amplify (by a factor of at least 1620 packets) a DDoS attack to take down targeted websites.</li><li>BIAS - Bluetooth Impersonation AttackS is nothing less than a complete collapse of Bluetooth security.</li><li>Is eBay port scanning its user's computers? Kinda.</li><li>Security Now trivia: Steve Gibson helped develop the Speak &amp; Spell! It did voice synthesis with only a 4K bits (0.5K bytes) processor.</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-768-Notes.pdf">https://www.grc.com/sn/SN-768-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
<li><a href="http://Wasabi.com">Wasabi.com  offer code SECURITYNOW</a></li>
<li><a href="http://extrahop.com/SECURITYNOW">extrahop.com/SECURITYNOW</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/817144/hero/sn768_thumbnail_v2_1.jpg?itok=ExtCsvMG"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0768/sn0768.mp3</guid>
			<itunes:duration>1:50:34</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0768/sn0768.mp3" length="53188358" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0768/sn0768.mp3" fileSize="53188358" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 768: Contact Tracing Apps R.I.P.</media:title>
				<media:description type="plain">Contact Tracing Apps Are Not Going to Work</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Contact Tracing, contact tracing apps, human contact tracing, Unc0ver, Unc0ver jailbreak, Unc0ver iOS Jailbreak, iOS jailbreak, ios 13.5 jailbreak, Firefox 77 security, Chome 83 features, Chrome tab grouping,</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/817144/hero/sn768_thumbnail_v2_1.jpg?itok=LEJz_dJI" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 767: WiFi 6</title>
			<itunes:title>WiFi 6</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 19 May 2020 17:00:00 PDT</pubDate>
			<itunes:episode>767</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/767</link>
			<comments>https://twit.tv/shows/security-now/episodes/767</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>WiFi 6, Apple vs. FBI, Face Masks</itunes:subtitle>
			<itunes:keywords>SecuritySecurity Now, TWiT, steve gibson, Leo Laporte, Windows Patch Tuesday, apple vs DOJ, apple vs fbi, face masks, london face masks, face masks lfr, london live face recognition, utah contact tracing app, Wifi 6, wi-fi 6</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>WiFi 6, Apple vs. FBI, face masks.</p><ul><li>Last Tuesday's Windows patch Tuesday was not the biggest ever, but it was the 3rd largest in Microsoft's history, weighing in with a whopping 111 CVE-tracked bug fixes, 16 of which were rated CRITICAL and all but one of which enabled Remote Code Execution by an attacker.</li><li>The DOJ and FBI again criticize Apple over encryption</li><li>When is a fix not a fix?</li><li>Face masks have thwarted the London police's LFR rollout</li><li>Utah chooses to roll their own contact tracing app</li><li>Everything you need to know about WiFi 6</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-767-Notes.pdf" target="_blank">https://www.grc.com/sn/SN-767-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>WiFi 6, Apple vs. FBI, face masks.</p><ul><li>Last Tuesday's Windows patch Tuesday was not the biggest ever, but it was the 3rd largest in Microsoft's history, weighing in with a whopping 111 CVE-tracked bug fixes, 16 of which were rated CRITICAL and all but one of which enabled Remote Code Execution by an attacker.</li><li>The DOJ and FBI again criticize Apple over encryption</li><li>When is a fix not a fix?</li><li>Face masks have thwarted the London police's LFR rollout</li><li>Utah chooses to roll their own contact tracing app</li><li>Everything you need to know about WiFi 6</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-767-Notes.pdf" target="_blank">https://www.grc.com/sn/SN-767-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>WiFi 6, Apple vs. FBI, face masks.</p><ul><li>Last Tuesday's Windows patch Tuesday was not the biggest ever, but it was the 3rd largest in Microsoft's history, weighing in with a whopping 111 CVE-tracked bug fixes, 16 of which were rated CRITICAL and all but one of which enabled Remote Code Execution by an attacker.</li><li>The DOJ and FBI again criticize Apple over encryption</li><li>When is a fix not a fix?</li><li>Face masks have thwarted the London police's LFR rollout</li><li>Utah chooses to roll their own contact tracing app</li><li>Everything you need to know about WiFi 6</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-767-Notes.pdf" target="_blank">https://www.grc.com/sn/SN-767-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://canary.tools/twit">canary.tools/twit - use code: TWIT</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/816913/hero/sn767_tn.jpg?itok=4S8DoQSi"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0767/sn0767.mp3</guid>
			<itunes:duration>2:00:03</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0767/sn0767.mp3" length="57738262" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0767/sn0767.mp3" fileSize="57738262" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 767: WiFi 6</media:title>
				<media:description type="plain">WiFi 6, Apple vs. FBI, Face Masks</media:description>
				<media:keywords>SecuritySecurity Now, TWiT, steve gibson, Leo Laporte, Windows Patch Tuesday, apple vs DOJ, apple vs fbi, face masks, london face masks, face masks lfr, london live face recognition, utah contact tracing app, Wifi 6, wi-fi 6</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/816913/hero/sn767_tn.jpg?itok=x5KmPIuH" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 766: ThunderSpy</title>
			<itunes:title>ThunderSpy</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 12 May 2020 15:00:00 PDT</pubDate>
			<itunes:episode>766</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/766</link>
			<comments>https://twit.tv/shows/security-now/episodes/766</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>Thunderbolt Security Flaw, Zoom Buys Keybase</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Thunderspy, thunderbolt flae, thunderbolt security, zoom keybase, zoom, Keybase, Firefox 76, Firefox Amazon, edge notifications, WordPress plugin bugs, vbullitin patch, samsung patch, defcon 2020 canceled</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>Thunderbolt security flaw, Zoom buys Keybase.</p><ul><li> Why the ThunderSpy Thunderbolt security flaw is such a big deal</li><li>Zoom purchases Keybase to fix encryption</li><li>Firefox 76 released with new features</li><li>But Firefox 76 broke Amazon's Assistant!</li><li>Hallelujah!! Edge moves to silence those annoying notification requests.</li><li>Critical WordPress plugin bugs present on over one million sites</li><li>Critical vBulletin patch</li><li>Samsung has patched a CRITICAL bug affecting the past 6 years of Smartphones</li><li>DefCon and Black Hat 2020 go virtual</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-766-Notes.pdf">https://www.grc.com/sn/SN-766-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://manscaped.com">manscaped.com code SECURITYNOW</a></li>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>Thunderbolt security flaw, Zoom buys Keybase.</p><ul><li> Why the ThunderSpy Thunderbolt security flaw is such a big deal</li><li>Zoom purchases Keybase to fix encryption</li><li>Firefox 76 released with new features</li><li>But Firefox 76 broke Amazon's Assistant!</li><li>Hallelujah!! Edge moves to silence those annoying notification requests.</li><li>Critical WordPress plugin bugs present on over one million sites</li><li>Critical vBulletin patch</li><li>Samsung has patched a CRITICAL bug affecting the past 6 years of Smartphones</li><li>DefCon and Black Hat 2020 go virtual</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-766-Notes.pdf">https://www.grc.com/sn/SN-766-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://manscaped.com">manscaped.com code SECURITYNOW</a></li>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>Thunderbolt security flaw, Zoom buys Keybase.</p><ul><li> Why the ThunderSpy Thunderbolt security flaw is such a big deal</li><li>Zoom purchases Keybase to fix encryption</li><li>Firefox 76 released with new features</li><li>But Firefox 76 broke Amazon's Assistant!</li><li>Hallelujah!! Edge moves to silence those annoying notification requests.</li><li>Critical WordPress plugin bugs present on over one million sites</li><li>Critical vBulletin patch</li><li>Samsung has patched a CRITICAL bug affecting the past 6 years of Smartphones</li><li>DefCon and Black Hat 2020 go virtual</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-766-Notes.pdf">https://www.grc.com/sn/SN-766-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://manscaped.com">manscaped.com code SECURITYNOW</a></li>
<li><a href="http://itpro.tv/securitynow">itpro.tv/securitynow  promo code SN30</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/816667/hero/sn0766_thumbnail.jpg?itok=XF8YBVor"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0766/sn0766.mp3</guid>
			<itunes:duration>1:57:48</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0766/sn0766.mp3" length="56658047" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0766/sn0766.mp3" fileSize="56658047" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 766: ThunderSpy</media:title>
				<media:description type="plain">Thunderbolt Security Flaw, Zoom Buys Keybase</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, Thunderspy, thunderbolt flae, thunderbolt security, zoom keybase, zoom, Keybase, Firefox 76, Firefox Amazon, edge notifications, WordPress plugin bugs, vbullitin patch, samsung patch, defcon 2020 canceled</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/816667/hero/sn0766_thumbnail.jpg?itok=mKGMZTSj" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
		<item>
			<title>SN 765: An Authoritarian Internet? --Joy.Joy.Title</title>
			<itunes:title>An Authoritarian Internet? --YES.YES.itunes</itunes:title>
			<itunes:episodeType>full</itunes:episodeType>
			<pubDate>Tue, 05 May 2020 19:00:00 PDT</pubDate>
			<itunes:episode>765</itunes:episode>
			<link>https://twit.tv/shows/security-now/episodes/765</link>
			<comments>https://twit.tv/shows/security-now/episodes/765</comments>
			<itunes:author>TWiT</itunes:author>
			<category>Technology</category>
			<category>Security</category>
			<itunes:explicit>clean</itunes:explicit>
			<itunes:subtitle>China Wants to Rebuild the Internet</itunes:subtitle>
			<itunes:keywords>Security Now, TWiT, steve gibson, Leo Laporte, China, Bruce Scheier, cybersecurity, DHS, Power Supplay, saltstack, devs</itunes:keywords>
			<description>
<![CDATA[
<p><img src="https://elroy.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg?itok=3OrHmJod" align="right" hspace="20" vspace="20" border="0" width="300" height="300" title="Security Now (Audio)" alt="Security Now (Audio)"/></p><p>China wants to rebuild the Internet.</p><ul><li>China's proposal to rebuild the internet is an authoritarian nightmare</li><li>Bruce Schneier on COVID-19 Contact Tracing Apps</li><li>Political Correctness hits cybersecurity</li><li>DHS's CISA says no to 3rd-party DoH</li><li>"POWER-SUPPLaY: Leaking Data from Air-Gapped Systems by Turning the Power-Supplies Into Speakers"</li><li>An authorization bypass in SaltStack</li><li>Adobe's Big Last Tuesday, Non-Patch Tuesday, Update</li><li>Google has announced its impending clean-up of the Chrome Web Store</li><li>Warning about RDP is not crying wolf</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-765-Notes.pdf">https://www.grc.com/sn/SN-765-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]>
			</description>
			<itunes:summary><![CDATA[
<p>China wants to rebuild the Internet.</p><ul><li>China's proposal to rebuild the internet is an authoritarian nightmare</li><li>Bruce Schneier on COVID-19 Contact Tracing Apps</li><li>Political Correctness hits cybersecurity</li><li>DHS's CISA says no to 3rd-party DoH</li><li>"POWER-SUPPLaY: Leaking Data from Air-Gapped Systems by Turning the Power-Supplies Into Speakers"</li><li>An authorization bypass in SaltStack</li><li>Adobe's Big Last Tuesday, Non-Patch Tuesday, Update</li><li>Google has announced its impending clean-up of the Chrome Web Store</li><li>Warning about RDP is not crying wolf</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-765-Notes.pdf">https://www.grc.com/sn/SN-765-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]></itunes:summary>
			<content:encoded><![CDATA[
<p>China wants to rebuild the Internet.</p><ul><li>China's proposal to rebuild the internet is an authoritarian nightmare</li><li>Bruce Schneier on COVID-19 Contact Tracing Apps</li><li>Political Correctness hits cybersecurity</li><li>DHS's CISA says no to 3rd-party DoH</li><li>"POWER-SUPPLaY: Leaking Data from Air-Gapped Systems by Turning the Power-Supplies Into Speakers"</li><li>An authorization bypass in SaltStack</li><li>Adobe's Big Last Tuesday, Non-Patch Tuesday, Update</li><li>Google has announced its impending clean-up of the Chrome Web Store</li><li>Warning about RDP is not crying wolf</li></ul><p>We invite you to read our show notes at <a href="https://www.grc.com/sn/SN-765-Notes.pdf">https://www.grc.com/sn/SN-765-Notes.pdf</a></p> 
<p><strong>Hosts:</strong> <a href="https://twit.tv/people/steve-gibson">Steve Gibson</a> and <a href="https://twit.tv/people/leo-laporte">Leo Laporte</a></p> 
<p>Download or subscribe to this show at <a href="https://twit.tv/shows/security-now">https://twit.tv/shows/security-now</a>.</p> 
<p>You can submit a question to Security Now! at the <a href="https://www.grc.com/feedback.htm" target="_blank">GRC Feedback Page</a>.</p> 
<p>For 16kbps versions, transcripts, and notes (including fixes), visit Steve's site: <a href="https://www.grc.com/securitynow.htm" target="_blank">grc.com</a>, also the home of the best disk maintenance and recovery utility ever written <a href="https://www.grc.com/sr/spinrite.htm" target="_blank">Spinrite 6</a>.</p>
<p><strong>Sponsors:</strong><ul>
<li><a href="http://LastPass.com/twit">LastPass.com/twit</a></li>
<li><a href="http://WWT.COM/TWIT">WWT.COM/TWIT</a></li>
<li><a href="http://barracuda.com/securitynow">barracuda.com/securitynow</a></li>
</ul></p>
]]></content:encoded>
				<itunes:image href="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_600x450/public/images/episodes/816417/hero/sn765_tn.jpg?itok=LEyssuQJ"/>
			<guid isPermaLink="false">https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0765/sn0765.mp3</guid>
			<itunes:duration>1:58:10</itunes:duration>
			<enclosure url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0765/sn0765.mp3" length="56835470" type="audio/mpeg"/>
			<media:content url="https://chtbl.com/track/E91833/cdn.twit.tv/audio/sn/sn0765/sn0765.mp3" fileSize="56835470" type="audio/mpeg" medium="audio">
				<media:title type="plain">SN 765: An Authoritarian Internet?</media:title>
				<media:description type="plain">China Wants to Rebuild the Internet</media:description>
				<media:keywords>Security Now, TWiT, steve gibson, Leo Laporte, China, Bruce Scheier, cybersecurity, DHS, Power Supplay, saltstack, devs</media:keywords>
				<media:thumbnail url="https://elroy.twit.tv/sites/default/files/styles/twit_slideshow_400x300/public/images/episodes/816417/hero/sn765_tn.jpg?itok=ClBHcm9E" width="400" height="225" />
				<media:rating scheme="urn:simple">nonadult</media:rating>
				<media:rating scheme="urn:v-chip">tv-g</media:rating>
				<media:category scheme="urn:iab:categories" label="Technology &amp; Computing">IAB19</media:category>		<media:credit role="anchor person">Steve Gibson</media:credit>
		<media:credit role="anchor person">Leo Laporte</media:credit>
			</media:content>
		</item>
	</channel>
</rss>`
	rss := Rss{}
	xml.Unmarshal([]byte(x), &rss)
	//fmt.Println(r)

    for _, item := range rss.Channel.Items {
        //fmt.Println(item)
        fmt.Println(item.Title)
        fmt.Println(item.Link)
    }


}


//https://stackoverflow.com/questions/34975837/parsing-rss-feed-in-go

/*
            <title>SN 765: An Authoritarian Internet? --Joy.Joy.Title</title>
            <itunes:title>An Authoritarian Internet? --YES.YES.itunes</itunes:title>
*/



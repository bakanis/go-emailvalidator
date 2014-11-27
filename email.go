/*
 * Copyright 2014 Branimir Karadzic. All rights reserved.
 * https://github.com/bkaradzic/go-emailvalidator/blob/master/LICENSE
 */

package emailvalidator

import (
	"regexp"
	"strings"
)

var (
	// List from:
	// github.com/abgoyal/disposable-email-domains
	disposableEmailDomains = []string{
		"e4ward.com",
		"mailexpire.com",
		"0815.ru",
		"0clickemail.com",
		"0wnd.net",
		"0wnd.org",
		"10minutemail.com",
		"20minutemail.com",
		"2prong.com",
		"9ox.net",
		"a-bc.net",
		"amiri.net",
		"amiriindustries.com",
		"anonymbox.com",
		"antichef.com",
		"antichef.net",
		"antispam.de",
		"baxomale.ht.cx",
		"beefmilk.com",
		"binkmail.com",
		"bio-muesli.net",
		"blogmyway.org",
		"bobmail.info",
		"bofthew.com",
		"brefmail.com",
		"bsnow.net",
		"bspamfree.org",
		"bugmenot.com",
		"casualdx.com",
		"centermail.com",
		"centermail.net",
		"chogmail.com",
		"choicemail1.com",
		"cool.fr.nf",
		"correo.blogos.net",
		"courriel.fr.nf",
		"courrieltemporaire.com",
		"cust.in",
		"dacoolest.com",
		"deadaddress.com",
		"deadspam.com",
		"despam.it",
		"despammed.com",
		"devnullmail.com",
		"dfgh.net",
		"digitalsanctuary.com",
		"disposableaddress.com",
		"disposemail.com",
		"dispostable.com",
		"dm.w3internet.co.uk",
		"dodgeit.com",
		"dodgit.com",
		"dontreg.com",
		"dontsendmespam.de",
		"dumpandjunk.com",
		"dumpyemail.com",
		"e4ward.com",
		"email60.com",
		"emailias.com",
		"emailmiser.com",
		"emailtemporario.com.br",
		"emailthe.net",
		"emailwarden.com",
		"etranquil.com",
		"explodemail.com",
		"fakeinbox.com",
		"fakeinformation.com",
		"fakemailz.com",
		"filzmail.com",
		"footard.com",
		"forgetmail.com",
		"garliclife.com",
		"getonemail.com",
		"gishpuppy.com",
		"gsrv.co.uk",
		"guerrillamail.biz",
		"guerrillamail.com",
		"guerrillamail.de",
		"guerrillamail.net",
		"guerrillamail.org",
		"guerrillamailblock.com",
		"hidemail.de",
		"hotpop.com",
		"ichimail.com",
		"iheartspam.org",
		"imails.info",
		"inboxclean.com",
		"inboxclean.org",
		"irish2me.com",
		"jetable.fr.nf",
		"jetable.org",
		"jnxjn.com",
		"junk1e.com",
		"kasmail.com",
		"kaspop.com",
		"killmail.com",
		"killmail.net",
		"klassmaster.com",
		"klzlk.com",
		"kulturbetrieb.info",
		"kurzepost.de",
		"lifebyfood.com",
		"link2mail.net",
		"lookugly.com",
		"lortemail.dk",
		"lr78.com",
		"mail.by",
		"mail333.com",
		"mail4trash.com",
		"mailbidon.com",
		"mailblocks.com",
		"mailcatch.com",
		"mailexpire.com",
		"mailfreeonline.com",
		"mailin8r.com",
		"mailinater.com",
		"mailinator.com",
		"mailinator.net",
		"mailinator2.com",
		"mailincubator.com",
		"mailme.lv",
		"mailmetrash.com",
		"mailmoat.com",
		"mailnator.com",
		"mailnesia.com",
		"mailnull.com",
		"mailquack.com",
		"mailshell.com",
		"mailsiphon.com",
		"mailslapping.com",
		"mailzilla.com",
		"mbx.cc",
		"mega.zik.dj",
		"meinspamschutz.de",
		"meltmail.com",
		"messagebeamer.de",
		"mintemail.com",
		"mmmmail.com",
		"moncourrier.fr.nf",
		"monemail.fr.nf",
		"monmail.fr.nf",
		"mt2009.com",
		"mycleaninbox.net",
		"mytrashmail.com",
		"neomailbox.com",
		"nepwk.com",
		"nervmich.net",
		"nervtmich.net",
		"netmails.net",
		"neverbox.com",
		"no-spam.hu",
		"no-spam.ws",
		"noclickemail.com",
		"nomail.xl.cx",
		"nomail2me.com",
		"nospam.ze.tc",
		"nospam4.us",
		"nospamfor.us",
		"nowmymail.com",
		"nurfuerspam.de",
		"objectmail.com",
		"oneoffemail.com",
		"oneoffmail.com",
		"onewaymail.com",
		"ordinaryamerican.net",
		"otherinbox.com",
		"owlpic.com",
		"pancakemail.com",
		"poofy.org",
		"pookmail.com",
		"privacy.net",
		"proxymail.eu",
		"punkass.com",
		"putthisinyourspamdatabase.com",
		"rcpt.at",
		"recode.me",
		"recursor.net",
		"rtrtr.com",
		"safersignup.de",
		"safetymail.info",
		"sharklasers.com",
		"shiftmail.com",
		"shitmail.me",
		"shortmail.net",
		"sibmail.com",
		"skeefmail.com",
		"slaskpost.se",
		"smellfear.com",
		"sneakemail.com",
		"sogetthis.com",
		"soodonims.com",
		"spamavert.com",
		"spambob.net",
		"spambob.org",
		"spambog.ru",
		"spambox.info",
		"spambox.us",
		"spamcannon.com",
		"spamcannon.net",
		"spamcon.org",
		"spamcorptastic.com",
		"spamcowboy.com",
		"spamcowboy.net",
		"spamcowboy.org",
		"spamday.com",
		"spamex.com",
		"spamfree24.com",
		"spamfree24.org",
		"spamgourmet.com",
		"spamgourmet.net",
		"spamgourmet.org",
		"spamherelots.com",
		"spamhereplease.com",
		"spamify.com",
		"spaml.com",
		"spaml.de",
		"spammotel.com",
		"spamobox.com",
		"spamslicer.com",
		"spamspot.com",
		"spamthis.co.uk",
		"spamtrail.com",
		"speed.1s.fr",
		"spoofmail.de",
		"suremail.info",
		"tempemail.net",
		"tempinbox.co.uk",
		"tempinbox.com",
		"tempomail.fr",
		"temporaryemail.net",
		"temporaryforwarding.com",
		"temporaryinbox.com",
		"thankyou2010.com",
		"thisisnotmyrealemail.com",
		"tmailinator.com",
		"tradermail.info",
		"trash-mail.at",
		"trash-mail.com",
		"trash-mail.de",
		"trash2009.com",
		"trashdevil.com",
		"trashemail.de",
		"trashmail.at",
		"trashmail.de",
		"trashmail.me",
		"trashmail.net",
		"trashmailer.com",
		"trashymail.com",
		"twinmail.de",
		"tyldd.com",
		"venompen.com",
		"walala.org",
		"wegwerfadresse.de",
		"wegwerfmail.de",
		"wegwerfmail.net",
		"wegwerfmail.org",
		"wh4f.org",
		"whopy.com",
		"whyspam.me",
		"winemaven.info",
		"wronghead.com",
		"wuzupmail.net",
		"wwwnew.eu",
		"xagloo.com",
		"xemaps.com",
		"xents.com",
		"xmaily.com",
		"xoxy.net",
		"yep.it",
		"yogamaven.com",
		"yopmail.com",
		"yopmail.fr",
		"yopmail.net",
		"yuurok.com",
		"zippymail.info",
		"objectmail.net",
	}
)

func IsDisposable(email string) bool {

	for _, suffix := range disposableEmailDomains {
		if strings.HasSuffix(email, suffix) {
			return true
		}
	}

	return false
}

// Reference:
// http://www.linuxjournal.com/article/9585?page=0,0
// http://www.regular-expressions.info/email.html
// http://fightingforalostcause.net/misc/2006/compare-email-regex.php

var (
	dotsCheck   = regexp.MustCompile("\\.\\.")
	domainCheck = regexp.MustCompile("^[A-Za-z0-9\\-\\.]+$")
	localCheck  = regexp.MustCompile("^(\\\\.|[A-Za-z0-9!#%&`_=\\/$'*+-?^{}|~.])+$")
	quotedCheck = regexp.MustCompile("^\"(\\\\\"|[^\"])+\"$")
)

func IsValid(email string) bool {

	at := strings.LastIndex(email, "@")

	if at == -1 {
		return false
	}

	local := email[0:at]
	localLen := len(local)
	domain := email[at+1:]
	domainLen := len(domain)

	if localLen < 1 || localLen > 64 {
		return false
	}

	if domainLen < 1 || domainLen > 255 {
		return false
	}

	if local[0] == '.' || local[localLen-1] == '.' {
		return false
	}

	if domain[0] == '.' || domain[domainLen-1] == '.' {
		return false
	}

	if dotsCheck.MatchString(local) || dotsCheck.MatchString(domain) {
		return false
	}

	if !domainCheck.MatchString(domain) {
		return false
	}

	local = strings.Replace(local, "\\\\", "", -1)
	if !localCheck.MatchString(local) && !quotedCheck.MatchString(local) {
		return false
	}

	return true
}

package services_test

import (
	"testing"

	"github.com/Haato3o/poogie/core/services"
)

func TestPatreonIsWebHookValid(t *testing.T) {
	secret := "c5KJMJCKK-jnKtgEgfQu0_oT-tJLCF9JCKuWLwrdAFp_xuyy7VbtcfG-yxhkNgkc"
	payload := `{"data":{"attributes":{"access_expires_at":null,"campaign_currency":"USD","campaign_lifetime_support_cents":12345,"campaign_pledge_amount_cents":150,"full_name":"Haato","is_follower":false,"last_charge_date":"2014-04-01T00:00:00.000+00:00","last_charge_status":"Paid","lifetime_support_cents":12345,"patron_status":"active_patron","pledge_amount_cents":150,"pledge_relationship_start":"2014-03-14T00:00:00.000+00:00"},"id":null,"relationships":{"address":{"data":null},"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}},"user":{"data":{"id":"31515460","type":"user"},"links":{"related":"https://www.patreon.com/api/user/31515460"}}},"type":"member"},"included":[{"attributes":{"avatar_photo_image_urls":{"default":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJ3Ijo2MjB9/5.png?token-time=1663459200&token-hash=KWnqHTt1vd21lu3G5BPAWcFw8mJscWdbG-ptXFFLKHg%3D","default_small":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJ3IjozNjB9/5.png?token-time=1663459200&token-hash=vnolK17BPVRse57u30AYwQ-7BD76hNNqFNCsRBg47wU%3D","original":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJxIjoxMDAsIndlYnAiOjB9/5.png?token-time=1663459200&token-hash=sdVcXZjXD2LZxj-w80_nUIoFRDdor7fiGTdcbIguvQM%3D","thumbnail":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJoIjozNjAsInciOjM2MH0%3D/5.png?token-time=1663459200&token-hash=Ev9d4wEqc0GYZ14HFFtRHwfLtHjXCcTNGSVdXAKls6A%3D","thumbnail_large":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJoIjoxMDgwLCJ3IjoxMDgwfQ%3D%3D/5.png?token-time=1663459200&token-hash=NcL9xd_alA1z4WdDVr6WqLWO78nLENkfZsvn1xBkTCE%3D"},"avatar_photo_url":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/cd129a5be815406e80af84753e78d7f3/eyJ3IjoyMDB9/5.png?token-time=2145916800&token-hash=zAIobDSf5kKro5TJbKJeqxIKRAcC4xGF4yP5LBQFTVQ%3D","campaign_pledge_sum":50805,"cover_photo_url":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/4.png?token-time=1664841600&token-hash=dOtJoC5kWBj4T9yjZM4YGwxN2rodkgWIWbKSwJZ7CNs%3D","cover_photo_url_sizes":{"large":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/4.png?token-time=1664841600&token-hash=dOtJoC5kWBj4T9yjZM4YGwxN2rodkgWIWbKSwJZ7CNs%3D","medium":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3Ijo5NjAsIndlIjoxfQ%3D%3D/4.png?token-time=1664841600&token-hash=8VooO_mMbdlxt6T9KkHIjzSIV_kHOBFtpTmHu4-8gTI%3D","small":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3Ijo2MjAsIndlIjoxfQ%3D%3D/4.png?token-time=1664841600&token-hash=-gpsNcGtvy-rsm2RSDWGK2NulbyjC5yjxv4wF3Z1ano%3D"},"created_at":"2020-03-06T14:22:37.000+00:00","creation_count":13,"creation_name":"HunterPie","currency":"USD","discord_server_id":"678286768046342147","display_patron_goals":false,"earnings_visibility":"public","image_small_url":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/4.png?token-time=1664841600&token-hash=dOtJoC5kWBj4T9yjZM4YGwxN2rodkgWIWbKSwJZ7CNs%3D","image_url":"https://c10.patreonusercontent.com/4/patreon-media/p/campaign/4077998/b059781e39a441178efbcd132034b439/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/4.png?token-time=1664841600&token-hash=dOtJoC5kWBj4T9yjZM4YGwxN2rodkgWIWbKSwJZ7CNs%3D","is_charge_upfront":true,"is_charged_immediately":true,"is_monthly":true,"is_nsfw":false,"is_plural":false,"main_video_embed":null,"main_video_url":null,"name":"Haato","one_liner":null,"outstanding_payment_amount_cents":0,"patron_count":348,"pay_per_name":"month","pledge_sum":264832,"pledge_sum_currency":"BRL","pledge_url":"/join/HunterPie","published_at":"2020-03-14T16:08:09.000+00:00","summary":"HunterPie is a simple, easy to use and modern overlay with Discord Rich Presence support for Monster Hunter: World.<br>","thanks_embed":null,"thanks_msg":null,"thanks_video_url":null,"url":"https://www.patreon.com/HunterPie"},"id":"4077998","relationships":{"creator":{"data":{"id":"31515460","type":"user"},"links":{"related":"https://www.patreon.com/api/user/31515460"}},"goals":{"data":[]},"rewards":{"data":[{"id":"-1","type":"reward"},{"id":"0","type":"reward"},{"id":"4734490","type":"reward"},{"id":"4734489","type":"reward"},{"id":"4734488","type":"reward"},{"id":"8176687","type":"reward"}]}},"type":"campaign"},{"attributes":{"about":null,"apple_id":null,"can_see_nsfw":true,"created":"2020-03-06T14:14:30.000+00:00","default_country_code":null,"discord_id":"183067754784358400","email":"matheus_mrp@hotmail.com","facebook":null,"facebook_id":null,"first_name":"Haato","full_name":"Haato","gender":0,"google_id":"106525201352814938830","has_password":true,"image_url":"https://c8.patreon.com/2/200/31515460","is_deleted":false,"is_email_verified":true,"is_nuked":false,"is_suspended":false,"last_name":"","patron_currency":"BRL","social_connections":{"deviantart":null,"discord":{"scopes":["guilds","guilds.join","identify"],"url":null,"user_id":"183067754784358400"},"facebook":null,"google":null,"instagram":null,"reddit":null,"spotify":null,"twitch":null,"twitter":null,"vimeo":null,"youtube":null},"thumb_url":"https://c8.patreon.com/2/200/31515460","twitch":null,"twitter":null,"url":"https://www.patreon.com/HunterPie","vanity":"HunterPie","youtube":null},"id":"31515460","relationships":{"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}}},"type":"user"},{"attributes":{"amount":0,"amount_cents":0,"created_at":null,"description":"Everyone","patron_currency":"BRL","remaining":0,"requires_shipping":false,"url":null,"user_limit":null},"id":"-1","type":"reward"},{"attributes":{"amount":1,"amount_cents":1,"created_at":null,"description":"Patrons Only","patron_currency":"BRL","remaining":0,"requires_shipping":false,"url":null,"user_limit":null},"id":"0","type":"reward"},{"attributes":{"amount":100,"amount_cents":100,"created_at":"2020-03-14T21:21:14.589+00:00","currency":"USD","description":"You'll be helping HunterPie development!<br><br><strong>What you'll get:<br></strong><ul><li><strong>Tier 1</strong> Role on Discord</li><li>Access to alpha releases</li></ul>","discord_role_ids":["932640302135656518","685555720615100515"],"edited_at":"2022-06-22T00:24:30.741+00:00","image_url":"https://c10.patreonusercontent.com/4/patreon-media/p/reward/4734490/57d0ab5db4694afc93c2f5b81848eed1/eyJ3Ijo0MDB9/3.png?token-time=2145916800&token-hash=TxsqQS5DAnWRf0C6FjcTe4Q1fTrUGl_c0JB6C0Xn2sU%3D","patron_amount_cents":750,"patron_count":350,"patron_currency":"BRL","post_count":0,"published":true,"published_at":"2022-06-22T00:24:30.720+00:00","remaining":null,"requires_shipping":false,"title":"Low Rank","unpublished_at":null,"url":"/join/HunterPie/checkout?rid=4734490","user_limit":null,"welcome_message":null,"welcome_message_unsafe":null,"welcome_video_embed":null,"welcome_video_url":null},"id":"4734490","relationships":{"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}}},"type":"reward"},{"attributes":{"amount":500,"amount_cents":500,"created_at":"2020-03-14T21:21:14.589+00:00","currency":"USD","description":"You'll be helping HunterPie development!<br><br><strong>What you'll get:</strong><br><ul><li><strong>Tier 2</strong> Role on Discord</li><li>Access to alpha releases</li></ul>","discord_role_ids":["932640041518366800","685555720615100515"],"edited_at":"2022-06-22T00:25:23.103+00:00","image_url":"https://c10.patreonusercontent.com/4/patreon-media/p/reward/4734489/35425807749d446cac219a810ffb9236/eyJ3Ijo0MDB9/2.png?token-time=2145916800&token-hash=xyEwwAPi66nIcBxt9q2C-yqqQpHPf85khlTDz1MdY7o%3D","patron_amount_cents":3000,"patron_count":34,"patron_currency":"BRL","post_count":0,"published":true,"published_at":"2022-06-22T00:25:23.066+00:00","remaining":null,"requires_shipping":false,"title":"High Rank","unpublished_at":null,"url":"/join/HunterPie/checkout?rid=4734489","user_limit":null,"welcome_message":null,"welcome_message_unsafe":null,"welcome_video_embed":null,"welcome_video_url":null},"id":"4734489","relationships":{"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}}},"type":"reward"},{"attributes":{"amount":1000,"amount_cents":1000,"created_at":"2020-03-14T21:21:14.581+00:00","currency":"USD","description":"You'll be helping HunterPie development!<br><br><strong>What you'll get:<br></strong><ul><li><strong>Tier 3</strong> Role on Discord</li><li>Access to alpha releases</li></ul>","discord_role_ids":["932639855614251028","685555720615100515"],"edited_at":"2022-06-22T00:25:36.781+00:00","image_url":"https://c10.patreonusercontent.com/4/patreon-media/p/reward/4734488/d48fe9d4d54f45d6b557c306bedf719f/eyJ3Ijo0MDB9/2.png?token-time=2145916800&token-hash=8Y7TUKuWGcsLFv50ojhlHyyAk181BNNmQGG9lSLuTTU%3D","patron_amount_cents":5750,"patron_count":13,"patron_currency":"BRL","post_count":0,"published":true,"published_at":"2022-06-22T00:25:36.758+00:00","remaining":null,"requires_shipping":false,"title":"Tempered","unpublished_at":null,"url":"/join/HunterPie/checkout?rid=4734488","user_limit":null,"welcome_message":null,"welcome_message_unsafe":null,"welcome_video_embed":null,"welcome_video_url":null},"id":"4734488","relationships":{"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}}},"type":"reward"},{"attributes":{"amount":2000,"amount_cents":2000,"created_at":"2022-01-17T14:22:08.461+00:00","currency":"USD","description":"You'll be supporting HunterPie's development!<br><br><strong>What you'll get:<br></strong><ul><li><strong>Tier 4 </strong>Role on Discord</li><li>Access to alpha releases</li></ul>","discord_role_ids":["932639681462554696","685555720615100515"],"edited_at":"2022-06-22T00:25:51.582+00:00","image_url":"https://c10.patreonusercontent.com/4/patreon-media/p/reward/8176687/f2abcacc094742458c0f3795042c415c/eyJ3Ijo0MDB9/1.png?token-time=2145916800&token-hash=SCVYJW3ylmxlxqbMptumnnPRHtzRe6EezH9g8sx1kYw%3D","patron_amount_cents":11500,"patron_count":4,"patron_currency":"BRL","post_count":0,"published":true,"published_at":"2022-06-22T00:25:51.564+00:00","remaining":null,"requires_shipping":false,"title":"Arch Tempered","unpublished_at":null,"url":"/join/HunterPie/checkout?rid=8176687","user_limit":null,"welcome_message":null,"welcome_message_unsafe":null,"welcome_video_embed":null,"welcome_video_url":null},"id":"8176687","relationships":{"campaign":{"data":{"id":"4077998","type":"campaign"},"links":{"related":"https://www.patreon.com/api/campaigns/4077998"}}},"type":"reward"}],"links":{"self":"https://www.patreon.com/api/members/None"}}`
	service := services.NewPatreonService(secret)

	isValid := service.IsWebhookValid("bcf743afe6c39e860abdf8a9926006b5", []byte(payload))

	if !isValid {
		t.Errorf("got %t, expected %t", isValid, true)
	}
}

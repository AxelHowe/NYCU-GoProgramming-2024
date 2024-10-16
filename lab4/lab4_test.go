package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculator(t *testing.T) {

	var tests = []struct {
		query         string
		expectedWrong bool
		exp           string
		res           string
	}{
		{"/?op=add&num1=58270622&num2=58178886", false, "58270622 + 58178886", "116449508"},
		{"/?op=gcd&num1=6627265&num2=77682628", false, "GCD(6627265, 77682628)", "1"},
		{"/?op=gcd&num1=2520&num2=217728", false, "GCD(2520, 217728)", "504"},
		{"/?op=div&num1=73161081&num2=239", false, "73161081 / 239", "306113"},
		{"/?op=lcm&num1=2178870&num2=2264715", false, "LCM(2178870, 2264715)", "5575728330"},
		{"/?op=div&num1=5740658&num2=880", false, "5740658 / 880", "6523"},
		{"/?op=sub&num1=10357427&num2=61397368", false, "10357427 - 61397368", "-51039941"},
		{"/?op=add&num1=59164249&num2=99217347", false, "59164249 + 99217347", "158381596"},
		{"/?op=gcd&num1=4998952&num2=2164925", false, "GCD(4998952, 2164925)", "623"},
		{"/?op=mul&num1=29154570&num2=21713233", false, "29154570 * 21713233", "633039971424810"},
		{"/?op=gcd&num1=1526503&num2=5907671", false, "GCD(1526503, 5907671)", "803"},
		{"/?op=add&num1=72935478&num2=43105088", false, "72935478 + 43105088", "116040566"},
		{"/?op=div&num1=97504070&num2=721", false, "97504070 / 721", "135234"},
		{"/?op=gcd&num1=2398600&num2=1966584", false, "GCD(2398600, 1966584)", "536"},
		{"/?op=add&num1=98722180&num2=59536118", false, "98722180 + 59536118", "158258298"},
		{"/?op=lcm&num1=368544&num2=135828", false, "LCM(368544, 135828)", "379231776"},
		{"/?op=add&num1=76706194&num2=30858274", false, "76706194 + 30858274", "107564468"},
		{"/?op=add&num1=6764883&num2=65616199", false, "6764883 + 65616199", "72381082"},
		{"/?op=lcm&num1=413381&num2=843404", false, "LCM(413381, 843404)", "2220682732"},
		{"/?op=mul&num1=35080682&num2=42037792", false, "35080682 * 42037792", "1474714413134144"},
		{"/?op=lcm&num1=64014&num2=987000", false, "LCM(64014, 987000)", "224049000"},
		{"/?op=lcm&num1=4446409&num2=4393191", false, "LCM(4446409, 4393191)", "43312470069"},
		{"/?op=lcm&num1=2610312&num2=2962038", false, "LCM(2610312, 2962038)", "21125255016"},
		{"/?op=sub&num1=61516052&num2=20669659", false, "61516052 - 20669659", "40846393"},
		{"/?op=sub&num1=7597860&num2=99194032", false, "7597860 - 99194032", "-91596172"},
		{"/?op=lcm&num1=328245&num2=659176", false, "LCM(328245, 659176)", "2738876280"},
		{"/?op=sub&num1=29578428&num2=61491493", false, "29578428 - 61491493", "-31913065"},
		{"/?op=lcm&num1=223979&num2=78155", false, "LCM(223979, 78155)", "357246505"},
		{"/?op=gcd&num1=801504&num2=879648", false, "GCD(801504, 879648)", "1056"},
		{"/?op=sub&num1=46808093&num2=74920999", false, "46808093 - 74920999", "-28112906"},
		{"/?op=add&num1=8224503&num2=74546964", false, "8224503 + 74546964", "82771467"},
		{"/?op=gcd&num1=413646&num2=299052", false, "GCD(413646, 299052)", "426"},
		{"/?op=gcd&num1=35074&num2=15613", false, "GCD(35074, 15613)", "13"},
		{"/?op=gcd&num1=976472&num2=1927716", false, "GCD(976472, 1927716)", "1484"},
		{"/?op=mul&num1=68425179&num2=18965701", false, "68425179 * 18965701", "1297731485785479"},
		{"/?op=div&num1=17482447&num2=606", false, "17482447 / 606", "28848"},
		{"/?op=lcm&num1=4260438&num2=1355988", false, "LCM(4260438, 1355988)", "3331662516"},
		{"/?op=add&num1=99661159&num2=13844561", false, "99661159 + 13844561", "113505720"},
		{"/?op=gcd&num1=3403050&num2=8400672", false, "GCD(3403050, 8400672)", "19446"},
		{"/?op=gcd&num1=1231271&num2=909074", false, "GCD(1231271, 909074)", "509"},
		{"/?op=add&num1=45690037&num2=45611061", false, "45690037 + 45611061", "91301098"},
		{"/?op=gcd&num1=245220&num2=2272908", false, "GCD(245220, 2272908)", "804"},
		{"/?op=gcd&num1=6627265&num2=77682628", false, "GCD(6627265, 77682628)", "1"},
		{"/?op=lcm&num1=7621691&num2=594859", false, "LCM(7621691, 594859)", "4717826729"},
		{"/?op=gcd&num1=4125568&num2=5504988", false, "GCD(4125568, 5504988)", "668"},
		{"/?op=mul&num1=80536133&num2=94163376", false, "80536133 * 94163376", "7583554173265008"},
		{"/?op=lcm&num1=2284100&num2=1958320", false, "LCM(2284100, 1958320)", "2457691600"},
		{"/?op=lcm&num1=654108&num2=2760030", false, "LCM(654108, 2760030)", "3306515940"},
		{"/?op=gcd&num1=77616&num2=2207205", false, "GCD(77616, 2207205)", "4851"},
		{"/?op=div&num1=20596979&num2=724", false, "20596979 / 724", "28448"},
		{"/?op=sub&num1=62709732&num2=1701229", false, "62709732 - 1701229", "61008503"},
		{"/?op=sub&num1=5831827&num2=71624192", false, "5831827 - 71624192", "-65792365"},
		{"/?op=sub&num1=54951166&num2=24597974", false, "54951166 - 24597974", "30353192"},
		{"/?op=sub&num1=44087479&num2=4128157", false, "44087479 - 4128157", "39959322"},
		{"/?op=lcm&num1=620686&num2=587934", false, "LCM(620686, 587934)", "2050125858"},
		{"/?op=div&num1=50783652&num2=920", false, "50783652 / 920", "55199"},
		{"/?op=lcm&num1=3889424&num2=4441360", false, "LCM(3889424, 4441360)", "2003053360"},
		{"/?op=gcd&num1=1488880&num2=1463720", false, "GCD(1488880, 1463720)", "1480"},
		{"/?op=lcm&num1=2215026&num2=3924558", false, "LCM(2215026, 3924558)", "14634676782"},
		{"/?op=gcd&num1=323570&num2=2905214", false, "GCD(323570, 2905214)", "494"},
		{"/?op=sub&num1=9049910&num2=16822009", false, "9049910 - 16822009", "-7772099"},
		{"/?op=add&num1=68852091&num2=61068057", false, "68852091 + 61068057", "129920148"},
		{"/?op=mul&num1=68267108&num2=90345968", false, "68267108 * 90345968", "6167657954820544"},
		{"/?op=gcd&num1=8070&num2=221550", false, "GCD(8070, 221550)", "30"},
		{"/?op=lcm&num1=1278648&num2=1282088", false, "LCM(1278648, 1282088)", "4765521096"},
		{"/?op=lcm&num1=2750280&num2=5094855", false, "LCM(2750280, 5094855)", "21724461720"},
		{"/?op=lcm&num1=1225440&num2=4522806", false, "LCM(1225440, 4522806)", "8321963040"},
		{"/?op=gcd&num1=1521432&num2=5988774", false, "GCD(1521432, 5988774)", "7458"},
		{"/?op=sub&num1=16338218&num2=4672490", false, "16338218 - 4672490", "11665728"},
		{"/?op=add&num1=6627265&num2=77682628", false, "6627265 + 77682628", "84309893"},
		{"/?op=lcm&num1=4807151&num2=5083519", false, "LCM(4807151, 5083519)", "32539605119"},
		{"/?op=add&num1=32608227&num2=74206780", false, "32608227 + 74206780", "106815007"},
		{"/?op=div&num1=77995456&num2=508", false, "77995456 / 508", "153534"},
		{"/?op=add&num1=83632880&num2=77216163", false, "83632880 + 77216163", "160849043"},
		{"/?op=mul&num1=60731823&num2=21184468", false, "60731823 * 21184468", "1286571360925164"},
		{"/?op=div&num1=75019105&num2=51", false, "75019105 / 51", "1470962"},
		{"/?op=div&num1=4351043&num2=908", false, "4351043 / 908", "4791"},
		{"/?op=div&num1=85082565&num2=895", false, "85082565 / 895", "95064"},
		{"/?op=lcm&num1=866592&num2=2636064", false, "LCM(866592, 2636064)", "2643972192"},
		{"/?op=add&num1=89083762&num2=2507411", false, "89083762 + 2507411", "91591173"},
		{"/?op=lcm&num1=3998008&num2=6742624", false, "LCM(3998008, 6742624)", "15528263072"},
		{"/?op=sub&num1=75050538&num2=50957188", false, "75050538 - 50957188", "24093350"},
		{"/?op=div&num1=60387026&num2=379", false, "60387026 / 379", "159332"},
		{"/?op=lcm&num1=7407720&num2=5522445", false, "LCM(7407720, 5522445)", "15948821160"},
		{"/?op=gcd&num1=2216718&num2=2338182", false, "GCD(2216718, 2338182)", "30366"},
		{"/?op=gcd&num1=2353880&num2=3042448", false, "GCD(2353880, 3042448)", "664"},
		{"/?op=lcm&num1=111690&num2=2743290", false, "LCM(111690, 2743290)", "200260170"},
		{"/?op=lcm&num1=5501174&num2=4048268", false, "LCM(5501174, 4048268)", "1793382724"},
		{"/?op=gcd&num1=7696649&num2=9219834", false, "GCD(7696649, 9219834)", "961"},
		{"/?op=lcm&num1=709464&num2=5499207", false, "LCM(709464, 5499207)", "4531346568"},
		{"/?op=div&num1=73544141&num2=28", false, "73544141 / 28", "2626576"},
		{"/?op=lcm&num1=5210280&num2=6414716", false, "LCM(5210280, 6414716)", "23670302040"},
		{"/?op=div&num1=63967382&num2=486", false, "63967382 / 486", "131620"},
		{"/?op=div&num1=17375336&num2=67", false, "17375336 / 67", "259333"},
		{"/?op=mul&num1=98750550&num2=14664916", false, "98750550 * 14664916", "1448168520703800"},
		{"/?op=gcd&num1=357544&num2=65164", false, "GCD(357544, 65164)", "44"},
		{"/?op=lcm&num1=352352&num2=4189640", false, "LCM(352352, 4189640)", "2027785760"},
		{"/?op=lcm&num1=553926&num2=446512", false, "LCM(553926, 446512)", "2875983792"},
		{"/?op=sub&num1=11716613&num2=25785145", false, "11716613 - 25785145", "-14068532"},
		{"/?op=lcm&num1=8045480&num2=7934175", false, "LCM(8045480, 7934175)", "64806341400"},
		{"/?op=sub&num1=52997871&num2=9604950", false, "52997871 - 9604950", "43392921"},
		{"/?op=sub&num1=45798915&num2=17510154", false, "45798915 - 17510154", "28288761"},
		{"/?op=gcd", true, "", ""},
		{"/?op=po&num1=4547151&num2=5087519", true, "", ""},
		{"/?op=div&cat=meow&num2=5087519", true, "", ""},
		{"/?op=lcm&num1=&num2=", true, "", ""},
		{"/?op=div&num1=8547403num2=932541", true, "", ""},
		{"/?op=&num1=&num2=", true, "", ""},
		{"/?op=div&num1=&num2=0", true, "", ""},
		{"/?op=mul&num1=seven&num2=eleven", true, "", ""},
		{"/?op=mul&num1=8&num2=seven", true, "", ""},
		{"/?op=sub&num1=58270622", true, "", ""},
		{"/?op=gcd&a=58178886&b=123", true, "", ""},
		{"/?op=div&num1=58270622&num2=0", true, "", ""},
		{"/?skill=starbust-stream&debuff=cant-logout", true, "", ""},
		{"/?name=jerry&enemy=tom", true, "", ""},
		{"/", true, "", ""},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.query, nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		Calculator(res, req)
		if res.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Code)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !test.expectedWrong {
			assert.Equal(t, test.exp, doc.Find("body .container .expression").Text())
			assert.Equal(t, test.res, doc.Find("body .container .result").Text())
		} else {
			assert.Equal(t, "Error!", doc.Find("body .container h1").Text())
		}
	}
}
package jma

var (
	Soya                      = Area{Code: "011000", Name: "宗谷地方", NameEn: "Soya", Kana: "", OfficeName: "稚内地方気象台", Parent: "010100", Children: []string{"011000"}}
	KamikawaRumoi             = Area{Code: "012000", Name: "上川・留萌地方", NameEn: "Kamikawa Rumoi", Kana: "", OfficeName: "旭川地方気象台", Parent: "010100", Children: []string{"012010", "012020"}}
	AbashiriKitamiMombetsu    = Area{Code: "013000", Name: "網走・北見・紋別地方", NameEn: "Abashiri Kitami Mombetsu", Kana: "", OfficeName: "網走地方気象台", Parent: "010100", Children: []string{"013010", "013020", "013030"}}
	Tokachi                   = Area{Code: "014030", Name: "十勝地方", NameEn: "Tokachi", Kana: "", OfficeName: "帯広測候所", Parent: "010100", Children: []string{"014030"}}
	KushiroNemuro             = Area{Code: "014100", Name: "釧路・根室地方", NameEn: "Kushiro Nemuro", Kana: "", OfficeName: "釧路地方気象台", Parent: "010100", Children: []string{"014010", "014020"}}
	IburiHidaka               = Area{Code: "015000", Name: "胆振・日高地方", NameEn: "Iburi Hidaka", Kana: "", OfficeName: "室蘭地方気象台", Parent: "010100", Children: []string{"015010", "015020"}}
	IshikariSorachiShiribeshi = Area{Code: "016000", Name: "石狩・空知・後志地方", NameEn: "Ishikari Sorachi Shiribeshi", Kana: "", OfficeName: "札幌管区気象台", Parent: "010100", Children: []string{"016010", "016020", "016030"}}
	OshimaHiyama              = Area{Code: "017000", Name: "渡島・檜山地方", NameEn: "Oshima Hiyama", Kana: "", OfficeName: "函館地方気象台", Parent: "010100", Children: []string{"017010", "017020"}}
	Aomori                    = Area{Code: "020000", Name: "青森県", NameEn: "Aomori", Kana: "", OfficeName: "青森地方気象台", Parent: "010200", Children: []string{"020010", "020020", "020030"}}
	Iwate                     = Area{Code: "030000", Name: "岩手県", NameEn: "Iwate", Kana: "", OfficeName: "盛岡地方気象台", Parent: "010200", Children: []string{"030010", "030020", "030030"}}
	Miyagi                    = Area{Code: "040000", Name: "宮城県", NameEn: "Miyagi", Kana: "", OfficeName: "仙台管区気象台", Parent: "010200", Children: []string{"040010", "040020"}}
	Akita                     = Area{Code: "050000", Name: "秋田県", NameEn: "Akita", Kana: "", OfficeName: "秋田地方気象台", Parent: "010200", Children: []string{"050010", "050020"}}
	Yamagata                  = Area{Code: "060000", Name: "山形県", NameEn: "Yamagata", Kana: "", OfficeName: "山形地方気象台", Parent: "010200", Children: []string{"060010", "060020", "060030", "060040"}}
	Fukushima                 = Area{Code: "070000", Name: "福島県", NameEn: "Fukushima", Kana: "", OfficeName: "福島地方気象台", Parent: "010200", Children: []string{"070010", "070020", "070030"}}
	Ibaraki                   = Area{Code: "080000", Name: "茨城県", NameEn: "Ibaraki", Kana: "", OfficeName: "水戸地方気象台", Parent: "010300", Children: []string{"080010", "080020"}}
	Tochigi                   = Area{Code: "090000", Name: "栃木県", NameEn: "Tochigi", Kana: "", OfficeName: "宇都宮地方気象台", Parent: "010300", Children: []string{"090010", "090020"}}
	Gunma                     = Area{Code: "100000", Name: "群馬県", NameEn: "Gunma", Kana: "", OfficeName: "前橋地方気象台", Parent: "010300", Children: []string{"100010", "100020"}}
	Saitama                   = Area{Code: "110000", Name: "埼玉県", NameEn: "Saitama", Kana: "", OfficeName: "熊谷地方気象台", Parent: "010300", Children: []string{"110010", "110020", "110030"}}
	Chiba                     = Area{Code: "120000", Name: "千葉県", NameEn: "Chiba", Kana: "", OfficeName: "銚子地方気象台", Parent: "010300", Children: []string{"120010", "120020", "120030"}}
	Tokyo                     = Area{Code: "130000", Name: "東京都", NameEn: "Tokyo", Kana: "", OfficeName: "気象庁", Parent: "010300", Children: []string{"130010", "130020", "130030", "130040"}}
	Kanagawa                  = Area{Code: "140000", Name: "神奈川県", NameEn: "Kanagawa", Kana: "", OfficeName: "横浜地方気象台", Parent: "010300", Children: []string{"140010", "140020"}}
	Niigata                   = Area{Code: "150000", Name: "新潟県", NameEn: "Niigata", Kana: "", OfficeName: "新潟地方気象台", Parent: "010500", Children: []string{"150010", "150020", "150030", "150040"}}
	Toyama                    = Area{Code: "160000", Name: "富山県", NameEn: "Toyama", Kana: "", OfficeName: "富山地方気象台", Parent: "010500", Children: []string{"160010", "160020"}}
	Ishikawa                  = Area{Code: "170000", Name: "石川県", NameEn: "Ishikawa", Kana: "", OfficeName: "金沢地方気象台", Parent: "010500", Children: []string{"170010", "170020"}}
	Fukui                     = Area{Code: "180000", Name: "福井県", NameEn: "Fukui", Kana: "", OfficeName: "福井地方気象台", Parent: "010500", Children: []string{"180010", "180020"}}
	Yamanashi                 = Area{Code: "190000", Name: "山梨県", NameEn: "Yamanashi", Kana: "", OfficeName: "甲府地方気象台", Parent: "010300", Children: []string{"190010", "190020"}}
	Nagano                    = Area{Code: "200000", Name: "長野県", NameEn: "Nagano", Kana: "", OfficeName: "長野地方気象台", Parent: "010300", Children: []string{"200010", "200020", "200030"}}
	Gifu                      = Area{Code: "210000", Name: "岐阜県", NameEn: "Gifu", Kana: "", OfficeName: "岐阜地方気象台", Parent: "010400", Children: []string{"210010", "210020"}}
	Shizuoka                  = Area{Code: "220000", Name: "静岡県", NameEn: "Shizuoka", Kana: "", OfficeName: "静岡地方気象台", Parent: "010400", Children: []string{"220010", "220020", "220030", "220040"}}
	Aichi                     = Area{Code: "230000", Name: "愛知県", NameEn: "Aichi", Kana: "", OfficeName: "名古屋地方気象台", Parent: "010400", Children: []string{"230010", "230020"}}
	Mie                       = Area{Code: "240000", Name: "三重県", NameEn: "Mie", Kana: "", OfficeName: "津地方気象台", Parent: "010400", Children: []string{"240010", "240020"}}
	Shiga                     = Area{Code: "250000", Name: "滋賀県", NameEn: "Shiga", Kana: "", OfficeName: "彦根地方気象台", Parent: "010600", Children: []string{"250010", "250020"}}
	Kyoto                     = Area{Code: "260000", Name: "京都府", NameEn: "Kyoto", Kana: "", OfficeName: "京都地方気象台", Parent: "010600", Children: []string{"260010", "260020"}}
	Osaka                     = Area{Code: "270000", Name: "大阪府", NameEn: "Osaka", Kana: "", OfficeName: "大阪管区気象台", Parent: "010600", Children: []string{"270000"}}
	Hyogo                     = Area{Code: "280000", Name: "兵庫県", NameEn: "Hyogo", Kana: "", OfficeName: "神戸地方気象台", Parent: "010600", Children: []string{"280010", "280020"}}
	Nara                      = Area{Code: "290000", Name: "奈良県", NameEn: "Nara", Kana: "", OfficeName: "奈良地方気象台", Parent: "010600", Children: []string{"290010", "290020"}}
	Wakayama                  = Area{Code: "300000", Name: "和歌山県", NameEn: "Wakayama", Kana: "", OfficeName: "和歌山地方気象台", Parent: "010600", Children: []string{"300010", "300020"}}
	Tottori                   = Area{Code: "310000", Name: "鳥取県", NameEn: "Tottori", Kana: "", OfficeName: "鳥取地方気象台", Parent: "010700", Children: []string{"310010", "310020"}}
	Shimane                   = Area{Code: "320000", Name: "島根県", NameEn: "Shimane", Kana: "", OfficeName: "松江地方気象台", Parent: "010700", Children: []string{"320010", "320020", "320030"}}
	Okayama                   = Area{Code: "330000", Name: "岡山県", NameEn: "Okayama", Kana: "", OfficeName: "岡山地方気象台", Parent: "010700", Children: []string{"330010", "330020"}}
	Hiroshima                 = Area{Code: "340000", Name: "広島県", NameEn: "Hiroshima", Kana: "", OfficeName: "広島地方気象台", Parent: "010700", Children: []string{"340010", "340020"}}
	Yamaguchi                 = Area{Code: "350000", Name: "山口県", NameEn: "Yamaguchi", Kana: "", OfficeName: "下関地方気象台", Parent: "010900", Children: []string{"350010", "350020", "350030", "350040"}}
	Tokushima                 = Area{Code: "360000", Name: "徳島県", NameEn: "Tokushima", Kana: "", OfficeName: "徳島地方気象台", Parent: "010800", Children: []string{"360010", "360020"}}
	Kagawa                    = Area{Code: "370000", Name: "香川県", NameEn: "Kagawa", Kana: "", OfficeName: "高松地方気象台", Parent: "010800", Children: []string{"370000"}}
	Ehime                     = Area{Code: "380000", Name: "愛媛県", NameEn: "Ehime", Kana: "", OfficeName: "松山地方気象台", Parent: "010800", Children: []string{"380010", "380020", "380030"}}
	Kochi                     = Area{Code: "390000", Name: "高知県", NameEn: "Kochi", Kana: "", OfficeName: "高知地方気象台", Parent: "010800", Children: []string{"390010", "390020", "390030"}}
	Fukuoka                   = Area{Code: "400000", Name: "福岡県", NameEn: "Fukuoka", Kana: "", OfficeName: "福岡管区気象台", Parent: "010900", Children: []string{"400010", "400020", "400030", "400040"}}
	Saga                      = Area{Code: "410000", Name: "佐賀県", NameEn: "Saga", Kana: "", OfficeName: "佐賀地方気象台", Parent: "010900", Children: []string{"410010", "410020"}}
	Nagasaki                  = Area{Code: "420000", Name: "長崎県", NameEn: "Nagasaki", Kana: "", OfficeName: "長崎地方気象台", Parent: "010900", Children: []string{"420010", "420020", "420030", "420040"}}
	Kumamoto                  = Area{Code: "430000", Name: "熊本県", NameEn: "Kumamoto", Kana: "", OfficeName: "熊本地方気象台", Parent: "010900", Children: []string{"430010", "430020", "430030", "430040"}}
	Oita                      = Area{Code: "440000", Name: "大分県", NameEn: "Oita", Kana: "", OfficeName: "大分地方気象台", Parent: "010900", Children: []string{"440010", "440020", "440030", "440040"}}
	Miyazaki                  = Area{Code: "450000", Name: "宮崎県", NameEn: "Miyazaki", Kana: "", OfficeName: "宮崎地方気象台", Parent: "011000", Children: []string{"450010", "450020", "450030", "450040"}}
	Amami                     = Area{Code: "460040", Name: "奄美地方", NameEn: "Amami", Kana: "", OfficeName: "名瀬測候所", Parent: "011000", Children: []string{"460040"}}
	KagoshimaExcludingAmami   = Area{Code: "460100", Name: "鹿児島県（奄美地方除く）", NameEn: "Kagoshima (Excluding Amami)", Kana: "", OfficeName: "鹿児島地方気象台", Parent: "011000", Children: []string{"460010", "460020", "460030"}}
	OkinawaMainIsland         = Area{Code: "471000", Name: "沖縄本島地方", NameEn: "Okinawa Main Island", Kana: "", OfficeName: "沖縄気象台", Parent: "011100", Children: []string{"471010", "471020", "471030"}}
	Daitojima                 = Area{Code: "472000", Name: "大東島地方", NameEn: "Daitojima", Kana: "", OfficeName: "南大東島地方気象台", Parent: "011100", Children: []string{"472000"}}
	Miyakojima                = Area{Code: "473000", Name: "宮古島地方", NameEn: "Miyakojima", Kana: "", OfficeName: "宮古島地方気象台", Parent: "011100", Children: []string{"473000"}}
	Yaeyama                   = Area{Code: "474000", Name: "八重山地方", NameEn: "Yaeyama", Kana: "", OfficeName: "石垣島地方気象台", Parent: "011100", Children: []string{"474010", "474020"}}
)

var Offices = []Area{
	Soya,
	KamikawaRumoi,
	AbashiriKitamiMombetsu,
	Tokachi,
	KushiroNemuro,
	IburiHidaka,
	IshikariSorachiShiribeshi,
	OshimaHiyama,
	Aomori,
	Iwate,
	Miyagi,
	Akita,
	Yamagata,
	Fukushima,
	Ibaraki,
	Tochigi,
	Gunma,
	Saitama,
	Chiba,
	Tokyo,
	Kanagawa,
	Niigata,
	Toyama,
	Ishikawa,
	Fukui,
	Yamanashi,
	Nagano,
	Gifu,
	Shizuoka,
	Aichi,
	Mie,
	Shiga,
	Kyoto,
	Osaka,
	Hyogo,
	Nara,
	Wakayama,
	Tottori,
	Shimane,
	Okayama,
	Hiroshima,
	Yamaguchi,
	Tokushima,
	Kagawa,
	Ehime,
	Kochi,
	Fukuoka,
	Saga,
	Nagasaki,
	Kumamoto,
	Oita,
	Miyazaki,
	Amami,
	KagoshimaExcludingAmami,
	OkinawaMainIsland,
	Daitojima,
	Miyakojima,
	Yaeyama,
}

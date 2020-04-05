package types

import "fmt"

const (
	LocaleAfZa  Locale = "af-ZA"
	LocaleAmEt  Locale = "am-ET"
	LocaleArAe  Locale = "ar-AE"
	LocaleArBh  Locale = "ar-BH"
	LocaleArDz  Locale = "ar-DZ"
	LocaleArEg  Locale = "ar-EG"
	LocaleArIq  Locale = "ar-IQ"
	LocaleArJo  Locale = "ar-JO"
	LocaleArKw  Locale = "ar-KW"
	LocaleArLb  Locale = "ar-LB"
	LocaleArLy  Locale = "ar-LY"
	LocaleArMa  Locale = "ar-MA"
	LocaleArnCl Locale = "arn-CL"
	LocaleArOm  Locale = "ar-OM"
	LocaleArQa  Locale = "ar-QA"
	LocaleArSa  Locale = "ar-SA"
	LocaleArSy  Locale = "ar-SY"
	LocaleArTn  Locale = "ar-TN"
	LocaleArYe  Locale = "ar-YE"
	LocaleAsIn  Locale = "as-IN"
	LocaleAzAz  Locale = "az-AZ"
	LocaleBaRu  Locale = "ba-RU"
	LocaleBeBy  Locale = "be-BY"
	LocaleBgBg  Locale = "bg-BG"
	LocaleBnBd  Locale = "bn-BD"
	LocaleBnIn  Locale = "bn-IN"
	LocaleBoCn  Locale = "bo-CN"
	LocaleBrFr  Locale = "br-FR"
	LocaleBsBa  Locale = "bs-BA"
	LocaleCaEs  Locale = "ca-ES"
	LocaleCoFr  Locale = "co-FR"
	LocaleCsCz  Locale = "cs-CZ"
	LocaleCyGb  Locale = "cy-GB"
	LocaleDaDk  Locale = "da-DK"
	LocaleDeAt  Locale = "de-AT"
	LocaleDeCh  Locale = "de-CH"
	LocaleDeDe  Locale = "de-DE"
	LocaleDeLi  Locale = "de-LI"
	LocaleDeLu  Locale = "de-LU"
	LocaleDsbDe Locale = "dsb-DE"
	LocaleDvMv  Locale = "dv-MV"
	LocaleElGr  Locale = "el-GR"
	LocaleEn029 Locale = "en-029"
	LocaleEnAu  Locale = "en-AU"
	LocaleEnBz  Locale = "en-BZ"
	LocaleEnCa  Locale = "en-CA"
	LocaleEnGb  Locale = "en-GB"
	LocaleEnIe  Locale = "en-IE"
	LocaleEnIn  Locale = "en-IN"
	LocaleEnJm  Locale = "en-JM"
	LocaleEnMy  Locale = "en-MY"
	LocaleEnNz  Locale = "en-NZ"
	LocaleEnPh  Locale = "en-PH"
	LocaleEnSg  Locale = "en-SG"
	LocaleEnTt  Locale = "en-TT"
	LocaleEnUs  Locale = "en-US"
	LocaleEnZa  Locale = "en-ZA"
	LocaleEnZw  Locale = "en-ZW"
	LocaleEsAr  Locale = "es-AR"
	LocaleEsBo  Locale = "es-BO"
	LocaleEsCl  Locale = "es-CL"
	LocaleEsCo  Locale = "es-CO"
	LocaleEsCr  Locale = "es-CR"
	LocaleEsDo  Locale = "es-DO"
	LocaleEsEc  Locale = "es-EC"
	LocaleEsEs  Locale = "es-ES"
	LocaleEsGt  Locale = "es-GT"
	LocaleEsHn  Locale = "es-HN"
	LocaleEsMx  Locale = "es-MX"
	LocaleEsNi  Locale = "es-NI"
	LocaleEsPa  Locale = "es-PA"
	LocaleEsPe  Locale = "es-PE"
	LocaleEsPr  Locale = "es-PR"
	LocaleEsPy  Locale = "es-PY"
	LocaleEsSv  Locale = "es-SV"
	LocaleEsUs  Locale = "es-US"
	LocaleEsUy  Locale = "es-UY"
	LocaleEsVe  Locale = "es-VE"
	LocaleEtEe  Locale = "et-EE"
	LocaleEuEs  Locale = "eu-ES"
	LocaleFaIr  Locale = "fa-IR"
	LocaleFiFi  Locale = "fi-FI"
	LocaleFilPh Locale = "fil-PH"
	LocaleFoFo  Locale = "fo-FO"
	LocaleFrBe  Locale = "fr-BE"
	LocaleFrCa  Locale = "fr-CA"
	LocaleFrCh  Locale = "fr-CH"
	LocaleFrFr  Locale = "fr-FR"
	LocaleFrLu  Locale = "fr-LU"
	LocaleFrMc  Locale = "fr-MC"
	LocaleFyNl  Locale = "fy-NL"
	LocaleGaIe  Locale = "ga-IE"
	LocaleGdGb  Locale = "gd-GB"
	LocaleGlEs  Locale = "gl-ES"
	LocaleGswFr Locale = "gsw-FR"
	LocaleGuIn  Locale = "gu-IN"
	LocaleHaNg  Locale = "ha-NG"
	LocaleHeIl  Locale = "he-IL"
	LocaleHiIn  Locale = "hi-IN"
	LocaleHrBa  Locale = "hr-BA"
	LocaleHrHr  Locale = "hr-HR"
	LocaleHsbDe Locale = "hsb-DE"
	LocaleHuHu  Locale = "hu-HU"
	LocaleHyAm  Locale = "hy-AM"
	LocaleIdId  Locale = "id-ID"
	LocaleIgNg  Locale = "ig-NG"
	LocaleIiCn  Locale = "ii-CN"
	LocaleIsIs  Locale = "is-IS"
	LocaleItCh  Locale = "it-CH"
	LocaleItIt  Locale = "it-IT"
	LocaleIuCa  Locale = "iu-CA"
	LocaleJaJp  Locale = "ja-JP"
	LocaleKaGe  Locale = "ka-GE"
	LocaleKkKz  Locale = "kk-KZ"
	LocaleKlGl  Locale = "kl-GL"
	LocaleKmKh  Locale = "km-KH"
	LocaleKnIn  Locale = "kn-IN"
	LocaleKokIn Locale = "kok-IN"
	LocaleKoKr  Locale = "ko-KR"
	LocaleKyKg  Locale = "ky-KG"
	LocaleLbLu  Locale = "lb-LU"
	LocaleLoLa  Locale = "lo-LA"
	LocaleLtLt  Locale = "lt-LT"
	LocaleLvLv  Locale = "lv-LV"
	LocaleMiNz  Locale = "mi-NZ"
	LocaleMkMk  Locale = "mk-MK"
	LocaleMlIn  Locale = "ml-IN"
	LocaleMnMn  Locale = "mn-MN"
	LocaleMnCn  Locale = "mn-CN"
	LocaleMohCa Locale = "moh-CA"
	LocaleMrIn  Locale = "mr-IN"
	LocaleMsBn  Locale = "ms-BN"
	LocaleMsMy  Locale = "ms-MY"
	LocaleMtMt  Locale = "mt-MT"
	LocaleNbNo  Locale = "nb-NO"
	LocaleNeNp  Locale = "ne-NP"
	LocaleNlBe  Locale = "nl-BE"
	LocaleNlNl  Locale = "nl-NL"
	LocaleNnNo  Locale = "nn-NO"
	LocaleNsoZa Locale = "nso-ZA"
	LocaleOcFr  Locale = "oc-FR"
	LocaleOrIn  Locale = "or-IN"
	LocalePaIn  Locale = "pa-IN"
	LocalePlPl  Locale = "pl-PL"
	LocalePrsAf Locale = "prs-AF"
	LocalePsAf  Locale = "ps-AF"
	LocalePtBr  Locale = "pt-BR"
	LocalePtPt  Locale = "pt-PT"
	LocaleQutGt Locale = "qut-GT"
	LocaleQuzBo Locale = "quz-BO"
	LocaleQuzEc Locale = "quz-EC"
	LocaleQuzPe Locale = "quz-PE"
	LocaleRmCh  Locale = "rm-CH"
	LocaleRoRo  Locale = "ro-RO"
	LocaleRuRu  Locale = "ru-RU"
	LocaleRwRw  Locale = "rw-RW"
	LocaleSahRu Locale = "sah-RU"
	LocaleSaIn  Locale = "sa-IN"
	LocaleSeFi  Locale = "se-FI"
	LocaleSeNo  Locale = "se-NO"
	LocaleSeSe  Locale = "se-SE"
	LocaleSiLk  Locale = "si-LK"
	LocaleSkSk  Locale = "sk-SK"
	LocaleSlSi  Locale = "sl-SI"
	LocaleSmaNo Locale = "sma-NO"
	LocaleSmaSe Locale = "sma-SE"
	LocaleSmjNo Locale = "smj-NO"
	LocaleSmjSe Locale = "smj-SE"
	LocaleSmnFi Locale = "smn-FI"
	LocaleSmsFi Locale = "sms-FI"
	LocaleSqAl  Locale = "sq-AL"
	LocaleSrBa  Locale = "sr-BA"
	LocaleSrCs  Locale = "sr-CS"
	LocaleSrMe  Locale = "sr-ME"
	LocaleSrRs  Locale = "sr-RS"
	LocaleSvFi  Locale = "sv-FI"
	LocaleSvSe  Locale = "sv-SE"
	LocaleSwKe  Locale = "sw-KE"
	LocaleSyrSy Locale = "syr-SY"
	LocaleTaIn  Locale = "ta-IN"
	LocaleTeIn  Locale = "te-IN"
	LocaleTgTj  Locale = "tg-TJ"
	LocaleThTh  Locale = "th-TH"
	LocaleTkTm  Locale = "tk-TM"
	LocaleTnZa  Locale = "tn-ZA"
	LocaleTrTr  Locale = "tr-TR"
	LocaleTtRu  Locale = "tt-RU"
	LocaleTzmDz Locale = "tzm-DZ"
	LocaleUgCn  Locale = "ug-CN"
	LocaleUkUa  Locale = "uk-UA"
	LocaleUrPk  Locale = "ur-PK"
	LocaleUzUz  Locale = "uz-UZ"
	LocaleViVn  Locale = "vi-VN"
	LocaleWoSn  Locale = "wo-SN"
	LocaleXhZa  Locale = "xh-ZA"
	LocaleYoNg  Locale = "yo-NG"
	LocaleZhCn  Locale = "zh-CN"
	LocaleZhHk  Locale = "zh-HK"
	LocaleZhMo  Locale = "zh-MO"
	LocaleZhSg  Locale = "zh-SG"
	LocaleZhTw  Locale = "zh-TW"
	LocaleZuZa  Locale = "zu-ZA"
)

type Locale string

func (e Locale) Validate() Fault {
	if e == LocaleAfZa {
		return nil
	}
	if e == LocaleAmEt {
		return nil
	}
	if e == LocaleArAe {
		return nil
	}
	if e == LocaleArBh {
		return nil
	}
	if e == LocaleArDz {
		return nil
	}
	if e == LocaleArEg {
		return nil
	}
	if e == LocaleArIq {
		return nil
	}
	if e == LocaleArJo {
		return nil
	}
	if e == LocaleArKw {
		return nil
	}
	if e == LocaleArLb {
		return nil
	}
	if e == LocaleArLy {
		return nil
	}
	if e == LocaleArMa {
		return nil
	}
	if e == LocaleArnCl {
		return nil
	}
	if e == LocaleArOm {
		return nil
	}
	if e == LocaleArQa {
		return nil
	}
	if e == LocaleArSa {
		return nil
	}
	if e == LocaleArSy {
		return nil
	}
	if e == LocaleArTn {
		return nil
	}
	if e == LocaleArYe {
		return nil
	}
	if e == LocaleAsIn {
		return nil
	}
	if e == LocaleAzAz {
		return nil
	}
	if e == LocaleBaRu {
		return nil
	}
	if e == LocaleBeBy {
		return nil
	}
	if e == LocaleBgBg {
		return nil
	}
	if e == LocaleBnBd {
		return nil
	}
	if e == LocaleBnIn {
		return nil
	}
	if e == LocaleBoCn {
		return nil
	}
	if e == LocaleBrFr {
		return nil
	}
	if e == LocaleBsBa {
		return nil
	}
	if e == LocaleCaEs {
		return nil
	}
	if e == LocaleCoFr {
		return nil
	}
	if e == LocaleCsCz {
		return nil
	}
	if e == LocaleCyGb {
		return nil
	}
	if e == LocaleDaDk {
		return nil
	}
	if e == LocaleDeAt {
		return nil
	}
	if e == LocaleDeCh {
		return nil
	}
	if e == LocaleDeDe {
		return nil
	}
	if e == LocaleDeLi {
		return nil
	}
	if e == LocaleDeLu {
		return nil
	}
	if e == LocaleDsbDe {
		return nil
	}
	if e == LocaleDvMv {
		return nil
	}
	if e == LocaleElGr {
		return nil
	}
	if e == LocaleEn029 {
		return nil
	}
	if e == LocaleEnAu {
		return nil
	}
	if e == LocaleEnBz {
		return nil
	}
	if e == LocaleEnCa {
		return nil
	}
	if e == LocaleEnGb {
		return nil
	}
	if e == LocaleEnIe {
		return nil
	}
	if e == LocaleEnIn {
		return nil
	}
	if e == LocaleEnJm {
		return nil
	}
	if e == LocaleEnMy {
		return nil
	}
	if e == LocaleEnNz {
		return nil
	}
	if e == LocaleEnPh {
		return nil
	}
	if e == LocaleEnSg {
		return nil
	}
	if e == LocaleEnTt {
		return nil
	}
	if e == LocaleEnUs {
		return nil
	}
	if e == LocaleEnZa {
		return nil
	}
	if e == LocaleEnZw {
		return nil
	}
	if e == LocaleEsAr {
		return nil
	}
	if e == LocaleEsBo {
		return nil
	}
	if e == LocaleEsCl {
		return nil
	}
	if e == LocaleEsCo {
		return nil
	}
	if e == LocaleEsCr {
		return nil
	}
	if e == LocaleEsDo {
		return nil
	}
	if e == LocaleEsEc {
		return nil
	}
	if e == LocaleEsEs {
		return nil
	}
	if e == LocaleEsGt {
		return nil
	}
	if e == LocaleEsHn {
		return nil
	}
	if e == LocaleEsMx {
		return nil
	}
	if e == LocaleEsNi {
		return nil
	}
	if e == LocaleEsPa {
		return nil
	}
	if e == LocaleEsPe {
		return nil
	}
	if e == LocaleEsPr {
		return nil
	}
	if e == LocaleEsPy {
		return nil
	}
	if e == LocaleEsSv {
		return nil
	}
	if e == LocaleEsUs {
		return nil
	}
	if e == LocaleEsUy {
		return nil
	}
	if e == LocaleEsVe {
		return nil
	}
	if e == LocaleEtEe {
		return nil
	}
	if e == LocaleEuEs {
		return nil
	}
	if e == LocaleFaIr {
		return nil
	}
	if e == LocaleFiFi {
		return nil
	}
	if e == LocaleFilPh {
		return nil
	}
	if e == LocaleFoFo {
		return nil
	}
	if e == LocaleFrBe {
		return nil
	}
	if e == LocaleFrCa {
		return nil
	}
	if e == LocaleFrCh {
		return nil
	}
	if e == LocaleFrFr {
		return nil
	}
	if e == LocaleFrLu {
		return nil
	}
	if e == LocaleFrMc {
		return nil
	}
	if e == LocaleFyNl {
		return nil
	}
	if e == LocaleGaIe {
		return nil
	}
	if e == LocaleGdGb {
		return nil
	}
	if e == LocaleGlEs {
		return nil
	}
	if e == LocaleGswFr {
		return nil
	}
	if e == LocaleGuIn {
		return nil
	}
	if e == LocaleHaNg {
		return nil
	}
	if e == LocaleHeIl {
		return nil
	}
	if e == LocaleHiIn {
		return nil
	}
	if e == LocaleHrBa {
		return nil
	}
	if e == LocaleHrHr {
		return nil
	}
	if e == LocaleHsbDe {
		return nil
	}
	if e == LocaleHuHu {
		return nil
	}
	if e == LocaleHyAm {
		return nil
	}
	if e == LocaleIdId {
		return nil
	}
	if e == LocaleIgNg {
		return nil
	}
	if e == LocaleIiCn {
		return nil
	}
	if e == LocaleIsIs {
		return nil
	}
	if e == LocaleItCh {
		return nil
	}
	if e == LocaleItIt {
		return nil
	}
	if e == LocaleIuCa {
		return nil
	}
	if e == LocaleJaJp {
		return nil
	}
	if e == LocaleKaGe {
		return nil
	}
	if e == LocaleKkKz {
		return nil
	}
	if e == LocaleKlGl {
		return nil
	}
	if e == LocaleKmKh {
		return nil
	}
	if e == LocaleKnIn {
		return nil
	}
	if e == LocaleKokIn {
		return nil
	}
	if e == LocaleKoKr {
		return nil
	}
	if e == LocaleKyKg {
		return nil
	}
	if e == LocaleLbLu {
		return nil
	}
	if e == LocaleLoLa {
		return nil
	}
	if e == LocaleLtLt {
		return nil
	}
	if e == LocaleLvLv {
		return nil
	}
	if e == LocaleMiNz {
		return nil
	}
	if e == LocaleMkMk {
		return nil
	}
	if e == LocaleMlIn {
		return nil
	}
	if e == LocaleMnMn {
		return nil
	}
	if e == LocaleMnCn {
		return nil
	}
	if e == LocaleMohCa {
		return nil
	}
	if e == LocaleMrIn {
		return nil
	}
	if e == LocaleMsBn {
		return nil
	}
	if e == LocaleMsMy {
		return nil
	}
	if e == LocaleMtMt {
		return nil
	}
	if e == LocaleNbNo {
		return nil
	}
	if e == LocaleNeNp {
		return nil
	}
	if e == LocaleNlBe {
		return nil
	}
	if e == LocaleNlNl {
		return nil
	}
	if e == LocaleNnNo {
		return nil
	}
	if e == LocaleNsoZa {
		return nil
	}
	if e == LocaleOcFr {
		return nil
	}
	if e == LocaleOrIn {
		return nil
	}
	if e == LocalePaIn {
		return nil
	}
	if e == LocalePlPl {
		return nil
	}
	if e == LocalePrsAf {
		return nil
	}
	if e == LocalePsAf {
		return nil
	}
	if e == LocalePtBr {
		return nil
	}
	if e == LocalePtPt {
		return nil
	}
	if e == LocaleQutGt {
		return nil
	}
	if e == LocaleQuzBo {
		return nil
	}
	if e == LocaleQuzEc {
		return nil
	}
	if e == LocaleQuzPe {
		return nil
	}
	if e == LocaleRmCh {
		return nil
	}
	if e == LocaleRoRo {
		return nil
	}
	if e == LocaleRuRu {
		return nil
	}
	if e == LocaleRwRw {
		return nil
	}
	if e == LocaleSahRu {
		return nil
	}
	if e == LocaleSaIn {
		return nil
	}
	if e == LocaleSeFi {
		return nil
	}
	if e == LocaleSeNo {
		return nil
	}
	if e == LocaleSeSe {
		return nil
	}
	if e == LocaleSiLk {
		return nil
	}
	if e == LocaleSkSk {
		return nil
	}
	if e == LocaleSlSi {
		return nil
	}
	if e == LocaleSmaNo {
		return nil
	}
	if e == LocaleSmaSe {
		return nil
	}
	if e == LocaleSmjNo {
		return nil
	}
	if e == LocaleSmjSe {
		return nil
	}
	if e == LocaleSmnFi {
		return nil
	}
	if e == LocaleSmsFi {
		return nil
	}
	if e == LocaleSqAl {
		return nil
	}
	if e == LocaleSrBa {
		return nil
	}
	if e == LocaleSrCs {
		return nil
	}
	if e == LocaleSrMe {
		return nil
	}
	if e == LocaleSrRs {
		return nil
	}
	if e == LocaleSvFi {
		return nil
	}
	if e == LocaleSvSe {
		return nil
	}
	if e == LocaleSwKe {
		return nil
	}
	if e == LocaleSyrSy {
		return nil
	}
	if e == LocaleTaIn {
		return nil
	}
	if e == LocaleTeIn {
		return nil
	}
	if e == LocaleTgTj {
		return nil
	}
	if e == LocaleThTh {
		return nil
	}
	if e == LocaleTkTm {
		return nil
	}
	if e == LocaleTnZa {
		return nil
	}
	if e == LocaleTrTr {
		return nil
	}
	if e == LocaleTtRu {
		return nil
	}
	if e == LocaleTzmDz {
		return nil
	}
	if e == LocaleUgCn {
		return nil
	}
	if e == LocaleUkUa {
		return nil
	}
	if e == LocaleUrPk {
		return nil
	}
	if e == LocaleUzUz {
		return nil
	}
	if e == LocaleViVn {
		return nil
	}
	if e == LocaleWoSn {
		return nil
	}
	if e == LocaleXhZa {
		return nil
	}
	if e == LocaleYoNg {
		return nil
	}
	if e == LocaleZhCn {
		return nil
	}
	if e == LocaleZhHk {
		return nil
	}
	if e == LocaleZhMo {
		return nil
	}
	if e == LocaleZhSg {
		return nil
	}
	if e == LocaleZhTw {
		return nil
	}
	if e == LocaleZuZa {
		return nil
	}
	return NewFault("illegal_locale", fmt.Sprintf("illegal value '%s' for enum type Locale", e))
}

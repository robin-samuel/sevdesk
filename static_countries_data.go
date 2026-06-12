package sevdesk

import "strings"

// Country references for every [StaticCountry] sevdesk knows about.
// Drop directly into a [*Ref] field, e.g. [ListContactsParams.Country] or
// [ContactAddress.Country].
//
// Naming follows ISO 3166-1 alpha-2 codes: CountryDE, CountryUS, CountryGB, …
// Where sevdesk has multiple records for the same code (e.g. England + United
// Kingdom + Great Britain all under "gb"), the lowest-ID entry takes the bare
// name and the others are suffixed with their English name.
var (
	// CountryAD — Andorra.
	CountryAD = &Ref{ID: 1418, ObjectName: ObjectStaticCountry}
	// CountryAE — United Arab Emirates (Vereinigte Arabische Emirate).
	CountryAE = &Ref{ID: 44, ObjectName: ObjectStaticCountry}
	// CountryAF — Afghanistan.
	CountryAF = &Ref{ID: 4, ObjectName: ObjectStaticCountry}
	// CountryAG — Antigua and Barbuda (Antigua und Barbuda).
	CountryAG = &Ref{ID: 1449, ObjectName: ObjectStaticCountry}
	// CountryAI — Anguilla.
	CountryAI = &Ref{ID: 1467, ObjectName: ObjectStaticCountry}
	// CountryAL — Albania (Albanien).
	CountryAL = &Ref{ID: 1347, ObjectName: ObjectStaticCountry}
	// CountryAM — Armenia (Armenien).
	CountryAM = &Ref{ID: 1348, ObjectName: ObjectStaticCountry}
	// CountryAO — Angola.
	CountryAO = &Ref{ID: 1466, ObjectName: ObjectStaticCountry}
	// CountryAQ — Antarctica (Antarktis).
	CountryAQ = &Ref{ID: 1470, ObjectName: ObjectStaticCountry}
	// CountryAR — Argentina (Argentinien).
	CountryAR = &Ref{ID: 5, ObjectName: ObjectStaticCountry}
	// CountryAS — American Samoa (Amerikanisch-Samoa).
	CountryAS = &Ref{ID: 1469, ObjectName: ObjectStaticCountry}
	// CountryAT — Austria (Österreich).
	CountryAT = &Ref{ID: 3, ObjectName: ObjectStaticCountry}
	// CountryAU — Australia (Australien).
	CountryAU = &Ref{ID: 38, ObjectName: ObjectStaticCountry}
	// CountryAW — Aruba.
	CountryAW = &Ref{ID: 1465, ObjectName: ObjectStaticCountry}
	// CountryAX — Åland Islands (Ålandinseln).
	CountryAX = &Ref{ID: 1468, ObjectName: ObjectStaticCountry}
	// CountryAZ — Azerbaijan (Aserbaidschan).
	CountryAZ = &Ref{ID: 33, ObjectName: ObjectStaticCountry}
	// CountryBA — Bosnia and Herzegovina (Bosnien und Herzegowina).
	CountryBA = &Ref{ID: 1396, ObjectName: ObjectStaticCountry}
	// CountryBB — Barbados.
	CountryBB = &Ref{ID: 1445, ObjectName: ObjectStaticCountry}
	// CountryBD — Bangladesh (Bangladesch).
	CountryBD = &Ref{ID: 1440, ObjectName: ObjectStaticCountry}
	// CountryBE — Belgium (Belgien).
	CountryBE = &Ref{ID: 6, ObjectName: ObjectStaticCountry}
	// CountryBF — Burkina Faso.
	CountryBF = &Ref{ID: 1473, ObjectName: ObjectStaticCountry}
	// CountryBG — Bulgaria (Bulgarien).
	CountryBG = &Ref{ID: 7, ObjectName: ObjectStaticCountry}
	// CountryBH — Bahrain.
	CountryBH = &Ref{ID: 1427, ObjectName: ObjectStaticCountry}
	// CountryBI — Burundi.
	CountryBI = &Ref{ID: 1404, ObjectName: ObjectStaticCountry}
	// CountryBJ — Benin.
	CountryBJ = &Ref{ID: 1472, ObjectName: ObjectStaticCountry}
	// CountryBL — Saint Barthélemy (Saint-Barthélemy).
	CountryBL = &Ref{ID: 1474, ObjectName: ObjectStaticCountry}
	// CountryBM — Bermuda.
	CountryBM = &Ref{ID: 1476, ObjectName: ObjectStaticCountry}
	// CountryBN — Brunei.
	CountryBN = &Ref{ID: 1479, ObjectName: ObjectStaticCountry}
	// CountryBO — Bolivia (Bolivien).
	CountryBO = &Ref{ID: 1477, ObjectName: ObjectStaticCountry}
	// CountryBQ — Caribbean Netherlands (Karibische Niederlande).
	CountryBQ = &Ref{ID: 1478, ObjectName: ObjectStaticCountry}
	// CountryBR — Brazil (Brasilien).
	CountryBR = &Ref{ID: 73, ObjectName: ObjectStaticCountry}
	// CountryBS — Bahamas.
	CountryBS = &Ref{ID: 1426, ObjectName: ObjectStaticCountry}
	// CountryBT — Bhutan.
	CountryBT = &Ref{ID: 1350, ObjectName: ObjectStaticCountry}
	// CountryBV — Bouvet Island (Bouvetinsel).
	CountryBV = &Ref{ID: 1480, ObjectName: ObjectStaticCountry}
	// CountryBW — Botswana.
	CountryBW = &Ref{ID: 1481, ObjectName: ObjectStaticCountry}
	// CountryBY — Republic of Belarus (Republik Belarus).
	CountryBY = &Ref{ID: 76, ObjectName: ObjectStaticCountry}
	// CountryBZ — Belize.
	CountryBZ = &Ref{ID: 46, ObjectName: ObjectStaticCountry}
	// CountryCA — Canada (Kanada).
	CountryCA = &Ref{ID: 1423, ObjectName: ObjectStaticCountry}
	// CountryCC — Cocos (Keeling) Islands (Kokosinseln).
	CountryCC = &Ref{ID: 1483, ObjectName: ObjectStaticCountry}
	// CountryCD — DR Congo (Kongo (Dem. Rep.)).
	CountryCD = &Ref{ID: 1484, ObjectName: ObjectStaticCountry}
	// CountryCF — Central African Republic (Zentralafrikanische Republik).
	CountryCF = &Ref{ID: 1482, ObjectName: ObjectStaticCountry}
	// CountryCG — Republic of the Congo (Kongo).
	CountryCG = &Ref{ID: 1485, ObjectName: ObjectStaticCountry}
	// CountryCH — Switzerland (Schweiz).
	CountryCH = &Ref{ID: 2, ObjectName: ObjectStaticCountry}
	// CountryCI — Ivory Coast (Elfenbeinküste).
	CountryCI = &Ref{ID: 1381, ObjectName: ObjectStaticCountry}
	// CountryCK — Cook Islands (Cookinseln).
	CountryCK = &Ref{ID: 1486, ObjectName: ObjectStaticCountry}
	// CountryCL — Chile.
	CountryCL = &Ref{ID: 49, ObjectName: ObjectStaticCountry}
	// CountryCM — Cameroon (Kamerun).
	CountryCM = &Ref{ID: 1400, ObjectName: ObjectStaticCountry}
	// CountryCN — People's Republic of China (Volksrepublik China).
	CountryCN = &Ref{ID: 60, ObjectName: ObjectStaticCountry}
	// CountryCO — Colombia (Kolumbien).
	CountryCO = &Ref{ID: 1403, ObjectName: ObjectStaticCountry}
	// CountryCR — Costa Rica.
	CountryCR = &Ref{ID: 1421, ObjectName: ObjectStaticCountry}
	// CountryCU — Cuba (Kuba).
	CountryCU = &Ref{ID: 1488, ObjectName: ObjectStaticCountry}
	// CountryCV — Cape Verde (Kap Verde).
	CountryCV = &Ref{ID: 1487, ObjectName: ObjectStaticCountry}
	// CountryCW — Curaçao.
	CountryCW = &Ref{ID: 1455, ObjectName: ObjectStaticCountry}
	// CountryCX — Christmas Island (Weihnachtsinsel).
	CountryCX = &Ref{ID: 1489, ObjectName: ObjectStaticCountry}
	// CountryCY — Cyprus (Zypern).
	CountryCY = &Ref{ID: 1490, ObjectName: ObjectStaticCountry}
	// CountryCZ — Czech Republic (Tschechische Republik).
	CountryCZ = &Ref{ID: 30, ObjectName: ObjectStaticCountry}
	// CountryDE — Germany (Deutschland).
	CountryDE = &Ref{ID: 1, ObjectName: ObjectStaticCountry}
	// CountryDJ — Djibouti (Dschibuti).
	CountryDJ = &Ref{ID: 1405, ObjectName: ObjectStaticCountry}
	// CountryDK — Denmark (Dänemark).
	CountryDK = &Ref{ID: 8, ObjectName: ObjectStaticCountry}
	// CountryDM — Dominica.
	CountryDM = &Ref{ID: 1491, ObjectName: ObjectStaticCountry}
	// CountryDO — Dominican Republic (Dominikanische Republik).
	CountryDO = &Ref{ID: 1447, ObjectName: ObjectStaticCountry}
	// CountryDU — Dubai.
	CountryDU = &Ref{ID: 67, ObjectName: ObjectStaticCountry}
	// CountryDZ — Algeria (Algerien).
	CountryDZ = &Ref{ID: 1458, ObjectName: ObjectStaticCountry}
	// CountryEC — Ecuador.
	CountryEC = &Ref{ID: 1492, ObjectName: ObjectStaticCountry}
	// CountryEE — Estonia (Estland).
	CountryEE = &Ref{ID: 1390, ObjectName: ObjectStaticCountry}
	// CountryEG — Egypt (Ägypten).
	CountryEG = &Ref{ID: 1383, ObjectName: ObjectStaticCountry}
	// CountryEH — Western Sahara (Westsahara).
	CountryEH = &Ref{ID: 1493, ObjectName: ObjectStaticCountry}
	// CountryER — Eritrea.
	CountryER = &Ref{ID: 1406, ObjectName: ObjectStaticCountry}
	// CountryES — Spain (Spanien).
	CountryES = &Ref{ID: 29, ObjectName: ObjectStaticCountry}
	// CountryET — Ethiopia (Äthiopien).
	CountryET = &Ref{ID: 1399, ObjectName: ObjectStaticCountry}
	// CountryFI — Finland (Finnland).
	CountryFI = &Ref{ID: 10, ObjectName: ObjectStaticCountry}
	// CountryFJ — Fiji (Fidschi).
	CountryFJ = &Ref{ID: 1494, ObjectName: ObjectStaticCountry}
	// CountryFK — Falkland Islands (Falklandinseln).
	CountryFK = &Ref{ID: 1495, ObjectName: ObjectStaticCountry}
	// CountryFM — Micronesia (Mikronesien).
	CountryFM = &Ref{ID: 1497, ObjectName: ObjectStaticCountry}
	// CountryFO — Faroe Islands (Färöer-Inseln).
	CountryFO = &Ref{ID: 1496, ObjectName: ObjectStaticCountry}
	// CountryFR — France (Frankreich).
	CountryFR = &Ref{ID: 11, ObjectName: ObjectStaticCountry}
	// CountryGA — Gabon (Gabun).
	CountryGA = &Ref{ID: 1498, ObjectName: ObjectStaticCountry}
	// CountryGB — Great Britain (Großbritannien).
	CountryGB = &Ref{ID: 9, ObjectName: ObjectStaticCountry}
	// CountryGBEngland — England.
	CountryGBEngland = &Ref{ID: 74, ObjectName: ObjectStaticCountry}
	// CountryGBUnitedKingdom — United Kingdom (Vereinigtes Königreich).
	CountryGBUnitedKingdom = &Ref{ID: 77, ObjectName: ObjectStaticCountry}
	// CountryGD — Grenada.
	CountryGD = &Ref{ID: 1504, ObjectName: ObjectStaticCountry}
	// CountryGE — Georgia (Georgien).
	CountryGE = &Ref{ID: 1394, ObjectName: ObjectStaticCountry}
	// CountryGF — French Guiana (Französisch Guyana).
	CountryGF = &Ref{ID: 1507, ObjectName: ObjectStaticCountry}
	// CountryGG — Guernsey.
	CountryGG = &Ref{ID: 1499, ObjectName: ObjectStaticCountry}
	// CountryGH — Ghana.
	CountryGH = &Ref{ID: 1382, ObjectName: ObjectStaticCountry}
	// CountryGI — Gibraltar.
	CountryGI = &Ref{ID: 1500, ObjectName: ObjectStaticCountry}
	// CountryGL — Greenland (Grönland).
	CountryGL = &Ref{ID: 1505, ObjectName: ObjectStaticCountry}
	// CountryGM — Gambia.
	CountryGM = &Ref{ID: 1501, ObjectName: ObjectStaticCountry}
	// CountryGN — Guinea.
	CountryGN = &Ref{ID: 1376, ObjectName: ObjectStaticCountry}
	// CountryGP — Guadeloupe.
	CountryGP = &Ref{ID: 1438, ObjectName: ObjectStaticCountry}
	// CountryGQ — Equatorial Guinea (Äquatorialguinea).
	CountryGQ = &Ref{ID: 1503, ObjectName: ObjectStaticCountry}
	// CountryGR — Greece (Griechenland).
	CountryGR = &Ref{ID: 12, ObjectName: ObjectStaticCountry}
	// CountryGS — South Georgia (Südgeorgien und die Südlichen Sandwichinseln).
	CountryGS = &Ref{ID: 1537, ObjectName: ObjectStaticCountry}
	// CountryGT — Guatemala.
	CountryGT = &Ref{ID: 1506, ObjectName: ObjectStaticCountry}
	// CountryGU — Guam.
	CountryGU = &Ref{ID: 1508, ObjectName: ObjectStaticCountry}
	// CountryGW — Guinea-Bissau.
	CountryGW = &Ref{ID: 1502, ObjectName: ObjectStaticCountry}
	// CountryGY — Guyana.
	CountryGY = &Ref{ID: 1509, ObjectName: ObjectStaticCountry}
	// CountryHK — Hong Kong (Hongkong).
	CountryHK = &Ref{ID: 59, ObjectName: ObjectStaticCountry}
	// CountryHM — Heard Island and McDonald Islands (Heard und die McDonaldinseln).
	CountryHM = &Ref{ID: 1510, ObjectName: ObjectStaticCountry}
	// CountryHN — Honduras.
	CountryHN = &Ref{ID: 1511, ObjectName: ObjectStaticCountry}
	// CountryHR — Croatia (Kroatien).
	CountryHR = &Ref{ID: 48, ObjectName: ObjectStaticCountry}
	// CountryHT — Haiti.
	CountryHT = &Ref{ID: 1512, ObjectName: ObjectStaticCountry}
	// CountryHU — Hungary (Ungarn).
	CountryHU = &Ref{ID: 31, ObjectName: ObjectStaticCountry}
	// CountryID — Indonesia (Indonesien).
	CountryID = &Ref{ID: 1374, ObjectName: ObjectStaticCountry}
	// CountryIE — Ireland (Irland).
	CountryIE = &Ref{ID: 13, ObjectName: ObjectStaticCountry}
	// CountryIL — Israel.
	CountryIL = &Ref{ID: 36, ObjectName: ObjectStaticCountry}
	// CountryIM — Isle of Man.
	CountryIM = &Ref{ID: 1513, ObjectName: ObjectStaticCountry}
	// CountryIN — India (Indien).
	CountryIN = &Ref{ID: 1375, ObjectName: ObjectStaticCountry}
	// CountryIO — British Indian Ocean Territory (Britisches Territorium im Indischen Ozean).
	CountryIO = &Ref{ID: 1514, ObjectName: ObjectStaticCountry}
	// CountryIQ — Iraq (Irak).
	CountryIQ = &Ref{ID: 1419, ObjectName: ObjectStaticCountry}
	// CountryIR — Iran.
	CountryIR = &Ref{ID: 1373, ObjectName: ObjectStaticCountry}
	// CountryIS — Iceland (Island).
	CountryIS = &Ref{ID: 51, ObjectName: ObjectStaticCountry}
	// CountryIT — Italy (Italien).
	CountryIT = &Ref{ID: 14, ObjectName: ObjectStaticCountry}
	// CountryJE — Jersey.
	CountryJE = &Ref{ID: 1515, ObjectName: ObjectStaticCountry}
	// CountryJM — Jamaica (Jamaika).
	CountryJM = &Ref{ID: 15, ObjectName: ObjectStaticCountry}
	// CountryJO — Jordan (Jordanien).
	CountryJO = &Ref{ID: 1344, ObjectName: ObjectStaticCountry}
	// CountryJP — Japan.
	CountryJP = &Ref{ID: 40, ObjectName: ObjectStaticCountry}
	// CountryKE — Kenya (Kenia).
	CountryKE = &Ref{ID: 1389, ObjectName: ObjectStaticCountry}
	// CountryKG — Kyrgyzstan (Kirgisistan).
	CountryKG = &Ref{ID: 1430, ObjectName: ObjectStaticCountry}
	// CountryKH — Cambodia (Kambodscha).
	CountryKH = &Ref{ID: 1401, ObjectName: ObjectStaticCountry}
	// CountryKI — Kiribati.
	CountryKI = &Ref{ID: 1516, ObjectName: ObjectStaticCountry}
	// CountryKM — Comoros (Union der Komoren).
	CountryKM = &Ref{ID: 1407, ObjectName: ObjectStaticCountry}
	// CountryKN — Saint Kitts and Nevis (Saint Christopher und Nevis).
	CountryKN = &Ref{ID: 1433, ObjectName: ObjectStaticCountry}
	// CountryKP — North Korea (Nordkorea).
	CountryKP = &Ref{ID: 1453, ObjectName: ObjectStaticCountry}
	// CountryKR — South Korea (Südkorea).
	CountryKR = &Ref{ID: 41, ObjectName: ObjectStaticCountry}
	// CountryKW — Kuwait.
	CountryKW = &Ref{ID: 1369, ObjectName: ObjectStaticCountry}
	// CountryKY — Cayman Islands (Kaimaninseln).
	CountryKY = &Ref{ID: 1460, ObjectName: ObjectStaticCountry}
	// CountryKZ — Kazakhstan (Kasachstan).
	CountryKZ = &Ref{ID: 56, ObjectName: ObjectStaticCountry}
	// CountryLA — Laos.
	CountryLA = &Ref{ID: 1518, ObjectName: ObjectStaticCountry}
	// CountryLB — Lebanon (Libanon).
	CountryLB = &Ref{ID: 1428, ObjectName: ObjectStaticCountry}
	// CountryLC — Saint Lucia (St. Lucia).
	CountryLC = &Ref{ID: 1448, ObjectName: ObjectStaticCountry}
	// CountryLI — Liechtenstein.
	CountryLI = &Ref{ID: 42, ObjectName: ObjectStaticCountry}
	// CountryLK — Sri Lanka.
	CountryLK = &Ref{ID: 1356, ObjectName: ObjectStaticCountry}
	// CountryLR — Liberia.
	CountryLR = &Ref{ID: 1519, ObjectName: ObjectStaticCountry}
	// CountryLS — Lesotho.
	CountryLS = &Ref{ID: 1520, ObjectName: ObjectStaticCountry}
	// CountryLT — Lithuania (Litauen).
	CountryLT = &Ref{ID: 1346, ObjectName: ObjectStaticCountry}
	// CountryLU — Luxembourg (Luxemburg).
	CountryLU = &Ref{ID: 17, ObjectName: ObjectStaticCountry}
	// CountryLV — Latvia (Lettland).
	CountryLV = &Ref{ID: 16, ObjectName: ObjectStaticCountry}
	// CountryLY — Libya (Libyen).
	CountryLY = &Ref{ID: 1454, ObjectName: ObjectStaticCountry}
	// CountryMA — Morocco (Marokko).
	CountryMA = &Ref{ID: 1365, ObjectName: ObjectStaticCountry}
	// CountryMC — Monaco.
	CountryMC = &Ref{ID: 43, ObjectName: ObjectStaticCountry}
	// CountryMD — Republic of Moldova (Republik Moldau).
	CountryMD = &Ref{ID: 1362, ObjectName: ObjectStaticCountry}
	// CountryME — Montenegro.
	CountryME = &Ref{ID: 1343, ObjectName: ObjectStaticCountry}
	// CountryMF — Saint Martin.
	CountryMF = &Ref{ID: 1521, ObjectName: ObjectStaticCountry}
	// CountryMG — Madagascar (Madagaskar).
	CountryMG = &Ref{ID: 1395, ObjectName: ObjectStaticCountry}
	// CountryMH — Marshall Islands (Marshallinseln).
	CountryMH = &Ref{ID: 1523, ObjectName: ObjectStaticCountry}
	// CountryMK — North Macedonia (Nordmazedonien).
	CountryMK = &Ref{ID: 1367, ObjectName: ObjectStaticCountry}
	// CountryML — Mali.
	CountryML = &Ref{ID: 1461, ObjectName: ObjectStaticCountry}
	// CountryMM — Myanmar.
	CountryMM = &Ref{ID: 1359, ObjectName: ObjectStaticCountry}
	// CountryMN — Mongolia (Mongolei).
	CountryMN = &Ref{ID: 1360, ObjectName: ObjectStaticCountry}
	// CountryMO — Macao (Macau).
	CountryMO = &Ref{ID: 61, ObjectName: ObjectStaticCountry}
	// CountryMP — Northern Mariana Islands (Nördliche Marianen).
	CountryMP = &Ref{ID: 1524, ObjectName: ObjectStaticCountry}
	// CountryMQ — Martinique.
	CountryMQ = &Ref{ID: 1385, ObjectName: ObjectStaticCountry}
	// CountryMR — Mauritania (Mauretanien).
	CountryMR = &Ref{ID: 1364, ObjectName: ObjectStaticCountry}
	// CountryMS — Montserrat.
	CountryMS = &Ref{ID: 1363, ObjectName: ObjectStaticCountry}
	// CountryMT — Malta.
	CountryMT = &Ref{ID: 37, ObjectName: ObjectStaticCountry}
	// CountryMU — Mauritius.
	CountryMU = &Ref{ID: 1525, ObjectName: ObjectStaticCountry}
	// CountryMV — Maldives (Malediven).
	CountryMV = &Ref{ID: 1435, ObjectName: ObjectStaticCountry}
	// CountryMW — Malawi.
	CountryMW = &Ref{ID: 1408, ObjectName: ObjectStaticCountry}
	// CountryMX — Mexico (Mexiko).
	CountryMX = &Ref{ID: 64, ObjectName: ObjectStaticCountry}
	// CountryMY — Malaysia.
	CountryMY = &Ref{ID: 1387, ObjectName: ObjectStaticCountry}
	// CountryMZ — Mozambique (Mosambik).
	CountryMZ = &Ref{ID: 1410, ObjectName: ObjectStaticCountry}
	// CountryNA — Namibia.
	CountryNA = &Ref{ID: 1429, ObjectName: ObjectStaticCountry}
	// CountryNC — New Caledonia (Neukaledonien).
	CountryNC = &Ref{ID: 1526, ObjectName: ObjectStaticCountry}
	// CountryNE — Niger.
	CountryNE = &Ref{ID: 1527, ObjectName: ObjectStaticCountry}
	// CountryNF — Norfolk Island (Norfolkinsel).
	CountryNF = &Ref{ID: 1528, ObjectName: ObjectStaticCountry}
	// CountryNG — Nigeria.
	CountryNG = &Ref{ID: 32, ObjectName: ObjectStaticCountry}
	// CountryNI — Nicaragua.
	CountryNI = &Ref{ID: 1529, ObjectName: ObjectStaticCountry}
	// CountryNL — Netherlands (Niederlande).
	CountryNL = &Ref{ID: 18, ObjectName: ObjectStaticCountry}
	// CountryNO — Norway (Norwegen).
	CountryNO = &Ref{ID: 19, ObjectName: ObjectStaticCountry}
	// CountryNP — Nepal.
	CountryNP = &Ref{ID: 1462, ObjectName: ObjectStaticCountry}
	// CountryNR — Nauru.
	CountryNR = &Ref{ID: 1531, ObjectName: ObjectStaticCountry}
	// CountryNU — Niue.
	CountryNU = &Ref{ID: 1530, ObjectName: ObjectStaticCountry}
	// CountryNZ — New Zealand (Neuseeland).
	CountryNZ = &Ref{ID: 50, ObjectName: ObjectStaticCountry}
	// CountryOM — Oman.
	CountryOM = &Ref{ID: 1357, ObjectName: ObjectStaticCountry}
	// CountryPA — Panama.
	CountryPA = &Ref{ID: 1446, ObjectName: ObjectStaticCountry}
	// CountryPE — Peru.
	CountryPE = &Ref{ID: 69, ObjectName: ObjectStaticCountry}
	// CountryPF — French Polynesia (Französisch-Polynesien).
	CountryPF = &Ref{ID: 1351, ObjectName: ObjectStaticCountry}
	// CountryPG — Papua New Guinea (Papua-Neuguinea).
	CountryPG = &Ref{ID: 1534, ObjectName: ObjectStaticCountry}
	// CountryPH — Philippines (Philippinen).
	CountryPH = &Ref{ID: 1392, ObjectName: ObjectStaticCountry}
	// CountryPK — Pakistan.
	CountryPK = &Ref{ID: 55, ObjectName: ObjectStaticCountry}
	// CountryPL — Poland (Polen).
	CountryPL = &Ref{ID: 21, ObjectName: ObjectStaticCountry}
	// CountryPM — Saint Pierre and Miquelon (Saint-Pierre und Miquelon).
	CountryPM = &Ref{ID: 1540, ObjectName: ObjectStaticCountry}
	// CountryPN — Pitcairn Islands (Pitcairn).
	CountryPN = &Ref{ID: 1532, ObjectName: ObjectStaticCountry}
	// CountryPR — Puerto Rico.
	CountryPR = &Ref{ID: 1436, ObjectName: ObjectStaticCountry}
	// CountryPS — State of Palestine (Staat Palästina).
	CountryPS = &Ref{ID: 1464, ObjectName: ObjectStaticCountry}
	// CountryPT — Portugal.
	CountryPT = &Ref{ID: 22, ObjectName: ObjectStaticCountry}
	// CountryPW — Palau.
	CountryPW = &Ref{ID: 1533, ObjectName: ObjectStaticCountry}
	// CountryPY — Paraguay.
	CountryPY = &Ref{ID: 1393, ObjectName: ObjectStaticCountry}
	// CountryQA — Qatar (Katar).
	CountryQA = &Ref{ID: 1371, ObjectName: ObjectStaticCountry}
	// CountryRE — Réunion.
	CountryRE = &Ref{ID: 1535, ObjectName: ObjectStaticCountry}
	// CountryRO — Romania (Rumänien).
	CountryRO = &Ref{ID: 23, ObjectName: ObjectStaticCountry}
	// CountryRS — Serbia (Serbien).
	CountryRS = &Ref{ID: 1425, ObjectName: ObjectStaticCountry}
	// CountryRU — Russia (Russland).
	CountryRU = &Ref{ID: 24, ObjectName: ObjectStaticCountry}
	// CountryRW — Rwanda (Ruanda).
	CountryRW = &Ref{ID: 1411, ObjectName: ObjectStaticCountry}
	// CountrySA — Saudi Arabia (Saudi-Arabien).
	CountrySA = &Ref{ID: 57, ObjectName: ObjectStaticCountry}
	// CountrySB — Solomon Islands (Salomonen).
	CountrySB = &Ref{ID: 1539, ObjectName: ObjectStaticCountry}
	// CountrySC — Seychelles (Seychellen).
	CountrySC = &Ref{ID: 1388, ObjectName: ObjectStaticCountry}
	// CountrySD — Sudan.
	CountrySD = &Ref{ID: 1536, ObjectName: ObjectStaticCountry}
	// CountrySE — Sweden (Schweden).
	CountrySE = &Ref{ID: 25, ObjectName: ObjectStaticCountry}
	// CountrySG — Singapore (Singapur).
	CountrySG = &Ref{ID: 75, ObjectName: ObjectStaticCountry}
	// CountrySH — Saint Helena, Ascension and Tristan da Cunha (St. Helena, Ascension und Tristan da Cunha).
	CountrySH = &Ref{ID: 1475, ObjectName: ObjectStaticCountry}
	// CountrySI — Slovenia (Slowenien).
	CountrySI = &Ref{ID: 28, ObjectName: ObjectStaticCountry}
	// CountrySJ — Svalbard and Jan Mayen (Spitzbergen).
	CountrySJ = &Ref{ID: 1538, ObjectName: ObjectStaticCountry}
	// CountrySK — Slovakia (Slowakei).
	CountrySK = &Ref{ID: 27, ObjectName: ObjectStaticCountry}
	// CountrySL — Sierra Leone.
	CountrySL = &Ref{ID: 1378, ObjectName: ObjectStaticCountry}
	// CountrySM — San Marino.
	CountrySM = &Ref{ID: 1451, ObjectName: ObjectStaticCountry}
	// CountrySN — Senegal.
	CountrySN = &Ref{ID: 1439, ObjectName: ObjectStaticCountry}
	// CountrySO — Somalia.
	CountrySO = &Ref{ID: 1414, ObjectName: ObjectStaticCountry}
	// CountrySR — Suriname.
	CountrySR = &Ref{ID: 1542, ObjectName: ObjectStaticCountry}
	// CountrySS — South Sudan (Südsudan).
	CountrySS = &Ref{ID: 1415, ObjectName: ObjectStaticCountry}
	// CountryST — São Tomé and Príncipe (São Tomé und Príncipe).
	CountryST = &Ref{ID: 1541, ObjectName: ObjectStaticCountry}
	// CountrySV — El Salvador.
	CountrySV = &Ref{ID: 1450, ObjectName: ObjectStaticCountry}
	// CountrySX — Sint Maarten.
	CountrySX = &Ref{ID: 1544, ObjectName: ObjectStaticCountry}
	// CountrySY — Syria (Syrien).
	CountrySY = &Ref{ID: 1545, ObjectName: ObjectStaticCountry}
	// CountrySZ — Swaziland (Swasiland).
	CountrySZ = &Ref{ID: 1543, ObjectName: ObjectStaticCountry}
	// CountryTC — Turks and Caicos Islands (Turks-und Caicosinseln).
	CountryTC = &Ref{ID: 1546, ObjectName: ObjectStaticCountry}
	// CountryTD — Chad (Tschad).
	CountryTD = &Ref{ID: 1547, ObjectName: ObjectStaticCountry}
	// CountryTF — French Southern and Antarctic Lands (Französische Süd-und Antarktisgebiete).
	CountryTF = &Ref{ID: 1471, ObjectName: ObjectStaticCountry}
	// CountryTG — Togo.
	CountryTG = &Ref{ID: 1463, ObjectName: ObjectStaticCountry}
	// CountryTH — Thailand.
	CountryTH = &Ref{ID: 1380, ObjectName: ObjectStaticCountry}
	// CountryTJ — Tajikistan (Tadschikistan).
	CountryTJ = &Ref{ID: 1548, ObjectName: ObjectStaticCountry}
	// CountryTK — Tokelau.
	CountryTK = &Ref{ID: 1549, ObjectName: ObjectStaticCountry}
	// CountryTL — Timor-Leste.
	CountryTL = &Ref{ID: 1551, ObjectName: ObjectStaticCountry}
	// CountryTM — Turkmenistan.
	CountryTM = &Ref{ID: 1550, ObjectName: ObjectStaticCountry}
	// CountryTN — Tunisia (Tunesien).
	CountryTN = &Ref{ID: 1354, ObjectName: ObjectStaticCountry}
	// CountryTO — Tonga.
	CountryTO = &Ref{ID: 1552, ObjectName: ObjectStaticCountry}
	// CountryTR — Turkey (Türkei).
	CountryTR = &Ref{ID: 34, ObjectName: ObjectStaticCountry}
	// CountryTT — Trinidad and Tobago (Trinidad und Tobago).
	CountryTT = &Ref{ID: 1553, ObjectName: ObjectStaticCountry}
	// CountryTV — Tuvalu.
	CountryTV = &Ref{ID: 1554, ObjectName: ObjectStaticCountry}
	// CountryTW — Taiwan.
	CountryTW = &Ref{ID: 1424, ObjectName: ObjectStaticCountry}
	// CountryTZ — Tanzania (Tansania).
	CountryTZ = &Ref{ID: 1416, ObjectName: ObjectStaticCountry}
	// CountryUA — Ukraine.
	CountryUA = &Ref{ID: 66, ObjectName: ObjectStaticCountry}
	// CountryUG — Uganda.
	CountryUG = &Ref{ID: 1417, ObjectName: ObjectStaticCountry}
	// CountryUM — United States Minor Outlying Islands (Kleinere Inselbesitzungen der Vereinigten Staaten).
	CountryUM = &Ref{ID: 1555, ObjectName: ObjectStaticCountry}
	// CountryUnknown — Unbekannt.
	CountryUnknown = &Ref{ID: 1337, ObjectName: ObjectStaticCountry}
	// CountryUS — United States of America (Vereinigte Staaten von Amerika).
	CountryUS = &Ref{ID: 1353, ObjectName: ObjectStaticCountry}
	// CountryUY — Uruguay.
	CountryUY = &Ref{ID: 1384, ObjectName: ObjectStaticCountry}
	// CountryUZ — Uzbekistan (Usbekistan).
	CountryUZ = &Ref{ID: 1342, ObjectName: ObjectStaticCountry}
	// CountryVA — Vatican City (Vatikanstadt).
	CountryVA = &Ref{ID: 1556, ObjectName: ObjectStaticCountry}
	// CountryVC — Saint Vincent and the Grenadines (Saint Vincent und die Grenadinen).
	CountryVC = &Ref{ID: 1557, ObjectName: ObjectStaticCountry}
	// CountryVE — Venezuela.
	CountryVE = &Ref{ID: 1558, ObjectName: ObjectStaticCountry}
	// CountryVG — British Virgin Islands (Britische Jungferninseln).
	CountryVG = &Ref{ID: 1391, ObjectName: ObjectStaticCountry}
	// CountryVI — United States Virgin Islands (Amerikanische Jungferninseln).
	CountryVI = &Ref{ID: 1559, ObjectName: ObjectStaticCountry}
	// CountryVN — Vietnam.
	CountryVN = &Ref{ID: 1352, ObjectName: ObjectStaticCountry}
	// CountryVU — Vanuatu.
	CountryVU = &Ref{ID: 1560, ObjectName: ObjectStaticCountry}
	// CountryWF — Wallis and Futuna (Wallis und Futuna).
	CountryWF = &Ref{ID: 1561, ObjectName: ObjectStaticCountry}
	// CountryWS — Samoa.
	CountryWS = &Ref{ID: 1562, ObjectName: ObjectStaticCountry}
	// CountryXK — Kosovo.
	CountryXK = &Ref{ID: 1422, ObjectName: ObjectStaticCountry}
	// CountryYE — Yemen (Jemen).
	CountryYE = &Ref{ID: 1459, ObjectName: ObjectStaticCountry}
	// CountryYT — Mayotte.
	CountryYT = &Ref{ID: 1409, ObjectName: ObjectStaticCountry}
	// CountryZA — South Africa (Südafrika).
	CountryZA = &Ref{ID: 47, ObjectName: ObjectStaticCountry}
	// CountryZM — Zambia (Sambia).
	CountryZM = &Ref{ID: 1412, ObjectName: ObjectStaticCountry}
	// CountryZW — Zimbabwe (Simbabwe).
	CountryZW = &Ref{ID: 1413, ObjectName: ObjectStaticCountry}
)

// Country returns the canonical [StaticCountry] reference for the given
// ISO 3166-1 alpha-2 code.
//
// For codes with multiple sevdesk entries (notably "gb"), this returns the
// lowest-ID record. Use [CountryGBEngland] etc. to reach the variants.
func Country(code string) *Ref {
	if c, ok := countries[strings.ToLower(code)]; ok {
		return c
	}
	return CountryUnknown
}

var countries = map[string]*Ref{
	"ad": CountryAD,
	"ae": CountryAE,
	"af": CountryAF,
	"ag": CountryAG,
	"ai": CountryAI,
	"al": CountryAL,
	"am": CountryAM,
	"ao": CountryAO,
	"aq": CountryAQ,
	"ar": CountryAR,
	"as": CountryAS,
	"at": CountryAT,
	"au": CountryAU,
	"aw": CountryAW,
	"ax": CountryAX,
	"az": CountryAZ,
	"ba": CountryBA,
	"bb": CountryBB,
	"bd": CountryBD,
	"be": CountryBE,
	"bf": CountryBF,
	"bg": CountryBG,
	"bh": CountryBH,
	"bi": CountryBI,
	"bj": CountryBJ,
	"bl": CountryBL,
	"bm": CountryBM,
	"bn": CountryBN,
	"bo": CountryBO,
	"bq": CountryBQ,
	"br": CountryBR,
	"bs": CountryBS,
	"bt": CountryBT,
	"bv": CountryBV,
	"bw": CountryBW,
	"by": CountryBY,
	"bz": CountryBZ,
	"ca": CountryCA,
	"cc": CountryCC,
	"cd": CountryCD,
	"cf": CountryCF,
	"cg": CountryCG,
	"ch": CountryCH,
	"ci": CountryCI,
	"ck": CountryCK,
	"cl": CountryCL,
	"cm": CountryCM,
	"cn": CountryCN,
	"co": CountryCO,
	"cr": CountryCR,
	"cu": CountryCU,
	"cv": CountryCV,
	"cw": CountryCW,
	"cx": CountryCX,
	"cy": CountryCY,
	"cz": CountryCZ,
	"de": CountryDE,
	"dj": CountryDJ,
	"dk": CountryDK,
	"dm": CountryDM,
	"do": CountryDO,
	"du": CountryDU,
	"dz": CountryDZ,
	"ec": CountryEC,
	"ee": CountryEE,
	"eg": CountryEG,
	"eh": CountryEH,
	"er": CountryER,
	"es": CountryES,
	"et": CountryET,
	"fi": CountryFI,
	"fj": CountryFJ,
	"fk": CountryFK,
	"fm": CountryFM,
	"fo": CountryFO,
	"fr": CountryFR,
	"ga": CountryGA,
	"gb": CountryGB,
	"gd": CountryGD,
	"ge": CountryGE,
	"gf": CountryGF,
	"gg": CountryGG,
	"gh": CountryGH,
	"gi": CountryGI,
	"gl": CountryGL,
	"gm": CountryGM,
	"gn": CountryGN,
	"gp": CountryGP,
	"gq": CountryGQ,
	"gr": CountryGR,
	"gs": CountryGS,
	"gt": CountryGT,
	"gu": CountryGU,
	"gw": CountryGW,
	"gy": CountryGY,
	"hk": CountryHK,
	"hm": CountryHM,
	"hn": CountryHN,
	"hr": CountryHR,
	"ht": CountryHT,
	"hu": CountryHU,
	"id": CountryID,
	"ie": CountryIE,
	"il": CountryIL,
	"im": CountryIM,
	"in": CountryIN,
	"io": CountryIO,
	"iq": CountryIQ,
	"ir": CountryIR,
	"is": CountryIS,
	"it": CountryIT,
	"je": CountryJE,
	"jm": CountryJM,
	"jo": CountryJO,
	"jp": CountryJP,
	"ke": CountryKE,
	"kg": CountryKG,
	"kh": CountryKH,
	"ki": CountryKI,
	"km": CountryKM,
	"kn": CountryKN,
	"kp": CountryKP,
	"kr": CountryKR,
	"kw": CountryKW,
	"ky": CountryKY,
	"kz": CountryKZ,
	"la": CountryLA,
	"lb": CountryLB,
	"lc": CountryLC,
	"li": CountryLI,
	"lk": CountryLK,
	"lr": CountryLR,
	"ls": CountryLS,
	"lt": CountryLT,
	"lu": CountryLU,
	"lv": CountryLV,
	"ly": CountryLY,
	"ma": CountryMA,
	"mc": CountryMC,
	"md": CountryMD,
	"me": CountryME,
	"mf": CountryMF,
	"mg": CountryMG,
	"mh": CountryMH,
	"mk": CountryMK,
	"ml": CountryML,
	"mm": CountryMM,
	"mn": CountryMN,
	"mo": CountryMO,
	"mp": CountryMP,
	"mq": CountryMQ,
	"mr": CountryMR,
	"ms": CountryMS,
	"mt": CountryMT,
	"mu": CountryMU,
	"mv": CountryMV,
	"mw": CountryMW,
	"mx": CountryMX,
	"my": CountryMY,
	"mz": CountryMZ,
	"na": CountryNA,
	"nc": CountryNC,
	"ne": CountryNE,
	"nf": CountryNF,
	"ng": CountryNG,
	"ni": CountryNI,
	"nl": CountryNL,
	"no": CountryNO,
	"np": CountryNP,
	"nr": CountryNR,
	"nu": CountryNU,
	"nz": CountryNZ,
	"om": CountryOM,
	"pa": CountryPA,
	"pe": CountryPE,
	"pf": CountryPF,
	"pg": CountryPG,
	"ph": CountryPH,
	"pk": CountryPK,
	"pl": CountryPL,
	"pm": CountryPM,
	"pn": CountryPN,
	"pr": CountryPR,
	"ps": CountryPS,
	"pt": CountryPT,
	"pw": CountryPW,
	"py": CountryPY,
	"qa": CountryQA,
	"re": CountryRE,
	"ro": CountryRO,
	"rs": CountryRS,
	"ru": CountryRU,
	"rw": CountryRW,
	"sa": CountrySA,
	"sb": CountrySB,
	"sc": CountrySC,
	"sd": CountrySD,
	"se": CountrySE,
	"sg": CountrySG,
	"sh": CountrySH,
	"si": CountrySI,
	"sj": CountrySJ,
	"sk": CountrySK,
	"sl": CountrySL,
	"sm": CountrySM,
	"sn": CountrySN,
	"so": CountrySO,
	"sr": CountrySR,
	"ss": CountrySS,
	"st": CountryST,
	"sv": CountrySV,
	"sx": CountrySX,
	"sy": CountrySY,
	"sz": CountrySZ,
	"tc": CountryTC,
	"td": CountryTD,
	"tf": CountryTF,
	"tg": CountryTG,
	"th": CountryTH,
	"tj": CountryTJ,
	"tk": CountryTK,
	"tl": CountryTL,
	"tm": CountryTM,
	"tn": CountryTN,
	"to": CountryTO,
	"tr": CountryTR,
	"tt": CountryTT,
	"tv": CountryTV,
	"tw": CountryTW,
	"tz": CountryTZ,
	"ua": CountryUA,
	"ug": CountryUG,
	"um": CountryUM,
	"us": CountryUS,
	"uy": CountryUY,
	"uz": CountryUZ,
	"va": CountryVA,
	"vc": CountryVC,
	"ve": CountryVE,
	"vg": CountryVG,
	"vi": CountryVI,
	"vn": CountryVN,
	"vu": CountryVU,
	"wf": CountryWF,
	"ws": CountryWS,
	"xk": CountryXK,
	"ye": CountryYE,
	"yt": CountryYT,
	"za": CountryZA,
	"zm": CountryZM,
	"zw": CountryZW,
}

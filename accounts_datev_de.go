package sevdesk

// AccountDatev references for the German DATEV accounts a business books to
// most often, for use with [VoucherPosCreate.AccountDatev] on sevdesk-Update 2.0:
//
//	pos.AccountDatev = sevdesk.AccountDatevBuerobedarf
//
// Identifiers follow the German DATEV labels with umlauts transliterated
// (ue, oe, ae, ss), so they can be matched against the account list in the
// sevdesk UI or an accountant's chart without translating anything. Each doc
// comment carries the DATEV number, the exact label and whether the account
// takes revenue or expense postings.
//
// This is a curated subset. All 652 accounts of the bundled chart are reachable
// by number with [AccountDatev] — sevdesk.AccountDatev(7364) — and the accounts
// your own client may actually book come from [ReceiptGuidanceService], which
// also reports the tax rules and rates each one accepts.
//
// Where DATEV reserves a range for one line (4000-4099 are all "Umsatzerlöse"),
// the name below is the first number of the range; see [AccountDatev].
var (
	// Erlöse

	// AccountDatevUmsatzerloese — 4000 Umsatzerlöse (revenue).
	AccountDatevUmsatzerloese = &Ref{ID: 3510, ObjectName: ObjectAccountDatev}
	// AccountDatevErloese — 4200 Erlöse (revenue).
	AccountDatevErloese = &Ref{ID: 3631, ObjectName: ObjectAccountDatev}
	// AccountDatevErloese7USt — 4300 Erlöse 7 % USt (revenue).
	AccountDatevErloese7USt = &Ref{ID: 5186, ObjectName: ObjectAccountDatev}
	// AccountDatevErloese19USt — 4400 Erlöse 19 % USt (revenue).
	AccountDatevErloese19USt = &Ref{ID: 5192, ObjectName: ObjectAccountDatev}
	// AccountDatevSteuerfreieUmsaetzeP4 — 4100 Steuerfreie Umsätze § 4 Nr. 8ff UStG (revenue).
	AccountDatevSteuerfreieUmsaetzeP4 = &Ref{ID: 3610, ObjectName: ObjectAccountDatev}
	// AccountDatevSteuerfreieErloeseDrittland — 4120 Steuerfreie Erlöse Drittland (revenue).
	AccountDatevSteuerfreieErloeseDrittland = &Ref{ID: 3617, ObjectName: ObjectAccountDatev}
	// AccountDatevSteuerfreieEUErloese — 4125 steuerfreie EU Erlöse (revenue).
	AccountDatevSteuerfreieEUErloese = &Ref{ID: 3618, ObjectName: ObjectAccountDatev}
	// AccountDatevErloeseLeistungen13bUStG — 4337 Erlöse aus Leistungen nach § 13 b UStG (revenue).
	AccountDatevErloeseLeistungen13bUStG = &Ref{ID: 3650, ObjectName: ObjectAccountDatev}
	// AccountDatevProvisionsumsaetze — 4560 Provisionsumsätze (revenue).
	AccountDatevProvisionsumsaetze = &Ref{ID: 3675, ObjectName: ObjectAccountDatev}
	// AccountDatevErloesschmaelerungen — 4700 Erlösschmälerungen (revenue).
	AccountDatevErloesschmaelerungen = &Ref{ID: 3712, ObjectName: ObjectAccountDatev}
	// AccountDatevGewaehrteSkonti — 4730 Gewährte Skonti (revenue).
	AccountDatevGewaehrteSkonti = &Ref{ID: 3725, ObjectName: ObjectAccountDatev}
	// AccountDatevSonstigeBetrieblicheErtraege — 4830 Sonstige betriebliche Erträge (revenue).
	AccountDatevSonstigeBetrieblicheErtraege = &Ref{ID: 3753, ObjectName: ObjectAccountDatev}
	// AccountDatevErloeseVermietungVerpachtung19USt — 4862 Erlöse aus Vermietung und Verpachtung 19 % USt (revenue).
	AccountDatevErloeseVermietungVerpachtung19USt = &Ref{ID: 5251, ObjectName: ObjectAccountDatev}

	// Wareneingang und Fremdleistungen

	// AccountDatevAufwendungenRHB — 5000 Aufwendungen für Roh-, Hilfs- und Betriebsstoffe und für bezogene Waren (expense).
	AccountDatevAufwendungenRHB = &Ref{ID: 6131, ObjectName: ObjectAccountDatev}
	// AccountDatevEinkaufRHB — 5100 Einkauf Roh-, Hilfs und Betriebsstoffe (expense).
	AccountDatevEinkaufRHB = &Ref{ID: 3924, ObjectName: ObjectAccountDatev}
	// AccountDatevWareneingang — 5200 Wareneingang (expense).
	AccountDatevWareneingang = &Ref{ID: 3939, ObjectName: ObjectAccountDatev}
	// AccountDatevWareneingang7VSt — 5300 Wareneingang 7 % VSt (expense).
	AccountDatevWareneingang7VSt = &Ref{ID: 5283, ObjectName: ObjectAccountDatev}
	// AccountDatevWareneingang19Vorsteuer — 5400 Wareneingang 19 % Vorsteuer (expense).
	AccountDatevWareneingang19Vorsteuer = &Ref{ID: 5285, ObjectName: ObjectAccountDatev}
	// AccountDatevErhalteneSkonti — 5730 Erhaltene Skonti (expense).
	AccountDatevErhalteneSkonti = &Ref{ID: 4011, ObjectName: ObjectAccountDatev}
	// AccountDatevBezugsnebenkosten — 5800 Bezugsnebenkosten (expense).
	AccountDatevBezugsnebenkosten = &Ref{ID: 4049, ObjectName: ObjectAccountDatev}
	// AccountDatevFremdleistungen — 5900 Fremdleistungen (expense).
	AccountDatevFremdleistungen = &Ref{ID: 4056, ObjectName: ObjectAccountDatev}
	// AccountDatevLeistungen13bMitVorsteuerabzug — 5960 Leistungen § 13b mit Vorsteuerabzug (expense).
	AccountDatevLeistungen13bMitVorsteuerabzug = &Ref{ID: 4078, ObjectName: ObjectAccountDatev}
	// AccountDatevLeistungen13bOhneVorsteuerabzug — 5965 Leistungen § 13b ohne Vorsteuerabzug (expense).
	AccountDatevLeistungen13bOhneVorsteuerabzug = &Ref{ID: 4079, ObjectName: ObjectAccountDatev}

	// Personal

	// AccountDatevLoehneUndGehaelter — 6000 Löhne und Gehälter (expense).
	AccountDatevLoehneUndGehaelter = &Ref{ID: 4085, ObjectName: ObjectAccountDatev}
	// AccountDatevGehaelter — 6020 Gehälter (expense).
	AccountDatevGehaelter = &Ref{ID: 4087, ObjectName: ObjectAccountDatev}
	// AccountDatevAushilfsloehne — 6030 Aushilfslöhne (expense).
	AccountDatevAushilfsloehne = &Ref{ID: 4093, ObjectName: ObjectAccountDatev}
	// AccountDatevLoehneFuerMinijobs — 6035 Löhne für Minijobs (expense).
	AccountDatevLoehneFuerMinijobs = &Ref{ID: 4094, ObjectName: ObjectAccountDatev}
	// AccountDatevGesetzlicheSozialaufwendungen — 6110 Gesetzliche Sozialaufwendungen (expense).
	AccountDatevGesetzlicheSozialaufwendungen = &Ref{ID: 4120, ObjectName: ObjectAccountDatev}
	// AccountDatevBeitraegeBerufsgenossenschaft — 6120 Beiträge zur Berufsgenossenschaft (expense).
	AccountDatevBeitraegeBerufsgenossenschaft = &Ref{ID: 4122, ObjectName: ObjectAccountDatev}

	// Abschreibungen

	// AccountDatevAbschreibungenSachanlagen — 6220 Abschreibungen auf Sachanlagen (expense).
	AccountDatevAbschreibungenSachanlagen = &Ref{ID: 4138, ObjectName: ObjectAccountDatev}
	// AccountDatevSofortabschreibungGWG — 6260 Sofortabschreibung GWG (expense).
	AccountDatevSofortabschreibungGWG = &Ref{ID: 4153, ObjectName: ObjectAccountDatev}

	// Raum und Betrieb

	// AccountDatevSonstigeBetrieblicheAufwendungen — 6300 sonstige betriebliche Aufwendungen (expense).
	AccountDatevSonstigeBetrieblicheAufwendungen = &Ref{ID: 4166, ObjectName: ObjectAccountDatev}
	// AccountDatevRaumkosten — 6305 Raumkosten (expense).
	AccountDatevRaumkosten = &Ref{ID: 4170, ObjectName: ObjectAccountDatev}
	// AccountDatevMiete — 6310 Miete, unbewegliche Wirtschaftsgüter (expense).
	AccountDatevMiete = &Ref{ID: 4171, ObjectName: ObjectAccountDatev}
	// AccountDatevMietUndPachtnebenkosten — 6318 Miet- und Pachtnebenkosten (expense).
	AccountDatevMietUndPachtnebenkosten = &Ref{ID: 4178, ObjectName: ObjectAccountDatev}
	// AccountDatevGasStromWasser — 6325 Gas, Strom, Wasser (expense).
	AccountDatevGasStromWasser = &Ref{ID: 4181, ObjectName: ObjectAccountDatev}
	// AccountDatevReinigung — 6330 Reinigung (expense).
	AccountDatevReinigung = &Ref{ID: 4182, ObjectName: ObjectAccountDatev}
	// AccountDatevInstandhaltungBetrieblicherRaeume — 6335 Instandhaltung betrieblicher Räume (expense).
	AccountDatevInstandhaltungBetrieblicherRaeume = &Ref{ID: 4183, ObjectName: ObjectAccountDatev}
	// AccountDatevArbeitszimmer — 6348 Aufwendungen für ein häusliches Arbeitszimmer (Abziehbarer Anteil) (expense).
	AccountDatevArbeitszimmer = &Ref{ID: 4186, ObjectName: ObjectAccountDatev}
	// AccountDatevVersicherung — 6400 Versicherung (expense).
	AccountDatevVersicherung = &Ref{ID: 4198, ObjectName: ObjectAccountDatev}
	// AccountDatevBeitraege — 6420 Beiträge (expense).
	AccountDatevBeitraege = &Ref{ID: 4201, ObjectName: ObjectAccountDatev}
	// AccountDatevWartungskostenHardUndSoftware — 6495 Wartungskosten für Hard- und Software (expense).
	AccountDatevWartungskostenHardUndSoftware = &Ref{ID: 4212, ObjectName: ObjectAccountDatev}

	// Fahrzeug

	// AccountDatevFahrzeugkosten — 6500 Fahrzeugkosten (expense).
	AccountDatevFahrzeugkosten = &Ref{ID: 4214, ObjectName: ObjectAccountDatev}
	// AccountDatevKfzVersicherung — 6520 KFZ-Versicherung (expense).
	AccountDatevKfzVersicherung = &Ref{ID: 4215, ObjectName: ObjectAccountDatev}
	// AccountDatevLaufendeKfzBetriebskosten — 6530 Laufende KFZ-Betriebskosten (expense).
	AccountDatevLaufendeKfzBetriebskosten = &Ref{ID: 4216, ObjectName: ObjectAccountDatev}
	// AccountDatevKfzReparaturen — 6540 KFZ-Reparaturen (expense).
	AccountDatevKfzReparaturen = &Ref{ID: 4217, ObjectName: ObjectAccountDatev}
	// AccountDatevGaragenmieten — 6550 Garagenmieten (expense).
	AccountDatevGaragenmieten = &Ref{ID: 4218, ObjectName: ObjectAccountDatev}
	// AccountDatevMietleasingKfz — 6560 Mietleasing Kfz (expense).
	AccountDatevMietleasingKfz = &Ref{ID: 4219, ObjectName: ObjectAccountDatev}
	// AccountDatevSonstigeKfzKosten — 6570 Sonstige Kfz-Kosten (expense).
	AccountDatevSonstigeKfzKosten = &Ref{ID: 4221, ObjectName: ObjectAccountDatev}

	// Vertrieb und Werbung

	// AccountDatevWerbekosten — 6600 Werbekosten (expense).
	AccountDatevWerbekosten = &Ref{ID: 4225, ObjectName: ObjectAccountDatev}
	// AccountDatevReisekostenArbeitnehmer — 6650 Reisekosten Arbeitnehmer (expense).
	AccountDatevReisekostenArbeitnehmer = &Ref{ID: 4242, ObjectName: ObjectAccountDatev}
	// AccountDatevReisekostenUnternehmer — 6670 Reisekosten Unternehmer (expense).
	AccountDatevReisekostenUnternehmer = &Ref{ID: 6042, ObjectName: ObjectAccountDatev}
	// AccountDatevVerpackungsmaterial — 6710 Verpackungsmaterial (expense).
	AccountDatevVerpackungsmaterial = &Ref{ID: 4257, ObjectName: ObjectAccountDatev}
	// AccountDatevAusgangsfrachten — 6740 Ausgangsfrachten (expense).
	AccountDatevAusgangsfrachten = &Ref{ID: 4258, ObjectName: ObjectAccountDatev}
	// AccountDatevVerkaufsprovisionen — 6770 Verkaufsprovisionen (expense).
	AccountDatevVerkaufsprovisionen = &Ref{ID: 4260, ObjectName: ObjectAccountDatev}

	// Büro und Beratung

	// AccountDatevPorto — 6800 Porto (expense).
	AccountDatevPorto = &Ref{ID: 4263, ObjectName: ObjectAccountDatev}
	// AccountDatevTelefon — 6805 Telefon (expense).
	AccountDatevTelefon = &Ref{ID: 4264, ObjectName: ObjectAccountDatev}
	// AccountDatevTelefaxUndInternetkosten — 6810 Telefax und Internetkosten (expense).
	AccountDatevTelefaxUndInternetkosten = &Ref{ID: 4265, ObjectName: ObjectAccountDatev}
	// AccountDatevBuerobedarf — 6815 Bürobedarf (expense).
	AccountDatevBuerobedarf = &Ref{ID: 4266, ObjectName: ObjectAccountDatev}
	// AccountDatevFachliteratur — 6820 Zeitschriften, Bücher (Fachliteratur) (expense).
	AccountDatevFachliteratur = &Ref{ID: 4267, ObjectName: ObjectAccountDatev}
	// AccountDatevFortbildungskosten — 6821 Fortbildungskosten (expense).
	AccountDatevFortbildungskosten = &Ref{ID: 4268, ObjectName: ObjectAccountDatev}
	// AccountDatevRechtsUndBeratungskosten — 6825 Rechts- und Beratungskosten (expense).
	AccountDatevRechtsUndBeratungskosten = &Ref{ID: 4272, ObjectName: ObjectAccountDatev}
	// AccountDatevAbschlussUndPruefungskosten — 6827 Abschluss- und Prüfungskosten (expense).
	AccountDatevAbschlussUndPruefungskosten = &Ref{ID: 4273, ObjectName: ObjectAccountDatev}
	// AccountDatevBuchfuehrungskosten — 6830 Buchführungskosten (expense).
	AccountDatevBuchfuehrungskosten = &Ref{ID: 4274, ObjectName: ObjectAccountDatev}
	// AccountDatevLaufendeLizenzgebuehren — 6837 laufende Gebühren für Lizenzen (expense).
	AccountDatevLaufendeLizenzgebuehren = &Ref{ID: 4279, ObjectName: ObjectAccountDatev}
	// AccountDatevWerkzeugeUndKleingeraete — 6845 Werkzeuge und Kleingeräte (expense).
	AccountDatevWerkzeugeUndKleingeraete = &Ref{ID: 4282, ObjectName: ObjectAccountDatev}
	// AccountDatevSonstigerBetriebsbedarf — 6850 Sonstiger Betriebsbedarf (expense).
	AccountDatevSonstigerBetriebsbedarf = &Ref{ID: 4283, ObjectName: ObjectAccountDatev}

	// Finanzen

	// AccountDatevNebenkostenDesGeldverkehrs — 6855 Nebenkosten des Geldverkehrs (expense).
	AccountDatevNebenkostenDesGeldverkehrs = &Ref{ID: 4285, ObjectName: ObjectAccountDatev}
	// AccountDatevSonstigeZinsertraege — 7100 Sonstige Zinsen und ähnliche Erträge (revenue).
	AccountDatevSonstigeZinsertraege = &Ref{ID: 4368, ObjectName: ObjectAccountDatev}
	// AccountDatevZinsaufwendungen — 7300 Zinsen und ähnliche Aufwendungen (expense).
	AccountDatevZinsaufwendungen = &Ref{ID: 4402, ObjectName: ObjectAccountDatev}
	// AccountDatevZinsenKurzfristigeVerbindlichkeiten — 7310 Zinsaufwendungen für kurzfristige Verbindlichkeiten (expense).
	AccountDatevZinsenKurzfristigeVerbindlichkeiten = &Ref{ID: 4411, ObjectName: ObjectAccountDatev}
	// AccountDatevZinsenKontokorrent — 7318 Zinsen auf Kontokorrentkonten (expense).
	AccountDatevZinsenKontokorrent = &Ref{ID: 4416, ObjectName: ObjectAccountDatev}

	// Privat und Durchlauf

	// AccountDatevKautionen — 1350 Kautionen (expense).
	AccountDatevKautionen = &Ref{ID: 2571, ObjectName: ObjectAccountDatev}
	// AccountDatevDurchlaufendePosten — 1370 Durchlaufende Posten (expense, revenue).
	AccountDatevDurchlaufendePosten = &Ref{ID: 2578, ObjectName: ObjectAccountDatev}
	// AccountDatevEinfuhrumsatzsteuer — 1433 Einfuhrumsatzsteuer (expense).
	AccountDatevEinfuhrumsatzsteuer = &Ref{ID: 2624, ObjectName: ObjectAccountDatev}
	// AccountDatevPrivatentnahmen — 2100 Privatentnahmen allgemein (expense).
	AccountDatevPrivatentnahmen = &Ref{ID: 2838, ObjectName: ObjectAccountDatev}
	// AccountDatevPrivateinlagen — 2180 Privateinlagen (revenue).
	AccountDatevPrivateinlagen = &Ref{ID: 5641, ObjectName: ObjectAccountDatev}
)

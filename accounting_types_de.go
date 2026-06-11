package sevdesk

// AccountingType references for the German SKR03 / SKR04 charts.
// Each value is a [*Ref] ready to drop into a [VoucherPosCreate.AccountingType]
// or [CreditNoteCreateFields.AccountingType] field.
//
// Snapshot taken from sevdesk's /AccountingType endpoint. IDs are permanent;
// sevdesk may add new entries over time. For non-German charts (Austria,
// Switzerland, etc.) or to refresh, use [AccountingTypesService].
var (
	// AccountingTypeOtherIncome — Sonstige Erträge (SKR03 8625, SKR04 4830, kind E).
	AccountingTypeOtherIncome = &Ref{ID: 1, ObjectName: ObjectAccountingType}
	// AccountingTypeOtherCosts — Sonstiges (SKR03 4900, SKR04 6300, kind IC).
	AccountingTypeOtherCosts = &Ref{ID: 2, ObjectName: ObjectAccountingType}
	// AccountingTypeSalesTaxCosts — bezahlte Umsatzsteuer (SKR03 1770, SKR04 3800, kind VAT).
	AccountingTypeSalesTaxCosts = &Ref{ID: 3, ObjectName: ObjectAccountingType}
	// AccountingTypeVehicle — Fahrzeug (SKR03 4500, SKR04 6500, kind IC).
	AccountingTypeVehicle = &Ref{ID: 4, ObjectName: ObjectAccountingType}
	// AccountingTypePetrol — Benzin (SKR03 4530, SKR04 6530).
	AccountingTypePetrol = &Ref{ID: 5, ObjectName: ObjectAccountingType}
	// AccountingTypeInspectionRepair — Inspektion/Reparatur (SKR03 4540, SKR04 6540).
	AccountingTypeInspectionRepair = &Ref{ID: 6, ObjectName: ObjectAccountingType}
	// AccountingTypeVehicleTax — KFZ-Steuer (SKR03 4510, SKR04 7685).
	AccountingTypeVehicleTax = &Ref{ID: 7, ObjectName: ObjectAccountingType}
	// AccountingTypeVehicleInsurance — KFZ-Versicherung (SKR03 4520, SKR04 6520).
	AccountingTypeVehicleInsurance = &Ref{ID: 8, ObjectName: ObjectAccountingType}
	// AccountingTypeVehicleLeasingRental — Leasing/Mietwagen (SKR03 4595, SKR04 6595).
	AccountingTypeVehicleLeasingRental = &Ref{ID: 9, ObjectName: ObjectAccountingType}
	// AccountingTypeParkingRent — Stellplatz/Garagenmiete (SKR03 4550, SKR04 6550).
	AccountingTypeParkingRent = &Ref{ID: 10, ObjectName: ObjectAccountingType}
	// AccountingTypeCarMaintenance — Wagenpflege (SKR03 4530, SKR04 6530).
	AccountingTypeCarMaintenance = &Ref{ID: 11, ObjectName: ObjectAccountingType}
	// AccountingTypeOtherVehicleCosts — Sonstige KFZ-Kosten (SKR03 4580, SKR04 6570).
	AccountingTypeOtherVehicleCosts = &Ref{ID: 12, ObjectName: ObjectAccountingType}
	// AccountingTypeCredit — Banken / Finanzen (SKR03 2120, SKR04 7320, kind IC).
	AccountingTypeCredit = &Ref{ID: 13, ObjectName: ObjectAccountingType}
	// AccountingTypeRepayment — Darlehen & Tilgung (SKR03 640, SKR04 1365, kind NONE).
	AccountingTypeRepayment = &Ref{ID: 14, ObjectName: ObjectAccountingType}
	// AccountingTypeLendingRates — Kreditzinsen (SKR03 2120, SKR04 7320).
	AccountingTypeLendingRates = &Ref{ID: 15, ObjectName: ObjectAccountingType}
	// AccountingTypeCreditFees — Kreditgebühren (SKR03 2120, SKR04 7320).
	AccountingTypeCreditFees = &Ref{ID: 16, ObjectName: ObjectAccountingType}
	// AccountingTypeMaterialGoods — Material/Waren (SKR03 3200, SKR04 5200, kind DC).
	AccountingTypeMaterialGoods = &Ref{ID: 17, ObjectName: ObjectAccountingType}
	// AccountingTypeMaterialPurchase — Materialeinkauf (SKR03 3000, SKR04 5100).
	AccountingTypeMaterialPurchase = &Ref{ID: 18, ObjectName: ObjectAccountingType}
	// AccountingTypeGoodsPurchase — Wareneinkauf (SKR03 3200, SKR04 5200).
	AccountingTypeGoodsPurchase = &Ref{ID: 19, ObjectName: ObjectAccountingType}
	// AccountingTypeCostsReductions — Aufwandsminderungen (SKR03 3700, SKR04 5700).
	AccountingTypeCostsReductions = &Ref{ID: 20, ObjectName: ObjectAccountingType}
	// AccountingTypeLawyer — Rechtsanwalt (SKR03 4950, SKR04 6825).
	AccountingTypeLawyer = &Ref{ID: 22, ObjectName: ObjectAccountingType}
	// AccountingTypeTaxConsultant — Steuerberater (SKR03 4957, SKR04 6827).
	AccountingTypeTaxConsultant = &Ref{ID: 23, ObjectName: ObjectAccountingType}
	// AccountingTypeSales — Umsätze (SKR03 8200, SKR04 4200, kind E).
	AccountingTypeSales = &Ref{ID: 24, ObjectName: ObjectAccountingType}
	// AccountingTypeSalesTaxIncome — erhaltene Umsatzsteuer (SKR03 1770, SKR04 3800, kind VAT).
	AccountingTypeSalesTaxIncome = &Ref{ID: 25, ObjectName: ObjectAccountingType}
	// AccountingTypeRevenue — Einnahmen / Erlöse / Verkäufe (SKR03 8200, SKR04 4200, kind E).
	AccountingTypeRevenue = &Ref{ID: 26, ObjectName: ObjectAccountingType}
	// AccountingTypeSalesDeductions — Erlösminderung (SKR03 8700, SKR04 4700).
	AccountingTypeSalesDeductions = &Ref{ID: 27, ObjectName: ObjectAccountingType}
	// AccountingTypeCommission — Provision / Courtage (SKR03 8510, SKR04 4560, kind E).
	AccountingTypeCommission = &Ref{ID: 31, ObjectName: ObjectAccountingType}
	// AccountingTypePowerConsumption — Eigenverbrauch (SKR03 8906, SKR04 4619).
	AccountingTypePowerConsumption = &Ref{ID: 34, ObjectName: ObjectAccountingType}
	// AccountingTypePatentLicense — Patent- und Lizenzverträge (SKR03 8570, SKR04 4570).
	AccountingTypePatentLicense = &Ref{ID: 36, ObjectName: ObjectAccountingType}
	// AccountingTypePrivateDeposit — Privateinlagen (SKR03 1890, SKR04 2180, kind EQUITYIN).
	AccountingTypePrivateDeposit = &Ref{ID: 37, ObjectName: ObjectAccountingType}
	// AccountingTypeOverdueFine — Mahngebühren (SKR03 2709, SKR04 4839).
	AccountingTypeOverdueFine = &Ref{ID: 38, ObjectName: ObjectAccountingType}
	// AccountingTypePassingItem — Durchlaufende Posten (SKR03 1590, SKR04 1370, kind NONE).
	AccountingTypePassingItem = &Ref{ID: 39, ObjectName: ObjectAccountingType}
	// AccountingTypeMonetaryTransit — Geldtransit (SKR03 1360, SKR04 1460, kind NONE).
	AccountingTypeMonetaryTransit = &Ref{ID: 40, ObjectName: ObjectAccountingType}
	// AccountingTypeRoundingDifferences — Rundungsdifferenzen (SKR03 2660, SKR04 4840).
	AccountingTypeRoundingDifferences = &Ref{ID: 41, ObjectName: ObjectAccountingType}
	// AccountingTypeBusinessMeeting — Betriebliche Besprechungen (SKR03 4653, SKR04 6643).
	AccountingTypeBusinessMeeting = &Ref{ID: 43, ObjectName: ObjectAccountingType}
	// AccountingTypeBusinessLunch — Geschäftsessen (SKR03 4650, SKR04 6640).
	AccountingTypeBusinessLunch = &Ref{ID: 44, ObjectName: ObjectAccountingType}
	// AccountingTypeForeignServices — Dienstleistung / Beratung (SKR03 3100, SKR04 5900, kind IC).
	AccountingTypeForeignServices = &Ref{ID: 45, ObjectName: ObjectAccountingType}
	// AccountingTypeServiceProvider — Dienstleister / Agenturen / Freelancer (SKR03 3100, SKR04 5900).
	AccountingTypeServiceProvider = &Ref{ID: 46, ObjectName: ObjectAccountingType}
	// AccountingTypeSubcontractor — Subunternehmer (SKR03 3100, SKR04 5900).
	AccountingTypeSubcontractor = &Ref{ID: 49, ObjectName: ObjectAccountingType}
	// AccountingTypeLeasingMachine — Leasing für Geräte (SKR03 4965, SKR04 6840).
	AccountingTypeLeasingMachine = &Ref{ID: 51, ObjectName: ObjectAccountingType}
	// AccountingTypeRentHouse — Miete / Pacht (SKR03 4210, SKR04 6310).
	AccountingTypeRentHouse = &Ref{ID: 52, ObjectName: ObjectAccountingType}
	// AccountingTypePowerWaterGas — Strom, Wasser, Gas (SKR03 4240, SKR04 6325).
	AccountingTypePowerWaterGas = &Ref{ID: 53, ObjectName: ObjectAccountingType}
	// AccountingTypeTrashTaxes — Müllgebühren (SKR03 4969, SKR04 6859).
	AccountingTypeTrashTaxes = &Ref{ID: 54, ObjectName: ObjectAccountingType}
	// AccountingTypeStaff — Personal (SKR03 4100, SKR04 6000, kind DC).
	AccountingTypeStaff = &Ref{ID: 55, ObjectName: ObjectAccountingType}
	// AccountingTypeHelpWage — Aushilfslohn (SKR03 4190, SKR04 6030).
	AccountingTypeHelpWage = &Ref{ID: 56, ObjectName: ObjectAccountingType}
	// AccountingTypeHealthInsurance — Krankenkasse (SKR03 4130, SKR04 6110).
	AccountingTypeHealthInsurance = &Ref{ID: 57, ObjectName: ObjectAccountingType}
	// AccountingTypeWageSalary — Lohn / Gehalt (SKR03 4100, SKR04 6000).
	AccountingTypeWageSalary = &Ref{ID: 58, ObjectName: ObjectAccountingType}
	// AccountingTypeHelpTaxes — Pauschale Steuer für Aushilfen (SKR03 4199, SKR04 6040).
	AccountingTypeHelpTaxes = &Ref{ID: 59, ObjectName: ObjectAccountingType}
	// AccountingTypeProvision — Prämie / Provision (SKR03 4100, SKR04 6000).
	AccountingTypeProvision = &Ref{ID: 60, ObjectName: ObjectAccountingType}
	// AccountingTypeTravel — Reisen / Verpflegung (SKR03 4670, SKR04 6670, kind IC).
	AccountingTypeTravel = &Ref{ID: 61, ObjectName: ObjectAccountingType}
	// AccountingTypeTrainPlaneTicket — Bahn- / Flugticket, Mietwagen (SKR03 4673, SKR04 6673).
	AccountingTypeTrainPlaneTicket = &Ref{ID: 62, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelCosts — Fahrtkosten (SKR03 4673, SKR04 6673).
	AccountingTypeTravelCosts = &Ref{ID: 63, ObjectName: ObjectAccountingType}
	// AccountingTypePublicTransport — Öffentliche Verkehrsmittel (SKR03 4673, SKR04 6673).
	AccountingTypePublicTransport = &Ref{ID: 64, ObjectName: ObjectAccountingType}
	// AccountingTypeTaxi — Taxi (SKR03 4673, SKR04 6673).
	AccountingTypeTaxi = &Ref{ID: 65, ObjectName: ObjectAccountingType}
	// AccountingTypeAccommodationBreakfast — Übernachtungskosten / Frühstück (SKR03 4676, SKR04 6680).
	AccountingTypeAccommodationBreakfast = &Ref{ID: 66, ObjectName: ObjectAccountingType}
	// AccountingTypeIncidentals — Kleingeräte (SKR03 4985, SKR04 6845).
	AccountingTypeIncidentals = &Ref{ID: 68, ObjectName: ObjectAccountingType}
	// AccountingTypeRenovationMaintenance — Instandhaltung Räume / Gebäude (SKR03 4260, SKR04 6335).
	AccountingTypeRenovationMaintenance = &Ref{ID: 69, ObjectName: ObjectAccountingType}
	// AccountingTypePurchases — Sonstige Anschaffungen (SKR03 4980, SKR04 6850).
	AccountingTypePurchases = &Ref{ID: 71, ObjectName: ObjectAccountingType}
	// AccountingTypeOfficeSupplies — Bürobedarf (SKR03 4930, SKR04 6815).
	AccountingTypeOfficeSupplies = &Ref{ID: 72, ObjectName: ObjectAccountingType}
	// AccountingTypeAccountRates — Girokonto Zinsen (SKR03 2110, SKR04 7310).
	AccountingTypeAccountRates = &Ref{ID: 73, ObjectName: ObjectAccountingType}
	// AccountingTypeAccountCardFee — Kontoführung / Kartengebühren (SKR03 4970, SKR04 6855).
	AccountingTypeAccountCardFee = &Ref{ID: 74, ObjectName: ObjectAccountingType}
	// AccountingTypePostage — Porto (SKR03 4910, SKR04 6800).
	AccountingTypePostage = &Ref{ID: 75, ObjectName: ObjectAccountingType}
	// AccountingTypePersonalDrawings — Privatentnahmen (SKR03 1800, SKR04 2100, kind EQUITYOUT).
	AccountingTypePersonalDrawings = &Ref{ID: 76, ObjectName: ObjectAccountingType}
	// AccountingTypeCleaning — Reinigung / Reinigungsmittel (SKR03 4250, SKR04 6330).
	AccountingTypeCleaning = &Ref{ID: 77, ObjectName: ObjectAccountingType}
	// AccountingTypeMagazineBooks — Zeitschriften / Bücher (SKR03 4940, SKR04 6820).
	AccountingTypeMagazineBooks = &Ref{ID: 78, ObjectName: ObjectAccountingType}
	// AccountingTypeReminderFees — Mahngebühren (SKR03 2309, SKR04 6969).
	AccountingTypeReminderFees = &Ref{ID: 79, ObjectName: ObjectAccountingType}
	// AccountingTypeCashTransit — Geldtransit (SKR03 1360, SKR04 1460, kind NONE).
	AccountingTypeCashTransit = &Ref{ID: 81, ObjectName: ObjectAccountingType}
	// AccountingTypeRoundingDifferences2 — Rundungsdifferenzen (SKR03 2150, SKR04 4840).
	AccountingTypeRoundingDifferences2 = &Ref{ID: 82, ObjectName: ObjectAccountingType}
	// AccountingTypeTaxes — Steuer (SKR03 1810, SKR04 2150, kind TAX).
	AccountingTypeTaxes = &Ref{ID: 83, ObjectName: ObjectAccountingType}
	// AccountingTypeSalesTax — Umsatzsteuer-Vorauszahlungen, -Nachzahlungen, -Erstattungen (SKR03 1780, SKR04 3820, kind VATPAY).
	AccountingTypeSalesTax = &Ref{ID: 84, ObjectName: ObjectAccountingType}
	// AccountingTypeIncomeTax — Einkommensteuer (SKR03 1810, SKR04 2150).
	AccountingTypeIncomeTax = &Ref{ID: 85, ObjectName: ObjectAccountingType}
	// AccountingTypeTradeTax — Gewerbesteuer (SKR03 4320, SKR04 7610).
	AccountingTypeTradeTax = &Ref{ID: 86, ObjectName: ObjectAccountingType}
	// AccountingTypeLandline — Festnetz (SKR03 4920, SKR04 6805).
	AccountingTypeLandline = &Ref{ID: 88, ObjectName: ObjectAccountingType}
	// AccountingTypeInternet — Internet (SKR03 4925, SKR04 6810).
	AccountingTypeInternet = &Ref{ID: 89, ObjectName: ObjectAccountingType}
	// AccountingTypeMobile — Mobil (SKR03 4920, SKR04 6805).
	AccountingTypeMobile = &Ref{ID: 90, ObjectName: ObjectAccountingType}
	// AccountingTypeInsuranceDues — Versicherungen / Beiträge (SKR03 4360, SKR04 6400, kind IC).
	AccountingTypeInsuranceDues = &Ref{ID: 91, ObjectName: ObjectAccountingType}
	// AccountingTypePublicLiability — Betriebshaftpflicht (SKR03 4360, SKR04 6400).
	AccountingTypePublicLiability = &Ref{ID: 92, ObjectName: ObjectAccountingType}
	// AccountingTypeGuildAssociationFees — Innungs- und Verbandsbeiträge (SKR03 4380, SKR04 6420).
	AccountingTypeGuildAssociationFees = &Ref{ID: 93, ObjectName: ObjectAccountingType}
	// AccountingTypeLegalProtection — Rechtschutz (SKR03 4360, SKR04 6400).
	AccountingTypeLegalProtection = &Ref{ID: 94, ObjectName: ObjectAccountingType}
	// AccountingTypeTransportInsurance — Transportversicherung (SKR03 4750, SKR04 6760).
	AccountingTypeTransportInsurance = &Ref{ID: 95, ObjectName: ObjectAccountingType}
	// AccountingTypeCompanyInsurance — Firmenversicherung (SKR03 4360, SKR04 6400).
	AccountingTypeCompanyInsurance = &Ref{ID: 96, ObjectName: ObjectAccountingType}
	// AccountingTypeAdvertising — Werbung (SKR03 4600, SKR04 6600, kind IC).
	AccountingTypeAdvertising = &Ref{ID: 97, ObjectName: ObjectAccountingType}
	// AccountingTypeBusinessCards — Geschäftspapier / Visitenkarten (SKR03 4600, SKR04 6600).
	AccountingTypeBusinessCards = &Ref{ID: 100, ObjectName: ObjectAccountingType}
	// AccountingTypeMarketing — Marketing / Werbekosten (SKR03 4600, SKR04 6600).
	AccountingTypeMarketing = &Ref{ID: 101, ObjectName: ObjectAccountingType}
	// AccountingTypeTradeFairCosts — Messekosten (SKR03 4600, SKR04 6600).
	AccountingTypeTradeFairCosts = &Ref{ID: 102, ObjectName: ObjectAccountingType}
	// AccountingTypeCorporateGifts — Werbegeschenke / Sponsoring (SKR03 4600, SKR04 6600).
	AccountingTypeCorporateGifts = &Ref{ID: 103, ObjectName: ObjectAccountingType}
	// AccountingTypeDonations — Spenden (SKR03 2382, SKR04 6391).
	AccountingTypeDonations = &Ref{ID: 104, ObjectName: ObjectAccountingType}
	// AccountingTypeImportSalesTax — Einfuhrumsatzsteuer (SKR03 1588, SKR04 1433, kind VATIMPORT).
	AccountingTypeImportSalesTax = &Ref{ID: 106, ObjectName: ObjectAccountingType}
	// AccountingTypeTransportFreight — Transport / Fracht (SKR03 4730, SKR04 6740, kind IC).
	AccountingTypeTransportFreight = &Ref{ID: 107, ObjectName: ObjectAccountingType}
	// AccountingTypeIncomePriceGain — Erträge aus Kursgewinnen (SKR03 2660, SKR04 4840).
	AccountingTypeIncomePriceGain = &Ref{ID: 108, ObjectName: ObjectAccountingType}
	// AccountingTypeSalesDeduction — Erlösschmälerung (SKR03 8700, SKR04 4700).
	AccountingTypeSalesDeduction = &Ref{ID: 110, ObjectName: ObjectAccountingType}
	// AccountingTypePerDiems — Verpflegungsmehraufwand (SKR03 4674, SKR04 6674).
	AccountingTypePerDiems = &Ref{ID: 1597, ObjectName: ObjectAccountingType}
	// AccountingTypeNonDeductibleCosts — Nicht abzugsfähige Bewirtung (30%) (SKR03 4654, SKR04 6644).
	AccountingTypeNonDeductibleCosts = &Ref{ID: 1598, ObjectName: ObjectAccountingType}
	// AccountingTypeMachines — Kauf einer Maschine (SKR03 210, SKR04 440).
	AccountingTypeMachines = &Ref{ID: 2809, ObjectName: ObjectAccountingType}
	// AccountingTypeBuildings — Kauf eines Gebäudes (SKR03 80, SKR04 230).
	AccountingTypeBuildings = &Ref{ID: 2811, ObjectName: ObjectAccountingType}
	// AccountingTypeCars — PKW (SKR03 320, SKR04 520).
	AccountingTypeCars = &Ref{ID: 2812, ObjectName: ObjectAccountingType}
	// AccountingTypeCostsBookkeeping — Buchführungskosten (SKR03 4955, SKR04 6830).
	AccountingTypeCostsBookkeeping = &Ref{ID: 2816, ObjectName: ObjectAccountingType}
	// AccountingTypeSoftware — Software-Miete / Lizenzen (SKR03 4964, SKR04 6837).
	AccountingTypeSoftware = &Ref{ID: 2819, ObjectName: ObjectAccountingType}
	// AccountingTypeHosting — Web-Hosting / Domains (SKR03 4964, SKR04 6837).
	AccountingTypeHosting = &Ref{ID: 2820, ObjectName: ObjectAccountingType}
	// AccountingTypeAdvancedEducation — Fortbildung / Weiterbildung (SKR03 4945, SKR04 6821).
	AccountingTypeAdvancedEducation = &Ref{ID: 2821, ObjectName: ObjectAccountingType}
	// AccountingTypeMaintenanceMaschine — Instandhaltung Maschinen (SKR03 4800, SKR04 6460).
	AccountingTypeMaintenanceMaschine = &Ref{ID: 2822, ObjectName: ObjectAccountingType}
	// AccountingTypeLiabilities — Verbindlichkeiten (SKR03 1700, SKR04 3500, kind TAX).
	AccountingTypeLiabilities = &Ref{ID: 25223, ObjectName: ObjectAccountingType}
	// AccountingTypeWage — Löhne (SKR03 4110, SKR04 6010, kind DC).
	AccountingTypeWage = &Ref{ID: 25224, ObjectName: ObjectAccountingType}
	// AccountingTypeSalary — Gehälter (SKR03 4120, SKR04 6020, kind DC).
	AccountingTypeSalary = &Ref{ID: 25225, ObjectName: ObjectAccountingType}
	// AccountingTypeSalaryManager — Geschäftsführergehälter (SKR03 4127, SKR04 6027, kind DC).
	AccountingTypeSalaryManager = &Ref{ID: 25226, ObjectName: ObjectAccountingType}
	// AccountingTypeMinijobTax — Pauschale Steuern für Minijobber (SKR03 4194, SKR04 6036, kind DC).
	AccountingTypeMinijobTax = &Ref{ID: 25227, ObjectName: ObjectAccountingType}
	// AccountingTypeMinijobWage — Löhne für Minijobs (SKR03 4195, SKR04 6035, kind DC).
	AccountingTypeMinijobWage = &Ref{ID: 25228, ObjectName: ObjectAccountingType}
	// AccountingTypeLiabilitiesWageSalary — Verbindlichkeiten aus Lohn und Gehalt (SKR03 1740, SKR04 3720).
	AccountingTypeLiabilitiesWageSalary = &Ref{ID: 25229, ObjectName: ObjectAccountingType}
	// AccountingTypeLiabilitiesWageChurchTax — Verbindlichkeiten aus Lohn- und Kirchensteuer (SKR03 1741, SKR04 3730).
	AccountingTypeLiabilitiesWageChurchTax = &Ref{ID: 25230, ObjectName: ObjectAccountingType}
	// AccountingTypeLiabilitiesSocialSecurity — Verbindlichkeiten im Rahmen der sozialen Sicherheit (SKR03 1742, SKR04 3740).
	AccountingTypeLiabilitiesSocialSecurity = &Ref{ID: 25231, ObjectName: ObjectAccountingType}
	// AccountingTypeWageSalaryBilling — Lohn- und Gehaltsverrechnung (SKR03 1755, SKR04 3790).
	AccountingTypeWageSalaryBilling = &Ref{ID: 25232, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelExpensesEmployee — Reisekosten Fahrtkosten Arbeitnehmener (SKR03 4663, SKR04 6663).
	AccountingTypeTravelExpensesEmployee = &Ref{ID: 25233, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelFoodEmployee — Reisekosten Verpflegungsmehraufwand Arbeitnehmer (SKR03 4664, SKR04 6664).
	AccountingTypeTravelFoodEmployee = &Ref{ID: 25234, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelOvernightEmployee — Reisekosten Übernachtungskosten Arbeitnehmer (SKR03 4666, SKR04 6660).
	AccountingTypeTravelOvernightEmployee = &Ref{ID: 25235, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelDrivingRefundEmployee — Reisekosten Kilometergelderstattung Arbeitnehmer (SKR03 4668, SKR04 6668).
	AccountingTypeTravelDrivingRefundEmployee = &Ref{ID: 25236, ObjectName: ObjectAccountingType}
	// AccountingTypeTravelDrivingRefundManager — Reisekosten Kilometergelderstattung Unternehmer (SKR03 4673, SKR04 6673).
	AccountingTypeTravelDrivingRefundManager = &Ref{ID: 25237, ObjectName: ObjectAccountingType}
	// AccountingTypeCnForeignServices — Fremdarbeiten/Fremdleistungen (SKR03 4909, SKR04 6303, kind IC).
	AccountingTypeCnForeignServices = &Ref{ID: 32890, ObjectName: ObjectAccountingType}
	// AccountingTypeCnSaleProvision — Verkaufsprovision (SKR03 4760, SKR04 6770, kind DC).
	AccountingTypeCnSaleProvision = &Ref{ID: 32891, ObjectName: ObjectAccountingType}
	// AccountingTypeNegativeSoldTangibleTaxAsset — Erlöse aus Verkäufen Sachanlagevermögen 19 % USt (SKR03 8801, SKR04 6885, kind E).
	AccountingTypeNegativeSoldTangibleTaxAsset = &Ref{ID: 33831, ObjectName: ObjectAccountingType}
	// AccountingTypeAccountingPositiveFixedAssetDisposals — Anlageabgänge Sachanlagen (SKR03 2315, SKR04 4855, kind E).
	AccountingTypeAccountingPositiveFixedAssetDisposals = &Ref{ID: 33832, ObjectName: ObjectAccountingType}
	// AccountingTypeNegativeFixedAssetDisposals — Anlageabgänge Sachanlagen (SKR03 2310, SKR04 6895, kind IC).
	AccountingTypeNegativeFixedAssetDisposals = &Ref{ID: 33833, ObjectName: ObjectAccountingType}
	// AccountingTypeOtherDues — Sonstige Abgaben (SKR03 4390, SKR04 6430).
	AccountingTypeOtherDues = &Ref{ID: 65231, ObjectName: ObjectAccountingType}
	// AccountingTypeFacilities — Mieten für Einrichtungen (bewegliche Wirtschaftsgüter) (SKR03 4960, SKR04 6835).
	AccountingTypeFacilities = &Ref{ID: 65232, ObjectName: ObjectAccountingType}
	// AccountingTypeVehicleLeasing — Mietleasing Kfz (SKR03 4570, SKR04 6560).
	AccountingTypeVehicleLeasing = &Ref{ID: 65233, ObjectName: ObjectAccountingType}
	// AccountingTypeCoronaSupport — Corona-Hilfe (SKR03 2743, SKR04 4975).
	AccountingTypeCoronaSupport = &Ref{ID: 662054, ObjectName: ObjectAccountingType}
)

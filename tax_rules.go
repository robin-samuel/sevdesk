package sevdesk

// TaxRule references for sevdesk-Update 2.0.
//
// A tax rule is the VAT regulation behind a document. In 2.0 it replaces the
// taxType/taxSet pair on vouchers, invoices, orders and credit notes; taxType
// "custom" (and the TaxSet that went with it) is no longer supported at all.
//
// Which rules you may use depends on the document side and on the client's tax
// setting:
//
//   - Revenue documents (invoices, orders, credit notes) take the revenue rules.
//   - Vouchers take the expense rules.
//   - Kleinunternehmer clients use [TaxRuleSmallBusiness] on the revenue side and
//     [TaxRuleNonDeductibleExpense] / [TaxRuleReverseCharge13bNonDeductible] on
//     the expense side.
//
// The rule also constrains the tax rates allowed on positions, noted per rule
// below. sevdesk answers HTTP 422 when rule, rate and booking account don't fit
// together; [ReceiptGuidanceService] reports the combinations each account
// accepts.
//
// Each value is a [*Ref] ready to drop into a TaxRule field, e.g.
// [VoucherCreateFields.TaxRule] or [InvoiceCreateFields.TaxRule].
//
// The ids and internal codes below were verified against
// /ReceiptGuidance/forAllAccounts on 2026-07-29 (a German Regelbesteuerer
// account). Rules 10, 11 and 18–20 are documented by sevdesk but were not
// offered by that client, so their ids come from the OpenAPI spec alone and
// their internal codes are unknown — see [TaxRuleName].
var (
	// TaxRuleTaxableRevenue — Umsatzsteuerpflichtige Umsätze (id 1, rates
	// 0/7/19), code USTPFL_UMS_EINN. The default for domestic sales; replaces
	// `"taxType": "default"` on revenue documents.
	TaxRuleTaxableRevenue = &Ref{ID: 1, ObjectName: ObjectTaxRule}
	// TaxRuleExports — Ausfuhren ins Drittland (id 2, rate 0), code AUSFUHREN.
	TaxRuleExports = &Ref{ID: 2, ObjectName: ObjectTaxRule}
	// TaxRuleIntraCommunitySupply — Lieferungen ins EU-Ausland (id 3, rates
	// 0/7/19), code INNERGEM_LIEF. Replaces `"taxType": "eu"`.
	TaxRuleIntraCommunitySupply = &Ref{ID: 3, ObjectName: ObjectTaxRule}
	// TaxRuleTaxFreeRevenue — Steuerfreie Umsätze §4 UStG (id 4, rate 0), code
	// STFREIE_UMS_P4.
	TaxRuleTaxFreeRevenue = &Ref{ID: 4, ObjectName: ObjectTaxRule}
	// TaxRuleReverseCharge13b — Reverse Charge gem. §13b UStG (id 5, rate 0),
	// code REV_CHARGE_13B_1. Field 60 of the VAT return.
	TaxRuleReverseCharge13b = &Ref{ID: 5, ObjectName: ObjectTaxRule}
	// TaxRuleNotTaxableDomestically — Nicht im Inland steuerbare Leistung
	// (id 17, rate 0), code NICHT_IM_INLAND_STEUERBAR — outside the EU, e.g.
	// Switzerland. Replaces `"taxType": "noteu"`. Not usable for advance,
	// partial or final invoices.
	TaxRuleNotTaxableDomestically = &Ref{ID: 17, ObjectName: ObjectTaxRule}
	// TaxRuleReverseChargeEUServices — services to businesses in other EU
	// states (id 21, rate 0), code REV_CHARGE_13B_1_USTG. sevdesk's written
	// documentation labels this one "Reverse Charge gem. §18b UStG (field 21 in
	// VAT return)". Not usable in vouchers or for advance, partial and final
	// invoices.
	TaxRuleReverseChargeEUServices = &Ref{ID: 21, ObjectName: ObjectTaxRule}
	// TaxRuleNonTaxableRevenue — nicht steuerbare Einnahmen (id 22, rate 0),
	// code NICHT_STEUERBAR_REVENUE, e.g. durchlaufende Posten. Not in sevdesk's
	// published tables; observed in the receipt guidance.
	TaxRuleNonTaxableRevenue = &Ref{ID: 22, ObjectName: ObjectTaxRule}
	// TaxRuleOneStopShopGoods — One Stop Shop, goods (id 18). Allowed rates
	// depend on the destination country. Not usable in vouchers, e-invoices,
	// invoices with a custom revenue account, or advance/partial/final invoices.
	TaxRuleOneStopShopGoods = &Ref{ID: 18, ObjectName: ObjectTaxRule}
	// TaxRuleOneStopShopElectronic — One Stop Shop, electronic service (id 19).
	// Same rate and usage caveats as [TaxRuleOneStopShopGoods].
	TaxRuleOneStopShopElectronic = &Ref{ID: 19, ObjectName: ObjectTaxRule}
	// TaxRuleOneStopShopOther — One Stop Shop, other service (id 20). Same rate
	// and usage caveats as [TaxRuleOneStopShopGoods].
	TaxRuleOneStopShopOther = &Ref{ID: 20, ObjectName: ObjectTaxRule}

	// TaxRuleIntraCommunityAcquisition — Innergemeinschaftliche Erwerbe (id 8,
	// rate 0), code INNERGEM_ERWERB. Expense side.
	TaxRuleIntraCommunityAcquisition = &Ref{ID: 8, ObjectName: ObjectTaxRule}
	// TaxRuleDeductibleExpense — Vorsteuerabziehbare Aufwendungen (id 9, rates
	// 0/7/19), code VORST_ABZUGSF_AUFW. The usual choice for a supplier bill;
	// replaces `"taxType": "default"` on vouchers.
	TaxRuleDeductibleExpense = &Ref{ID: 9, ObjectName: ObjectTaxRule}
	// TaxRuleNonDeductibleExpense — Nicht vorsteuerabziehbare Aufwendungen
	// (id 10, rate 0). Expense side; also replaces `"taxType": "ss"` for
	// Kleinunternehmer.
	TaxRuleNonDeductibleExpense = &Ref{ID: 10, ObjectName: ObjectTaxRule}
	// TaxRuleReverseCharge13b2Deductible — Reverse Charge gem. §13b Abs. 2 UStG
	// mit Vorsteuerabzug 0% (19%) (id 12, rate 0), code
	// REV_CHARGE_13B_MIT_VORST_ABZUG_0, e.g. services from outside the EU or
	// construction work. Expense side.
	TaxRuleReverseCharge13b2Deductible = &Ref{ID: 12, ObjectName: ObjectTaxRule}
	// TaxRuleReverseCharge13bNonDeductible — Reverse Charge gem. §13b UStG ohne
	// Vorsteuerabzug 0% (19%) (id 13, rate 0), code
	// REV_CHARGE_13B_OHNE_VORST_ABZUG_0, e.g. Kleinunternehmer, insurance
	// agents, doctors. Expense side.
	TaxRuleReverseCharge13bNonDeductible = &Ref{ID: 13, ObjectName: ObjectTaxRule}
	// TaxRuleReverseCharge13b1EU — Reverse Charge gem. §13b Abs. 1 EU Umsätze
	// 0% (19%) (id 14, rate 0), code REV_CHARGE_13B_EU_0, e.g. services from EU
	// states. Expense side.
	TaxRuleReverseCharge13b1EU = &Ref{ID: 14, ObjectName: ObjectTaxRule}
	// TaxRuleNonTaxableExpense — nicht steuerbare Aufwendungen (id 16, rate 0),
	// code NICHT_STEUERBAR_EXPENSE, e.g. Privateinlagen or Eigenverbrauch. Not
	// in sevdesk's published tables; observed in the receipt guidance.
	TaxRuleNonTaxableExpense = &Ref{ID: 16, ObjectName: ObjectTaxRule}
	// TaxRuleNonTaxableVAT — nicht steuerbare Steuerbewegungen (id 15, rate 0),
	// code NICHT_STEUERBAR_TAX, e.g. VAT prepayments and refunds. Not in
	// sevdesk's published tables; observed in the receipt guidance.
	TaxRuleNonTaxableVAT = &Ref{ID: 15, ObjectName: ObjectTaxRule}

	// TaxRuleSmallBusiness — Steuer nicht erhoben nach §19 UStG (id 11, rate 0).
	// The only revenue rule available to Kleinunternehmer; replaces
	// `"taxType": "ss"`.
	TaxRuleSmallBusiness = &Ref{ID: 11, ObjectName: ObjectTaxRule}
)

// TaxRuleName is sevdesk's internal code for a tax rule, as reported in
// [ReceiptGuideTaxRule.Name] and accepted by
// [ReceiptGuidanceService.ForTaxRule].
type TaxRuleName string

// TaxRuleName values, verified against /ReceiptGuidance/forAllAccounts on
// 2026-07-29. Each corresponds to the like-named rule above; codes for
// [TaxRuleNonDeductibleExpense], [TaxRuleSmallBusiness] and the One Stop Shop
// rules were not observed and so are not listed.
const (
	TaxRuleNameTaxableRevenue                = TaxRuleName("USTPFL_UMS_EINN")
	TaxRuleNameExports                       = TaxRuleName("AUSFUHREN")
	TaxRuleNameIntraCommunitySupply          = TaxRuleName("INNERGEM_LIEF")
	TaxRuleNameTaxFreeRevenue                = TaxRuleName("STFREIE_UMS_P4")
	TaxRuleNameReverseCharge13b              = TaxRuleName("REV_CHARGE_13B_1")
	TaxRuleNameNotTaxableDomestically        = TaxRuleName("NICHT_IM_INLAND_STEUERBAR")
	TaxRuleNameReverseChargeEUServices       = TaxRuleName("REV_CHARGE_13B_1_USTG")
	TaxRuleNameNonTaxableRevenue             = TaxRuleName("NICHT_STEUERBAR_REVENUE")
	TaxRuleNameIntraCommunityAcquisition     = TaxRuleName("INNERGEM_ERWERB")
	TaxRuleNameDeductibleExpense             = TaxRuleName("VORST_ABZUGSF_AUFW")
	TaxRuleNameReverseCharge13b2Deductible   = TaxRuleName("REV_CHARGE_13B_MIT_VORST_ABZUG_0")
	TaxRuleNameReverseCharge13bNonDeductible = TaxRuleName("REV_CHARGE_13B_OHNE_VORST_ABZUG_0")
	TaxRuleNameReverseCharge13b1EU           = TaxRuleName("REV_CHARGE_13B_EU_0")
	TaxRuleNameNonTaxableExpense             = TaxRuleName("NICHT_STEUERBAR_EXPENSE")
	TaxRuleNameNonTaxableVAT                 = TaxRuleName("NICHT_STEUERBAR_TAX")
)

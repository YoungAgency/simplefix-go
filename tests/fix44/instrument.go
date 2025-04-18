// Code generated by fixgen. DO NOT EDIT.

package fix44

import (
	"github.com/YoungAgency/simplefix-go/fix"
)

type Instrument struct {
	*fix.Component
}

func makeInstrument() *Instrument {
	return &Instrument{fix.NewComponent(
		fix.NewKeyValue(FieldSymbol, &fix.String{}),
		fix.NewKeyValue(FieldSymbolSfx, &fix.String{}),
		fix.NewKeyValue(FieldSecurityID, &fix.String{}),
		fix.NewKeyValue(FieldSecurityIDSource, &fix.String{}),
		NewSecurityAltIDGrp().Group,
		fix.NewKeyValue(FieldProduct, &fix.String{}),
		fix.NewKeyValue(FieldCFICode, &fix.String{}),
		fix.NewKeyValue(FieldSecurityType, &fix.String{}),
		fix.NewKeyValue(FieldSecuritySubType, &fix.String{}),
		fix.NewKeyValue(FieldMaturityMonthYear, &fix.String{}),
		fix.NewKeyValue(FieldMaturityDate, &fix.String{}),
		fix.NewKeyValue(FieldCouponPaymentDate, &fix.String{}),
		fix.NewKeyValue(FieldIssueDate, &fix.String{}),
		fix.NewKeyValue(FieldRepoCollateralSecurityType, &fix.Int{}),
		fix.NewKeyValue(FieldRepurchaseTerm, &fix.Int{}),
		fix.NewKeyValue(FieldRepurchaseRate, &fix.Float{}),
		fix.NewKeyValue(FieldFactor, &fix.Float{}),
		fix.NewKeyValue(FieldCreditRating, &fix.String{}),
		fix.NewKeyValue(FieldInstrRegistry, &fix.String{}),
		fix.NewKeyValue(FieldCountryOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldStateOrProvinceOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldLocaleOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldRedemptionDate, &fix.String{}),
		fix.NewKeyValue(FieldStrikePrice, &fix.Float{}),
		fix.NewKeyValue(FieldStrikeCurrency, &fix.String{}),
		fix.NewKeyValue(FieldOptAttribute, &fix.String{}),
		fix.NewKeyValue(FieldContractMultiplier, &fix.Float{}),
		fix.NewKeyValue(FieldCouponRate, &fix.Float{}),
		fix.NewKeyValue(FieldSecurityExchange, &fix.String{}),
		fix.NewKeyValue(FieldIssuer, &fix.String{}),
		fix.NewKeyValue(FieldEncodedIssuerLen, &fix.Int{}),
		fix.NewKeyValue(FieldEncodedIssuer, &fix.String{}),
		fix.NewKeyValue(FieldSecurityDesc, &fix.String{}),
		fix.NewKeyValue(FieldEncodedSecurityDescLen, &fix.Int{}),
		fix.NewKeyValue(FieldEncodedSecurityDesc, &fix.String{}),
		fix.NewKeyValue(FieldPool, &fix.String{}),
		fix.NewKeyValue(FieldContractSettlMonth, &fix.String{}),
		fix.NewKeyValue(FieldCPProgram, &fix.String{}),
		fix.NewKeyValue(FieldCPRegType, &fix.String{}),
		NewEventsGrp().Group,
		fix.NewKeyValue(FieldDatedDate, &fix.String{}),
		fix.NewKeyValue(FieldInterestAccrualDate, &fix.String{}),
	)}
}

func NewInstrument() *Instrument {
	return makeInstrument()
}

func (instrument *Instrument) Symbol() string {
	kv := instrument.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSymbol(symbol string) *Instrument {
	kv := instrument.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(symbol)
	return instrument
}

func (instrument *Instrument) SymbolSfx() string {
	kv := instrument.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSymbolSfx(symbolSfx string) *Instrument {
	kv := instrument.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(symbolSfx)
	return instrument
}

func (instrument *Instrument) SecurityID() string {
	kv := instrument.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecurityID(securityID string) *Instrument {
	kv := instrument.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(securityID)
	return instrument
}

func (instrument *Instrument) SecurityIDSource() string {
	kv := instrument.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecurityIDSource(securityIDSource string) *Instrument {
	kv := instrument.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(securityIDSource)
	return instrument
}

func (instrument *Instrument) SecurityAltIDGrp() *SecurityAltIDGrp {
	group := instrument.Get(4).(*fix.Group)

	return &SecurityAltIDGrp{group}
}

func (instrument *Instrument) SetSecurityAltIDGrp(noSecurityAltID *SecurityAltIDGrp) *Instrument {
	instrument.Set(4, noSecurityAltID.Group)

	return instrument
}

func (instrument *Instrument) Product() string {
	kv := instrument.Get(5)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetProduct(product string) *Instrument {
	kv := instrument.Get(5).(*fix.KeyValue)
	_ = kv.Load().Set(product)
	return instrument
}

func (instrument *Instrument) CFICode() string {
	kv := instrument.Get(6)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCFICode(cFICode string) *Instrument {
	kv := instrument.Get(6).(*fix.KeyValue)
	_ = kv.Load().Set(cFICode)
	return instrument
}

func (instrument *Instrument) SecurityType() string {
	kv := instrument.Get(7)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecurityType(securityType string) *Instrument {
	kv := instrument.Get(7).(*fix.KeyValue)
	_ = kv.Load().Set(securityType)
	return instrument
}

func (instrument *Instrument) SecuritySubType() string {
	kv := instrument.Get(8)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecuritySubType(securitySubType string) *Instrument {
	kv := instrument.Get(8).(*fix.KeyValue)
	_ = kv.Load().Set(securitySubType)
	return instrument
}

func (instrument *Instrument) MaturityMonthYear() string {
	kv := instrument.Get(9)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetMaturityMonthYear(maturityMonthYear string) *Instrument {
	kv := instrument.Get(9).(*fix.KeyValue)
	_ = kv.Load().Set(maturityMonthYear)
	return instrument
}

func (instrument *Instrument) MaturityDate() string {
	kv := instrument.Get(10)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetMaturityDate(maturityDate string) *Instrument {
	kv := instrument.Get(10).(*fix.KeyValue)
	_ = kv.Load().Set(maturityDate)
	return instrument
}

func (instrument *Instrument) CouponPaymentDate() string {
	kv := instrument.Get(11)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCouponPaymentDate(couponPaymentDate string) *Instrument {
	kv := instrument.Get(11).(*fix.KeyValue)
	_ = kv.Load().Set(couponPaymentDate)
	return instrument
}

func (instrument *Instrument) IssueDate() string {
	kv := instrument.Get(12)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetIssueDate(issueDate string) *Instrument {
	kv := instrument.Get(12).(*fix.KeyValue)
	_ = kv.Load().Set(issueDate)
	return instrument
}

func (instrument *Instrument) RepoCollateralSecurityType() int {
	kv := instrument.Get(13)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrument *Instrument) SetRepoCollateralSecurityType(repoCollateralSecurityType int) *Instrument {
	kv := instrument.Get(13).(*fix.KeyValue)
	_ = kv.Load().Set(repoCollateralSecurityType)
	return instrument
}

func (instrument *Instrument) RepurchaseTerm() int {
	kv := instrument.Get(14)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrument *Instrument) SetRepurchaseTerm(repurchaseTerm int) *Instrument {
	kv := instrument.Get(14).(*fix.KeyValue)
	_ = kv.Load().Set(repurchaseTerm)
	return instrument
}

func (instrument *Instrument) RepurchaseRate() float64 {
	kv := instrument.Get(15)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrument *Instrument) SetRepurchaseRate(repurchaseRate float64) *Instrument {
	kv := instrument.Get(15).(*fix.KeyValue)
	_ = kv.Load().Set(repurchaseRate)
	return instrument
}

func (instrument *Instrument) Factor() float64 {
	kv := instrument.Get(16)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrument *Instrument) SetFactor(factor float64) *Instrument {
	kv := instrument.Get(16).(*fix.KeyValue)
	_ = kv.Load().Set(factor)
	return instrument
}

func (instrument *Instrument) CreditRating() string {
	kv := instrument.Get(17)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCreditRating(creditRating string) *Instrument {
	kv := instrument.Get(17).(*fix.KeyValue)
	_ = kv.Load().Set(creditRating)
	return instrument
}

func (instrument *Instrument) InstrRegistry() string {
	kv := instrument.Get(18)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetInstrRegistry(instrRegistry string) *Instrument {
	kv := instrument.Get(18).(*fix.KeyValue)
	_ = kv.Load().Set(instrRegistry)
	return instrument
}

func (instrument *Instrument) CountryOfIssue() string {
	kv := instrument.Get(19)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCountryOfIssue(countryOfIssue string) *Instrument {
	kv := instrument.Get(19).(*fix.KeyValue)
	_ = kv.Load().Set(countryOfIssue)
	return instrument
}

func (instrument *Instrument) StateOrProvinceOfIssue() string {
	kv := instrument.Get(20)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetStateOrProvinceOfIssue(stateOrProvinceOfIssue string) *Instrument {
	kv := instrument.Get(20).(*fix.KeyValue)
	_ = kv.Load().Set(stateOrProvinceOfIssue)
	return instrument
}

func (instrument *Instrument) LocaleOfIssue() string {
	kv := instrument.Get(21)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetLocaleOfIssue(localeOfIssue string) *Instrument {
	kv := instrument.Get(21).(*fix.KeyValue)
	_ = kv.Load().Set(localeOfIssue)
	return instrument
}

func (instrument *Instrument) RedemptionDate() string {
	kv := instrument.Get(22)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetRedemptionDate(redemptionDate string) *Instrument {
	kv := instrument.Get(22).(*fix.KeyValue)
	_ = kv.Load().Set(redemptionDate)
	return instrument
}

func (instrument *Instrument) StrikePrice() float64 {
	kv := instrument.Get(23)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrument *Instrument) SetStrikePrice(strikePrice float64) *Instrument {
	kv := instrument.Get(23).(*fix.KeyValue)
	_ = kv.Load().Set(strikePrice)
	return instrument
}

func (instrument *Instrument) StrikeCurrency() string {
	kv := instrument.Get(24)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetStrikeCurrency(strikeCurrency string) *Instrument {
	kv := instrument.Get(24).(*fix.KeyValue)
	_ = kv.Load().Set(strikeCurrency)
	return instrument
}

func (instrument *Instrument) OptAttribute() string {
	kv := instrument.Get(25)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetOptAttribute(optAttribute string) *Instrument {
	kv := instrument.Get(25).(*fix.KeyValue)
	_ = kv.Load().Set(optAttribute)
	return instrument
}

func (instrument *Instrument) ContractMultiplier() float64 {
	kv := instrument.Get(26)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrument *Instrument) SetContractMultiplier(contractMultiplier float64) *Instrument {
	kv := instrument.Get(26).(*fix.KeyValue)
	_ = kv.Load().Set(contractMultiplier)
	return instrument
}

func (instrument *Instrument) CouponRate() float64 {
	kv := instrument.Get(27)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrument *Instrument) SetCouponRate(couponRate float64) *Instrument {
	kv := instrument.Get(27).(*fix.KeyValue)
	_ = kv.Load().Set(couponRate)
	return instrument
}

func (instrument *Instrument) SecurityExchange() string {
	kv := instrument.Get(28)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecurityExchange(securityExchange string) *Instrument {
	kv := instrument.Get(28).(*fix.KeyValue)
	_ = kv.Load().Set(securityExchange)
	return instrument
}

func (instrument *Instrument) Issuer() string {
	kv := instrument.Get(29)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetIssuer(issuer string) *Instrument {
	kv := instrument.Get(29).(*fix.KeyValue)
	_ = kv.Load().Set(issuer)
	return instrument
}

func (instrument *Instrument) EncodedIssuerLen() int {
	kv := instrument.Get(30)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrument *Instrument) SetEncodedIssuerLen(encodedIssuerLen int) *Instrument {
	kv := instrument.Get(30).(*fix.KeyValue)
	_ = kv.Load().Set(encodedIssuerLen)
	return instrument
}

func (instrument *Instrument) EncodedIssuer() string {
	kv := instrument.Get(31)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetEncodedIssuer(encodedIssuer string) *Instrument {
	kv := instrument.Get(31).(*fix.KeyValue)
	_ = kv.Load().Set(encodedIssuer)
	return instrument
}

func (instrument *Instrument) SecurityDesc() string {
	kv := instrument.Get(32)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetSecurityDesc(securityDesc string) *Instrument {
	kv := instrument.Get(32).(*fix.KeyValue)
	_ = kv.Load().Set(securityDesc)
	return instrument
}

func (instrument *Instrument) EncodedSecurityDescLen() int {
	kv := instrument.Get(33)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrument *Instrument) SetEncodedSecurityDescLen(encodedSecurityDescLen int) *Instrument {
	kv := instrument.Get(33).(*fix.KeyValue)
	_ = kv.Load().Set(encodedSecurityDescLen)
	return instrument
}

func (instrument *Instrument) EncodedSecurityDesc() string {
	kv := instrument.Get(34)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetEncodedSecurityDesc(encodedSecurityDesc string) *Instrument {
	kv := instrument.Get(34).(*fix.KeyValue)
	_ = kv.Load().Set(encodedSecurityDesc)
	return instrument
}

func (instrument *Instrument) Pool() string {
	kv := instrument.Get(35)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetPool(pool string) *Instrument {
	kv := instrument.Get(35).(*fix.KeyValue)
	_ = kv.Load().Set(pool)
	return instrument
}

func (instrument *Instrument) ContractSettlMonth() string {
	kv := instrument.Get(36)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetContractSettlMonth(contractSettlMonth string) *Instrument {
	kv := instrument.Get(36).(*fix.KeyValue)
	_ = kv.Load().Set(contractSettlMonth)
	return instrument
}

func (instrument *Instrument) CPProgram() string {
	kv := instrument.Get(37)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCPProgram(cPProgram string) *Instrument {
	kv := instrument.Get(37).(*fix.KeyValue)
	_ = kv.Load().Set(cPProgram)
	return instrument
}

func (instrument *Instrument) CPRegType() string {
	kv := instrument.Get(38)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetCPRegType(cPRegType string) *Instrument {
	kv := instrument.Get(38).(*fix.KeyValue)
	_ = kv.Load().Set(cPRegType)
	return instrument
}

func (instrument *Instrument) EventsGrp() *EventsGrp {
	group := instrument.Get(39).(*fix.Group)

	return &EventsGrp{group}
}

func (instrument *Instrument) SetEventsGrp(noEvents *EventsGrp) *Instrument {
	instrument.Set(39, noEvents.Group)

	return instrument
}

func (instrument *Instrument) DatedDate() string {
	kv := instrument.Get(40)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetDatedDate(datedDate string) *Instrument {
	kv := instrument.Get(40).(*fix.KeyValue)
	_ = kv.Load().Set(datedDate)
	return instrument
}

func (instrument *Instrument) InterestAccrualDate() string {
	kv := instrument.Get(41)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrument *Instrument) SetInterestAccrualDate(interestAccrualDate string) *Instrument {
	kv := instrument.Get(41).(*fix.KeyValue)
	_ = kv.Load().Set(interestAccrualDate)
	return instrument
}

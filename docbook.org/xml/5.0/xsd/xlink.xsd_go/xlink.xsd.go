//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		docbook.org/xml/5.0/xsd/xlink.xsd
package go_Xlink



import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasAttr_Href struct {
	Href xsdt.String `xml:"http://www.w3.org/1999/xlink href,attr"`

}

type XsdGoPkgHasAttr_Type struct {
	Type xsdt.String `xml:"http://www.w3.org/1999/xlink type,attr"`

}

type XsdGoPkgHasAttr_Role struct {
	Role xsdt.String `xml:"http://www.w3.org/1999/xlink role,attr"`

}

type XsdGoPkgHasAttr_Arcrole struct {
	Arcrole xsdt.String `xml:"http://www.w3.org/1999/xlink arcrole,attr"`

}

type XsdGoPkgHasAttr_Title struct {
	Title xsdt.String `xml:"http://www.w3.org/1999/xlink title,attr"`

}

type TxsdShow xsdt.Token

//	This convenience method just performs a simple type conversion to TxsdShow's alias type xsdt.Token.
func (me TxsdShow) ToXsdtToken () xsdt.Token { return xsdt.Token(me) }

//	Returns true if the value of this enumerated TxsdShow is "other".
func (me TxsdShow) IsOther () bool { return me == "other" }

//	Returns true if the value of this enumerated TxsdShow is "none".
func (me TxsdShow) IsNone () bool { return me == "none" }

//	Returns true if the value of this enumerated TxsdShow is "new".
func (me TxsdShow) IsNew () bool { return me == "new" }

//	Since TxsdShow is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdShow) SetFromString (s string)  { (*xsdt.Token)(me).SetFromString(s) }

//	Since TxsdShow is just a simple String type, this merely returns the current string value.
func (me TxsdShow) String () string { return xsdt.Token(me).String() }

//	Returns true if the value of this enumerated TxsdShow is "embed".
func (me TxsdShow) IsEmbed () bool { return me == "embed" }

//	Returns true if the value of this enumerated TxsdShow is "replace".
func (me TxsdShow) IsReplace () bool { return me == "replace" }

type XsdGoPkgHasAttr_Show struct {
	Show TxsdShow `xml:"http://www.w3.org/1999/xlink show,attr"`

}

type TxsdActuate xsdt.Token

//	Since TxsdActuate is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdActuate) SetFromString (s string)  { (*xsdt.Token)(me).SetFromString(s) }

//	Since TxsdActuate is just a simple String type, this merely returns the current string value.
func (me TxsdActuate) String () string { return xsdt.Token(me).String() }

//	Returns true if the value of this enumerated TxsdActuate is "none".
func (me TxsdActuate) IsNone () bool { return me == "none" }

//	This convenience method just performs a simple type conversion to TxsdActuate's alias type xsdt.Token.
func (me TxsdActuate) ToXsdtToken () xsdt.Token { return xsdt.Token(me) }

//	Returns true if the value of this enumerated TxsdActuate is "other".
func (me TxsdActuate) IsOther () bool { return me == "other" }

//	Returns true if the value of this enumerated TxsdActuate is "onRequest".
func (me TxsdActuate) IsOnRequest () bool { return me == "onRequest" }

//	Returns true if the value of this enumerated TxsdActuate is "onLoad".
func (me TxsdActuate) IsOnLoad () bool { return me == "onLoad" }

type XsdGoPkgHasAttr_Actuate struct {
	Actuate TxsdActuate `xml:"http://www.w3.org/1999/xlink actuate,attr"`

}

type XsdGoPkgHasAttr_Label struct {
	Label xsdt.Nmtoken `xml:"http://www.w3.org/1999/xlink label,attr"`

}

type XsdGoPkgHasAttr_From struct {
	From xsdt.Nmtoken `xml:"http://www.w3.org/1999/xlink from,attr"`

}

type XsdGoPkgHasAttr_To struct {
	To xsdt.Nmtoken `xml:"http://www.w3.org/1999/xlink to,attr"`

}

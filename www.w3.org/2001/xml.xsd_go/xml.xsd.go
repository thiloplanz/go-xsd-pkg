//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/2001/xml.xsd
package go_Xml



import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type TxsdLang xsdt.String

//	Since TxsdLang is just a simple String type, this merely returns the current string value.
func (me TxsdLang) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdLang's alias type xsdt.String.
func (me TxsdLang) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Since TxsdLang is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdLang) SetFromString (s string)  { (*xsdt.String)(me).SetFromString(s) }

type XsdGoPkgHasAttr_Lang struct {
	Lang TxsdLang `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`

}

type TxsdSpace xsdt.NCName

//	Returns true if the value of this enumerated TxsdSpace is "preserve".
func (me TxsdSpace) IsPreserve () bool { return me == "preserve" }

//	This convenience method just performs a simple type conversion to TxsdSpace's alias type xsdt.NCName.
func (me TxsdSpace) ToXsdtNCName () xsdt.NCName { return xsdt.NCName(me) }

//	Since TxsdSpace is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdSpace) SetFromString (s string)  { (*xsdt.NCName)(me).SetFromString(s) }

//	Returns true if the value of this enumerated TxsdSpace is "default".
func (me TxsdSpace) IsDefault () bool { return me == "default" }

//	Since TxsdSpace is just a simple String type, this merely returns the current string value.
func (me TxsdSpace) String () string { return xsdt.NCName(me).String() }

type XsdGoPkgHasAttr_Space struct {
	Space TxsdSpace `xml:"http://www.w3.org/XML/1998/namespace space,attr"`

}

type XsdGoPkgHasAttr_Base struct {
	Base xsdt.AnyURI `xml:"http://www.w3.org/XML/1998/namespace base,attr"`

}

type XsdGoPkgHasAttr_Id struct {
	Id xsdt.Id `xml:"http://www.w3.org/XML/1998/namespace id,attr"`

}

type XsdGoPkgHasAtts_SpecialAttrs struct {
	XsdGoPkgHasAttr_Base

	XsdGoPkgHasAttr_Lang

	XsdGoPkgHasAttr_Space

	XsdGoPkgHasAttr_Id

}

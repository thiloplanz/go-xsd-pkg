//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/common/xlink-href.xsd
package go_XlinkHref

//	This schema provides the XLink href attribute for use in the MathML2
//	schema. Written by Max Froumentin, W3C.

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasAttr_Href struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/1999/xlink href,attr"`
}

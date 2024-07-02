package signxml

import "encoding/xml"

type Signature struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue"`
	KeyInfo        KeyInfo
}

type Algorithm struct {
	Algorithm string `xml:",attr"`
}

type SignedInfo struct {
	XMLName                xml.Name  `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo"`
	CanonicalizationMethod Algorithm `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod"`
	SignatureMethod        Algorithm `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod"`
	Reference              Reference
}

type Reference struct {
	XMLName      xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Reference"`
	URI          string   `xml:",attr,omitempty"`
	Transforms   Transforms
	DigestMethod Algorithm `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	DigestValue  string    `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
}

type Transforms struct {
	XMLName   xml.Name    `xml:"http://www.w3.org/2000/09/xmldsig# Transforms"`
	Transform []Algorithm `xml:"http://www.w3.org/2000/09/xmldsig# Transform"`
}

type KeyInfo struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo"`
	X509Data *X509Data
	Children []interface{}
}

type X509Data struct {
	XMLName         xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`
	X509Certificate string   `xml:"http://www.w3.org/2000/09/xmldsig# X509Certificate"`
}

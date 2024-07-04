package signxml

import "encoding/xml"

type Signature struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        KeyInfo
}

type Algorithm struct {
	Algorithm string `xml:",attr"`
}

type SignedInfo struct {
	XMLName                xml.Name  `xml:"SignedInfo"`
	CanonicalizationMethod Algorithm `xml:"CanonicalizationMethod"`
	SignatureMethod        Algorithm `xml:"SignatureMethod"`
	Reference              Reference
}

type Reference struct {
	XMLName      xml.Name `xml:"Reference"`
	URI          string   `xml:",attr,omitempty"`
	Transforms   Transforms
	DigestMethod Algorithm `xml:"DigestMethod"`
	DigestValue  string    `xml:"DigestValue"`
}

type Transforms struct {
	XMLName   xml.Name    `xml:"Transforms"`
	Transform []Algorithm `xml:"Transform"`
}

type KeyInfo struct {
	XMLName  xml.Name `xml:"KeyInfo"`
	X509Data *X509Data
	Children []interface{}
}

type X509Data struct {
	XMLName         xml.Name `xml:"X509Data"`
	X509Certificate string   `xml:"X509Certificate"`
}

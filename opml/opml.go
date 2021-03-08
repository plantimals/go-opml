// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package opml provides all the required structures and functions for parsing
OPML files, as defined by the specification of the OPML format:

	[OPML 1.0] http://dev.opml.org/spec1.html
	[OPML 2.0] http://dev.opml.org/spec2.html

It is able to parse both, OPML 1.0 and OPML 2.0, files.
*/
package opml

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// OPML is the root node of an OPML document. It only has a single required
// attribute: the version.
type OPML struct {
	XMLName xml.Name `xml:"opml" json:"opml"`
	Version string   `xml:"version,attr" json:"version"`
	Head    Head     `xml:"head" json:"head"`
	Body    Body     `xml:"body" json:"body"`
}

// Head holds some meta information about the document.
type Head struct {
	Title           string `xml:"title" json:"title"`
	DateCreated     string `xml:"dateCreated,omitempty" json:"dateCreated,omitempty"`
	DateModified    string `xml:"dateModified,omitempty" json:"dateModified,omitempty"`
	OwnerName       string `xml:"ownerName,omitempty" json:"ownerName,omitempty"`
	OwnerEmail      string `xml:"ownerEmail,omitempty" json:"ownerEmail,omitempty"`
	OwnerID         string `xml:"ownerId,omitempty" json:"ownerId,omitempty"`
	Docs            string `xml:"docs,omitempty" json:"docs,omitempty"`
	ExpansionState  string `xml:"expansionState,omitempty" json:"expansionState,omitempty"`
	VertScrollState string `xml:"vertScrollState,omitempty json:"verScrollState,omitempty"`
	WindowTop       string `xml:"windowTop,omitempty" json:"windowTop,omitempty"`
	WindowBottom    string `xml:"windowBottom,omitempty" json:"windowBottom,omitempty"`
	WindowLeft      string `xml:"windowLeft,omitempty" json:"windowLeft,omitempty"`
	WindowRight     string `xml:"windowRight,omitempty" json:"windowRight,omitempty"`
}

// Body is the parent structure of all outlines.
type Body struct {
	Outlines []Outline `xml:"outline" json:"outline"`
}

// Outline holds all information about an outline.
type Outline struct {
	Outlines     []Outline `xml:"outline" json:"outline"`
	Text         string    `xml:"text,attr" json:"text"`
	Type         string    `xml:"type,attr,omitempty" json:"type,omitempty"`
	IsComment    string    `xml:"isComment,attr,omitempty" json:"isComment,omitempty"`
	IsBreakpoint string    `xml:"isBreakpoint,attr,omitempty" json:"isBreakpoint,omitempty"`
	Created      string    `xml:"created,attr,omitempty" json:"created,omitempty"`
	Category     string    `xml:"category,attr,omitempty" json:"category,omitempty"`
	XMLURL       string    `xml:"xmlUrl,attr,omitempty" json:"xmlUrl,omitempty"`
	HTMLURL      string    `xml:"htmlUrl,attr,omitempty" json:"htmlUrl,omitempty"`
	URL          string    `xml:"url,attr,omitempty" json:"url,omitempty"`
	Language     string    `xml:"language,attr,omitempty" json:"language,omitempty"`
	Title        string    `xml:"title,attr,omitempty" json:"title,omitempty"`
	Version      string    `xml:"version,attr,omitempty" json:"version,omitempty"`
	Description  string    `xml:"description,attr,omitempty" json:"description,omitempty"`
}

// NewOPML creates a new OPML structure from a slice of bytes.
func NewOPML(b []byte) (*OPML, error) {
	var root OPML
	err := xml.Unmarshal(b, &root)
	if err != nil {
		return nil, err
	}

	return &root, nil
}

// NewOPMLFromURL creates a new OPML structure from an URL.
func NewOPMLFromURL(url string) (*OPML, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return NewOPML(b)
}

// NewOPMLFromFile creates a new OPML structure from a file.
func NewOPMLFromFile(filePath string) (*OPML, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return NewOPML(b)
}

// Outlines returns a slice of the outlines.
func (doc OPML) Outlines() []Outline {
	return doc.Body.Outlines
}

// XML exports the OPML document to a XML string.
func (doc OPML) XML() (string, error) {
	b, err := xml.MarshalIndent(doc, "", "\t")
	return xml.Header + string(b), err
}

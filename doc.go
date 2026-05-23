// Package sevdesk is a client for the sevdesk REST API (https://api.sevdesk.de/).
//
// Create a client with an API key and call methods on its services:
//
//	c := sevdesk.New("YOUR_API_KEY")
//	contact, err := c.Contacts.Get(ctx, 42)
//
// Optional fields on *Params types are pointers; use the [String], [Int64],
// [Bool], [Float64], and [Time] helpers to set them.
package sevdesk

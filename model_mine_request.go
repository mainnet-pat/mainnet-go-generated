/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in active development, breaking changes may be made prior to official release of version 1.  **Important:** This library is in active development 
 *
 * API version: 0.0.2
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MineRequest struct for MineRequest
type MineRequest struct {
	Cashaddr string `json:"cashaddr,omitempty"`
	// the number of blocks to mine
	Blocks float32 `json:"blocks,omitempty"`
}

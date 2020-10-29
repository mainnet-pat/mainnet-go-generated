/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in active development, breaking changes may be made prior to official release of version 1.  **Important:** This library is in active development 
 *
 * API version: 0.0.1-rc
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Utxo struct for Utxo
type Utxo struct {
	Index float32 `json:"index,omitempty"`
	// The hash of a transaction
	TxId string `json:"txId"`
	Value float32 `json:"value,omitempty"`
	// Unit of account.
	Unit string `json:"unit,omitempty"`
	// serialized outpoint
	UtxoId string `json:"utxoId"`
}
// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// *************************************
// *** GENERATED FILE: DO NOT MODIFY ***
// *************************************

// generated by:   generate-fault.sh
// generated on:   2019-10-15 18:14:49

package fault

import "errors"

func e(s string) error {
	return errors.New(s)
}
// auto generated error vars: *** DO NOT MODIFY ***

var (
	AddressIsNil                          = e("address is nil")
	AlreadyInitialised                    = e("already initialised")
	AssetFingerprintIsRequired            = e("asset fingerprint is required")
	AssetIsNotIndexed                     = e("asset is not indexed")
	AssetMetadataIsRequired               = e("asset metadata is required")
	AssetMetadataMustBeMap                = e("asset metadata must be map")
	AssetNotFound                         = e("asset not found")
	AssetsAlreadyRegistered               = e("assets already registered")
	BitcoinAddressForWrongNetwork         = e("bitcoin address for wrong network")
	BitcoinAddressIsNotSupported          = e("bitcoin address is not supported")
	BlockAlreadyProcessed                 = e("block already processed")
	BlockEndEarlierThanBegin              = e("block end earlier than begin")
	BlockHeaderNotFound                   = e("block header not found")
	BlockHeightNotFound                   = e("block height not found")
	BlockIsTooOld                         = e("block is too old")
	BlockNotFound                         = e("block not found")
	BlockVersionMustNotDecrease           = e("block version must not decrease")
	BufferCapacityLimit                   = e("buffer capacity limit")
	CannotConvertSharesBackToAssets       = e("cannot convert shares back to assets")
	CannotDecodeAccount                   = e("cannot decode account")
	CannotDecodePrivateKey                = e("cannot decode private key")
	CannotDecodeSeed                      = e("cannot decode seed")
	CanOnlyConvertAssetsToShares          = e("can only convert assets to shares")
	CertificateFileAlreadyExists          = e("certificate file already exists")
	ChecksumMismatch                      = e("checksum mismatch")
	ClientSocketNotConnected              = e("client socket not connected")
	ClientSocketNotCreated                = e("client socket not created")
	ConnectingToSelfForbidden             = e("connecting to self forbidden")
	ConnectIsRequired                     = e("connect is required")
	ConnectRequiresPortNumberSuffix       = e("connect requires port number suffix")
	CryptoFailed                          = e("crypto failed")
	CurrencyAddressIsRequired             = e("currency address is required")
	CurrencyIsNotSupportedByProofer       = e("currency is not supported by proofer")
	DatabaseIsNotSet                      = e("database is not set")
	DataInconsistent                      = e("data inconsistent")
	DescriptionIsRequired                 = e("description is required")
	DifficultyDoesNotMatchCalculated      = e("difficulty does not match calculated")
	DoubleTransferAttempt                 = e("double transfer attempt")
	FileDoesNotExist                      = e("file does not exist")
	FileNameIsRequired                    = e("file name is required")
	FingerprintTooLong                    = e("fingerprint too long")
	FingerprintTooShort                   = e("fingerprint too short")
	HashCannotBeNil                       = e("hash cannot be nil")
	HashNotFound                          = e("hash not found")
	HeightOutOfSequence                   = e("height out of sequence")
	IdentityNameAlreadyExists             = e("identity name already exists")
	IdentityNameIsRequired                = e("identity name is required")
	IdentityNameNotFound                  = e("identity name not found")
	IncompatibleOptions                   = e("incompatible options")
	IncorrectBlockRangeToRollback         = e("incorrect block range to rollback")
	IncorrectChain                        = e("incorrect chain")
	InsufficientShares                    = e("insufficient shares")
	InvalidBitcoinAddress                 = e("invalid bitcoin address")
	InvalidBlockHeaderDifficulty          = e("invalid block header difficulty")
	InvalidBlockHeaderSize                = e("invalid block header size")
	InvalidBlockHeaderTimestamp           = e("invalid block header timestamp")
	InvalidBlockHeaderVersion             = e("invalid block header version")
	InvalidBuffer                         = e("invalid buffer")
	InvalidChain                          = e("invalid chain")
	InvalidCount                          = e("invalid count")
	InvalidCurrency                       = e("invalid currency")
	InvalidCurrencyAddress                = e("invalid currency address")
	InvalidCursor                         = e("invalid cursor")
	InvalidDnsTxtRecord                   = e("invalid dns txt record")
	InvalidFingerprint                    = e("invalid fingerprint")
	InvalidIdentityName                   = e("invalid identity name")
	InvalidIpAddress                      = e("invalid ip address")
	InvalidItem                           = e("invalid item")
	InvalidKeyLength                      = e("invalid key length")
	InvalidKeyType                        = e("invalid key type")
	InvalidLength                         = e("invalid length")
	InvalidLitecoinAddress                = e("invalid litecoin address")
	InvalidNodeDomain                     = e("invalid node domain")
	InvalidNonce                          = e("invalid nonce")
	InvalidOwnerOrRegistrant              = e("invalid owner or registrant")
	InvalidPasswordLength                 = e("invalid password length")
	InvalidPaymentVersion                 = e("invalid payment version")
	InvalidPeerResponse                   = e("invalid peer response")
	InvalidPortNumber                     = e("invalid port number")
	InvalidPrivateKey                     = e("invalid private key")
	InvalidProofSigningKey                = e("invalid proof signing key")
	InvalidPublicKey                      = e("invalid public key")
	InvalidRecoveryPhraseLength           = e("invalid recovery phrase length")
	InvalidSecretKeyLength                = e("invalid secret key length")
	InvalidSeedHeader                     = e("invalid seed header")
	InvalidSeedLength                     = e("invalid seed length")
	InvalidSignature                      = e("invalid signature")
	InvalidTimestamp                      = e("invalid timestamp")
	KeyFileAlreadyExists                  = e("key file already exists")
	LinkToInvalidOrUnconfirmedTransaction = e("link to invalid or unconfirmed transaction")
	LitecoinAddressForWrongNetwork        = e("litecoin address for wrong network")
	LitecoinAddressIsNotSupported         = e("litecoin address is not supported")
	MakeBlockTransferFailed               = e("make block transfer failed")
	MakeGrantFailed                       = e("make grant failed")
	MakeIssueFailed                       = e("make issue failed")
	MakeShareFailed                       = e("make share failed")
	MakeSwapFailed                        = e("make swap failed")
	MakeTransferFailed                    = e("make transfer failed")
	MerkleRootDoesNotMatch                = e("merkle root does not match")
	MetadataIsNotMap                      = e("metadata is not map")
	MetadataTooLong                       = e("metadata too long")
	MissingBlockOwner                     = e("missing block owner")
	MissingOwnerData                      = e("missing owner data")
	MissingParameters                     = e("missing parameters")
	MissingPaymentBitcoinSection          = e("missing payment bitcoin section")
	MissingPaymentDiscoverySection        = e("missing payment discovery section")
	MissingPaymentLitecoinSection         = e("missing payment litecoin section")
	MissingPreviousBlockHeader            = e("missing previous block header")
	NameTooLong                           = e("name too long")
	NoAddressToReturn                     = e("no address to return")
	NoConnectionsAvailable                = e("no connections available")
	NoNewBlockHeadersFromPeer             = e("no new block headers from peer")
	NoNewTransactions                     = e("no new transactions")
	NotACountersignableRecord             = e("not a countersignable record")
	NotAPayId                             = e("not a pay id")
	NotAPayNonce                          = e("not a pay nonce")
	NotAssetId                            = e("not asset id")
	NotAssetIdentifier                    = e("not asset identifier")
	NotAvailableDuringSynchronise         = e("not available during synchronise")
	NotConnected                          = e("not connected")
	NotInitialised                        = e("not initialised")
	NotLink                               = e("not link")
	NotOwnedItem                          = e("not owned item")
	NotOwnerDataPack                      = e("not owner data pack")
	NotPrivateKey                         = e("not private key")
	NotPublicKey                          = e("not public key")
	NotTransactionPack                    = e("not transaction pack")
	OwnershipIsNotIndexed                 = e("ownership is not indexed")
	PasswordMismatch                      = e("password mismatch")
	PayIdAlreadyUsed                      = e("pay id already used")
	PaymentAddressTooLong                 = e("payment address too long")
	PreviousBlockDigestDoesNotMatch       = e("previous block digest does not match")
	PreviousOwnershipWasNotDeleted        = e("previous ownership was not deleted")
	PreviousTransactionWasNotDeleted      = e("previous transaction was not deleted")
	ProcessStopping                       = e("process stopping")
	RateLimiting                          = e("rate limiting")
	RecordHasExpired                      = e("record has expired")
	ShareIdsCannotBeIdentical             = e("share ids cannot be identical")
	ShareQuantityTooSmall                 = e("share quantity too small")
	SignatureTooLong                      = e("signature too long")
	TimeoutWaitingForHeader               = e("timeout waiting for header")
	TooManyItemsToProcess                 = e("too many items to process")
	TransactionAlreadyExists              = e("transaction already exists")
	TransactionCountOutOfRange            = e("transaction count out of range")
	TransactionHexDataIsRequired          = e("transaction hex data is required")
	TransactionIdIsRequired               = e("transaction id is required")
	TransactionIsNotAnAsset               = e("transaction is not an asset")
	TransactionIsNotAnIssue               = e("transaction is not an issue")
	TransactionIsNotATransfer             = e("transaction is not a transfer")
	TransactionIsNotIndexed               = e("transaction is not indexed")
	TransactionLinksToSelf                = e("transaction links to self")
	UnexpectedTransactionRecord           = e("unexpected transaction record")
	UnmarshalTextFailed                   = e("unmarshal text failed")
	UnsupportedCurrency                   = e("unsupported currency")
	VotesInsufficient                     = e("votes insufficient")
	VotesWithEmptyWinner                  = e("votes with empty winner")
	VotesWithZeroCount                    = e("votes with zero count")
	VotesWithZeroHeight                   = e("votes with zero height")
	WrongNetworkForPublicKey              = e("wrong network for public key")
	WrongPassword                         = e("wrong password")
)

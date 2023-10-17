// @generated by protoc-gen-connect-es v1.1.2
// @generated from file rpc.proto (package land.gno.gnomobile.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddressFromBech32Request, AddressFromBech32Response, AddressToBech32Request, AddressToBech32Response, CallRequest, CallResponse, CreateAccountRequest, CreateAccountResponse, DeleteAccountRequest, DeleteAccountResponse, GenerateRecoveryPhraseRequest, GenerateRecoveryPhraseResponse, GetActiveAccountRequest, GetActiveAccountResponse, HelloRequest, HelloResponse, ListKeyInfoRequest, ListKeyInfoResponse, QueryAccountRequest, QueryAccountResponse, QueryRequest, QueryResponse, SelectAccountRequest, SelectAccountResponse, SetChainIDRequest, SetChainIDResponse, SetPasswordRequest, SetPasswordResponse, SetRemoteRequest, SetRemoteResponse } from "./gnomobiletypes_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * GnomobileService is the service to interact with the Gno blockchain
 *
 * @generated from service land.gno.gnomobile.v1.GnomobileService
 */
export declare const GnomobileService: {
  readonly typeName: "land.gno.gnomobile.v1.GnomobileService",
  readonly methods: {
    /**
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetRemote
     */
    readonly setRemote: {
      readonly name: "SetRemote",
      readonly I: typeof SetRemoteRequest,
      readonly O: typeof SetRemoteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetChainID
     */
    readonly setChainID: {
      readonly name: "SetChainID",
      readonly I: typeof SetChainIDRequest,
      readonly O: typeof SetChainIDResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the
     * crypto library random number generator. This can be used as the mnemonic in
     * CreateAccount.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase
     */
    readonly generateRecoveryPhrase: {
      readonly name: "GenerateRecoveryPhrase",
      readonly I: typeof GenerateRecoveryPhraseRequest,
      readonly O: typeof GenerateRecoveryPhraseResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Get the keys informations in the keybase
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.ListKeyInfo
     */
    readonly listKeyInfo: {
      readonly name: "ListKeyInfo",
      readonly I: typeof ListKeyInfoRequest,
      readonly O: typeof ListKeyInfoResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.CreateAccount
     */
    readonly createAccount: {
      readonly name: "CreateAccount",
      readonly I: typeof CreateAccountRequest,
      readonly O: typeof CreateAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * SelectAccount selects the active account to use for later operations
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SelectAccount
     */
    readonly selectAccount: {
      readonly name: "SelectAccount",
      readonly I: typeof SelectAccountRequest,
      readonly O: typeof SelectAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Set the password for the account in the keybase, used for later operations
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetPassword
     */
    readonly setPassword: {
      readonly name: "SetPassword",
      readonly I: typeof SetPasswordRequest,
      readonly O: typeof SetPasswordResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the
     * length of the result.)
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.GetActiveAccount
     */
    readonly getActiveAccount: {
      readonly name: "GetActiveAccount",
      readonly I: typeof GetActiveAccountRequest,
      readonly O: typeof GetActiveAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * QueryAccount retrieves account information from the blockchain for a given address.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.QueryAccount
     */
    readonly queryAccount: {
      readonly name: "QueryAccount",
      readonly I: typeof QueryAccountRequest,
      readonly O: typeof QueryAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * DeleteAccount deletes the account with the given name, using the password to
     * ensure access. However, if skip_password is true, then ignore the password.
     * If the account doesn't exist, then return ErrCryptoKeyNotFound.
     * If the password is wrong, then return ErrDecryptionFailed.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.DeleteAccount
     */
    readonly deleteAccount: {
      readonly name: "DeleteAccount",
      readonly I: typeof DeleteAccountRequest,
      readonly O: typeof DeleteAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Make an ABCI query to the remote node.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Query
     */
    readonly query: {
      readonly name: "Query",
      readonly I: typeof QueryRequest,
      readonly O: typeof QueryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Call a specific realm function.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Call
     */
    readonly call: {
      readonly name: "Call",
      readonly I: typeof CallRequest,
      readonly O: typeof CallResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Convert a byte array address to a bech32 string address.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.AddressToBech32
     */
    readonly addressToBech32: {
      readonly name: "AddressToBech32",
      readonly I: typeof AddressToBech32Request,
      readonly O: typeof AddressToBech32Response,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Convert a bech32 string address to a byte array address.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.AddressFromBech32
     */
    readonly addressFromBech32: {
      readonly name: "AddressFromBech32",
      readonly I: typeof AddressFromBech32Request,
      readonly O: typeof AddressFromBech32Response,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Hello is for debug purposes
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Hello
     */
    readonly hello: {
      readonly name: "Hello",
      readonly I: typeof HelloRequest,
      readonly O: typeof HelloResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};


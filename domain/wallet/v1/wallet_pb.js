// @generated by protoc-gen-es v1.1.1
// @generated from file wallet/v1/wallet.proto (package wallet.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Transaction } from "../../transaction/v1/transaction_pb.js";

/**
 * @generated from message wallet.v1.Wallet
 */
export const Wallet = proto3.makeMessageType(
  "wallet.v1.Wallet",
  () => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "network_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "language", kind: "enum", T: proto3.getEnumType(Wallet_Language) },
    { no: 4, name: "theme", kind: "enum", T: proto3.getEnumType(Wallet_Theme) },
    { no: 5, name: "transactions", kind: "message", T: Transaction, repeated: true },
    { no: 6, name: "collectibles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * An enumeration of the supported languages for the wallet.
 *
 * @generated from enum wallet.v1.Wallet.Language
 */
export const Wallet_Language = proto3.makeEnum(
  "wallet.v1.Wallet.Language",
  [
    {no: 0, name: "LANGUAGE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "LANGUAGE_EN", localName: "EN"},
    {no: 2, name: "LANGUAGE_DE", localName: "DE"},
    {no: 3, name: "LANGUAGE_ES", localName: "ES"},
  ],
);

/**
 * An enumeration of the supported themes for the wallet.
 *
 * @generated from enum wallet.v1.Wallet.Theme
 */
export const Wallet_Theme = proto3.makeEnum(
  "wallet.v1.Wallet.Theme",
  [
    {no: 0, name: "THEME_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "THEME_LIGHT", localName: "LIGHT"},
    {no: 2, name: "THEME_DARK", localName: "DARK"},
  ],
);


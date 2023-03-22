// @generated by protoc-gen-es v1.1.1
// @generated from file transaction/v1/transaction.proto (package transaction.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message transaction.v1.Transaction
 */
export declare class Transaction extends Message<Transaction> {
  /**
   * @generated from field: optional transaction.v1.Transaction.Metadata metadata = 1;
   */
  metadata?: Transaction_Metadata;

  /**
   * @generated from field: string contract = 2;
   */
  contract: string;

  /**
   * @generated from field: uint32 network = 3;
   */
  network: number;

  /**
   * @generated from oneof transaction.v1.Transaction.data
   */
  data: {
    /**
     * @generated from field: transaction.v1.Transaction.ERC1155.SafeTransferFrom safe_transfer_from = 4;
     */
    value: Transaction_ERC1155_SafeTransferFrom;
    case: "safeTransferFrom";
  } | {
    /**
     * @generated from field: transaction.v1.Transaction.ERC1155.SafeBatchTransferFrom safe_batch_transfer_from = 5;
     */
    value: Transaction_ERC1155_SafeBatchTransferFrom;
    case: "safeBatchTransferFrom";
  } | {
    /**
     * @generated from field: transaction.v1.Transaction.ERC1155.SetApprovalForAll safe_approval_for_all = 6;
     */
    value: Transaction_ERC1155_SetApprovalForAll;
    case: "safeApprovalForAll";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<Transaction>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction;

  static equals(a: Transaction | PlainMessage<Transaction> | undefined, b: Transaction | PlainMessage<Transaction> | undefined): boolean;
}

/**
 * @generated from message transaction.v1.Transaction.ERC1155
 */
export declare class Transaction_ERC1155 extends Message<Transaction_ERC1155> {
  constructor(data?: PartialMessage<Transaction_ERC1155>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction.ERC1155";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction_ERC1155;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction_ERC1155;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction_ERC1155;

  static equals(a: Transaction_ERC1155 | PlainMessage<Transaction_ERC1155> | undefined, b: Transaction_ERC1155 | PlainMessage<Transaction_ERC1155> | undefined): boolean;
}

/**
 * @generated from message transaction.v1.Transaction.ERC1155.SafeTransferFrom
 */
export declare class Transaction_ERC1155_SafeTransferFrom extends Message<Transaction_ERC1155_SafeTransferFrom> {
  /**
   * @generated from field: string from = 1;
   */
  from: string;

  /**
   * @generated from field: string to = 2;
   */
  to: string;

  /**
   * @generated from field: uint64 id = 3;
   */
  id: bigint;

  /**
   * @generated from field: uint64 value = 4;
   */
  value: bigint;

  /**
   * @generated from field: bytes data = 5;
   */
  data: Uint8Array;

  constructor(data?: PartialMessage<Transaction_ERC1155_SafeTransferFrom>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction.ERC1155.SafeTransferFrom";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction_ERC1155_SafeTransferFrom;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SafeTransferFrom;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SafeTransferFrom;

  static equals(a: Transaction_ERC1155_SafeTransferFrom | PlainMessage<Transaction_ERC1155_SafeTransferFrom> | undefined, b: Transaction_ERC1155_SafeTransferFrom | PlainMessage<Transaction_ERC1155_SafeTransferFrom> | undefined): boolean;
}

/**
 * @generated from message transaction.v1.Transaction.ERC1155.SafeBatchTransferFrom
 */
export declare class Transaction_ERC1155_SafeBatchTransferFrom extends Message<Transaction_ERC1155_SafeBatchTransferFrom> {
  /**
   * @generated from field: string from = 1;
   */
  from: string;

  /**
   * @generated from field: string to = 2;
   */
  to: string;

  /**
   * @generated from field: repeated uint64 ids = 3;
   */
  ids: bigint[];

  /**
   * @generated from field: repeated uint64 values = 4;
   */
  values: bigint[];

  /**
   * @generated from field: bytes data = 5;
   */
  data: Uint8Array;

  constructor(data?: PartialMessage<Transaction_ERC1155_SafeBatchTransferFrom>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction.ERC1155.SafeBatchTransferFrom";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction_ERC1155_SafeBatchTransferFrom;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SafeBatchTransferFrom;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SafeBatchTransferFrom;

  static equals(a: Transaction_ERC1155_SafeBatchTransferFrom | PlainMessage<Transaction_ERC1155_SafeBatchTransferFrom> | undefined, b: Transaction_ERC1155_SafeBatchTransferFrom | PlainMessage<Transaction_ERC1155_SafeBatchTransferFrom> | undefined): boolean;
}

/**
 * @generated from message transaction.v1.Transaction.ERC1155.SetApprovalForAll
 */
export declare class Transaction_ERC1155_SetApprovalForAll extends Message<Transaction_ERC1155_SetApprovalForAll> {
  /**
   * @generated from field: string operator = 1;
   */
  operator: string;

  /**
   * @generated from field: bool approved = 2;
   */
  approved: boolean;

  constructor(data?: PartialMessage<Transaction_ERC1155_SetApprovalForAll>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction.ERC1155.SetApprovalForAll";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction_ERC1155_SetApprovalForAll;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SetApprovalForAll;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction_ERC1155_SetApprovalForAll;

  static equals(a: Transaction_ERC1155_SetApprovalForAll | PlainMessage<Transaction_ERC1155_SetApprovalForAll> | undefined, b: Transaction_ERC1155_SetApprovalForAll | PlainMessage<Transaction_ERC1155_SetApprovalForAll> | undefined): boolean;
}

/**
 * @generated from message transaction.v1.Transaction.Metadata
 */
export declare class Transaction_Metadata extends Message<Transaction_Metadata> {
  /**
   * @generated from field: string hash = 1;
   */
  hash: string;

  /**
   * @generated from field: google.protobuf.Timestamp date = 2;
   */
  date?: Timestamp;

  /**
   * @generated from field: optional string reverted = 3;
   */
  reverted?: string;

  constructor(data?: PartialMessage<Transaction_Metadata>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "transaction.v1.Transaction.Metadata";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction_Metadata;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction_Metadata;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction_Metadata;

  static equals(a: Transaction_Metadata | PlainMessage<Transaction_Metadata> | undefined, b: Transaction_Metadata | PlainMessage<Transaction_Metadata> | undefined): boolean;
}


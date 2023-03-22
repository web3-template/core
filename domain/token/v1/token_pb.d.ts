// @generated by protoc-gen-es v1.1.1
// @generated from file token/v1/token.proto (package token.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message token.v1.Token
 */
export declare class Token extends Message<Token> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string image = 3;
   */
  image: string;

  /**
   * @generated from field: string external_url = 4;
   */
  externalUrl: string;

  /**
   * @generated from field: repeated token.v1.Token.Attribute attributes = 5;
   */
  attributes: Token_Attribute[];

  constructor(data?: PartialMessage<Token>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "token.v1.Token";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Token;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Token;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Token;

  static equals(a: Token | PlainMessage<Token> | undefined, b: Token | PlainMessage<Token> | undefined): boolean;
}

/**
 * @generated from message token.v1.Token.Attribute
 */
export declare class Token_Attribute extends Message<Token_Attribute> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string value = 2;
   */
  value: string;

  constructor(data?: PartialMessage<Token_Attribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "token.v1.Token.Attribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Token_Attribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Token_Attribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Token_Attribute;

  static equals(a: Token_Attribute | PlainMessage<Token_Attribute> | undefined, b: Token_Attribute | PlainMessage<Token_Attribute> | undefined): boolean;
}

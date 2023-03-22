// @generated by protoc-gen-es v1.1.1
// @generated from file cmd/v1/cmd.proto (package cmd.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message cmd.v1.Command
 */
export declare class Command extends Message<Command> {
  /**
   * @generated from oneof cmd.v1.Command.command
   */
  command: {
    /**
     * @generated from field: cmd.v1.Command.AddCollectible add_collectible = 1;
     */
    value: Command_AddCollectible;
    case: "addCollectible";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<Command>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cmd.v1.Command";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Command;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Command;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Command;

  static equals(a: Command | PlainMessage<Command> | undefined, b: Command | PlainMessage<Command> | undefined): boolean;
}

/**
 * @generated from message cmd.v1.Command.AddCollectible
 */
export declare class Command_AddCollectible extends Message<Command_AddCollectible> {
  /**
   * @generated from field: string address = 1;
   */
  address: string;

  constructor(data?: PartialMessage<Command_AddCollectible>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cmd.v1.Command.AddCollectible";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Command_AddCollectible;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Command_AddCollectible;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Command_AddCollectible;

  static equals(a: Command_AddCollectible | PlainMessage<Command_AddCollectible> | undefined, b: Command_AddCollectible | PlainMessage<Command_AddCollectible> | undefined): boolean;
}


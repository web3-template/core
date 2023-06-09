// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: wallet/v1/wallet.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

struct Wallet_V1_Wallet {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// The account associated with the wallet.
  var account: String = String()

  /// The network ID associated with the wallet.
  var networkID: UInt64 = 0

  /// The language of the wallet.
  var language: Wallet_V1_Wallet.Language = .unspecified

  /// The theme of the wallet.
  var theme: Wallet_V1_Wallet.Theme = .unspecified

  /// A list of transactions associated with the wallet.
  var transactions: [Transaction_V1_Transaction] = []

  /// A list of collectibles associated with the wallet.
  var collectibles: [String] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  /// An enumeration of the supported languages for the wallet.
  enum Language: SwiftProtobuf.Enum {
    typealias RawValue = Int

    /// The language is not specified.
    case unspecified // = 0

    /// English language.
    case en // = 1

    /// German language.
    case de // = 2

    /// Spanish language.
    case es // = 3
    case UNRECOGNIZED(Int)

    init() {
      self = .unspecified
    }

    init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .unspecified
      case 1: self = .en
      case 2: self = .de
      case 3: self = .es
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    var rawValue: Int {
      switch self {
      case .unspecified: return 0
      case .en: return 1
      case .de: return 2
      case .es: return 3
      case .UNRECOGNIZED(let i): return i
      }
    }

  }

  /// An enumeration of the supported themes for the wallet.
  enum Theme: SwiftProtobuf.Enum {
    typealias RawValue = Int

    /// The theme is not specified.
    case unspecified // = 0

    /// Light theme.
    case light // = 1

    /// Dark theme.
    case dark // = 2
    case UNRECOGNIZED(Int)

    init() {
      self = .unspecified
    }

    init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .unspecified
      case 1: self = .light
      case 2: self = .dark
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    var rawValue: Int {
      switch self {
      case .unspecified: return 0
      case .light: return 1
      case .dark: return 2
      case .UNRECOGNIZED(let i): return i
      }
    }

  }

  init() {}
}

#if swift(>=4.2)

extension Wallet_V1_Wallet.Language: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Wallet_V1_Wallet.Language] = [
    .unspecified,
    .en,
    .de,
    .es,
  ]
}

extension Wallet_V1_Wallet.Theme: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Wallet_V1_Wallet.Theme] = [
    .unspecified,
    .light,
    .dark,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Wallet_V1_Wallet: @unchecked Sendable {}
extension Wallet_V1_Wallet.Language: @unchecked Sendable {}
extension Wallet_V1_Wallet.Theme: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "wallet.v1"

extension Wallet_V1_Wallet: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Wallet"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "account"),
    2: .standard(proto: "network_id"),
    3: .same(proto: "language"),
    4: .same(proto: "theme"),
    5: .same(proto: "transactions"),
    6: .same(proto: "collectibles"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.account) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.networkID) }()
      case 3: try { try decoder.decodeSingularEnumField(value: &self.language) }()
      case 4: try { try decoder.decodeSingularEnumField(value: &self.theme) }()
      case 5: try { try decoder.decodeRepeatedMessageField(value: &self.transactions) }()
      case 6: try { try decoder.decodeRepeatedStringField(value: &self.collectibles) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.account.isEmpty {
      try visitor.visitSingularStringField(value: self.account, fieldNumber: 1)
    }
    if self.networkID != 0 {
      try visitor.visitSingularUInt64Field(value: self.networkID, fieldNumber: 2)
    }
    if self.language != .unspecified {
      try visitor.visitSingularEnumField(value: self.language, fieldNumber: 3)
    }
    if self.theme != .unspecified {
      try visitor.visitSingularEnumField(value: self.theme, fieldNumber: 4)
    }
    if !self.transactions.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.transactions, fieldNumber: 5)
    }
    if !self.collectibles.isEmpty {
      try visitor.visitRepeatedStringField(value: self.collectibles, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Wallet_V1_Wallet, rhs: Wallet_V1_Wallet) -> Bool {
    if lhs.account != rhs.account {return false}
    if lhs.networkID != rhs.networkID {return false}
    if lhs.language != rhs.language {return false}
    if lhs.theme != rhs.theme {return false}
    if lhs.transactions != rhs.transactions {return false}
    if lhs.collectibles != rhs.collectibles {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Wallet_V1_Wallet.Language: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "LANGUAGE_UNSPECIFIED"),
    1: .same(proto: "LANGUAGE_EN"),
    2: .same(proto: "LANGUAGE_DE"),
    3: .same(proto: "LANGUAGE_ES"),
  ]
}

extension Wallet_V1_Wallet.Theme: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "THEME_UNSPECIFIED"),
    1: .same(proto: "THEME_LIGHT"),
    2: .same(proto: "THEME_DARK"),
  ]
}

import { Buffer } from "buffer";

export function encode(data, base = `base64`) {
  return Buffer.from(data).toString(base);
}

export function decode(encoded, base = `base64`) {
  return Uint8Array.from(Buffer.from(encoded, base));
}

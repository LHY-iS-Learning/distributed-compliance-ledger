/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface RejectedCertificate {
  subject: string
  subjectKeyId: string
  pemCert: string
  serialNumber: string
  owner: string
  approvals: string[]
}

const baseRejectedCertificate: object = { subject: '', subjectKeyId: '', pemCert: '', serialNumber: '', owner: '', approvals: '' }

export const RejectedCertificate = {
  encode(message: RejectedCertificate, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    if (message.pemCert !== '') {
      writer.uint32(26).string(message.pemCert)
    }
    if (message.serialNumber !== '') {
      writer.uint32(34).string(message.serialNumber)
    }
    if (message.owner !== '') {
      writer.uint32(42).string(message.owner)
    }
    for (const v of message.approvals) {
      writer.uint32(50).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RejectedCertificate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRejectedCertificate } as RejectedCertificate
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        case 3:
          message.pemCert = reader.string()
          break
        case 4:
          message.serialNumber = reader.string()
          break
        case 5:
          message.owner = reader.string()
          break
        case 6:
          message.approvals.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RejectedCertificate {
    const message = { ...baseRejectedCertificate } as RejectedCertificate
    message.approvals = []
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = String(object.subject)
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = String(object.subjectKeyId)
    } else {
      message.subjectKeyId = ''
    }
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = String(object.pemCert)
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = String(object.serialNumber)
    } else {
      message.serialNumber = ''
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(String(e))
      }
    }
    return message
  },

  toJSON(message: RejectedCertificate): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.pemCert !== undefined && (obj.pemCert = message.pemCert)
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber)
    message.owner !== undefined && (obj.owner = message.owner)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e)
    } else {
      obj.approvals = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<RejectedCertificate>): RejectedCertificate {
    const message = { ...baseRejectedCertificate } as RejectedCertificate
    message.approvals = []
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = object.subject
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = object.subjectKeyId
    } else {
      message.subjectKeyId = ''
    }
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = object.pemCert
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber
    } else {
      message.serialNumber = ''
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(e)
      }
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

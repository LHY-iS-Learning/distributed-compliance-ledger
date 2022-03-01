/* eslint-disable */
import { Plan } from '../cosmos/upgrade/v1beta1/upgrade';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseApprovedUpgrade = { creator: '', approvals: '' };
export const ApprovedUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.plan !== undefined) {
            Plan.encode(message.plan, writer.uint32(10).fork()).ldelim();
        }
        if (message.creator !== '') {
            writer.uint32(18).string(message.creator);
        }
        for (const v of message.approvals) {
            writer.uint32(26).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.plan = Plan.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                case 3:
                    message.approvals.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = Plan.fromJSON(object.plan);
        }
        else {
            message.plan = undefined;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.plan !== undefined && (obj.plan = message.plan ? Plan.toJSON(message.plan) : undefined);
        message.creator !== undefined && (obj.creator = message.creator);
        if (message.approvals) {
            obj.approvals = message.approvals.map((e) => e);
        }
        else {
            obj.approvals = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = Plan.fromPartial(object.plan);
        }
        else {
            message.plan = undefined;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(e);
            }
        }
        return message;
    }
};

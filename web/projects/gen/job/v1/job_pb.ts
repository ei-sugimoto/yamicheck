// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file job/v1/job.proto (package job.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message job.v1.CheckRequest
 */
export class CheckRequest extends Message<CheckRequest> {
  /**
   * @generated from field: string company_name = 1;
   */
  companyName = "";

  /**
   * @generated from field: int64 hourly_wage = 2;
   */
  hourlyWage = protoInt64.zero;

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  constructor(data?: PartialMessage<CheckRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "job.v1.CheckRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "company_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "hourly_wage", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckRequest {
    return new CheckRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckRequest {
    return new CheckRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckRequest {
    return new CheckRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CheckRequest | PlainMessage<CheckRequest> | undefined, b: CheckRequest | PlainMessage<CheckRequest> | undefined): boolean {
    return proto3.util.equals(CheckRequest, a, b);
  }
}

/**
 * @generated from message job.v1.CheckResponse
 */
export class CheckResponse extends Message<CheckResponse> {
  /**
   * @generated from field: int64 level = 1;
   */
  level = protoInt64.zero;

  /**
   * @generated from field: string message = 2;
   */
  message = "";

  constructor(data?: PartialMessage<CheckResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "job.v1.CheckResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "level", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckResponse {
    return new CheckResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckResponse {
    return new CheckResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckResponse {
    return new CheckResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CheckResponse | PlainMessage<CheckResponse> | undefined, b: CheckResponse | PlainMessage<CheckResponse> | undefined): boolean {
    return proto3.util.equals(CheckResponse, a, b);
  }
}


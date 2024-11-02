// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts"
// @generated from file job/v1/job.proto (package job.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CheckRequest, CheckResponse } from "./job_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service job.v1.JobService
 */
export const JobService = {
  typeName: "job.v1.JobService",
  methods: {
    /**
     * @generated from rpc job.v1.JobService.Check
     */
    check: {
      name: "Check",
      I: CheckRequest,
      O: CheckResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.1
 * source: driver/driver.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../common/ack";
import * as dependency_2 from "./../common/query";
import * as grpc_1 from "@grpc/grpc-js";
export namespace driver.driver {
    interface GrpcUnaryServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    }
    interface GrpcStreamServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
        (message: P, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
    }
    interface GrpWritableServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    }
    interface GrpcChunkServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
        (options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
    }
    interface GrpcPromiseServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): Promise<R>;
        (message: P, options?: grpc_1.CallOptions): Promise<R>;
    }
    export abstract class UnimplementedDriverCommunicationService {
        static definition = {
            RequestDriverState: {
                path: "/driver.driver.DriverCommunication/RequestDriverState",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: dependency_2.common.query.Query) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => dependency_2.common.query.Query.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: dependency_1.common.ack.Ack) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => dependency_1.common.ack.Ack.deserialize(new Uint8Array(bytes))
            }
        };
        [method: string]: grpc_1.UntypedHandleCall;
        abstract RequestDriverState(call: grpc_1.ServerUnaryCall<dependency_2.common.query.Query, dependency_1.common.ack.Ack>, callback: grpc_1.sendUnaryData<dependency_1.common.ack.Ack>): void;
    }
    export class DriverCommunicationClient extends grpc_1.makeGenericClientConstructor(UnimplementedDriverCommunicationService.definition, "DriverCommunication", {}) {
        constructor(address: string, credentials: grpc_1.ChannelCredentials, options?: Partial<grpc_1.ChannelOptions>) {
            super(address, credentials, options);
        }
        RequestDriverState: GrpcUnaryServiceInterface<dependency_2.common.query.Query, dependency_1.common.ack.Ack> = (message: dependency_2.common.query.Query, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<dependency_1.common.ack.Ack>, options?: grpc_1.CallOptions | grpc_1.requestCallback<dependency_1.common.ack.Ack>, callback?: grpc_1.requestCallback<dependency_1.common.ack.Ack>): grpc_1.ClientUnaryCall => {
            return super.RequestDriverState(message, metadata, options, callback);
        };
    }
}

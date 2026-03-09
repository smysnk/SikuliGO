// package: sikuli.v1
// file: sikuli/v1/sikuli.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class GrayImage extends jspb.Message { 
    getName(): string;
    setName(value: string): GrayImage;
    getWidth(): number;
    setWidth(value: number): GrayImage;
    getHeight(): number;
    setHeight(value: number): GrayImage;
    getPix(): Uint8Array | string;
    getPix_asU8(): Uint8Array;
    getPix_asB64(): string;
    setPix(value: Uint8Array | string): GrayImage;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GrayImage.AsObject;
    static toObject(includeInstance: boolean, msg: GrayImage): GrayImage.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GrayImage, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GrayImage;
    static deserializeBinaryFromReader(message: GrayImage, reader: jspb.BinaryReader): GrayImage;
}

export namespace GrayImage {
    export type AsObject = {
        name: string,
        width: number,
        height: number,
        pix: Uint8Array | string,
    }
}

export class Point extends jspb.Message { 
    getX(): number;
    setX(value: number): Point;
    getY(): number;
    setY(value: number): Point;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Point.AsObject;
    static toObject(includeInstance: boolean, msg: Point): Point.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Point, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Point;
    static deserializeBinaryFromReader(message: Point, reader: jspb.BinaryReader): Point;
}

export namespace Point {
    export type AsObject = {
        x: number,
        y: number,
    }
}

export class Rect extends jspb.Message { 
    getX(): number;
    setX(value: number): Rect;
    getY(): number;
    setY(value: number): Rect;
    getW(): number;
    setW(value: number): Rect;
    getH(): number;
    setH(value: number): Rect;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Rect.AsObject;
    static toObject(includeInstance: boolean, msg: Rect): Rect.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Rect, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Rect;
    static deserializeBinaryFromReader(message: Rect, reader: jspb.BinaryReader): Rect;
}

export namespace Rect {
    export type AsObject = {
        x: number,
        y: number,
        w: number,
        h: number,
    }
}

export class Pattern extends jspb.Message { 

    hasImage(): boolean;
    clearImage(): void;
    getImage(): GrayImage | undefined;
    setImage(value?: GrayImage): Pattern;

    hasSimilarity(): boolean;
    clearSimilarity(): void;
    getSimilarity(): number | undefined;
    setSimilarity(value: number): Pattern;

    hasExact(): boolean;
    clearExact(): void;
    getExact(): boolean | undefined;
    setExact(value: boolean): Pattern;

    hasTargetOffset(): boolean;
    clearTargetOffset(): void;
    getTargetOffset(): Point | undefined;
    setTargetOffset(value?: Point): Pattern;

    hasResizeFactor(): boolean;
    clearResizeFactor(): void;
    getResizeFactor(): number | undefined;
    setResizeFactor(value: number): Pattern;

    hasMask(): boolean;
    clearMask(): void;
    getMask(): GrayImage | undefined;
    setMask(value?: GrayImage): Pattern;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Pattern.AsObject;
    static toObject(includeInstance: boolean, msg: Pattern): Pattern.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Pattern, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Pattern;
    static deserializeBinaryFromReader(message: Pattern, reader: jspb.BinaryReader): Pattern;
}

export namespace Pattern {
    export type AsObject = {
        image?: GrayImage.AsObject,
        similarity?: number,
        exact?: boolean,
        targetOffset?: Point.AsObject,
        resizeFactor?: number,
        mask?: GrayImage.AsObject,
    }
}

export class Match extends jspb.Message { 

    hasRect(): boolean;
    clearRect(): void;
    getRect(): Rect | undefined;
    setRect(value?: Rect): Match;
    getScore(): number;
    setScore(value: number): Match;

    hasTarget(): boolean;
    clearTarget(): void;
    getTarget(): Point | undefined;
    setTarget(value?: Point): Match;
    getIndex(): number;
    setIndex(value: number): Match;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Match.AsObject;
    static toObject(includeInstance: boolean, msg: Match): Match.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Match, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Match;
    static deserializeBinaryFromReader(message: Match, reader: jspb.BinaryReader): Match;
}

export namespace Match {
    export type AsObject = {
        rect?: Rect.AsObject,
        score: number,
        target?: Point.AsObject,
        index: number,
    }
}

export class TextMatch extends jspb.Message { 

    hasRect(): boolean;
    clearRect(): void;
    getRect(): Rect | undefined;
    setRect(value?: Rect): TextMatch;
    getText(): string;
    setText(value: string): TextMatch;
    getConfidence(): number;
    setConfidence(value: number): TextMatch;
    getIndex(): number;
    setIndex(value: number): TextMatch;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TextMatch.AsObject;
    static toObject(includeInstance: boolean, msg: TextMatch): TextMatch.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TextMatch, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TextMatch;
    static deserializeBinaryFromReader(message: TextMatch, reader: jspb.BinaryReader): TextMatch;
}

export namespace TextMatch {
    export type AsObject = {
        rect?: Rect.AsObject,
        text: string,
        confidence: number,
        index: number,
    }
}

export class OCRParams extends jspb.Message { 
    getLanguage(): string;
    setLanguage(value: string): OCRParams;
    getTrainingDataPath(): string;
    setTrainingDataPath(value: string): OCRParams;

    hasMinConfidence(): boolean;
    clearMinConfidence(): void;
    getMinConfidence(): number | undefined;
    setMinConfidence(value: number): OCRParams;

    hasTimeoutMillis(): boolean;
    clearTimeoutMillis(): void;
    getTimeoutMillis(): number | undefined;
    setTimeoutMillis(value: number): OCRParams;

    hasCaseSensitive(): boolean;
    clearCaseSensitive(): void;
    getCaseSensitive(): boolean | undefined;
    setCaseSensitive(value: boolean): OCRParams;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OCRParams.AsObject;
    static toObject(includeInstance: boolean, msg: OCRParams): OCRParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OCRParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OCRParams;
    static deserializeBinaryFromReader(message: OCRParams, reader: jspb.BinaryReader): OCRParams;
}

export namespace OCRParams {
    export type AsObject = {
        language: string,
        trainingDataPath: string,
        minConfidence?: number,
        timeoutMillis?: number,
        caseSensitive?: boolean,
    }
}

export class InputOptions extends jspb.Message { 

    hasDelayMillis(): boolean;
    clearDelayMillis(): void;
    getDelayMillis(): number | undefined;
    setDelayMillis(value: number): InputOptions;
    getButton(): string;
    setButton(value: string): InputOptions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InputOptions.AsObject;
    static toObject(includeInstance: boolean, msg: InputOptions): InputOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InputOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InputOptions;
    static deserializeBinaryFromReader(message: InputOptions, reader: jspb.BinaryReader): InputOptions;
}

export namespace InputOptions {
    export type AsObject = {
        delayMillis?: number,
        button: string,
    }
}

export class ObserveOptions extends jspb.Message { 

    hasIntervalMillis(): boolean;
    clearIntervalMillis(): void;
    getIntervalMillis(): number | undefined;
    setIntervalMillis(value: number): ObserveOptions;

    hasTimeoutMillis(): boolean;
    clearTimeoutMillis(): void;
    getTimeoutMillis(): number | undefined;
    setTimeoutMillis(value: number): ObserveOptions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ObserveOptions.AsObject;
    static toObject(includeInstance: boolean, msg: ObserveOptions): ObserveOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ObserveOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ObserveOptions;
    static deserializeBinaryFromReader(message: ObserveOptions, reader: jspb.BinaryReader): ObserveOptions;
}

export namespace ObserveOptions {
    export type AsObject = {
        intervalMillis?: number,
        timeoutMillis?: number,
    }
}

export class AppOptions extends jspb.Message { 

    hasTimeoutMillis(): boolean;
    clearTimeoutMillis(): void;
    getTimeoutMillis(): number | undefined;
    setTimeoutMillis(value: number): AppOptions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AppOptions.AsObject;
    static toObject(includeInstance: boolean, msg: AppOptions): AppOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AppOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AppOptions;
    static deserializeBinaryFromReader(message: AppOptions, reader: jspb.BinaryReader): AppOptions;
}

export namespace AppOptions {
    export type AsObject = {
        timeoutMillis?: number,
    }
}

export class WindowQuery extends jspb.Message { 
    getId(): string;
    setId(value: string): WindowQuery;
    getTitleExact(): string;
    setTitleExact(value: string): WindowQuery;
    getTitleContains(): string;
    setTitleContains(value: string): WindowQuery;
    getFocusedOnly(): boolean;
    setFocusedOnly(value: boolean): WindowQuery;

    hasIndex(): boolean;
    clearIndex(): void;
    getIndex(): number | undefined;
    setIndex(value: number): WindowQuery;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WindowQuery.AsObject;
    static toObject(includeInstance: boolean, msg: WindowQuery): WindowQuery.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WindowQuery, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WindowQuery;
    static deserializeBinaryFromReader(message: WindowQuery, reader: jspb.BinaryReader): WindowQuery;
}

export namespace WindowQuery {
    export type AsObject = {
        id: string,
        titleExact: string,
        titleContains: string,
        focusedOnly: boolean,
        index?: number,
    }
}

export class Window extends jspb.Message { 
    getId(): string;
    setId(value: string): Window;
    getApp(): string;
    setApp(value: string): Window;
    getPid(): number;
    setPid(value: number): Window;
    getTitle(): string;
    setTitle(value: string): Window;

    hasBounds(): boolean;
    clearBounds(): void;
    getBounds(): Rect | undefined;
    setBounds(value?: Rect): Window;
    getFocused(): boolean;
    setFocused(value: boolean): Window;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Window.AsObject;
    static toObject(includeInstance: boolean, msg: Window): Window.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Window, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Window;
    static deserializeBinaryFromReader(message: Window, reader: jspb.BinaryReader): Window;
}

export namespace Window {
    export type AsObject = {
        id: string,
        app: string,
        pid: number,
        title: string,
        bounds?: Rect.AsObject,
        focused: boolean,
    }
}

export class ScreenDescriptor extends jspb.Message { 
    getId(): number;
    setId(value: number): ScreenDescriptor;
    getName(): string;
    setName(value: string): ScreenDescriptor;

    hasBounds(): boolean;
    clearBounds(): void;
    getBounds(): Rect | undefined;
    setBounds(value?: Rect): ScreenDescriptor;
    getPrimary(): boolean;
    setPrimary(value: boolean): ScreenDescriptor;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ScreenDescriptor.AsObject;
    static toObject(includeInstance: boolean, msg: ScreenDescriptor): ScreenDescriptor.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ScreenDescriptor, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ScreenDescriptor;
    static deserializeBinaryFromReader(message: ScreenDescriptor, reader: jspb.BinaryReader): ScreenDescriptor;
}

export namespace ScreenDescriptor {
    export type AsObject = {
        id: number,
        name: string,
        bounds?: Rect.AsObject,
        primary: boolean,
    }
}

export class ObserveEvent extends jspb.Message { 
    getType(): string;
    setType(value: string): ObserveEvent;

    hasMatch(): boolean;
    clearMatch(): void;
    getMatch(): Match | undefined;
    setMatch(value?: Match): ObserveEvent;
    getTimestampUnixMillis(): number;
    setTimestampUnixMillis(value: number): ObserveEvent;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ObserveEvent.AsObject;
    static toObject(includeInstance: boolean, msg: ObserveEvent): ObserveEvent.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ObserveEvent, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ObserveEvent;
    static deserializeBinaryFromReader(message: ObserveEvent, reader: jspb.BinaryReader): ObserveEvent;
}

export namespace ObserveEvent {
    export type AsObject = {
        type: string,
        match?: Match.AsObject,
        timestampUnixMillis: number,
    }
}

export class FindRequest extends jspb.Message { 

    hasSource(): boolean;
    clearSource(): void;
    getSource(): GrayImage | undefined;
    setSource(value?: GrayImage): FindRequest;

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): FindRequest;
    getMatcherEngine(): MatcherEngine;
    setMatcherEngine(value: MatcherEngine): FindRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FindRequest): FindRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindRequest;
    static deserializeBinaryFromReader(message: FindRequest, reader: jspb.BinaryReader): FindRequest;
}

export namespace FindRequest {
    export type AsObject = {
        source?: GrayImage.AsObject,
        pattern?: Pattern.AsObject,
        matcherEngine: MatcherEngine,
    }
}

export class FindResponse extends jspb.Message { 

    hasMatch(): boolean;
    clearMatch(): void;
    getMatch(): Match | undefined;
    setMatch(value?: Match): FindResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FindResponse): FindResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindResponse;
    static deserializeBinaryFromReader(message: FindResponse, reader: jspb.BinaryReader): FindResponse;
}

export namespace FindResponse {
    export type AsObject = {
        match?: Match.AsObject,
    }
}

export class FindAllResponse extends jspb.Message { 
    clearMatchesList(): void;
    getMatchesList(): Array<Match>;
    setMatchesList(value: Array<Match>): FindAllResponse;
    addMatches(value?: Match, index?: number): Match;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindAllResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FindAllResponse): FindAllResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindAllResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindAllResponse;
    static deserializeBinaryFromReader(message: FindAllResponse, reader: jspb.BinaryReader): FindAllResponse;
}

export namespace FindAllResponse {
    export type AsObject = {
        matchesList: Array<Match.AsObject>,
    }
}

export class ScreenQueryOptions extends jspb.Message { 

    hasRegion(): boolean;
    clearRegion(): void;
    getRegion(): Rect | undefined;
    setRegion(value?: Rect): ScreenQueryOptions;

    hasTimeoutMillis(): boolean;
    clearTimeoutMillis(): void;
    getTimeoutMillis(): number | undefined;
    setTimeoutMillis(value: number): ScreenQueryOptions;

    hasIntervalMillis(): boolean;
    clearIntervalMillis(): void;
    getIntervalMillis(): number | undefined;
    setIntervalMillis(value: number): ScreenQueryOptions;
    getMatcherEngine(): MatcherEngine;
    setMatcherEngine(value: MatcherEngine): ScreenQueryOptions;

    hasScreenId(): boolean;
    clearScreenId(): void;
    getScreenId(): number | undefined;
    setScreenId(value: number): ScreenQueryOptions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ScreenQueryOptions.AsObject;
    static toObject(includeInstance: boolean, msg: ScreenQueryOptions): ScreenQueryOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ScreenQueryOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ScreenQueryOptions;
    static deserializeBinaryFromReader(message: ScreenQueryOptions, reader: jspb.BinaryReader): ScreenQueryOptions;
}

export namespace ScreenQueryOptions {
    export type AsObject = {
        region?: Rect.AsObject,
        timeoutMillis?: number,
        intervalMillis?: number,
        matcherEngine: MatcherEngine,
        screenId?: number,
    }
}

export class ListScreensRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListScreensRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListScreensRequest): ListScreensRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListScreensRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListScreensRequest;
    static deserializeBinaryFromReader(message: ListScreensRequest, reader: jspb.BinaryReader): ListScreensRequest;
}

export namespace ListScreensRequest {
    export type AsObject = {
    }
}

export class ListScreensResponse extends jspb.Message { 
    clearScreensList(): void;
    getScreensList(): Array<ScreenDescriptor>;
    setScreensList(value: Array<ScreenDescriptor>): ListScreensResponse;
    addScreens(value?: ScreenDescriptor, index?: number): ScreenDescriptor;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListScreensResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListScreensResponse): ListScreensResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListScreensResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListScreensResponse;
    static deserializeBinaryFromReader(message: ListScreensResponse, reader: jspb.BinaryReader): ListScreensResponse;
}

export namespace ListScreensResponse {
    export type AsObject = {
        screensList: Array<ScreenDescriptor.AsObject>,
    }
}

export class GetPrimaryScreenRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetPrimaryScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetPrimaryScreenRequest): GetPrimaryScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetPrimaryScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetPrimaryScreenRequest;
    static deserializeBinaryFromReader(message: GetPrimaryScreenRequest, reader: jspb.BinaryReader): GetPrimaryScreenRequest;
}

export namespace GetPrimaryScreenRequest {
    export type AsObject = {
    }
}

export class GetPrimaryScreenResponse extends jspb.Message { 

    hasScreen(): boolean;
    clearScreen(): void;
    getScreen(): ScreenDescriptor | undefined;
    setScreen(value?: ScreenDescriptor): GetPrimaryScreenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetPrimaryScreenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetPrimaryScreenResponse): GetPrimaryScreenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetPrimaryScreenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetPrimaryScreenResponse;
    static deserializeBinaryFromReader(message: GetPrimaryScreenResponse, reader: jspb.BinaryReader): GetPrimaryScreenResponse;
}

export namespace GetPrimaryScreenResponse {
    export type AsObject = {
        screen?: ScreenDescriptor.AsObject,
    }
}

export class CaptureScreenRequest extends jspb.Message { 

    hasScreenId(): boolean;
    clearScreenId(): void;
    getScreenId(): number | undefined;
    setScreenId(value: number): CaptureScreenRequest;

    hasRegion(): boolean;
    clearRegion(): void;
    getRegion(): Rect | undefined;
    setRegion(value?: Rect): CaptureScreenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CaptureScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CaptureScreenRequest): CaptureScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CaptureScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CaptureScreenRequest;
    static deserializeBinaryFromReader(message: CaptureScreenRequest, reader: jspb.BinaryReader): CaptureScreenRequest;
}

export namespace CaptureScreenRequest {
    export type AsObject = {
        screenId?: number,
        region?: Rect.AsObject,
    }
}

export class CaptureScreenResponse extends jspb.Message { 

    hasImage(): boolean;
    clearImage(): void;
    getImage(): GrayImage | undefined;
    setImage(value?: GrayImage): CaptureScreenResponse;

    hasScreen(): boolean;
    clearScreen(): void;
    getScreen(): ScreenDescriptor | undefined;
    setScreen(value?: ScreenDescriptor): CaptureScreenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CaptureScreenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CaptureScreenResponse): CaptureScreenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CaptureScreenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CaptureScreenResponse;
    static deserializeBinaryFromReader(message: CaptureScreenResponse, reader: jspb.BinaryReader): CaptureScreenResponse;
}

export namespace CaptureScreenResponse {
    export type AsObject = {
        image?: GrayImage.AsObject,
        screen?: ScreenDescriptor.AsObject,
    }
}

export class FindOnScreenRequest extends jspb.Message { 

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): FindOnScreenRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ScreenQueryOptions | undefined;
    setOpts(value?: ScreenQueryOptions): FindOnScreenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindOnScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FindOnScreenRequest): FindOnScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindOnScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindOnScreenRequest;
    static deserializeBinaryFromReader(message: FindOnScreenRequest, reader: jspb.BinaryReader): FindOnScreenRequest;
}

export namespace FindOnScreenRequest {
    export type AsObject = {
        pattern?: Pattern.AsObject,
        opts?: ScreenQueryOptions.AsObject,
    }
}

export class ExistsOnScreenRequest extends jspb.Message { 

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): ExistsOnScreenRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ScreenQueryOptions | undefined;
    setOpts(value?: ScreenQueryOptions): ExistsOnScreenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExistsOnScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ExistsOnScreenRequest): ExistsOnScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExistsOnScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExistsOnScreenRequest;
    static deserializeBinaryFromReader(message: ExistsOnScreenRequest, reader: jspb.BinaryReader): ExistsOnScreenRequest;
}

export namespace ExistsOnScreenRequest {
    export type AsObject = {
        pattern?: Pattern.AsObject,
        opts?: ScreenQueryOptions.AsObject,
    }
}

export class ExistsOnScreenResponse extends jspb.Message { 
    getExists(): boolean;
    setExists(value: boolean): ExistsOnScreenResponse;

    hasMatch(): boolean;
    clearMatch(): void;
    getMatch(): Match | undefined;
    setMatch(value?: Match): ExistsOnScreenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExistsOnScreenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ExistsOnScreenResponse): ExistsOnScreenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExistsOnScreenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExistsOnScreenResponse;
    static deserializeBinaryFromReader(message: ExistsOnScreenResponse, reader: jspb.BinaryReader): ExistsOnScreenResponse;
}

export namespace ExistsOnScreenResponse {
    export type AsObject = {
        exists: boolean,
        match?: Match.AsObject,
    }
}

export class WaitOnScreenRequest extends jspb.Message { 

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): WaitOnScreenRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ScreenQueryOptions | undefined;
    setOpts(value?: ScreenQueryOptions): WaitOnScreenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WaitOnScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: WaitOnScreenRequest): WaitOnScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WaitOnScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WaitOnScreenRequest;
    static deserializeBinaryFromReader(message: WaitOnScreenRequest, reader: jspb.BinaryReader): WaitOnScreenRequest;
}

export namespace WaitOnScreenRequest {
    export type AsObject = {
        pattern?: Pattern.AsObject,
        opts?: ScreenQueryOptions.AsObject,
    }
}

export class ClickOnScreenRequest extends jspb.Message { 

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): ClickOnScreenRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ScreenQueryOptions | undefined;
    setOpts(value?: ScreenQueryOptions): ClickOnScreenRequest;

    hasClickOpts(): boolean;
    clearClickOpts(): void;
    getClickOpts(): InputOptions | undefined;
    setClickOpts(value?: InputOptions): ClickOnScreenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClickOnScreenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ClickOnScreenRequest): ClickOnScreenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClickOnScreenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClickOnScreenRequest;
    static deserializeBinaryFromReader(message: ClickOnScreenRequest, reader: jspb.BinaryReader): ClickOnScreenRequest;
}

export namespace ClickOnScreenRequest {
    export type AsObject = {
        pattern?: Pattern.AsObject,
        opts?: ScreenQueryOptions.AsObject,
        clickOpts?: InputOptions.AsObject,
    }
}

export class ReadTextRequest extends jspb.Message { 

    hasSource(): boolean;
    clearSource(): void;
    getSource(): GrayImage | undefined;
    setSource(value?: GrayImage): ReadTextRequest;

    hasParams(): boolean;
    clearParams(): void;
    getParams(): OCRParams | undefined;
    setParams(value?: OCRParams): ReadTextRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadTextRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ReadTextRequest): ReadTextRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadTextRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadTextRequest;
    static deserializeBinaryFromReader(message: ReadTextRequest, reader: jspb.BinaryReader): ReadTextRequest;
}

export namespace ReadTextRequest {
    export type AsObject = {
        source?: GrayImage.AsObject,
        params?: OCRParams.AsObject,
    }
}

export class ReadTextResponse extends jspb.Message { 
    getText(): string;
    setText(value: string): ReadTextResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadTextResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ReadTextResponse): ReadTextResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadTextResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadTextResponse;
    static deserializeBinaryFromReader(message: ReadTextResponse, reader: jspb.BinaryReader): ReadTextResponse;
}

export namespace ReadTextResponse {
    export type AsObject = {
        text: string,
    }
}

export class FindTextRequest extends jspb.Message { 

    hasSource(): boolean;
    clearSource(): void;
    getSource(): GrayImage | undefined;
    setSource(value?: GrayImage): FindTextRequest;
    getQuery(): string;
    setQuery(value: string): FindTextRequest;

    hasParams(): boolean;
    clearParams(): void;
    getParams(): OCRParams | undefined;
    setParams(value?: OCRParams): FindTextRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindTextRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FindTextRequest): FindTextRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindTextRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindTextRequest;
    static deserializeBinaryFromReader(message: FindTextRequest, reader: jspb.BinaryReader): FindTextRequest;
}

export namespace FindTextRequest {
    export type AsObject = {
        source?: GrayImage.AsObject,
        query: string,
        params?: OCRParams.AsObject,
    }
}

export class FindTextResponse extends jspb.Message { 
    clearMatchesList(): void;
    getMatchesList(): Array<TextMatch>;
    setMatchesList(value: Array<TextMatch>): FindTextResponse;
    addMatches(value?: TextMatch, index?: number): TextMatch;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindTextResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FindTextResponse): FindTextResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindTextResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindTextResponse;
    static deserializeBinaryFromReader(message: FindTextResponse, reader: jspb.BinaryReader): FindTextResponse;
}

export namespace FindTextResponse {
    export type AsObject = {
        matchesList: Array<TextMatch.AsObject>,
    }
}

export class MoveMouseRequest extends jspb.Message { 
    getX(): number;
    setX(value: number): MoveMouseRequest;
    getY(): number;
    setY(value: number): MoveMouseRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): InputOptions | undefined;
    setOpts(value?: InputOptions): MoveMouseRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MoveMouseRequest.AsObject;
    static toObject(includeInstance: boolean, msg: MoveMouseRequest): MoveMouseRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MoveMouseRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MoveMouseRequest;
    static deserializeBinaryFromReader(message: MoveMouseRequest, reader: jspb.BinaryReader): MoveMouseRequest;
}

export namespace MoveMouseRequest {
    export type AsObject = {
        x: number,
        y: number,
        opts?: InputOptions.AsObject,
    }
}

export class ClickRequest extends jspb.Message { 
    getX(): number;
    setX(value: number): ClickRequest;
    getY(): number;
    setY(value: number): ClickRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): InputOptions | undefined;
    setOpts(value?: InputOptions): ClickRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClickRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ClickRequest): ClickRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClickRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClickRequest;
    static deserializeBinaryFromReader(message: ClickRequest, reader: jspb.BinaryReader): ClickRequest;
}

export namespace ClickRequest {
    export type AsObject = {
        x: number,
        y: number,
        opts?: InputOptions.AsObject,
    }
}

export class TypeTextRequest extends jspb.Message { 
    getText(): string;
    setText(value: string): TypeTextRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): InputOptions | undefined;
    setOpts(value?: InputOptions): TypeTextRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TypeTextRequest.AsObject;
    static toObject(includeInstance: boolean, msg: TypeTextRequest): TypeTextRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TypeTextRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TypeTextRequest;
    static deserializeBinaryFromReader(message: TypeTextRequest, reader: jspb.BinaryReader): TypeTextRequest;
}

export namespace TypeTextRequest {
    export type AsObject = {
        text: string,
        opts?: InputOptions.AsObject,
    }
}

export class HotkeyRequest extends jspb.Message { 
    clearKeysList(): void;
    getKeysList(): Array<string>;
    setKeysList(value: Array<string>): HotkeyRequest;
    addKeys(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HotkeyRequest.AsObject;
    static toObject(includeInstance: boolean, msg: HotkeyRequest): HotkeyRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: HotkeyRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HotkeyRequest;
    static deserializeBinaryFromReader(message: HotkeyRequest, reader: jspb.BinaryReader): HotkeyRequest;
}

export namespace HotkeyRequest {
    export type AsObject = {
        keysList: Array<string>,
    }
}

export class ScrollWheelRequest extends jspb.Message { 
    getX(): number;
    setX(value: number): ScrollWheelRequest;
    getY(): number;
    setY(value: number): ScrollWheelRequest;
    getDirection(): string;
    setDirection(value: string): ScrollWheelRequest;
    getSteps(): number;
    setSteps(value: number): ScrollWheelRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): InputOptions | undefined;
    setOpts(value?: InputOptions): ScrollWheelRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ScrollWheelRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ScrollWheelRequest): ScrollWheelRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ScrollWheelRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ScrollWheelRequest;
    static deserializeBinaryFromReader(message: ScrollWheelRequest, reader: jspb.BinaryReader): ScrollWheelRequest;
}

export namespace ScrollWheelRequest {
    export type AsObject = {
        x: number,
        y: number,
        direction: string,
        steps: number,
        opts?: InputOptions.AsObject,
    }
}

export class ActionResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ActionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ActionResponse): ActionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ActionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ActionResponse;
    static deserializeBinaryFromReader(message: ActionResponse, reader: jspb.BinaryReader): ActionResponse;
}

export namespace ActionResponse {
    export type AsObject = {
    }
}

export class ObserveRequest extends jspb.Message { 

    hasSource(): boolean;
    clearSource(): void;
    getSource(): GrayImage | undefined;
    setSource(value?: GrayImage): ObserveRequest;

    hasRegion(): boolean;
    clearRegion(): void;
    getRegion(): Rect | undefined;
    setRegion(value?: Rect): ObserveRequest;

    hasPattern(): boolean;
    clearPattern(): void;
    getPattern(): Pattern | undefined;
    setPattern(value?: Pattern): ObserveRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ObserveOptions | undefined;
    setOpts(value?: ObserveOptions): ObserveRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ObserveRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ObserveRequest): ObserveRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ObserveRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ObserveRequest;
    static deserializeBinaryFromReader(message: ObserveRequest, reader: jspb.BinaryReader): ObserveRequest;
}

export namespace ObserveRequest {
    export type AsObject = {
        source?: GrayImage.AsObject,
        region?: Rect.AsObject,
        pattern?: Pattern.AsObject,
        opts?: ObserveOptions.AsObject,
    }
}

export class ObserveChangeRequest extends jspb.Message { 

    hasSource(): boolean;
    clearSource(): void;
    getSource(): GrayImage | undefined;
    setSource(value?: GrayImage): ObserveChangeRequest;

    hasRegion(): boolean;
    clearRegion(): void;
    getRegion(): Rect | undefined;
    setRegion(value?: Rect): ObserveChangeRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): ObserveOptions | undefined;
    setOpts(value?: ObserveOptions): ObserveChangeRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ObserveChangeRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ObserveChangeRequest): ObserveChangeRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ObserveChangeRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ObserveChangeRequest;
    static deserializeBinaryFromReader(message: ObserveChangeRequest, reader: jspb.BinaryReader): ObserveChangeRequest;
}

export namespace ObserveChangeRequest {
    export type AsObject = {
        source?: GrayImage.AsObject,
        region?: Rect.AsObject,
        opts?: ObserveOptions.AsObject,
    }
}

export class ObserveResponse extends jspb.Message { 
    clearEventsList(): void;
    getEventsList(): Array<ObserveEvent>;
    setEventsList(value: Array<ObserveEvent>): ObserveResponse;
    addEvents(value?: ObserveEvent, index?: number): ObserveEvent;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ObserveResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ObserveResponse): ObserveResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ObserveResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ObserveResponse;
    static deserializeBinaryFromReader(message: ObserveResponse, reader: jspb.BinaryReader): ObserveResponse;
}

export namespace ObserveResponse {
    export type AsObject = {
        eventsList: Array<ObserveEvent.AsObject>,
    }
}

export class AppActionRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): AppActionRequest;
    clearArgsList(): void;
    getArgsList(): Array<string>;
    setArgsList(value: Array<string>): AppActionRequest;
    addArgs(value: string, index?: number): string;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): AppOptions | undefined;
    setOpts(value?: AppOptions): AppActionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AppActionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AppActionRequest): AppActionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AppActionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AppActionRequest;
    static deserializeBinaryFromReader(message: AppActionRequest, reader: jspb.BinaryReader): AppActionRequest;
}

export namespace AppActionRequest {
    export type AsObject = {
        name: string,
        argsList: Array<string>,
        opts?: AppOptions.AsObject,
    }
}

export class WindowQueryRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): WindowQueryRequest;

    hasOpts(): boolean;
    clearOpts(): void;
    getOpts(): AppOptions | undefined;
    setOpts(value?: AppOptions): WindowQueryRequest;

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): WindowQuery | undefined;
    setQuery(value?: WindowQuery): WindowQueryRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WindowQueryRequest.AsObject;
    static toObject(includeInstance: boolean, msg: WindowQueryRequest): WindowQueryRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WindowQueryRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WindowQueryRequest;
    static deserializeBinaryFromReader(message: WindowQueryRequest, reader: jspb.BinaryReader): WindowQueryRequest;
}

export namespace WindowQueryRequest {
    export type AsObject = {
        name: string,
        opts?: AppOptions.AsObject,
        query?: WindowQuery.AsObject,
    }
}

export class IsAppRunningResponse extends jspb.Message { 
    getRunning(): boolean;
    setRunning(value: boolean): IsAppRunningResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): IsAppRunningResponse.AsObject;
    static toObject(includeInstance: boolean, msg: IsAppRunningResponse): IsAppRunningResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: IsAppRunningResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): IsAppRunningResponse;
    static deserializeBinaryFromReader(message: IsAppRunningResponse, reader: jspb.BinaryReader): IsAppRunningResponse;
}

export namespace IsAppRunningResponse {
    export type AsObject = {
        running: boolean,
    }
}

export class GetWindowResponse extends jspb.Message { 
    getFound(): boolean;
    setFound(value: boolean): GetWindowResponse;

    hasWindow(): boolean;
    clearWindow(): void;
    getWindow(): Window | undefined;
    setWindow(value?: Window): GetWindowResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetWindowResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetWindowResponse): GetWindowResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetWindowResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetWindowResponse;
    static deserializeBinaryFromReader(message: GetWindowResponse, reader: jspb.BinaryReader): GetWindowResponse;
}

export namespace GetWindowResponse {
    export type AsObject = {
        found: boolean,
        window?: Window.AsObject,
    }
}

export class ListWindowsResponse extends jspb.Message { 
    clearWindowsList(): void;
    getWindowsList(): Array<Window>;
    setWindowsList(value: Array<Window>): ListWindowsResponse;
    addWindows(value?: Window, index?: number): Window;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListWindowsResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListWindowsResponse): ListWindowsResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListWindowsResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListWindowsResponse;
    static deserializeBinaryFromReader(message: ListWindowsResponse, reader: jspb.BinaryReader): ListWindowsResponse;
}

export namespace ListWindowsResponse {
    export type AsObject = {
        windowsList: Array<Window.AsObject>,
    }
}

export enum MatcherEngine {
    MATCHER_ENGINE_UNSPECIFIED = 0,
    MATCHER_ENGINE_TEMPLATE = 1,
    MATCHER_ENGINE_ORB = 2,
    MATCHER_ENGINE_HYBRID = 3,
    MATCHER_ENGINE_AKAZE = 4,
    MATCHER_ENGINE_BRISK = 5,
    MATCHER_ENGINE_KAZE = 6,
    MATCHER_ENGINE_SIFT = 7,
}

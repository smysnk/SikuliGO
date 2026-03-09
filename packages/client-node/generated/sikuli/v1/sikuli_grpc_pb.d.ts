// package: sikuli.v1
// file: sikuli/v1/sikuli.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as sikuli_v1_sikuli_pb from "../../sikuli/v1/sikuli_pb";

interface ISikuliServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listScreens: ISikuliServiceService_IListScreens;
    getPrimaryScreen: ISikuliServiceService_IGetPrimaryScreen;
    captureScreen: ISikuliServiceService_ICaptureScreen;
    find: ISikuliServiceService_IFind;
    findAll: ISikuliServiceService_IFindAll;
    findOnScreen: ISikuliServiceService_IFindOnScreen;
    existsOnScreen: ISikuliServiceService_IExistsOnScreen;
    waitOnScreen: ISikuliServiceService_IWaitOnScreen;
    clickOnScreen: ISikuliServiceService_IClickOnScreen;
    readText: ISikuliServiceService_IReadText;
    findText: ISikuliServiceService_IFindText;
    moveMouse: ISikuliServiceService_IMoveMouse;
    click: ISikuliServiceService_IClick;
    typeText: ISikuliServiceService_ITypeText;
    pasteText: ISikuliServiceService_IPasteText;
    hotkey: ISikuliServiceService_IHotkey;
    mouseDown: ISikuliServiceService_IMouseDown;
    mouseUp: ISikuliServiceService_IMouseUp;
    keyDown: ISikuliServiceService_IKeyDown;
    keyUp: ISikuliServiceService_IKeyUp;
    scrollWheel: ISikuliServiceService_IScrollWheel;
    observeAppear: ISikuliServiceService_IObserveAppear;
    observeVanish: ISikuliServiceService_IObserveVanish;
    observeChange: ISikuliServiceService_IObserveChange;
    openApp: ISikuliServiceService_IOpenApp;
    focusApp: ISikuliServiceService_IFocusApp;
    closeApp: ISikuliServiceService_ICloseApp;
    isAppRunning: ISikuliServiceService_IIsAppRunning;
    listWindows: ISikuliServiceService_IListWindows;
    findWindows: ISikuliServiceService_IFindWindows;
    getWindow: ISikuliServiceService_IGetWindow;
    getFocusedWindow: ISikuliServiceService_IGetFocusedWindow;
}

interface ISikuliServiceService_IListScreens extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ListScreensRequest, sikuli_v1_sikuli_pb.ListScreensResponse> {
    path: "/sikuli.v1.SikuliService/ListScreens";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ListScreensRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ListScreensRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ListScreensResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ListScreensResponse>;
}
interface ISikuliServiceService_IGetPrimaryScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, sikuli_v1_sikuli_pb.GetPrimaryScreenResponse> {
    path: "/sikuli.v1.SikuliService/GetPrimaryScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.GetPrimaryScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.GetPrimaryScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.GetPrimaryScreenResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.GetPrimaryScreenResponse>;
}
interface ISikuliServiceService_ICaptureScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.CaptureScreenRequest, sikuli_v1_sikuli_pb.CaptureScreenResponse> {
    path: "/sikuli.v1.SikuliService/CaptureScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.CaptureScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.CaptureScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.CaptureScreenResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.CaptureScreenResponse>;
}
interface ISikuliServiceService_IFind extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.FindRequest, sikuli_v1_sikuli_pb.FindResponse> {
    path: "/sikuli.v1.SikuliService/Find";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindResponse>;
}
interface ISikuliServiceService_IFindAll extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.FindRequest, sikuli_v1_sikuli_pb.FindAllResponse> {
    path: "/sikuli.v1.SikuliService/FindAll";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindAllResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindAllResponse>;
}
interface ISikuliServiceService_IFindOnScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.FindOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse> {
    path: "/sikuli.v1.SikuliService/FindOnScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindOnScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindOnScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindResponse>;
}
interface ISikuliServiceService_IExistsOnScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ExistsOnScreenRequest, sikuli_v1_sikuli_pb.ExistsOnScreenResponse> {
    path: "/sikuli.v1.SikuliService/ExistsOnScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ExistsOnScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ExistsOnScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ExistsOnScreenResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ExistsOnScreenResponse>;
}
interface ISikuliServiceService_IWaitOnScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.WaitOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse> {
    path: "/sikuli.v1.SikuliService/WaitOnScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.WaitOnScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.WaitOnScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindResponse>;
}
interface ISikuliServiceService_IClickOnScreen extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ClickOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse> {
    path: "/sikuli.v1.SikuliService/ClickOnScreen";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ClickOnScreenRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ClickOnScreenRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindResponse>;
}
interface ISikuliServiceService_IReadText extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ReadTextRequest, sikuli_v1_sikuli_pb.ReadTextResponse> {
    path: "/sikuli.v1.SikuliService/ReadText";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ReadTextRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ReadTextRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ReadTextResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ReadTextResponse>;
}
interface ISikuliServiceService_IFindText extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.FindTextRequest, sikuli_v1_sikuli_pb.FindTextResponse> {
    path: "/sikuli.v1.SikuliService/FindText";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindTextRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindTextRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.FindTextResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.FindTextResponse>;
}
interface ISikuliServiceService_IMoveMouse extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.MoveMouseRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/MoveMouse";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.MoveMouseRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.MoveMouseRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IClick extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/Click";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ClickRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ClickRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_ITypeText extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.TypeTextRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/TypeText";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.TypeTextRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.TypeTextRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IPasteText extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.TypeTextRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/PasteText";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.TypeTextRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.TypeTextRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IHotkey extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/Hotkey";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IMouseDown extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/MouseDown";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ClickRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ClickRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IMouseUp extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/MouseUp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ClickRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ClickRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IKeyDown extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/KeyDown";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IKeyUp extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/KeyUp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.HotkeyRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IScrollWheel extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ScrollWheelRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/ScrollWheel";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ScrollWheelRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ScrollWheelRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IObserveAppear extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ObserveRequest, sikuli_v1_sikuli_pb.ObserveResponse> {
    path: "/sikuli.v1.SikuliService/ObserveAppear";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveResponse>;
}
interface ISikuliServiceService_IObserveVanish extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ObserveRequest, sikuli_v1_sikuli_pb.ObserveResponse> {
    path: "/sikuli.v1.SikuliService/ObserveVanish";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveResponse>;
}
interface ISikuliServiceService_IObserveChange extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.ObserveChangeRequest, sikuli_v1_sikuli_pb.ObserveResponse> {
    path: "/sikuli.v1.SikuliService/ObserveChange";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveChangeRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveChangeRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ObserveResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ObserveResponse>;
}
interface ISikuliServiceService_IOpenApp extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/OpenApp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IFocusApp extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/FocusApp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_ICloseApp extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse> {
    path: "/sikuli.v1.SikuliService/CloseApp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ActionResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ActionResponse>;
}
interface ISikuliServiceService_IIsAppRunning extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.IsAppRunningResponse> {
    path: "/sikuli.v1.SikuliService/IsAppRunning";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.IsAppRunningResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.IsAppRunningResponse>;
}
interface ISikuliServiceService_IListWindows extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ListWindowsResponse> {
    path: "/sikuli.v1.SikuliService/ListWindows";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ListWindowsResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ListWindowsResponse>;
}
interface ISikuliServiceService_IFindWindows extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.WindowQueryRequest, sikuli_v1_sikuli_pb.ListWindowsResponse> {
    path: "/sikuli.v1.SikuliService/FindWindows";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.WindowQueryRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.WindowQueryRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.ListWindowsResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.ListWindowsResponse>;
}
interface ISikuliServiceService_IGetWindow extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.WindowQueryRequest, sikuli_v1_sikuli_pb.GetWindowResponse> {
    path: "/sikuli.v1.SikuliService/GetWindow";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.WindowQueryRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.WindowQueryRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.GetWindowResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.GetWindowResponse>;
}
interface ISikuliServiceService_IGetFocusedWindow extends grpc.MethodDefinition<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.GetWindowResponse> {
    path: "/sikuli.v1.SikuliService/GetFocusedWindow";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    requestDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.AppActionRequest>;
    responseSerialize: grpc.serialize<sikuli_v1_sikuli_pb.GetWindowResponse>;
    responseDeserialize: grpc.deserialize<sikuli_v1_sikuli_pb.GetWindowResponse>;
}

export const SikuliServiceService: ISikuliServiceService;

export interface ISikuliServiceServer extends grpc.UntypedServiceImplementation {
    listScreens: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ListScreensRequest, sikuli_v1_sikuli_pb.ListScreensResponse>;
    getPrimaryScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, sikuli_v1_sikuli_pb.GetPrimaryScreenResponse>;
    captureScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.CaptureScreenRequest, sikuli_v1_sikuli_pb.CaptureScreenResponse>;
    find: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.FindRequest, sikuli_v1_sikuli_pb.FindResponse>;
    findAll: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.FindRequest, sikuli_v1_sikuli_pb.FindAllResponse>;
    findOnScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.FindOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse>;
    existsOnScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ExistsOnScreenRequest, sikuli_v1_sikuli_pb.ExistsOnScreenResponse>;
    waitOnScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.WaitOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse>;
    clickOnScreen: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ClickOnScreenRequest, sikuli_v1_sikuli_pb.FindResponse>;
    readText: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ReadTextRequest, sikuli_v1_sikuli_pb.ReadTextResponse>;
    findText: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.FindTextRequest, sikuli_v1_sikuli_pb.FindTextResponse>;
    moveMouse: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.MoveMouseRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    click: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    typeText: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.TypeTextRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    pasteText: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.TypeTextRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    hotkey: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    mouseDown: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    mouseUp: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ClickRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    keyDown: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    keyUp: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.HotkeyRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    scrollWheel: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ScrollWheelRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    observeAppear: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ObserveRequest, sikuli_v1_sikuli_pb.ObserveResponse>;
    observeVanish: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ObserveRequest, sikuli_v1_sikuli_pb.ObserveResponse>;
    observeChange: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.ObserveChangeRequest, sikuli_v1_sikuli_pb.ObserveResponse>;
    openApp: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    focusApp: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    closeApp: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ActionResponse>;
    isAppRunning: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.IsAppRunningResponse>;
    listWindows: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.ListWindowsResponse>;
    findWindows: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.WindowQueryRequest, sikuli_v1_sikuli_pb.ListWindowsResponse>;
    getWindow: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.WindowQueryRequest, sikuli_v1_sikuli_pb.GetWindowResponse>;
    getFocusedWindow: grpc.handleUnaryCall<sikuli_v1_sikuli_pb.AppActionRequest, sikuli_v1_sikuli_pb.GetWindowResponse>;
}

export interface ISikuliServiceClient {
    listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    find(request: sikuli_v1_sikuli_pb.FindRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    find(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    find(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    findAll(request: sikuli_v1_sikuli_pb.FindRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    findAll(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    findAll(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    findText(request: sikuli_v1_sikuli_pb.FindTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    findText(request: sikuli_v1_sikuli_pb.FindTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    findText(request: sikuli_v1_sikuli_pb.FindTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    click(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    click(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    click(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
}

export class SikuliServiceClient extends grpc.Client implements ISikuliServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    public listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    public listScreens(request: sikuli_v1_sikuli_pb.ListScreensRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListScreensResponse) => void): grpc.ClientUnaryCall;
    public getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    public getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    public getPrimaryScreen(request: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse) => void): grpc.ClientUnaryCall;
    public captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    public captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    public captureScreen(request: sikuli_v1_sikuli_pb.CaptureScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.CaptureScreenResponse) => void): grpc.ClientUnaryCall;
    public find(request: sikuli_v1_sikuli_pb.FindRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public find(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public find(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public findAll(request: sikuli_v1_sikuli_pb.FindRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    public findAll(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    public findAll(request: sikuli_v1_sikuli_pb.FindRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindAllResponse) => void): grpc.ClientUnaryCall;
    public findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public findOnScreen(request: sikuli_v1_sikuli_pb.FindOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    public existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    public existsOnScreen(request: sikuli_v1_sikuli_pb.ExistsOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ExistsOnScreenResponse) => void): grpc.ClientUnaryCall;
    public waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public waitOnScreen(request: sikuli_v1_sikuli_pb.WaitOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public clickOnScreen(request: sikuli_v1_sikuli_pb.ClickOnScreenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindResponse) => void): grpc.ClientUnaryCall;
    public readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    public readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    public readText(request: sikuli_v1_sikuli_pb.ReadTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ReadTextResponse) => void): grpc.ClientUnaryCall;
    public findText(request: sikuli_v1_sikuli_pb.FindTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    public findText(request: sikuli_v1_sikuli_pb.FindTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    public findText(request: sikuli_v1_sikuli_pb.FindTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.FindTextResponse) => void): grpc.ClientUnaryCall;
    public moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public moveMouse(request: sikuli_v1_sikuli_pb.MoveMouseRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public click(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public click(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public click(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public typeText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public pasteText(request: sikuli_v1_sikuli_pb.TypeTextRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public hotkey(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseDown(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public mouseUp(request: sikuli_v1_sikuli_pb.ClickRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyDown(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public keyUp(request: sikuli_v1_sikuli_pb.HotkeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public scrollWheel(request: sikuli_v1_sikuli_pb.ScrollWheelRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeAppear(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeVanish(request: sikuli_v1_sikuli_pb.ObserveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public observeChange(request: sikuli_v1_sikuli_pb.ObserveChangeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ObserveResponse) => void): grpc.ClientUnaryCall;
    public openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public openApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public focusApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public closeApp(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ActionResponse) => void): grpc.ClientUnaryCall;
    public isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    public isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    public isAppRunning(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.IsAppRunningResponse) => void): grpc.ClientUnaryCall;
    public listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public listWindows(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public findWindows(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.ListWindowsResponse) => void): grpc.ClientUnaryCall;
    public getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    public getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    public getWindow(request: sikuli_v1_sikuli_pb.WindowQueryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    public getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    public getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
    public getFocusedWindow(request: sikuli_v1_sikuli_pb.AppActionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sikuli_v1_sikuli_pb.GetWindowResponse) => void): grpc.ClientUnaryCall;
}

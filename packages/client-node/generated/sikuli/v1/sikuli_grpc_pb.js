// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var sikuli_v1_sikuli_pb = require('../../sikuli/v1/sikuli_pb.js');

function serialize_sikuli_v1_ActionResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ActionResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ActionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ActionResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ActionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_AppActionRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.AppActionRequest)) {
    throw new Error('Expected argument of type sikuli.v1.AppActionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_AppActionRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.AppActionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_CaptureScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.CaptureScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.CaptureScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_CaptureScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.CaptureScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_CaptureScreenResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.CaptureScreenResponse)) {
    throw new Error('Expected argument of type sikuli.v1.CaptureScreenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_CaptureScreenResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.CaptureScreenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ClickOnScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ClickOnScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ClickOnScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ClickOnScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ClickOnScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ClickRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ClickRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ClickRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ClickRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ClickRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ExistsOnScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ExistsOnScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ExistsOnScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ExistsOnScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ExistsOnScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ExistsOnScreenResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ExistsOnScreenResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ExistsOnScreenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ExistsOnScreenResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ExistsOnScreenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindAllResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindAllResponse)) {
    throw new Error('Expected argument of type sikuli.v1.FindAllResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindAllResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindAllResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindOnScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindOnScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.FindOnScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindOnScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindOnScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindRequest)) {
    throw new Error('Expected argument of type sikuli.v1.FindRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindResponse)) {
    throw new Error('Expected argument of type sikuli.v1.FindResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindTextRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindTextRequest)) {
    throw new Error('Expected argument of type sikuli.v1.FindTextRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindTextRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindTextRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_FindTextResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.FindTextResponse)) {
    throw new Error('Expected argument of type sikuli.v1.FindTextResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_FindTextResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.FindTextResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_GetPrimaryScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.GetPrimaryScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.GetPrimaryScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_GetPrimaryScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.GetPrimaryScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_GetPrimaryScreenResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.GetPrimaryScreenResponse)) {
    throw new Error('Expected argument of type sikuli.v1.GetPrimaryScreenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_GetPrimaryScreenResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.GetPrimaryScreenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_GetWindowResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.GetWindowResponse)) {
    throw new Error('Expected argument of type sikuli.v1.GetWindowResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_GetWindowResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.GetWindowResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_HotkeyRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.HotkeyRequest)) {
    throw new Error('Expected argument of type sikuli.v1.HotkeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_HotkeyRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.HotkeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_IsAppRunningResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.IsAppRunningResponse)) {
    throw new Error('Expected argument of type sikuli.v1.IsAppRunningResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_IsAppRunningResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.IsAppRunningResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ListScreensRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ListScreensRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ListScreensRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ListScreensRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ListScreensRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ListScreensResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ListScreensResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ListScreensResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ListScreensResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ListScreensResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ListWindowsResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ListWindowsResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ListWindowsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ListWindowsResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ListWindowsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_MoveMouseRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.MoveMouseRequest)) {
    throw new Error('Expected argument of type sikuli.v1.MoveMouseRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_MoveMouseRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.MoveMouseRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ObserveChangeRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ObserveChangeRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ObserveChangeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ObserveChangeRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ObserveChangeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ObserveRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ObserveRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ObserveRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ObserveRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ObserveRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ObserveResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ObserveResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ObserveResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ObserveResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ObserveResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ReadTextRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ReadTextRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ReadTextRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ReadTextRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ReadTextRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ReadTextResponse(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ReadTextResponse)) {
    throw new Error('Expected argument of type sikuli.v1.ReadTextResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ReadTextResponse(buffer_arg) {
  return sikuli_v1_sikuli_pb.ReadTextResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_ScrollWheelRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.ScrollWheelRequest)) {
    throw new Error('Expected argument of type sikuli.v1.ScrollWheelRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_ScrollWheelRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.ScrollWheelRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_TypeTextRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.TypeTextRequest)) {
    throw new Error('Expected argument of type sikuli.v1.TypeTextRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_TypeTextRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.TypeTextRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_WaitOnScreenRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.WaitOnScreenRequest)) {
    throw new Error('Expected argument of type sikuli.v1.WaitOnScreenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_WaitOnScreenRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.WaitOnScreenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sikuli_v1_WindowQueryRequest(arg) {
  if (!(arg instanceof sikuli_v1_sikuli_pb.WindowQueryRequest)) {
    throw new Error('Expected argument of type sikuli.v1.WindowQueryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sikuli_v1_WindowQueryRequest(buffer_arg) {
  return sikuli_v1_sikuli_pb.WindowQueryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var SikuliServiceService = exports.SikuliServiceService = {
  listScreens: {
    path: '/sikuli.v1.SikuliService/ListScreens',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ListScreensRequest,
    responseType: sikuli_v1_sikuli_pb.ListScreensResponse,
    requestSerialize: serialize_sikuli_v1_ListScreensRequest,
    requestDeserialize: deserialize_sikuli_v1_ListScreensRequest,
    responseSerialize: serialize_sikuli_v1_ListScreensResponse,
    responseDeserialize: deserialize_sikuli_v1_ListScreensResponse,
  },
  getPrimaryScreen: {
    path: '/sikuli.v1.SikuliService/GetPrimaryScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.GetPrimaryScreenRequest,
    responseType: sikuli_v1_sikuli_pb.GetPrimaryScreenResponse,
    requestSerialize: serialize_sikuli_v1_GetPrimaryScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_GetPrimaryScreenRequest,
    responseSerialize: serialize_sikuli_v1_GetPrimaryScreenResponse,
    responseDeserialize: deserialize_sikuli_v1_GetPrimaryScreenResponse,
  },
  captureScreen: {
    path: '/sikuli.v1.SikuliService/CaptureScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.CaptureScreenRequest,
    responseType: sikuli_v1_sikuli_pb.CaptureScreenResponse,
    requestSerialize: serialize_sikuli_v1_CaptureScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_CaptureScreenRequest,
    responseSerialize: serialize_sikuli_v1_CaptureScreenResponse,
    responseDeserialize: deserialize_sikuli_v1_CaptureScreenResponse,
  },
  find: {
    path: '/sikuli.v1.SikuliService/Find',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.FindRequest,
    responseType: sikuli_v1_sikuli_pb.FindResponse,
    requestSerialize: serialize_sikuli_v1_FindRequest,
    requestDeserialize: deserialize_sikuli_v1_FindRequest,
    responseSerialize: serialize_sikuli_v1_FindResponse,
    responseDeserialize: deserialize_sikuli_v1_FindResponse,
  },
  findAll: {
    path: '/sikuli.v1.SikuliService/FindAll',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.FindRequest,
    responseType: sikuli_v1_sikuli_pb.FindAllResponse,
    requestSerialize: serialize_sikuli_v1_FindRequest,
    requestDeserialize: deserialize_sikuli_v1_FindRequest,
    responseSerialize: serialize_sikuli_v1_FindAllResponse,
    responseDeserialize: deserialize_sikuli_v1_FindAllResponse,
  },
  findOnScreen: {
    path: '/sikuli.v1.SikuliService/FindOnScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.FindOnScreenRequest,
    responseType: sikuli_v1_sikuli_pb.FindResponse,
    requestSerialize: serialize_sikuli_v1_FindOnScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_FindOnScreenRequest,
    responseSerialize: serialize_sikuli_v1_FindResponse,
    responseDeserialize: deserialize_sikuli_v1_FindResponse,
  },
  existsOnScreen: {
    path: '/sikuli.v1.SikuliService/ExistsOnScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ExistsOnScreenRequest,
    responseType: sikuli_v1_sikuli_pb.ExistsOnScreenResponse,
    requestSerialize: serialize_sikuli_v1_ExistsOnScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_ExistsOnScreenRequest,
    responseSerialize: serialize_sikuli_v1_ExistsOnScreenResponse,
    responseDeserialize: deserialize_sikuli_v1_ExistsOnScreenResponse,
  },
  waitOnScreen: {
    path: '/sikuli.v1.SikuliService/WaitOnScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.WaitOnScreenRequest,
    responseType: sikuli_v1_sikuli_pb.FindResponse,
    requestSerialize: serialize_sikuli_v1_WaitOnScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_WaitOnScreenRequest,
    responseSerialize: serialize_sikuli_v1_FindResponse,
    responseDeserialize: deserialize_sikuli_v1_FindResponse,
  },
  clickOnScreen: {
    path: '/sikuli.v1.SikuliService/ClickOnScreen',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ClickOnScreenRequest,
    responseType: sikuli_v1_sikuli_pb.FindResponse,
    requestSerialize: serialize_sikuli_v1_ClickOnScreenRequest,
    requestDeserialize: deserialize_sikuli_v1_ClickOnScreenRequest,
    responseSerialize: serialize_sikuli_v1_FindResponse,
    responseDeserialize: deserialize_sikuli_v1_FindResponse,
  },
  readText: {
    path: '/sikuli.v1.SikuliService/ReadText',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ReadTextRequest,
    responseType: sikuli_v1_sikuli_pb.ReadTextResponse,
    requestSerialize: serialize_sikuli_v1_ReadTextRequest,
    requestDeserialize: deserialize_sikuli_v1_ReadTextRequest,
    responseSerialize: serialize_sikuli_v1_ReadTextResponse,
    responseDeserialize: deserialize_sikuli_v1_ReadTextResponse,
  },
  findText: {
    path: '/sikuli.v1.SikuliService/FindText',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.FindTextRequest,
    responseType: sikuli_v1_sikuli_pb.FindTextResponse,
    requestSerialize: serialize_sikuli_v1_FindTextRequest,
    requestDeserialize: deserialize_sikuli_v1_FindTextRequest,
    responseSerialize: serialize_sikuli_v1_FindTextResponse,
    responseDeserialize: deserialize_sikuli_v1_FindTextResponse,
  },
  moveMouse: {
    path: '/sikuli.v1.SikuliService/MoveMouse',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.MoveMouseRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_MoveMouseRequest,
    requestDeserialize: deserialize_sikuli_v1_MoveMouseRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  click: {
    path: '/sikuli.v1.SikuliService/Click',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ClickRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_ClickRequest,
    requestDeserialize: deserialize_sikuli_v1_ClickRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  typeText: {
    path: '/sikuli.v1.SikuliService/TypeText',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.TypeTextRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_TypeTextRequest,
    requestDeserialize: deserialize_sikuli_v1_TypeTextRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  pasteText: {
    path: '/sikuli.v1.SikuliService/PasteText',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.TypeTextRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_TypeTextRequest,
    requestDeserialize: deserialize_sikuli_v1_TypeTextRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  hotkey: {
    path: '/sikuli.v1.SikuliService/Hotkey',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.HotkeyRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_HotkeyRequest,
    requestDeserialize: deserialize_sikuli_v1_HotkeyRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  mouseDown: {
    path: '/sikuli.v1.SikuliService/MouseDown',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ClickRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_ClickRequest,
    requestDeserialize: deserialize_sikuli_v1_ClickRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  mouseUp: {
    path: '/sikuli.v1.SikuliService/MouseUp',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ClickRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_ClickRequest,
    requestDeserialize: deserialize_sikuli_v1_ClickRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  keyDown: {
    path: '/sikuli.v1.SikuliService/KeyDown',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.HotkeyRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_HotkeyRequest,
    requestDeserialize: deserialize_sikuli_v1_HotkeyRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  keyUp: {
    path: '/sikuli.v1.SikuliService/KeyUp',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.HotkeyRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_HotkeyRequest,
    requestDeserialize: deserialize_sikuli_v1_HotkeyRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  scrollWheel: {
    path: '/sikuli.v1.SikuliService/ScrollWheel',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ScrollWheelRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_ScrollWheelRequest,
    requestDeserialize: deserialize_sikuli_v1_ScrollWheelRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  observeAppear: {
    path: '/sikuli.v1.SikuliService/ObserveAppear',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ObserveRequest,
    responseType: sikuli_v1_sikuli_pb.ObserveResponse,
    requestSerialize: serialize_sikuli_v1_ObserveRequest,
    requestDeserialize: deserialize_sikuli_v1_ObserveRequest,
    responseSerialize: serialize_sikuli_v1_ObserveResponse,
    responseDeserialize: deserialize_sikuli_v1_ObserveResponse,
  },
  observeVanish: {
    path: '/sikuli.v1.SikuliService/ObserveVanish',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ObserveRequest,
    responseType: sikuli_v1_sikuli_pb.ObserveResponse,
    requestSerialize: serialize_sikuli_v1_ObserveRequest,
    requestDeserialize: deserialize_sikuli_v1_ObserveRequest,
    responseSerialize: serialize_sikuli_v1_ObserveResponse,
    responseDeserialize: deserialize_sikuli_v1_ObserveResponse,
  },
  observeChange: {
    path: '/sikuli.v1.SikuliService/ObserveChange',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.ObserveChangeRequest,
    responseType: sikuli_v1_sikuli_pb.ObserveResponse,
    requestSerialize: serialize_sikuli_v1_ObserveChangeRequest,
    requestDeserialize: deserialize_sikuli_v1_ObserveChangeRequest,
    responseSerialize: serialize_sikuli_v1_ObserveResponse,
    responseDeserialize: deserialize_sikuli_v1_ObserveResponse,
  },
  openApp: {
    path: '/sikuli.v1.SikuliService/OpenApp',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  focusApp: {
    path: '/sikuli.v1.SikuliService/FocusApp',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  closeApp: {
    path: '/sikuli.v1.SikuliService/CloseApp',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.ActionResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_ActionResponse,
    responseDeserialize: deserialize_sikuli_v1_ActionResponse,
  },
  isAppRunning: {
    path: '/sikuli.v1.SikuliService/IsAppRunning',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.IsAppRunningResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_IsAppRunningResponse,
    responseDeserialize: deserialize_sikuli_v1_IsAppRunningResponse,
  },
  listWindows: {
    path: '/sikuli.v1.SikuliService/ListWindows',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.ListWindowsResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_ListWindowsResponse,
    responseDeserialize: deserialize_sikuli_v1_ListWindowsResponse,
  },
  findWindows: {
    path: '/sikuli.v1.SikuliService/FindWindows',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.WindowQueryRequest,
    responseType: sikuli_v1_sikuli_pb.ListWindowsResponse,
    requestSerialize: serialize_sikuli_v1_WindowQueryRequest,
    requestDeserialize: deserialize_sikuli_v1_WindowQueryRequest,
    responseSerialize: serialize_sikuli_v1_ListWindowsResponse,
    responseDeserialize: deserialize_sikuli_v1_ListWindowsResponse,
  },
  getWindow: {
    path: '/sikuli.v1.SikuliService/GetWindow',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.WindowQueryRequest,
    responseType: sikuli_v1_sikuli_pb.GetWindowResponse,
    requestSerialize: serialize_sikuli_v1_WindowQueryRequest,
    requestDeserialize: deserialize_sikuli_v1_WindowQueryRequest,
    responseSerialize: serialize_sikuli_v1_GetWindowResponse,
    responseDeserialize: deserialize_sikuli_v1_GetWindowResponse,
  },
  getFocusedWindow: {
    path: '/sikuli.v1.SikuliService/GetFocusedWindow',
    requestStream: false,
    responseStream: false,
    requestType: sikuli_v1_sikuli_pb.AppActionRequest,
    responseType: sikuli_v1_sikuli_pb.GetWindowResponse,
    requestSerialize: serialize_sikuli_v1_AppActionRequest,
    requestDeserialize: deserialize_sikuli_v1_AppActionRequest,
    responseSerialize: serialize_sikuli_v1_GetWindowResponse,
    responseDeserialize: deserialize_sikuli_v1_GetWindowResponse,
  },
};

exports.SikuliServiceClient = grpc.makeGenericClientConstructor(SikuliServiceService, 'SikuliService');

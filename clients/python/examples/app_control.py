from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo_client.client import SikuliGrpcClient


def main() -> int:
    client = SikuliGrpcClient()
    app_name = "Calculator"
    try:
        client.open_app(pb.AppActionRequest(name=app_name))
        running = client.is_app_running(pb.AppActionRequest(name=app_name))
        print(f"is running => {running.running}")

        windows = client.list_windows(pb.AppActionRequest(name=app_name))
        print(f"windows => {len(windows.windows)}")

        client.focus_app(pb.AppActionRequest(name=app_name))
        client.close_app(pb.AppActionRequest(name=app_name))
        print("app control actions sent")
    finally:
        client.close()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

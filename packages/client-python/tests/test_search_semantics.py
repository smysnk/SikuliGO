from __future__ import annotations

import unittest

import grpc

from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo.client import SikuliError
from sikuligo.sikulix import Region


PNG_BYTES = (
    b"\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x01\x00\x00\x00\x01"
    b"\x08\x04\x00\x00\x00\xb5\x1c\x0c\x02\x00\x00\x00\x0bIDATx\xdac\xfc\xff"
    b"\x1f\x00\x03\x03\x02\x00\xed\xd9\xf1]\x00\x00\x00\x00IEND\xaeB`\x82"
)


def present_response() -> pb.ExistsOnScreenResponse:
    return pb.ExistsOnScreenResponse(
        exists=True,
        match=pb.Match(
            rect=pb.Rect(x=10, y=20, w=30, h=40),
            target=pb.Point(x=25, y=40),
            score=0.95,
            index=0,
        ),
    )


class FakeSession:
    def __init__(self) -> None:
        self.find_calls = 0
        self.wait_calls = 0
        self.exists_calls = 0

    def find_on_screen(self, *args, **kwargs):
        self.find_calls += 1
        raise NotImplementedError

    def wait_on_screen(self, *args, **kwargs):
        self.wait_calls += 1
        raise NotImplementedError

    def exists_on_screen(self, *args, **kwargs):
        self.exists_calls += 1
        raise NotImplementedError


class SearchSemanticsTest(unittest.TestCase):
    def test_find_preserves_transport_not_found_errors(self) -> None:
        session = FakeSession()
        err = SikuliError(grpc.StatusCode.NOT_FOUND, "sikuli: find failed")

        def find_on_screen(*args, **kwargs):
            session.find_calls += 1
            raise err

        session.find_on_screen = find_on_screen
        region = Region(session)

        with self.assertRaises(SikuliError) as ctx:
            region.find(PNG_BYTES)

        self.assertIs(ctx.exception, err)
        self.assertEqual(session.find_calls, 1)

    def test_wait_preserves_transport_deadline_errors(self) -> None:
        session = FakeSession()
        err = SikuliError(grpc.StatusCode.DEADLINE_EXCEEDED, "deadline exceeded")

        def wait_on_screen(*args, **kwargs):
            session.wait_calls += 1
            raise err

        session.wait_on_screen = wait_on_screen
        region = Region(session)

        with self.assertRaises(SikuliError) as ctx:
            region.wait(PNG_BYTES, timeout_millis=25)

        self.assertIs(ctx.exception, err)
        self.assertEqual(session.wait_calls, 1)

    def test_wait_vanish_returns_true_immediately_when_target_is_absent(self) -> None:
        session = FakeSession()

        def exists_on_screen(*args, **kwargs):
            session.exists_calls += 1
            return pb.ExistsOnScreenResponse(exists=False)

        session.exists_on_screen = exists_on_screen
        region = Region(session)

        vanished = region.wait_vanish(PNG_BYTES, timeout_millis=0)

        self.assertTrue(vanished)
        self.assertEqual(session.exists_calls, 1)

    def test_wait_vanish_retries_until_target_disappears(self) -> None:
        session = FakeSession()

        def exists_on_screen(*args, **kwargs):
            session.exists_calls += 1
            if session.exists_calls < 3:
                return present_response()
            return pb.ExistsOnScreenResponse(exists=False)

        session.exists_on_screen = exists_on_screen
        region = Region(session)

        vanished = region.wait_vanish(PNG_BYTES, timeout_millis=50, interval_millis=1)

        self.assertTrue(vanished)
        self.assertEqual(session.exists_calls, 3)

    def test_wait_vanish_returns_false_on_timeout_without_throwing(self) -> None:
        session = FakeSession()

        def exists_on_screen(*args, **kwargs):
            session.exists_calls += 1
            return present_response()

        session.exists_on_screen = exists_on_screen
        region = Region(session)

        vanished = region.wait_vanish(PNG_BYTES, timeout_millis=5, interval_millis=1)

        self.assertFalse(vanished)
        self.assertGreaterEqual(session.exists_calls, 1)


if __name__ == "__main__":
    unittest.main()

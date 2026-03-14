from __future__ import annotations

import ast
import sys
from pathlib import Path


def _step_expr(node: ast.stmt, counter: int) -> ast.Expr:
    expression = ast.Expr(
        value=ast.Call(
            func=ast.Name(id="__sikuli_step", ctx=ast.Load()),
            args=[
                ast.Constant(getattr(node, "lineno", None)),
                ast.Constant(getattr(node, "col_offset", 0) + 1 if hasattr(node, "col_offset") else None),
                ast.Constant(f"python-{getattr(node, 'lineno', 'na')}-{counter}"),
            ],
            keywords=[],
        )
    )
    return ast.copy_location(expression, node)


class StepInjector(ast.NodeTransformer):
    def __init__(self) -> None:
        self.counter = 0

    def _should_instrument(self, node: ast.stmt) -> bool:
        return not isinstance(
            node,
            (
                ast.Import,
                ast.ImportFrom,
                ast.FunctionDef,
                ast.AsyncFunctionDef,
                ast.ClassDef,
                ast.Global,
                ast.Nonlocal,
                ast.Pass,
            ),
        )

    def _instrument_body(self, body: list[ast.stmt], *, preserve_future_imports: bool = False) -> list[ast.stmt]:
        output: list[ast.stmt] = []
        start_index = 0

        if preserve_future_imports:
            while start_index < len(body):
                statement = body[start_index]
                if isinstance(statement, ast.ImportFrom) and statement.module == "__future__":
                    output.append(self.visit(statement))
                    start_index += 1
                    continue
                break

        for statement in body[start_index:]:
            visited = self.visit(statement)
            if visited is None:
                continue
            statements = visited if isinstance(visited, list) else [visited]
            for item in statements:
                if self._should_instrument(item):
                    self.counter += 1
                    output.append(_step_expr(item, self.counter))
                output.append(item)

        return output

    def visit_Module(self, node: ast.Module) -> ast.Module:
        node.body = self._instrument_body(node.body, preserve_future_imports=True)
        return node

    def visit_FunctionDef(self, node: ast.FunctionDef) -> ast.FunctionDef:
        node.body = self._instrument_body(node.body)
        return node

    def visit_AsyncFunctionDef(self, node: ast.AsyncFunctionDef) -> ast.AsyncFunctionDef:
        node.body = self._instrument_body(node.body)
        return node

    def visit_ClassDef(self, node: ast.ClassDef) -> ast.ClassDef:
        node.body = self._instrument_body(node.body)
        return node

    def visit_For(self, node: ast.For) -> ast.For:
        node.body = self._instrument_body(node.body)
        node.orelse = self._instrument_body(node.orelse)
        return node

    def visit_AsyncFor(self, node: ast.AsyncFor) -> ast.AsyncFor:
        node.body = self._instrument_body(node.body)
        node.orelse = self._instrument_body(node.orelse)
        return node

    def visit_While(self, node: ast.While) -> ast.While:
        node.body = self._instrument_body(node.body)
        node.orelse = self._instrument_body(node.orelse)
        return node

    def visit_If(self, node: ast.If) -> ast.If:
        node.body = self._instrument_body(node.body)
        node.orelse = self._instrument_body(node.orelse)
        return node

    def visit_With(self, node: ast.With) -> ast.With:
        node.body = self._instrument_body(node.body)
        return node

    def visit_AsyncWith(self, node: ast.AsyncWith) -> ast.AsyncWith:
        node.body = self._instrument_body(node.body)
        return node

    def visit_Try(self, node: ast.Try) -> ast.Try:
        node.body = self._instrument_body(node.body)
        node.orelse = self._instrument_body(node.orelse)
        node.finalbody = self._instrument_body(node.finalbody)
        node.handlers = [self.visit(handler) for handler in node.handlers]
        return node

    def visit_ExceptHandler(self, node: ast.ExceptHandler) -> ast.ExceptHandler:
        node.body = self._instrument_body(node.body)
        return node


def transform_source(source: str) -> str:
    tree = ast.parse(source)
    injector = StepInjector()
    transformed = injector.visit(tree)
    ast.fix_missing_locations(transformed)
    return f"{ast.unparse(transformed)}\n"


def main(argv: list[str]) -> int:
    if len(argv) != 3:
        print("usage: python-transform.py <input.py> <output.py>", file=sys.stderr)
        return 2

    input_path = Path(argv[1])
    output_path = Path(argv[2])
    output_path.write_text(transform_source(input_path.read_text(encoding="utf-8")), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))

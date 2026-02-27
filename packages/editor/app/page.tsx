"use client";

import { useEffect, useMemo, useRef, useState, type PointerEvent as ReactPointerEvent } from "react";
import styles from "./page.module.css";

const SIDEBAR_MIN = 220;
const SIDEBAR_MAX = 560;
const SIDEBAR_DEFAULT = 300;
const STORAGE_KEY = "sikuligo.editor.workflows.sidebar.width";

type Workflow = {
  id: string;
  name: string;
  updatedAt: string;
};

const WORKFLOWS: Workflow[] = [
  { id: "wf-1", name: "Login And Dashboard Check", updatedAt: "Updated 2m ago" },
  { id: "wf-2", name: "Visual Regression Snapshot", updatedAt: "Updated 14m ago" },
  { id: "wf-3", name: "Daily OCR Validation", updatedAt: "Updated 52m ago" }
];

function clampWidth(width: number): number {
  return Math.max(SIDEBAR_MIN, Math.min(SIDEBAR_MAX, width));
}

export default function Home() {
  const [sidebarWidth, setSidebarWidth] = useState(SIDEBAR_DEFAULT);
  const [dragging, setDragging] = useState(false);
  const dragState = useRef<{ startX: number; startWidth: number } | null>(null);

  useEffect(() => {
    const raw = window.localStorage.getItem(STORAGE_KEY);
    if (!raw) {
      return;
    }
    const parsed = Number(raw);
    if (Number.isFinite(parsed)) {
      setSidebarWidth(clampWidth(parsed));
    }
  }, []);

  useEffect(() => {
    window.localStorage.setItem(STORAGE_KEY, String(sidebarWidth));
  }, [sidebarWidth]);

  useEffect(() => {
    const onMove = (event: globalThis.PointerEvent) => {
      if (!dragState.current) {
        return;
      }
      const delta = event.clientX - dragState.current.startX;
      setSidebarWidth(clampWidth(dragState.current.startWidth + delta));
    };

    const onUp = () => {
      dragState.current = null;
      setDragging(false);
    };

    window.addEventListener("pointermove", onMove);
    window.addEventListener("pointerup", onUp);
    return () => {
      window.removeEventListener("pointermove", onMove);
      window.removeEventListener("pointerup", onUp);
    };
  }, []);

  const sidebarStyle = useMemo(() => ({ width: `${sidebarWidth}px` }), [sidebarWidth]);

  const onResizePointerDown = (event: ReactPointerEvent<HTMLDivElement>) => {
    dragState.current = { startX: event.clientX, startWidth: sidebarWidth };
    setDragging(true);
    event.currentTarget.setPointerCapture(event.pointerId);
  };

  return (
    <div className={styles.root} style={dragging ? { userSelect: "none" } : undefined}>
      <aside className={styles.sidebar} style={sidebarStyle}>
        <h2 className={styles.sidebarHeader}>Workflows</h2>
        <ul className={styles.workflowList}>
          {WORKFLOWS.map((workflow) => (
            <li key={workflow.id} className={styles.workflowItem}>
              <div className={styles.workflowItemTitle}>{workflow.name}</div>
              <div className={styles.workflowItemMeta}>{workflow.updatedAt}</div>
            </li>
          ))}
        </ul>
      </aside>

      <div
        className={styles.resizer}
        role="separator"
        aria-orientation="vertical"
        aria-label="Resize workflows sidebar"
        onPointerDown={onResizePointerDown}
        onDoubleClick={() => setSidebarWidth(SIDEBAR_DEFAULT)}
      />

      <main className={styles.main}>
        <section className={styles.panel}>
          <h1>SikuliGO Editor</h1>
          <p>
            The workflows sidebar is resizable. Drag the divider or double-click it to reset width.
          </p>
        </section>
      </main>
    </div>
  );
}

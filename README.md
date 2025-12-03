# Advent of Code — Solutions

This repository collects my Advent of Code solutions. The primary focus for now is **2025**, but the structure and checklist support adding solutions for any year.

Quick links

- Year folder: `2025/`
- Example day: `2025/day_1/`

Goals

- Keep one clear solution directory per year and day.
- Track progress with checklists (per-day and per-year).
- Make it easy to run and test solutions locally.

Repository layout

```go
README.md
2025/
    day_1/
        go.mod
        main.go
    day_2/
        go.mod      
        main.go
```

How I organize each day

- `year/day_N/` — contains one solution (or small module) per day.
- `input.txt` — (optional) put puzzle input here if needed.
- `main.go` or `solution.py` etc. — language-specific solution file.

Run a solution (Go example)

```go
cd 2025/day_1
go run .
```

Checklist & templates

Per-day checklist (use this in issue, project board, or in the day folder README):

- [ ] **Day NN: Title** — short description or link to puzzle
  - [ ] **Problem:** read/understand puzzle
  - [ ] **Part 1:** implemented
  - [ ] **Part 2:** implemented
  - [ ] **Tests:** unit tests added (if applicable)
  - [ ] **Input file:** `input.txt` present (if used)
  - [ ] **Runtime:** documented (command to run)
  - [ ] **Notes:** key observations, complexity, pitfalls
Per-year checklist (example for 2025)

- [ ] Year 2025 — complete all days 1..25
  - [ ] Day 01 — `2025/day_1/` — [ ] P1 [ ] P2
  - [ ] Day 02 — `2025/day_2/` — [ ] P1 [ ] P2
  - [ ] Day 03 — `2025/day_3/` — [ ] P1 [ ] P2
  - [ ] ...
  - [ ] Day 25 — `2025/day_25/` — [ ] P1 [ ] P2
Per-repo checklist (meta)

- [ ] Add year folders for each AoC year you want to store
- [ ] Add license and contribution guidelines (if collaborating)
- [ ] Add CI (optional) to run tests for each day

Conventions

- Name daily folders `day_1`, `day_2`, ... `day_25`.
- Put the puzzle input in `input.txt` (if not committed, add `.gitignore` entry for sensitive inputs).
- Keep solutions simple and readable — include short comments about approach and complexity.

Notes on inputs and privacy

- If you prefer not to commit your personal puzzle inputs, add them to `.gitignore` and keep a `README` note explaining how to add a local `input.txt` to run the day locally.

Example `day` README snippet (copy into `2025/day_1/README.md`):

```markdown
    # Day 01 — Calorie Counting

    Run:
    ```go
    go run .
    ```

    Checklist:

    - [ ] Part 1
    - [ ] Part 2

    Notes:

    - [ ] Approach: sort sums, use sliding window, etc.
    - [ ] Complexity: O(n log n) due to sorting
    - [ ] Pitfalls: watch for empty lines in input 
    - [ ] Input: place your puzzle input in `input.txt`
```

Next steps I can do for you

- Add a `2025/README.md` with an auto-generated progress table.
- Add a `.gitignore` entry for input files and a small run script.

## Advent of Code 2025 — Progress

This tracks progress for Advent of Code 2025. Each row links to the day folder and includes checkboxes for Part 1 and Part 2. Update the checkboxes as you complete parts.

How to run a day's solution (Go example):

```bash
cd 2025/day_1
go run .
```

Progress table

| Day | Folder | Part 1 | Part 2 | Notes |
|---:|:---|:---:|:---:|:---|
| 01 | `day_1/` | - [x] | - [x] | |
| 02 | `day_2/` | - [x] | - [x] | |
| 03 | `day_3/` | - [x] | - [ ] | |
| 04 | `day_4/` | - [ ] | - [ ] | |
| 05 | `day_5/` | - [ ] | - [ ] | |
| 06 | `day_6/` | - [ ] | - [ ] | |
| 07 | `day_7/` | - [ ] | - [ ] | |
| 08 | `day_8/` | - [ ] | - [ ] | |
| 09 | `day_9/` | - [ ] | - [ ] | |
| 10 | `day_10/` | - [ ] | - [ ] | |
| 11 | `day_11/` | - [ ] | - [ ] | |
| 12 | `day_12/` | - [ ] | - [ ] | |
| 13 | `day_13/` | - [ ] | - [ ] | |
| 14 | `day_14/` | - [ ] | - [ ] | |
| 15 | `day_15/` | - [ ] | - [ ] | |
| 16 | `day_16/` | - [ ] | - [ ] | |
| 17 | `day_17/` | - [ ] | - [ ] | |
| 18 | `day_18/` | - [ ] | - [ ] | |
| 19 | `day_19/` | - [ ] | - [ ] | |
| 20 | `day_20/` | - [ ] | - [ ] | |
| 21 | `day_21/` | - [ ] | - [ ] | |
| 22 | `day_22/` | - [ ] | - [ ] | |
| 23 | `day_23/` | - [ ] | - [ ] | |
| 24 | `day_24/` | - [ ] | - [ ] | |
| 25 | `day_25/` | - [ ] | - [ ] | |

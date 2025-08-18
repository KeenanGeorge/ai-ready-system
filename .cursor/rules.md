# Cursor Rule — Branching & PRs

**Goal:** Always branch from `features`, use a consistent name, and open PRs against `features` with Linear context.

## Branching
- **Base branch:** `features` (never `main`).
- **Branch name format:** `feature/<ISSUEKEY>-<slug>`
  - Examples: `feature/SMA-10-add-testmo`, `feature/ENG-123-google-login-fix`
- **ISSUEKEYs allowed:** `SMA-\d+` or `ENG-\d+`
- **Slug:** lowercase, `-` separated, letters/numbers only.

## When creating a branch from an issue
1. Include the Linear issue key in the branch name.
2. Ensure base = `features`.
3. Commit message starts with the key, e.g. `SMA-10: integrate Testmo automation`.
4. PR title/body include the key (e.g., `Fixes SMA-10`).

## Pull Requests
- **Target/base:** `features`
- Add a comment with the Testmo run link if available.
- Never open PRs directly to `main`.

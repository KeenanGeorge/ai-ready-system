# Go + Linear + Testmo Demo

This repo demonstrates a lightweight, AI-ready workflow:

- **Linear** for issues & PR automation
- **GitHub Actions** for CI (Go tests → JUnit)
- **Testmo** for test result collection and reporting



## 1) Prerequisites

- GitHub repository (this project)
- Linear workspace (create a demo issue, e.g. `ENG-123`)
- Testmo workspace with a project

Create these **GitHub Actions secrets** in your repo (**Settings → Secrets and variables → Actions**)::

- `TESTMO_TOKEN` — a Testmo API token
- `TESTMO_INSTANCE` — e.g. `https://yourcompany.testmo.net`
- `TESTMO_PROJECT_ID` — numeric id of your Testmo project



## 2) Linear ↔ GitHub integration

1. In Linear, enable the **GitHub integration** (Settings → Integrations → GitHub).
2. Select your repo and enable PR linking.
3. In PR titles or descriptions, include the issue key to auto-link **and** transition states, e.g.:  
   `Fixes ENG-123`



## 3) Run locally (optional)

```bash
go run .
# in another terminal
curl -s localhost:8080/health
# ok
```

Run tests locally:
```bash
go test ./...
```

Create JUnit locally (optional for demo):
```bash
go install gotest.tools/gotestsum@latest
mkdir -p reports
$(go env GOPATH)/bin/gotestsum --junitfile reports/unit-tests.xml -- ./...
```



## 4) CI/CD flow

- On PR or push to `main`, GitHub Actions will:
  1. Run Go unit tests
  2. Emit `reports/unit-tests.xml` (JUnit)
  3. Submit results to **Testmo** with `testmo automation:run:submit`

> Tip: Protect `main` and **require** the `ci` workflow to pass before merge.


## 5) Demo script

1. Create a Linear issue: `ENG-123 Add /health endpoint`.
2. Create a feature branch (you can use Linear's suggested branch name).
3. Open a PR with the title containing `Fixes ENG-123`.
4. Watch the GitHub Action complete; then open Testmo → Project → Automation to see the run.
5. Merge PR; Linear should move the issue to **Done** (depending on your team rules).



## 6) Troubleshooting

- **No results in Testmo**: Ensure secrets are set and the `--results` glob matches `reports/*.xml`.
- **Linear not updating on PR**: Check that the PR references the issue key (e.g., `Fixes ENG-123`) and the GitHub integration is enabled per team.
- **Go tests fail**: Run `go test ./...` locally to debug.



## 7) What to extend next

- Add integration tests (emit more JUnit files into `reports/`).
- Add labels/milestones in Testmo via CLI flags for nicer dashboards.
- Add a PR review step (linters or an AI reviewer) — keep the required checks as the final gate.

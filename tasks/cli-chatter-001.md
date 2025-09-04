# Task: cli-chatter-001 - Add command-line flags for server/client mode

**User Story**: As a developer, I want to run the app in server or client mode so I can set up a chat server or connect as a client.

**Acceptance Criteria**:
- Given the app is built, When I run `go run main.go --server`, Then the app starts in server mode listening on a port.
- Given the app is built, When I run `go run main.go --client <server_address>`, Then the app connects to the server as a client.
- Given invalid flags, When I run the app, Then it shows usage help.

**Dependencies / Assumptions / Risks**:
- Dependencies: Existing main.go structure.
- Assumptions: Use standard Go flag package.
- Risks: None major.

**Test Notes & Observability Hooks**:
- Manual test: Run with flags, check output.
- Observability: Log mode startup to stderr.

**Effort**: S (simple flag parsing, ~2 hours)

**Owners**: @engineer

**PRD References**: #3, #8, #9
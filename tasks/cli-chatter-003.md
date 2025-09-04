# Task: cli-chatter-002 - Implement TCP server for handling connections

**User Story**: As a server, I want to accept TCP connections from clients so users can join the chat.

**Acceptance Criteria**:
- Given server mode is active, When a client attempts to connect, Then the connection is accepted.
- Given multiple clients, When they connect simultaneously, Then all are handled concurrently.
- Given server stops, When connections are active, Then they are gracefully closed.

**Dependencies / Assumptions / Risks**:
- Dependencies: Go net package.
- Assumptions: Port 8080 default.
- Risks: Networking complexity if firewall issues.

**Test Notes & Observability Hooks**:
- Test: Run server, connect with netcat or another instance.
- Observability: Log connection events to stderr.

**Effort**: M (basic TCP server, ~4 hours)

**Owners**: @engineer

**PRD References**: #5 M, #8, #9
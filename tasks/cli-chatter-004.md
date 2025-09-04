# Task: cli-chatter-003 - Implement client connection logic

**User Story**: As a user, I want to connect to the chat server so I can participate in chats.

**Acceptance Criteria**:
- Given client mode with server address, When started, Then connects to the server.
- Given connection fails, When server is down, Then shows error and exits.
- Given connected, When I send input, Then it's forwarded to server.

**Dependencies / Assumptions / Risks**:
- Dependencies: Task 002.
- Assumptions: Server running.
- Risks: Connection timeouts.

**Test Notes & Observability Hooks**:
- Test: Start server, run client, check connection.
- Observability: Log connection status.

**Effort**: M (~4 hours)

**Owners**: @engineer

**PRD References**: #5 M, #8
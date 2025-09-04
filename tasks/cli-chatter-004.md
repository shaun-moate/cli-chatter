# Task: cli-chatter-004 - Implement user authentication (username)

**User Story**: As a user, I want to set a username on connect so I'm identified in the chat.

**Acceptance Criteria**:
- Given client connects, When prompted for username, Then enter and it's set.
- Given username taken, When another user tries same, Then denied or allowed with suffix.
- Given authenticated, When sending message, Then prefixed with username.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 002, 003.
- Assumptions: Simple username, no password.
- Risks: Username conflicts.

**Test Notes & Observability Hooks**:
- Test: Connect two clients, set usernames, send messages.
- Observability: Log username assignments.

**Effort**: S (~2 hours)

**Owners**: @engineer

**PRD References**: #5 S
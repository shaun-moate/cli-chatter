# Task: cli-chatter-011 - Add file-based persistence for messages

**User Story**: As a system, I want to persist messages to file so history survives restarts.

**Acceptance Criteria**:
- Given messages sent, When server restarts, Then history loaded from file.
- Given file not exist, When started, Then created.
- Given large history, When loaded, Then efficient.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 009.
- Assumptions: JSON or simple text file.
- Risks: File corruption.

**Test Notes & Observability Hooks**:
- Test: Send messages, restart, check history.
- Observability: Log persistence operations.

**Effort**: M (~4 hours)

**Owners**: @engineer

**PRD References**: #9

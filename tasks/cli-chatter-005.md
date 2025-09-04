# Task: cli-chatter-005 - Implement channel data structures

**User Story**: As a system, I want channels to organize chats so users can join specific topics.

**Acceptance Criteria**:
- Given server, When initialized, Then default channel #general exists.
- Given channels, When user joins, Then added to channel user list.
- Given message sent to channel, Then broadcasted to all members.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 002.
- Assumptions: In-memory storage for MVP.
- Risks: Concurrency issues with multiple goroutines.

**Test Notes & Observability Hooks**:
- Test: Join channel, send message, check received.
- Observability: Log channel joins.

**Effort**: M (~4 hours)

**Owners**: @engineer

**PRD References**: #5 M, #9
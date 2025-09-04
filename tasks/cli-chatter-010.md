# Task: cli-chatter-010 - Implement message history

**User Story**: As a user, I want to see message history when joining so I can catch up on conversations.

**Acceptance Criteria**:
- Given channel has messages, When I join, Then last 10 messages displayed.
- Given message sent, When stored, Then added to history.
- Given history, When requested /history, Then shows all.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 005, 007.
- Assumptions: In-memory for MVP.
- Risks: Memory usage for large history.

**Test Notes & Observability Hooks**:
- Test: Join channel with history, check displayed.
- Observability: Log history access.

**Effort**: S (~2 hours)

**Owners**: @engineer

**PRD References**: #5 S

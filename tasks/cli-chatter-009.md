# Task: cli-chatter-008 - Implement /msg command for private messages

**User Story**: As a user, I want to send private messages using /msg so I can communicate one-on-one.

**Acceptance Criteria**:
- Given connected users, When I type /msg username message, Then only that user receives it.
- Given user offline, When /msg, Then error message.
- Given received, When private message, Then displayed privately.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 004, 007.
- Assumptions: User list maintained.
- Risks: None.

**Test Notes & Observability Hooks**:
- Test: Send /msg, check only recipient sees.
- Observability: Log private messages.

**Effort**: S (~2 hours)

**Owners**: @engineer

**PRD References**: #5 M